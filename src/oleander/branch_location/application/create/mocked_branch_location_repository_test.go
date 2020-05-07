package create

import (
	"github.com/stretchr/testify/mock"
	"hexample.com/src/oleander/branch_location/domain"
	"hexample.com/src/shared/shared_domain"
)

type MockedBranchLocationRepository struct {
	mock.Mock
}

func NewMockedBranchLocationRepository() *MockedBranchLocationRepository {
	return &MockedBranchLocationRepository{}
}

func (m *MockedBranchLocationRepository) Begin() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockedBranchLocationRepository) Commit() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockedBranchLocationRepository) Rollback() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockedBranchLocationRepository) SearchByID(id *shared_domain.BranchLocationIDValueVO) (*domain.BranchLocationAG, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*domain.BranchLocationAG), args.Error(1)
	}
}

func (m *MockedBranchLocationRepository) Save(e *domain.BranchLocationAG) error {
	args := m.Called(e)
	return args.Error(0)
}
