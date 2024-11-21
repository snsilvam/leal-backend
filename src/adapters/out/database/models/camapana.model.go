package models

import "time"

type Campana struct {
	ID                 uint  `gorm:"primaryKey"`
	SucursarID         int16 `gorm:"not null"`
	ComercioID         int16 `gorm:"not null"`
	TipoBeneficioID    int16 `gorm:"not null"`
	ComprasRegistradas int32
	FechaInicio        time.Time
	FechaFin           time.Time
}
