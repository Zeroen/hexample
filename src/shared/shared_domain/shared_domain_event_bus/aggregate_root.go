package shared_domain_event_bus

type AggregateRoot struct {
	events []DomainEvent
}

func (ag *AggregateRoot) Record(event DomainEvent) {
	ag.events = append(ag.events, event)
}

func (ag *AggregateRoot) PullDomainEvents() []DomainEvent {
	eventsToBeReturned := ag.events
	ag.events = make([]DomainEvent, 0)
	return eventsToBeReturned
}
