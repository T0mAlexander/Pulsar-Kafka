package pulsar

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

func Consumer() {
	consumer, error := Client().
	Subscribe(pulsar.ConsumerOptions{
		Topic:            "persistent://public/default/golang",
		SubscriptionName: "Subscrição - Golang",
	})
	
	if error != nil {
		log.Fatalf("Não foi possível criar um consumidor do Apache Pulsar: %v", error)
	}
	
	for {
		message, error := consumer.Receive(context.Background())
		if error != nil {
			log.Fatal(error)
		}
	
		fmt.Printf("Mensagem recebida: '%s'\n", string(message.Payload()))
		consumer.Ack(message)
	}
}
