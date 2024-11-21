package in

import "time"

type CreateCampanaDTO struct {
	SucursarID      int16     `json:"sucursalID" validate:"required"`
	ComercioID      int16     `json:"comercioID" validate:"required"`
	TipoBeneficioID int16     `json:"tipoBeneficioID" validate:"required"`
	FechaInicio     time.Time `json:"fechaInicio" validate:"required"`
	FechaFin        time.Time `json:"fechaFin" validate:"required,gtfield=FechaInicio"`
}
