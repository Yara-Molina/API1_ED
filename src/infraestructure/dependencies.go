package infraestructure

import (
	"api/src/application/repositories"
	"api/src/application/services"

	"github.com/streadway/amqp"
)

func Setup() (*services.EventService, error) {
	conn, err := amqp.Dial("amqp://yara:noobmaster69@54.161.81.210:5672/")
	if err != nil {
		return nil, err
	}

	rabbitRepo := &repositories.RabbitRepository{Conn: conn}
	eventService := services.EventService{LoanRepo: rabbitRepo}

	return &eventService, nil
}
