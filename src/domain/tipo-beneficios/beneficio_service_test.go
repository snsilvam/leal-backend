package beneficios

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBeneficios_Success(t *testing.T) {
	mockRepo := new(MockBeneficiosRepository)
	service := ConstructorBeneficioService(mockRepo)

	ctx := context.Background()
	expectedBeneficios := []*TipoBeneficios{
		{
			ID:          1,
			Nombre:      "Beneficio A",
			Descripcion: "Beneficio A",
			TipoPuntos:  "Leal",
		},
		{
			ID:          2,
			Nombre:      "Beneficio B",
			Descripcion: "Beneficio B",
			TipoPuntos:  "Cashback",
		},
	}

	mockRepo.On("GetAllBeneficios", ctx).Return(expectedBeneficios, nil)

	result, err := service.GetAllBeneficios(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedBeneficios, result)
	mockRepo.AssertCalled(t, "GetAllBeneficios", ctx)
}

func TestGetAllBeneficios_Error(t *testing.T) {
	mockRepo := new(MockBeneficiosRepository)
	service := ConstructorBeneficioService(mockRepo)

	ctx := context.Background()

	mockRepo.On("GetAllBeneficios", ctx).Return(nil, errors.New("failed to fetch beneficios"))

	result, err := service.GetAllBeneficios(ctx)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "failed to fetch beneficios", err.Error())
	mockRepo.AssertCalled(t, "GetAllBeneficios", ctx)
}
