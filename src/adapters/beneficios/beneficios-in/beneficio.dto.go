package beneficioshandler

type TipoBeneficiosDTO struct {
	ID          int16  `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	TipoPuntos  string `json:"tipoPuntos"`
}
