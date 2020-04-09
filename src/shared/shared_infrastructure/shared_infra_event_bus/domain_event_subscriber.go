package shared_infra_event_bus

import (
	"hexample.com/src/shared/shared_domain/shared_domain_event_bus"
)

type DomainEventSubscriber interface {
	SubscribedTo() shared_domain_event_bus.DomainEvent
	ChannelName() string
	Consume(event []byte)
}
