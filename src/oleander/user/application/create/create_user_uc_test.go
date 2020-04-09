package create

import (
	"github.com/stretchr/testify/require"
	"hexample.com/src/oleander/user/application/test/fixtures"
	"hexample.com/src/oleander/user/application/test/mocks"
	"hexample.com/src/oleander/user/domain"
	shared_fixtures "hexample.com/src/shared/shared_domain/test/shared_fixtures"
	"hexample.com/src/shared/shared_infrastructure/shared_infra_event_bus"
	"testing"
)

func TestNewCreateUserUC(t *testing.T) {

	t.Run("Scenario 1: Create a new User.", func(t *testing.T) {

		uID := shared_fixtures.NewUserIDVOFixture1()
		name := fixtures.NewNameVOFixture1()
		age := fixtures.NewAgeVOFixture1()
		eID := shared_fixtures.NewEmailIDVOFixture1()
		user, err := domain.NewUser(*uID, name, age, *eID)
		require.NoError(t, err)

		ur := mocks.NewMockedUserRepository()
		ur.On("Begin").Return(nil)
		ur.On("SearchByID", *uID).Return(nil, nil)
		ur.On("Save", user).Return(nil)
		ur.On("Commit").Return(nil)

		eb, err := shared_infra_event_bus.NewMemoryEventBus(nil)
		require.NoError(t, err)

		uc := NewCreateUserUC(eb, ur)
		err = uc.Invoke(*uID, name, age, *eID)
		require.NoError(t, err)
		ur.AssertExpectations(t)
	})
}
