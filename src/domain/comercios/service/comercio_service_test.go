package comercioService

import (
	"context"
	"errors"
	"testing"

	domain "leal-backend/src/domain/comercios/entity"
	"leal-backend/src/domain/comercios/port"

	"github.com/stretchr/testify/assert"
)

func TestGetAllComercios_Success(t *testing.T) {
	mockRepo := new(port.MockComercioRepository)
	service := ConstructorComercioService(mockRepo)

	ctx := context.Background()
	expectedComercios := []*domain.Comercio{
		{ID: 1, Nombre: "Comercio A"},
		{ID: 2, Nombre: "Comercio B"},
	}

	mockRepo.On("GetAllComercios", ctx).Return(expectedComercios, nil)

	comercios, err := service.GetAllComercios(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedComercios, comercios)
	mockRepo.AssertCalled(t, "GetAllComercios", ctx)
}

func TestGetAllComercios_Error(t *testing.T) {
	mockRepo := new(port.MockComercioRepository)
	service := ConstructorComercioService(mockRepo)

	ctx := context.Background()
	mockRepo.On("GetAllComercios", ctx).Return(nil, errors.New("database error"))

	comercios, err := service.GetAllComercios(ctx)

	assert.Error(t, err)
	assert.Nil(t, comercios)
	assert.Equal(t, "database error", err.Error())
	mockRepo.AssertCalled(t, "GetAllComercios", ctx)
}
