# Overview

This is to send events from TiDB to Kafka with TiCDC. Then the example code is a Kafka consumer that prints some info about received events.

# setup

To use RedPanda as Kafka service

```
docker run -d --pull=always --name=redpanda-1 --rm \
-p 8081:8081 \
-p 8082:8082 \
-p 9092:9092 \
-p 9644:9644 \
docker.redpanda.com/redpandadata/redpanda:latest \
redpanda start \
--overprovisioned \
--smp 1  \
--memory 1G \
--reserve-memory 0M \
--node-id 0 \
--check=false
```

See the [docs](https://docs.redpanda.com/docs/21.11/quickstart/quick-start-docker/). Note that newer versions no longer recommend to run it like this directly.

When running TiUP Playground you need to set `--ticdc 1` like this:
```
tiup playground --ticdc 1 ...
```

TiCDC requires `max.message.bytes` to be set on the broker. This can be done with `rpk`:
```
wget https://github.com/redpanda-data/redpanda/releases/latest/download/rpk-linux-amd64.zip
unzip rpk-linux-amd64.zip 
rm rpk-linux-amd64.zip 
./rpk topic create ticker
./rpk topic alter-config ticker --set max.message.bytes=4294897000
```

Then to create the changefeed:
```
tiup cdc cli changefeed create --sink-uri="kafka://127.0.0.1:9092/ticker?protocol=canal-json&max-message-bytes=67108864" --changefeed-id="ticker"
```

To check the state:
```
$ tiup cdc cli changefeed list
tiup is checking updates for component cdc ...
Starting component `cdc`: /home/dvaneeden/.tiup/components/cdc/v7.2.0/cdc cli changefeed list
[
  {
    "id": "ticker",
    "namespace": "default",
    "summary": {
      "state": "normal",
      "tso": 442985056137379844,
      "checkpoint": "2023-07-20 13:51:22.360",
      "error": null
    }
  }
]
```

To check the contents of the topic:
```
./rpk topic consume ticker
```

Now you can run the example code:
```
go run tickerstat.go
```