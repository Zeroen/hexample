package datastore_counter

import (
	"encoding/json"
	"hexample.com/src/oleander/user/domain"
	"hexample.com/src/oleander/user/domain/events"
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/shared/shared_domain/shared_domain_event_bus"
)

type IncreaseDatastoreCounterOnDsCreatedEvent struct {
	increaseDsCounterUc *IncreaseDatastoreCounterUc
}

func NewIncreaseDatastoreCounterOnDsCreatedEvent(repository domain.UserRepository, bus shared_domain_event_bus.EventBus) *IncreaseDatastoreCounterOnDsCreatedEvent {
	return &IncreaseDatastoreCounterOnDsCreatedEvent{
		increaseDsCounterUc: NewIncreaseDatastoreCounterUc(repository, bus),
	}
}

func (uc *IncreaseDatastoreCounterOnDsCreatedEvent) SubscribedTo() shared_domain_event_bus.DomainEvent {
	return events.DatastoreCreatedEvent{}
}

func (uc *IncreaseDatastoreCounterOnDsCreatedEvent) ChannelName() string {
	return "IncreaseDatastoreCounterOnDatastoreCreatedQueue"
}

func (uc *IncreaseDatastoreCounterOnDsCreatedEvent) Consume(event []byte) {
	var dsce events.DatastoreCreatedEvent
	err := json.Unmarshal(event, &dsce)
	if err != nil {
		// TODO Log a message
	}

	usId, err := shared_domain.NewUserIDValueVO(dsce.UserId)
	if err != nil {
		// TODO
	}
	
	err = uc.increaseDsCounterUc.Invoke(usId)
	if err != nil {
		// TODO Si falla este, que leches hacemos
		// O Logeamos un error o reintentamos....
	}
}
