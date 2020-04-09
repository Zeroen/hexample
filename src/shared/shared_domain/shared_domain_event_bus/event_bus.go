package shared_domain_event_bus

type EventBus interface {
	Publish(events []DomainEvent) error
	Close() error
}
