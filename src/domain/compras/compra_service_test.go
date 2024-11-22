package compras

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateCompra_Success(t *testing.T) {
	mockRepo := new(MockComprasRepository)
	service := ConstructorCompraService(mockRepo)

	ctx := context.Background()
	compra := &Compras{
		ID:                    1,
		UsuarioID:             1,
		SucursalID:            1,
		CampanaID:             1,
		ComercioID:            1,
		PuntosLealGanados:     100,
		PuntosCashbackGanados: 100,
		FechaCompra:           time.Now(),
		ValorCompra:           200000,
	}

	mockRepo.On("CreateCompra", ctx, compra).Return(nil)

	err := service.CreateCompra(ctx, compra)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "CreateCompra", ctx, compra)
}

func TestCreateCompra_Error(t *testing.T) {
	mockRepo := new(MockComprasRepository)
	service := ConstructorCompraService(mockRepo)

	ctx := context.Background()
	compra := &Compras{
		ID:                    1,
		UsuarioID:             1,
		SucursalID:            1,
		CampanaID:             1,
		ComercioID:            1,
		PuntosLealGanados:     100,
		PuntosCashbackGanados: 100,
		FechaCompra:           time.Now(),
		ValorCompra:           200000,
	}

	mockRepo.On("CreateCompra", ctx, compra).Return(errors.New("failed to create compra"))

	err := service.CreateCompra(ctx, compra)

	assert.Error(t, err)
	assert.Equal(t, "failed to create compra", err.Error())
	mockRepo.AssertCalled(t, "CreateCompra", ctx, compra)
}

func TestAddPuntosTipoBeneficio1_Success(t *testing.T) {
	mockRepo := new(MockComprasRepository)
	service := ConstructorCompraService(mockRepo)

	ctx := context.Background()
	compra := &Compras{}

	mockRepo.On("AddPuntosTipoBeneficio1", ctx, compra).Return(nil)

	err := service.AddPuntosTipoBeneficio1(ctx, compra)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "AddPuntosTipoBeneficio1", ctx, compra)
}

func TestAddPuntosTipoBeneficio1_Error(t *testing.T) {
	mockRepo := new(MockComprasRepository)
	service := ConstructorCompraService(mockRepo)

	ctx := context.Background()
	compra := &Compras{}

	mockRepo.On("AddPuntosTipoBeneficio1", ctx, compra).Return(errors.New("failed to add puntos"))

	err := service.AddPuntosTipoBeneficio1(ctx, compra)

	assert.Error(t, err)
	assert.Equal(t, "failed to add puntos", err.Error())
	mockRepo.AssertCalled(t, "AddPuntosTipoBeneficio1", ctx, compra)
}

func TestCampanaActiva_Success(t *testing.T) {
	mockRepo := new(MockComprasRepository)
	service := ConstructorCompraService(mockRepo)

	ctx := context.Background()
	idCampana := int16(10)
	expectedResult := int16(1)

	mockRepo.On("CampanaActiva", ctx, idCampana).Return(expectedResult, nil)

	result, err := service.CampanaActiva(ctx, idCampana)

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	mockRepo.AssertCalled(t, "CampanaActiva", ctx, idCampana)
}

func TestCampanaActiva_Error(t *testing.T) {
	mockRepo := new(MockComprasRepository)
	service := ConstructorCompraService(mockRepo)

	ctx := context.Background()
	idCampana := int16(10)

	mockRepo.On("CampanaActiva", ctx, idCampana).Return(int16(0), errors.New("no active campana"))

	result, err := service.CampanaActiva(ctx, idCampana)

	assert.Error(t, err)
	assert.Equal(t, int16(0), result)
	assert.Equal(t, "no active campana", err.Error())
	mockRepo.AssertCalled(t, "CampanaActiva", ctx, idCampana)
}

func TestDeterminarQueTipoDeBeneficioEstaActivo_Success(t *testing.T) {
	mockRepo := new(MockComprasRepository)
	service := ConstructorCompraService(mockRepo)

	ctx := context.Background()
	beneficioID := int16(5)
	compra := &Compras{}

	mockRepo.On("DeterminarQueTipoDeBeneficioEstaActivo", ctx, beneficioID, compra).Return(nil)

	err := service.DeterminarQueTipoDeBeneficioEstaActivo(ctx, beneficioID, compra)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "DeterminarQueTipoDeBeneficioEstaActivo", ctx, beneficioID, compra)
}

func TestDeterminarQueTipoDeBeneficioEstaActivo_Error(t *testing.T) {
	mockRepo := new(MockComprasRepository)
	service := ConstructorCompraService(mockRepo)

	ctx := context.Background()
	beneficioID := int16(5)
	compra := &Compras{}

	mockRepo.On("DeterminarQueTipoDeBeneficioEstaActivo", ctx, beneficioID, compra).Return(errors.New("beneficio not active"))

	err := service.DeterminarQueTipoDeBeneficioEstaActivo(ctx, beneficioID, compra)

	assert.Error(t, err)
	assert.Equal(t, "beneficio not active", err.Error())
	mockRepo.AssertCalled(t, "DeterminarQueTipoDeBeneficioEstaActivo", ctx, beneficioID, compra)
}
