package shared_domain_event_bus

type DomainEvent interface {
	FullQualifiedEventName() string
}
