package models

type Sucursal struct {
	ID                 uint  `gorm:"primaryKey"`
	ComercioID         int16 `gorm:"not null"`
	ComprasRegistradas int32
	Ubicacion          string `gorm:"size:255"`
}
