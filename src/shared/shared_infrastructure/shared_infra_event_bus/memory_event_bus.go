package shared_infra_event_bus

import (
	"encoding/json"
	"hexample.com/src/shared/shared_domain/shared_domain_event_bus"
	"reflect"
)

type MemoryEventBus struct {
	queueMemoryMap map[string][]DomainEventSubscriber
}

func NewMemoryEventBus(subscribers []DomainEventSubscriber) (*MemoryEventBus, error) {
	bus := MemoryEventBus{
		queueMemoryMap: make(map[string][]DomainEventSubscriber),
	}
	for _, subscriber := range subscribers {
		err := bus.RegisterOnEventBus(subscriber)
		if err != nil {
			return nil, err
		}
	}
	return &bus, nil
}

func (meb *MemoryEventBus) RegisterOnEventBus(subscriber DomainEventSubscriber) error {
	eventName := meb.getName(subscriber.SubscribedTo())
	meb.queueMemoryMap[eventName] = append(meb.queueMemoryMap[eventName], subscriber)
	return nil
}

func (meb *MemoryEventBus) Publish(events []shared_domain_event_bus.DomainEvent) error {
	for _, event := range events {
		eventName := meb.getName(event)
		for _, mb := range meb.queueMemoryMap[eventName] {

			body, err := json.Marshal(event)
			if err != nil {
				return err
			}

			go func() {
				mb.Consume(body)
			}()
		}
	}
	return nil
}

func (meb *MemoryEventBus) Close() error {
	return nil
}

func (meb *MemoryEventBus) getName(subscriber shared_domain_event_bus.DomainEvent) string {
	if t := reflect.TypeOf(subscriber); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
