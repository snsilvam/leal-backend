package models

type Usuario struct {
	ID             uint  `gorm:"primaryKey"`
	PuntosLeal     int16 `gorm:"not null"`
	PuntosCashback int32
	Nombre         string `gorm:"size:255"`
}