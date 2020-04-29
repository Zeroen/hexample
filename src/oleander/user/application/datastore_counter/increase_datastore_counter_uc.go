package datastore_counter

import (
	domain "hexample.com/src/oleander/user/domain"
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/shared/shared_domain/shared_domain_event_bus"
)

type IncreaseDatastoreCounterUc struct {
	r             domain.UserRepository
	getUserByIdDs *domain.GetUserByIDDS
	eb            shared_domain_event_bus.EventBus
}

func NewIncreaseDatastoreCounterUc(r domain.UserRepository, eb shared_domain_event_bus.EventBus) *IncreaseDatastoreCounterUc {
	return &IncreaseDatastoreCounterUc{
		r:             r,
		getUserByIdDs: domain.NewGetUserByIDDS(r),
		eb:            eb,
	}
}

func (uc *IncreaseDatastoreCounterUc) Invoke(userId *shared_domain.UserIDValueVO) (rErr error) {
	err := shared_domain.BeginTX(uc.r)
	if err != nil {
		return err
	}
	defer shared_domain.EndTX(uc.r, &rErr)

	userToIncrement, err := uc.getUserByIdDs.Invoke(*userId)
	if err != nil {
		return err
	}

	userToIncrement.Increase()

	err = uc.r.Update(userToIncrement)
	if err != nil {
		return err
	}

	err = uc.eb.Publish(userToIncrement.PullDomainEvents())
	if err != nil {
		return err
	}

	return nil
}
