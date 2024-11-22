package sucursales

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllSucursales_Success(t *testing.T) {
	mockRepo := new(MockSucursalesRepository)
	service := ConstructorSucursalesService(mockRepo)

	ctx := context.Background()
	comercioId := int16(1)
	expectedSucursales := []*Sucursales{
		{
			ID:                 1,
			ComercioID:         1,
			ComprasRegistradas: 500,
			Ubicacion:          "Bogota",
		},
		{
			ID:                 1,
			ComercioID:         1,
			ComprasRegistradas: 600,
			Ubicacion:          "Cali",
		},
	}

	mockRepo.On("GetAllSucursales", ctx, comercioId).Return(expectedSucursales, nil)

	result, err := service.GetAllSucursales(ctx, comercioId)

	assert.NoError(t, err)
	assert.Equal(t, expectedSucursales, result)
	mockRepo.AssertCalled(t, "GetAllSucursales", ctx, comercioId)
}

func TestGetAllSucursales_Error(t *testing.T) {
	mockRepo := new(MockSucursalesRepository)
	service := ConstructorSucursalesService(mockRepo)

	ctx := context.Background()
	comercioId := int16(1)

	mockRepo.On("GetAllSucursales", ctx, comercioId).Return(nil, errors.New("failed to fetch sucursales"))

	result, err := service.GetAllSucursales(ctx, comercioId)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to fetch sucursales", err.Error())
	mockRepo.AssertCalled(t, "GetAllSucursales", ctx, comercioId)
}
