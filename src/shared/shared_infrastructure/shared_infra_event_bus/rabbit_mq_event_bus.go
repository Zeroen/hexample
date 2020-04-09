package shared_infra_event_bus

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"hexample.com/src/shared/shared_domain/shared_domain_event_bus"
	"reflect"
	"strconv"
)

type RabbitEventBus struct {
	user         string
	password     string
	hostname     string
	port         int
	rabbitBusMap map[string][]*RabbitQueue
}

type RabbitQueue struct {
	ch    *amqp.Channel
	conn  *amqp.Connection
	queue *amqp.Queue
}

func NewRabbitEventBus(
	user string,
	password string,
	hostname string,
	port int,
	subscribers []DomainEventSubscriber) (*RabbitEventBus, error) {

	bus := RabbitEventBus{
		user:         user,
		password:     password,
		hostname:     hostname,
		port:         port,
		rabbitBusMap: make(map[string][]*RabbitQueue),
	}
	for _, subscriber := range subscribers {
		err := bus.RegisterOnEventBus(subscriber)
		if err != nil {
			return nil, err
		}
	}
	return &bus, nil
}

func (reb *RabbitEventBus) RegisterOnEventBus(subscriber DomainEventSubscriber) error {
	queueName := subscriber.ChannelName()
	q, err := reb.createQueue(queueName)
	if err != nil {
		return err
	}

	eventName := reb.getName(subscriber.SubscribedTo())
	reb.rabbitBusMap[eventName] = append(reb.rabbitBusMap[eventName], q)

	go func() {
		msgs, err := q.ch.Consume(
			queueName,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			fmt.Printf("Error creating the RabbitMq consumer. error := %s", err.Error())
			return
		}

		for d := range msgs {
			subscriber.Consume(d.Body)
		}
	}()
	return nil
}

func (reb *RabbitEventBus) Publish(events []shared_domain_event_bus.DomainEvent) error {
	for _, event := range events {
		eventName := reb.getName(event)
		for _, rb := range reb.rabbitBusMap[eventName] {

			body, err := json.Marshal(event)
			if err != nil {
				return err
			}

			err = rb.ch.Publish(
				"",
				rb.queue.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
				})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (reb *RabbitEventBus) Close() error {
	for _, rbeb := range reb.rabbitBusMap {
		for _, rb := range rbeb {
			err := rb.conn.Close()
			if err != nil {
				fmt.Printf("Error closing the RabbitMq connection. error: %s", err.Error())
			}

			err = rb.ch.Close()
			if err != nil {
				fmt.Printf("Error closing the RabbitMq channel. error: %s", err.Error())
			}
		}
	}
	return nil
}

func (reb *RabbitEventBus) createQueue(queueName string) (*RabbitQueue, error) {
	portStr := strconv.FormatInt(int64(reb.port), 10)
	conn, err := amqp.Dial("amqp://" + reb.user + ":" + reb.password + "@" + reb.hostname + ":" + portStr + "/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitQueue{
		ch:    ch,
		conn:  conn,
		queue: &q,
	}, nil
}

func (reb *RabbitEventBus) getName(subscriber shared_domain_event_bus.DomainEvent) string {
	if t := reflect.TypeOf(subscriber); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
