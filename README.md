# Kafka Worker Demo

A fully local demo of [`kafka-worker`](https://github.com/maxwellpeterson/kafka-worker) and [`kafka-websocket-shim`](https://github.com/maxwellpeterson/kafka-websocket-shim), no Cloudflare account required.

The `go` folder shows a Kafka client written in Go using the Go package provided by `kafka-websocket-shim` to connect to a local version of `kafka-worker`.

The `python` folder shows a Kafka client written in Python using the TCP proxy provided by `kafka-websocket-shim` to connect to a local version of `kafka-worker`.

Both examples look and behave exactly the same to the user, despite using different client libraries and methods for connecting to the broker.

## Video

https://user-images.githubusercontent.com/64494795/192863065-29dc2ceb-ae00-4305-81b6-2ac322b01d2e.mp4
