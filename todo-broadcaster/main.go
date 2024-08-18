package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		log.Fatal("NATS_URL is not set")
	}

	discordURL := os.Getenv("DISCORD_URL")
	if discordURL == "" {
		log.Fatal("DISCORD_URL is not set")
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	_, err = nc.QueueSubscribe("todos", "workers", func(m *nats.Msg) {
		log.Printf("Message received: %s\n", string(m.Data))

		payload := map[string]string{"content": string(m.Data)}
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			log.Printf("Error marshaling message: %v", err)
			return
		}

		log.Println("Message received from NC:\n %s", string(payloadByte)s)
	})

	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}
	select {}
}
