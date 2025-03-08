package adapters

import (
	"api/src/domain"
	"encoding/json"

	"github.com/streadway/amqp"
)

type RabbitMQProducer struct {
	Connection *amqp.Connection
}

func (r *RabbitMQProducer) SendLoanEvent(loan *domain.Loan) error {
	channel, err := r.Connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	body, err := json.Marshal(loan)
	if err != nil {
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

	return err
}
