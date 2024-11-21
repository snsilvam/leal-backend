package entity

import "time"

type Campana struct {
	SucursarID         int16
	ComercioID         int16
	TipoBeneficioID    int16
	ComprasRegistradas int32
	FechaInicio        time.Time
	FechaFin           time.Time
	Estado             bool
}
