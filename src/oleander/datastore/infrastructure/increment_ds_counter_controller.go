package infrastructure

import (
	"hexample.com/src/oleander/user/application/datastore_counter"
	domain "hexample.com/src/oleander/user/domain"
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/shared/shared_domain/shared_domain_event_bus"
)

type IncreaseDsCounterPutController struct {
	uc *datastore_counter.IncreaseDatastoreCounterUc
}

func NewIncreaseDsCounterPutController(r domain.UserRepository, bus shared_domain_event_bus.EventBus) *IncreaseDsCounterPutController {
	return &IncreaseDsCounterPutController{
		uc: datastore_counter.NewIncreaseDatastoreCounterUc(r, bus),
	}
}

func (c* IncreaseDsCounterPutController) Invoke(userId *shared_domain.UserIDValueVO) error {
	err := c.uc.Invoke(userId)
	if err != nil {
		return err
	}
	return nil
}