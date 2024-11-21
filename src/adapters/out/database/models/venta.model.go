package models

type Ventas struct {
	ID         uint  `gorm:"primaryKey"`
	CompraID   int16 `gorm:"not null"`
	SucursalID int16 `gorm:"not null"`
	CampanaID  int16 `gorm:"default:null"` //atributo opcional
}
