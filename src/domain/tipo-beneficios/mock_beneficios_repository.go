package beneficios

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockBeneficiosRepository struct {
	mock.Mock
}

func (m *MockBeneficiosRepository) GetAllBeneficios(ctx context.Context) ([]*TipoBeneficios, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*TipoBeneficios), args.Error(1)
}
