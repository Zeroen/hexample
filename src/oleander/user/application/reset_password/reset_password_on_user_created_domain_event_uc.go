package reset_password

import (
	"encoding/json"
	"fmt"
	"hexample.com/src/oleander/user/domain"
	"hexample.com/src/oleander/user/domain/events"
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/shared/shared_domain/shared_domain_event_bus"
)

type ResetPasswordUCOnUserCreatedDomainEvent struct {
	resetPasswordUC *ResetPasswordUC
}

func NewResetPasswordUCOnUserCreatedDomainEvent(r domain.UserRepository) *ResetPasswordUCOnUserCreatedDomainEvent {
	return &ResetPasswordUCOnUserCreatedDomainEvent{
		resetPasswordUC: NewResetPasswordUC(r),
	}
}

func (uc *ResetPasswordUCOnUserCreatedDomainEvent) SubscribedTo() shared_domain_event_bus.DomainEvent {
	return events.UserCreatedDomainEvent{}
}

func (uc *ResetPasswordUCOnUserCreatedDomainEvent) ChannelName() string {
	return "ResetPasswordUCOnUserCreatedDomainEventOnUserCreatedDomainEvent"
}

func (uc *ResetPasswordUCOnUserCreatedDomainEvent) Consume(event []byte) {
	fmt.Printf("[APPLICATION] | [SUBSCRIBER] | ResetPasswordUCOnUserCreatedDomainEvent \n")

	var uce events.UserCreatedDomainEvent
	err := json.Unmarshal(event, &uce)
	if err != nil {
		// TODO
	}

	vo, err := shared_domain.NewUserIDValueVO(uce.Id)
	if err != nil {
		// TODO
	}

	if vo == nil {
		// Impossible case
		return
	}

	err = uc.resetPasswordUC.Invoke(*vo)
	if err != nil {
		// TODO
	}
}
