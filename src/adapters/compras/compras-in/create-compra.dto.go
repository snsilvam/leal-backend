package compraIn

type CreateCompraDTO struct {
	UsuarioID   int16   `json:"usuarioID" validate:"required"`
	SucursalID  int16   `json:"sucursalID" validate:"required"`
	CampanaID   *int16  `json:"campanaID"`
	ComercioID  int16   `json:"comercioID" validate:"required"`
	ValorCompra float64 `json:"ValorCompra" validate:"required"`
}
