package compras

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockComprasRepository struct {
	mock.Mock
}

func (m *MockComprasRepository) CreateCompra(ctx context.Context, compra *Compras) error {
	args := m.Called(ctx, compra)
	return args.Error(0)
}

func (m *MockComprasRepository) CampanaActiva(ctx context.Context, idCampana int16) (int16, error) {
	args := m.Called(ctx, idCampana)
	return args.Get(0).(int16), args.Error(1)
}

func (m *MockComprasRepository) DeterminarQueTipoDeBeneficioEstaActivo(ctx context.Context, beneficioID int16, compra *Compras) error {
	args := m.Called(ctx, beneficioID, compra)
	return args.Error(0)
}

func (m *MockComprasRepository) AddPuntosTipoBeneficio1(ctx context.Context, compra *Compras) error {
	args := m.Called(ctx, compra)
	return args.Error(0)
}

func (m *MockComprasRepository) AddPuntosTipoBeneficio2(ctx context.Context, compra *Compras) error {
	args := m.Called(ctx, compra)
	return args.Error(0)
}
