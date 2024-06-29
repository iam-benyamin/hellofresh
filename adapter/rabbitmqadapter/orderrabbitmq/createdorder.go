package orderrabbitmq

import (
	"context"
	"encoding/json"
	"github.com/iam-benyamin/hellofresh/param/orderparam"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func (a Adapter) PublishCreatedOrder(ctx context.Context, msg orderparam.Message, routingKey string) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = a.rmq.Channel.PublishWithContext(
		ctx,
		a.rmq.Exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s", body)
	return nil
}
