package amqp

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

// default routing
type Routing string

const (
	RegisterVerification Routing = "email.registration.otp"
)

type producer struct {
	channel      *amqp.Channel
	exchange     string
	exchangeType string
}

type Producer interface {
	// PublishMessage will publish message into queue with spesific
	// routing key
	PublishMessage(routingKey Routing, data interface{}) error
}

func New() Producer {
	amqpURI := fmt.Sprintf(
		"amqp://%s:%s@%s:%s%s",
		os.Getenv("RABBIT_USERNAME"),
		os.Getenv("RABBIT_PASSWORD"),
		os.Getenv("RABBIT_HOST"),
		os.Getenv("RABBIT_PORT"),
		os.Getenv("RABBIT_VH"),
	)

	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return &producer{
		channel:      ch,
		exchange:     os.Getenv("RABBIT_EXCHANGE"),
		exchangeType: "topic",
	}
}
