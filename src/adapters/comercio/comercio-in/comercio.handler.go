package comercioin

import (
	comercioService "leal-backend/src/domain/comercios/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ComercioHandler struct {
	service *comercioService.ComercioService
}

func ConstructorComercioHandler(service *comercioService.ComercioService) *ComercioHandler {
	return &ComercioHandler{
		service: service,
	}
}

func (h *ComercioHandler) RegisterComercioRoutes(router *gin.Engine) {
	comercioRoutes := router.Group("/comercios")
	{
		comercioRoutes.GET("", h.GetAllComercios)
	}
}

// @Summary Consultar los comercios disponibles.
// @Description Consultar los comercios disponibles, para crear una sucursal o una campaña.
// @Tags Comercio
// @Accept json
// @Produce json
// @Success 200 {object} Comercio "Información del beneficio"
// @Router /comercios [get]
func (h *ComercioHandler) GetAllComercios(c *gin.Context) {
	comercios, err := h.service.GetAllComercios(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error-message": "Error consultando los comercios.",
		})
		return
	}

	c.JSON(http.StatusOK, comercios)
}
