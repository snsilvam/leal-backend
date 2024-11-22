package compras

import "time"

type Compras struct {
	ID                    int16
	UsuarioID             int16
	SucursalID            int16
	CampanaID             int16
	ComercioID            int16
	PuntosLealGanados     int16
	PuntosCashbackGanados int16
	FechaCompra           time.Time
	ValorCompra           float64
}
