package beneficioshandler

import (
	beneficios "leal-backend/src/domain/tipo-beneficios"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BeneficiosHandler struct {
	serviceBeneficiosHandler *beneficios.BeneficioService
}

func ConstructorBeneficiosHandler(service *beneficios.BeneficioService) *BeneficiosHandler {
	return &BeneficiosHandler{
		serviceBeneficiosHandler: service,
	}
}

func (h *BeneficiosHandler) RegisterBeneficiosRoutes(router *gin.Engine) {
	beneficiosRoutes := router.Group("/beneficios")
	{
		beneficiosRoutes.GET("", h.GetAllBeneficios)
	}
}

func (h *BeneficiosHandler) GetAllBeneficios(c *gin.Context) {
	beneficios, err := h.serviceBeneficiosHandler.GetAllBeneficios(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error-message": "Error consultando los beneficios disponibles para las campañas.",
		})
		return
	}

	c.JSON(http.StatusOK, beneficios)
}
