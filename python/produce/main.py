import argparse
import sys

from kafka import KafkaProducer

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
producer = KafkaProducer(
    bootstrap_servers=[args.broker],
    security_protocol="SSL" if args.tls else "PLAINTEXT",
    compression_type=None,
    api_version=(0, 8, 0),
)

print(
    "Each line of input will be produced as a separate Kafka message "
    + "to the topic: {}\nPress Ctrl+D to stop\n\nWaiting for input:".format(args.topic)
)


for line in sys.stdin:
    msg = line.rstrip()
    future = producer.send(args.topic, msg.encode("utf-8"))
    msg_meta = future.get(timeout=10)
    print(
        "Produced (topic: {}, partition: {}, offset: {}): {}\nWaiting for input:".format(
            msg_meta.topic, msg_meta.partition, msg_meta.offset, msg
        )
    )
