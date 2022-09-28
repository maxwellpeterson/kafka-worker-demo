import argparse
import signal
import sys

from kafka import KafkaConsumer

parser = argparse.ArgumentParser()
parser.add_argument(
    "-broker", default="localhost:8787", help="the address of the broker"
)
parser.add_argument(
    "-tls", action="store_true", help="use tls for the broker connection"
)
parser.add_argument("-topic", default="test-topic", help="the topic to produce to")
args = parser.parse_args()

print("Connecting to broker: {}".format(args.broker))

consumer = KafkaConsumer(
    args.topic,
    bootstrap_servers=[args.broker],
    security_protocol="SSL" if args.tls else "PLAINTEXT",
    api_version=(0, 8, 0),
)

print("Consuming messages from topic: {}\nPress Ctrl+C to stop\n".format(args.topic))


def signal_handler(sig, frame):
    consumer.close()
    sys.exit(0)


signal.signal(signal.SIGINT, signal_handler)


for msg in consumer:
    print(
        "Consumed (topic: {}, partition: {}, offset: {}): {}".format(
            msg.topic, msg.partition, msg.offset, msg.value.decode("utf-8")
        )
    )
