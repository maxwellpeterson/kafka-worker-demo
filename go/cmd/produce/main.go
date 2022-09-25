package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/maxwellpeterson/kafka-websocket-shim/pkg/shim"
	"github.com/pkg/errors"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kversion"
)

var (
	broker = flag.String("broker", "localhost:8787", "the address of the broker")
	tls    = flag.Bool("tls", false, "use tls for the broker connection")
	topic  = flag.String("topic", "test-topic", "the topic to produce to")
)

func main() {
	flag.Parse()
	dialer := shim.NewDialer(shim.DialerConfig{TLS: *tls})

	fmt.Printf("Connecting to broker: %s\n", *broker)
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(*broker),
		kgo.MaxVersions(kversion.V0_8_0()),
		kgo.Dialer(dialer.DialContext),
		kgo.ProducerBatchCompression(kgo.NoCompression()),
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "create client failed"))
	}
	defer cl.Close()

	fmt.Printf("Each line of input will be produced as a separate Kafka message "+
		"to the topic: %s\nPress Ctrl+D to stop\n\nWaiting for input:\n", *topic)

	scanner := bufio.NewScanner(os.Stdin)
	ctx := context.Background()
	for scanner.Scan() {
		record := &kgo.Record{Topic: *topic, Value: scanner.Bytes()}
		if err := cl.ProduceSync(ctx, record).FirstErr(); err != nil {
			log.Fatal(errors.Wrap(err, "produce record failed"))
		}
		fmt.Printf("Produced message: %s\nWaiting for input:\n", string(record.Value))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(errors.Wrap(err, "scan input failed"))
	}
}
