package models

type PuntosLeal struct {
	ID             uint  `gorm:"primaryKey"`
	Comercio       int16 `gorm:"not null"`
	Usuario        int16 `gorm:"not null"`
	CantidadPuntos float64
}
