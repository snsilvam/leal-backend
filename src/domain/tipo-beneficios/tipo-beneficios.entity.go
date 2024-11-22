package beneficios

type TipoBeneficios struct {
	ID          int16
	Nombre      string
	Descripcion string
	TipoPuntos  string // pueden ser puntos_cashback o puntos_leal
}
