package repositories

import (
	"api/src/domain"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type RabbitRepository struct {
	Conn *amqp.Connection
}

func (r *RabbitRepository) CreateLoan(loan *domain.Loan) error {
	channel, err := r.Conn.Channel()
	if err != nil {
		log.Println("Error al abrir un canal en RabbitMQ:", err)
		return err
	}
	defer channel.Close()

	// Serializar el objeto Loan a JSON
	body, err := json.Marshal(loan)
	if err != nil {
		log.Println("Error al serializar el objeto Loan a JSON:", err)
		return err
	}

	log.Println("📤 Intentando enviar mensaje a RabbitMQ...")
	log.Println("📦 Contenido del mensaje:", string(body))

	// Asegurar que la cola existe antes de publicar
	_, err = channel.QueueDeclare(
		"loan1", // Asegúrate de que el consumidor está escuchando esta misma cola
		true,    // Durable
		false,   // Auto-delete
		false,   // Exclusive
		false,   // No-wait
		nil,     // Args
	)
	if err != nil {
		log.Println("Error al declarar la cola en RabbitMQ:", err)
		return err
	}

	err = channel.Publish(
		"",
		"loan1",
		true,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Println("Error al publicar mensaje en RabbitMQ:", err)
		return err
	}

	log.Println("Mensaje publicado exitosamente en RabbitMQ")
	return nil
}
