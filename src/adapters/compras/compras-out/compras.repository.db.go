package comprasout

import (
	"context"
	"fmt"
	"leal-backend/src/adapters/out/database"
	"leal-backend/src/adapters/out/database/models"
	"leal-backend/src/domain/compras"
	"log"
)

type ComprasPostgresRepository struct {
	database database.Database
}

func ConstructorComprasPostgresRepository(database database.Database) *ComprasPostgresRepository {
	return &ComprasPostgresRepository{
		database: database,
	}
}

func (r *ComprasPostgresRepository) CreateCompra(context context.Context, compra *compras.Compras) error {
	if compra.CampanaID == 0 {
		result := r.database.Db.Create(&compra)
		if result.Error != nil {
			log.Println("No se pudo insertar el registro de compra: ", result.Error)
			return result.Error
		}
		return nil
	}

	tipoBeneficioID, err := r.CampanaActiva(context, compra.CampanaID)
	if err != nil {
		return err
	}

	if tipoBeneficioID != 0 {
		err := r.DeterminarQueTipoDeBeneficioEstaActivo(context, tipoBeneficioID, compra)
		if err != nil {
			return err
		}
	}

	result := r.database.Db.Create(&compra)
	if result.Error != nil {
		log.Println("No se pudo insertar el registro de compra: ", result.Error)
		return result.Error
	}

	return nil
}

func (r *ComprasPostgresRepository) CampanaActiva(ctx context.Context, idCampana int16) (int16, error) {
	var campana models.Campana

	result := r.database.Db.First(&campana, idCampana)
	if result.Error != nil {
		return 0, result.Error
	}

	if campana.Estado {
		return campana.TipoBeneficioID, nil
	}

	return 0, nil
}

func (r *ComprasPostgresRepository) DeterminarQueTipoDeBeneficioEstaActivo(context context.Context, beneficioID int16, compra *compras.Compras) error {
	if beneficioID == 1 {
		err := r.AddPuntosTipoBeneficio1(context, compra)
		if err != nil {
			return err
		}
		return nil
	}

	if beneficioID == 2 {
		err := r.AddPuntosTipoBeneficio2(context, compra)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("el beneficio con CampanaID %d no se encuentra registrado en el sistema", beneficioID)
}

func (r *ComprasPostgresRepository) AddPuntosTipoBeneficio1(context context.Context, compra *compras.Compras) error {
	puntosLeal := r.PuntosGanados(compra.ValorCompra)
	puntosLeal = puntosLeal * 2

	createPuntosLeal := &models.PuntosLeal{
		Comercio:       compra.ComercioID,
		Usuario:        compra.UsuarioID,
		CantidadPuntos: puntosLeal,
	}

	result := r.database.Db.Create(&createPuntosLeal)
	if result.Error != nil {
		log.Println("No se pudo insertar los puntos leal: ", result.Error)
		return result.Error
	}
	return nil
}

func (r *ComprasPostgresRepository) AddPuntosTipoBeneficio2(context context.Context, compra *compras.Compras) error {
	puntosCashback := r.PuntosGanados(compra.ValorCompra)
	puntosCashback = r.CalcularPuntos(compra.ValorCompra, puntosCashback)

	updatePuntos := &models.Usuario{
		PuntosCashback: puntosCashback,
	}

	result := r.database.Db.Model(&models.Usuario{}).Where("id = ?", compra.UsuarioID).Updates(updatePuntos)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ComprasPostgresRepository) PuntosGanados(valorCompra float64) float64 {
	if valorCompra > 999 {
		puntosGanados := valorCompra / 1000
		return puntosGanados
	}

	return 0
}

func (r *ComprasPostgresRepository) CalcularPuntos(valorCompra float64, puntosGanados float64) float64 {
	if valorCompra > 20000 {
		puntosGanados += puntosGanados * 0.30
	}
	return puntosGanados
}
