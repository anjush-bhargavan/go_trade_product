package rabbitmq

import (
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
)



func PublishNotifications(data Messages) error {
	conn,err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	//Create a channel for communication
	ch,err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

		// Declare a queue
		q, err := ch.QueueDeclare(
			"notification_queue", // queue name
			false,                // durable
			false,                // delete when unused
			false,                // exclusive
			false,                // no-wait
			nil,                  // arguments
		)
		if err != nil {
			return err
		}
	
		// Convert notification details to JSON
		body, err := json.Marshal(data)
		if err != nil {
			return err
		}
	
		// Publish message to the queue
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		if err != nil {
			return err
		}
	
		return nil
}