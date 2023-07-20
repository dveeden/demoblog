package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type tidbMessage struct {
	Database string
	Table    string
	Data     []map[string]string
}

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "ticker",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(42)

	var totalHeartbeats uint64
	go func() {
		var lastNr uint64
		for {
			log.Printf(
				"Heartbeast Stats: Total: %d (%d in the last 10 seconds)",
				totalHeartbeats, totalHeartbeats-lastNr,
			)
			lastNr = totalHeartbeats
			time.Sleep(time.Second * 10)
		}
	}()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		var j tidbMessage
		err = json.Unmarshal(m.Value, &j)
		if err != nil {
			log.Print(err)
		}

		if j.Database == "blog" && j.Table == "ticker" {
			loc, _ := time.LoadLocation("Europe/Amsterdam")
			ts, _ := time.ParseInLocation("2006-01-02 15:04:05.999", j.Data[0]["ts"], loc)
			// \a is Terminal bell, to simulate heartbeat
			log.Printf(
				"\aReceived heartbeat at %s which is %s ago, total heartbeats: %d",
				ts, time.Since(ts), totalHeartbeats,
			)
			totalHeartbeats++
		} else if j.Database == "blog" && j.Table == "posts" {
			log.Printf("post: %#v", j)
		} else {
			log.Printf("Got event for %s", j.Table)
		}
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
