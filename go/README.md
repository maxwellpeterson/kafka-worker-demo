# Go Demo

This demo uses the [`franz-go`](https://github.com/twmb/franz-go) client and the Go package provided by [`kafka-websocket-shim`](https://github.com/maxwellpeterson/kafka-websocket-shim) to connect to a local version of [`kafka-worker`](https://github.com/maxwellpeterson/kafka-worker).

## Quick Start

This demo requires a local installation of Docker and Docker Compose.

Clone the repository and navigate to this directory:
```shell
git clone https://github.com/maxwellpeterson/kafka-worker-demo.git
cd kafka-worker-demo/go
```

Start the producer client program:
```shell
docker compose run producer
```

Open a second terminal window in the `kafka-worker-demo/go` directory, and start the consumer client program:
```shell
docker compose run consumer
```

Optionally, to view broker logs, open a third terminal window in the `kafka-worker-demo/go` directory and run:
```shell
docker compose logs --follow broker
```

Once you're done, clean things up with:
```shell
docker compose down
```
