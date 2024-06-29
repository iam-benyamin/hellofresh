package main

import (
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Order struct {
	OrderID          string  `json:"order_id"`
	CustomerFullname string  `json:"customer_fullname"`
	ProductName      string  `json:"product_name"`
	TotalAmount      float64 `json:"total_amount"`
	CreatedAt        string  `json:"created_at"`
}

type Payload struct {
	Order Order `json:"order"`
}

type Message struct {
	Producer string  `json:"producer"`
	SentAt   string  `json:"sent_at"`
	Type     string  `json:"type"`
	Payload  Payload `json:"payload"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://hellofresh:food@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare the exchange
	err = ch.ExchangeDeclare(
		"orders", // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// Create the message
	order := Order{
		OrderID:          "12345",
		CustomerFullname: "John Doe",
		ProductName:      "Widget",
		TotalAmount:      99.99,
		CreatedAt:        time.Now().Format(time.RFC3339),
	}

	payload := Payload{
		Order: order,
	}

	message := Message{
		Producer: "example_producer",
		SentAt:   time.Now().Format(time.RFC3339),
		Type:     "created_order",
		Payload:  payload,
	}

	// Serialize the message to JSON
	body, err := json.Marshal(message)
	failOnError(err, "Failed to marshal JSON")

	// Publish the message
	err = ch.Publish(
		"orders",        // exchange
		"created_order", // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
