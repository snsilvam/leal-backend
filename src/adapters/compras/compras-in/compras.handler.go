package compraIn

import (
	"leal-backend/src/domain/compras"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CompraHandler struct {
	service *compras.CompraService
}

func ConstructorCompraHandler(service *compras.CompraService) *CompraHandler {
	return &CompraHandler{
		service: service,
	}
}

func (h *CompraHandler) RegisterComprasRoutes(router *gin.Engine) {
	comprasRoutes := router.Group("/compras")
	{
		comprasRoutes.POST("", h.CreateCompra)
	}
}

// @Summary Crea una compra :D
// @Description Crea una compra en una sucursal/comercio y si tiene activa una campaña, dispara la acumulacion de puntos leal o cashback.
// @Accept json
// @Tags Compras
// @Param compra body CreateCompraDTO true "Datos para realizar una compra."
// @Produce json
// @Success 200 {object} map[string]string
// @Router /compras [post]
func (h *CompraHandler) CreateCompra(c *gin.Context) {
	var createCompraDTO CreateCompraDTO
	if err := c.BindJSON(&createCompraDTO); err != nil {
		log.Println("error al recibir el JSON de la compra: ", err)

		c.JSON(404, gin.H{
			"error message": "error en el modelo de datos de la compra.",
		})
		return
	}

	compraEntity := compras.Compras{
		UsuarioID:   createCompraDTO.UsuarioID,
		SucursalID:  createCompraDTO.SucursalID,
		CampanaID:   *createCompraDTO.CampanaID,
		ComercioID:  createCompraDTO.ComercioID,
		FechaCompra: time.Now(),
		ValorCompra: createCompraDTO.ValorCompra,
	}

	err := h.service.CreateCompra(c, &compraEntity)
	if err != nil {
		c.JSON(500, gin.H{
			"error message": "error creando la compra en el sistema.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, "Compra almacenada en la base de datos.")
}
