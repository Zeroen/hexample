package mocks

import (
	"github.com/stretchr/testify/mock"
	"hexample.com/src/oleander/user/domain"
	"hexample.com/src/shared/shared_domain"
	"hexample.com/src/shared/shared_domain/criteria"
)

type MockedUserRepository struct {
	mock.Mock
}

func NewMockedUserRepository() *MockedUserRepository {
	return &MockedUserRepository{}
}

func (m *MockedUserRepository) Begin() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockedUserRepository) Commit() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockedUserRepository) Rollback() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockedUserRepository) SearchAll() ([]*domain.UserAG, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).([]*domain.UserAG), args.Error(1)
	}
}

func (m *MockedUserRepository) Count() (int, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return 0, args.Error(1)
	} else {
		return args.Get(0).(int), args.Error(1)
	}
}

func (m *MockedUserRepository) SearchBy(criteria *criteria.Criteria) ([]*domain.UserAG, error) {
	args := m.Called(criteria)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).([]*domain.UserAG), args.Error(1)
	}
}

func (m *MockedUserRepository) SearchByID(id shared_domain.UserIDValueVO) (*domain.UserAG, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*domain.UserAG), args.Error(1)
	}
}

func (m *MockedUserRepository) Save(e *domain.UserAG) error {
	args := m.Called(e)
	return args.Error(0)
}

func (m *MockedUserRepository) Update(e *domain.UserAG) error {
	args := m.Called(e)
	return args.Error(1)
}

func (m *MockedUserRepository) Delete(e *domain.UserAG) error {
	args := m.Called(e)
	return args.Error(0)
}
