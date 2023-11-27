package pulsar

import (
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func Client() pulsar.Client {

	client, error := pulsar.NewClient(
		pulsar.ClientOptions{
			URL:               "pulsar://localhost:6650",
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		})

	if error != nil {
		log.Fatalf("Não foi possível iniciar o Apache Pulsar: %v", error)
	}

	return client
}
