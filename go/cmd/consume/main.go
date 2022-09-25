package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/maxwellpeterson/kafka-websocket-shim/pkg/shim"
	"github.com/pkg/errors"
	"github.com/twmb/franz-go/pkg/kerr"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kversion"
)

var (
	broker = flag.String("broker", "localhost:8787", "the address of the broker")
	tls    = flag.Bool("tls", false, "use tls for the broker connection")
	topic  = flag.String("topic", "test-topic", "the topic to consume from")
)

func main() {
	flag.Parse()
	dialer := shim.NewDialer(shim.DialerConfig{TLS: *tls})

	fmt.Printf("Connecting to broker: %s\n", *broker)
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(*broker),
		kgo.MaxVersions(kversion.V0_8_0()),
		kgo.ConsumeTopics(*topic),
		kgo.Dialer(dialer.DialContext),
		kgo.ProducerBatchCompression(kgo.NoCompression()),
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "create client failed"))
	}
	defer cl.Close()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sig
		cancel()
	}()

	fmt.Printf("Consuming messages from: %s\nPress Ctrl+C to stop\n\n", *topic)
	for {
		fetches := cl.PollFetches(ctx)
		for _, err := range fetches.Errors() {
			if errors.Is(err.Err, kerr.RequestTimedOut) {
				// Keep polling (isn't this a retriable error?)
				continue
			} else if errors.Is(err.Err, context.Canceled) {
				// User interrupt, stop the program
				return
			} else {
				log.Fatal(errors.Wrap(err.Err, "poll topic failed"))
			}
		}

		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			fmt.Printf("Consumed message: %s\n", string(record.Value))
		}
	}
}
