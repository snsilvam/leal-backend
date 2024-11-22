package campanaIn

import (
	"fmt"
	in "leal-backend/src/adapters/campanas/campanas-in/dto"
	"leal-backend/src/domain/campanas/entity"
	"leal-backend/src/domain/campanas/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampanaHandler struct {
	service *service.CampanaService
}

func ConstructorCampanaHandler(service *service.CampanaService) *CampanaHandler {
	return &CampanaHandler{
		service: service,
	}
}

func (h *CampanaHandler) RegisterCampanasRoutes(router *gin.Engine) {
	campanasRoutes := router.Group("/campanas")
	{
		campanasRoutes.POST("", h.CreateCampana)
		campanasRoutes.GET("/:comercioID/:sucursalID", h.GetAllCampanas)
	}
}

func (h *CampanaHandler) CreateCampana(c *gin.Context) {
	fmt.Println("campaña recibida. . . ")
	var createCampanaDTO in.CreateCampanaDTO
	if err := c.BindJSON(&createCampanaDTO); err != nil {
		log.Println("error al recibir el gasto: ", err)

		c.JSON(404, gin.H{
			"error message": "error en el modelo de datos de la campaña.",
		})
		return
	}

	campanaEntity := entity.Campana{
		SucursalID:         createCampanaDTO.SucursalID,
		ComercioID:         createCampanaDTO.ComercioID,
		TipoBeneficioID:    createCampanaDTO.TipoBeneficioID,
		ComprasRegistradas: 0,
		FechaInicio:        createCampanaDTO.FechaInicio,
		FechaFin:           createCampanaDTO.FechaFin,
		Estado:             createCampanaDTO.Estado,
	}

	err := h.service.CreateCampana(c, &campanaEntity)
	if err != nil {
		c.JSON(500, gin.H{
			"error message": "error creando la campaña en el sistema.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, "Campaña almacenada en la base de datos.")
}

func (h *CampanaHandler) GetAllCampanas(c *gin.Context) {
	comercioId := c.Param("comercioID")
	comercioID, err := strconv.Atoi(comercioId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ComercioID debe ser un entero.",
		})
		return
	}

	sucursalId := c.Param("sucursalID")
	sucursalID, err := strconv.Atoi(sucursalId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "SucursalID debe ser un entero.",
		})
		return
	}

	campanas, err := h.service.GetAllCampanasOfComercioAndSucursal(c, int16(comercioID), int16(sucursalID))
	if err != nil {
		c.JSON(500, gin.H{
			"error-message": "Error consultando las campañas.",
		})
		return
	}

	c.JSON(http.StatusOK, campanas)
}
