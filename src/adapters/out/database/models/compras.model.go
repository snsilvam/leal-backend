package models

import "time"

type Compras struct {
	ID                    uint  `gorm:"primaryKey"`
	UsuarioID             int16 `gorm:"not null"`
	SucursalID            int16 `gorm:"not null"`
	CampanaID             int16 `gorm:"default:null"`
	ComercioID            int16 `gorm:"default:null"`
	PuntosLealGanados     int16
	PuntosCashbackGanados int16
	FechaCompra           time.Time
	ValorCompra           float64
}
