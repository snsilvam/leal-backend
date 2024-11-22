package beneficios

import "context"

type BeneficioService struct {
	repositoryBeneficio BeneficiosRepository
}

func ConstructorBeneficioService(repository BeneficiosRepository) *BeneficioService {
	return &BeneficioService{
		repositoryBeneficio: repository,
	}
}

func (s *BeneficioService) GetAllBeneficios(context context.Context) ([]*TipoBeneficios, error) {
	return s.repositoryBeneficio.GetAllBeneficios(context)
}
