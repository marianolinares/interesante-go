package storagemock

import (
	"github.com/stretchr/testify/mock"
	"marian.com/interesante-go/code/internal"
)

type MyMockedRepo struct {
	mock.Mock
}

func (m *MyMockedRepo) GetEntities() ([]internal.Entity, error) {
	args := m.Called()
	return args.Get(0).([]internal.Entity), args.Error(1)
}

func (m *MyMockedRepo) SaveEntity(_ internal.Entity) {
	m.Called()
}
