package events

type UserCreatedDomainEvent struct {
	Id string `json:"id"`
}

func NewUserCreatedDomainEvent(id string) UserCreatedDomainEvent {
	return UserCreatedDomainEvent{Id: id}
}

func (ucde UserCreatedDomainEvent) FullQualifiedEventName() string {
	return "UserCreatedDomainEvent"
}
