package amqp

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

func (p *producer) PublishMessage(routingKey Routing, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.channel.Publish(
		p.exchange,
		string(routingKey),
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
