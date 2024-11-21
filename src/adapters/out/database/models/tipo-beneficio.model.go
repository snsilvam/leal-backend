package models

type TipoBeneficios struct {
	ID          uint   `gorm:"primaryKey"`
	Nombre      string `gorm:"size:255;not null"`
	Descripcion string `gorm:"size:255;not null"`
}
