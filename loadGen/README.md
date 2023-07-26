# To Run

`go run loadgen.go`

# Issues

- Expecting the [`AUTO_INCREMENT`](https://docs.pingcap.com/tidb/stable/auto-increment) to increase doesn't work on TiDB Cloud as it skips to 30002.
- Setting `AUTO_ID_CACHE=1` doesn't seem to work as expected on Serverless
- We could create an REST API to get the list of posts