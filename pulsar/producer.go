package pulsar

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

func Producer() {
	producer, error := Client().
		CreateProducer(pulsar.ProducerOptions{Topic: "topic-1"})

	if error != nil {
		log.Fatalf("Não foi possível criar um novo produtor do Apache Pulsar: %v", error)
	}

	id := 0
	for {
		message := fmt.Sprintf("Mensagem-%d", id)
		producer.Send( //<-- Para envio assíncrono, use o método "SendAsync"
			context.Background(),
			&pulsar.ProducerMessage{
				Payload: []byte(message),
			})
		id++
	}
}
