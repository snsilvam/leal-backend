package models

type Comercio struct {
	ID         uint       `gorm:"primaryKey"`
	Nombre     string     `gorm:"size:255;not null"`
	Sucursales []Sucursal `gorm:"foreignKey:ComercioID"`
}
