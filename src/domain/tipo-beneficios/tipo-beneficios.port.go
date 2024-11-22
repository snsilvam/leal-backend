package beneficios

import "context"

type BeneficiosRepository interface {
	GetAllBeneficios(context context.Context) ([]*TipoBeneficios, error)
}
