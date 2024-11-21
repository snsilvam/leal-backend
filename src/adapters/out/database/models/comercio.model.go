package models

type Comercio struct {
	ID     uint `gorm:"primaryKey"`
	Nombre string
}
