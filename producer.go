package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

const topic = "test"
const partition = 0

func main() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	conn.Close()
	fmt.Println("Wrote messages")
}
