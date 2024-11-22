package sucursaleshandler

import (
	"leal-backend/src/domain/sucursales"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SucursalesHandler struct {
	serviceSucursalesHandler *sucursales.SucursalesService
}

func ConstructorSucursalesHandler(service *sucursales.SucursalesService) *SucursalesHandler {
	return &SucursalesHandler{
		serviceSucursalesHandler: service,
	}
}

func (h *SucursalesHandler) RegisterSucursalesRoutes(router *gin.Engine) {
	sucursalesRoutes := router.Group("/sucursales")
	{
		sucursalesRoutes.GET("/:comercioId", h.GetAllSucursales)
	}
}

// @Summary Consultar las sucursales de un comercio.
// @Description Consultar todas las sucursales relacionadas a un comercio.
// @Tags Sucursales
// @Accept json
// @Produce json
// @Success 200 {object} SucursalesDTO "Sucursales registradas para un comercio."
// @Router /sucursales/:comercioId [get]
func (h *SucursalesHandler) GetAllSucursales(c *gin.Context) {
	comercioId := c.Param("comercioId")
	intID, err := strconv.Atoi(comercioId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID debe ser un entero.",
		})
		return
	}

	sucursales, err := h.serviceSucursalesHandler.GetAllSucursales(c, int16(intID))
	if err != nil {
		c.JSON(500, gin.H{
			"error-message": "Error consultando las sucursales.",
		})
		return
	}

	c.JSON(http.StatusOK, sucursales)
}
