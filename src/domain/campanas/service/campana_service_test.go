package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"leal-backend/src/domain/campanas/entity"
	"leal-backend/src/domain/campanas/port"

	"github.com/stretchr/testify/assert"
)

func TestCreateCampana_Success(t *testing.T) {
	mockRepo := new(port.MockCampanaRepository)
	service := ConstructorCampanaService(mockRepo)

	ctx := context.Background()
	campana := &entity.Campana{
		SucursalID:         1,
		ComercioID:         1,
		TipoBeneficioID:    1,
		ComprasRegistradas: 100000,
		FechaInicio:        time.Now(),
		FechaFin:           time.Now(),
		Estado:             true,
	}

	mockRepo.On("CreateCampana", ctx, campana).Return(nil)

	err := service.CreateCampana(ctx, campana)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "CreateCampana", ctx, campana)
}

func TestCreateCampana_Error(t *testing.T) {
	mockRepo := new(port.MockCampanaRepository)
	service := ConstructorCampanaService(mockRepo)

	ctx := context.Background()
	campana := &entity.Campana{
		SucursalID:         1,
		ComercioID:         1,
		TipoBeneficioID:    1,
		ComprasRegistradas: 100000,
		FechaInicio:        time.Now(),
		FechaFin:           time.Now(),
		Estado:             true,
	}

	mockRepo.On("CreateCampana", ctx, campana).Return(errors.New("failed to create campana"))

	err := service.CreateCampana(ctx, campana)

	assert.Error(t, err)
	assert.Equal(t, "failed to create campana", err.Error())
	mockRepo.AssertCalled(t, "CreateCampana", ctx, campana)
}

func TestGetAllCampanasOfComercioAndSucursal_Success(t *testing.T) {
	mockRepo := new(port.MockCampanaRepository)
	service := ConstructorCampanaService(mockRepo)

	ctx := context.Background()
	comercioID := int16(1)
	sucursalID := int16(2)
	expectedCampanas := []*entity.Campana{
		{
			SucursalID:         1,
			ComercioID:         1,
			TipoBeneficioID:    1,
			ComprasRegistradas: 100000,
			FechaInicio:        time.Now(),
			FechaFin:           time.Now(),
			Estado:             true,
		},
		{
			SucursalID:         2,
			ComercioID:         2,
			TipoBeneficioID:    2,
			ComprasRegistradas: 200000,
			FechaInicio:        time.Now(),
			FechaFin:           time.Now(),
			Estado:             true,
		},
	}

	mockRepo.On("GetAllCampanasOfComercioAndSucursal", ctx, comercioID, sucursalID).Return(expectedCampanas, nil)

	campanas, err := service.GetAllCampanasOfComercioAndSucursal(ctx, comercioID, sucursalID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCampanas, campanas)
	mockRepo.AssertCalled(t, "GetAllCampanasOfComercioAndSucursal", ctx, comercioID, sucursalID)
}

func TestGetAllCampanasOfComercioAndSucursal_Error(t *testing.T) {
	mockRepo := new(port.MockCampanaRepository)
	service := ConstructorCampanaService(mockRepo)

	ctx := context.Background()
	comercioID := int16(1)
	sucursalID := int16(2)

	mockRepo.On("GetAllCampanasOfComercioAndSucursal", ctx, comercioID, sucursalID).Return(nil, errors.New("failed to fetch campanas"))

	campanas, err := service.GetAllCampanasOfComercioAndSucursal(ctx, comercioID, sucursalID)

	assert.Error(t, err)
	assert.Nil(t, campanas)
	assert.Equal(t, "failed to fetch campanas", err.Error())
	mockRepo.AssertCalled(t, "GetAllCampanasOfComercioAndSucursal", ctx, comercioID, sucursalID)
}
