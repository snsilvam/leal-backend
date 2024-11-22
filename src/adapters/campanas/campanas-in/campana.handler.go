package campanaIn

import (
	"fmt"
	in "leal-backend/src/adapters/campanas/campanas-in/dto"
	"leal-backend/src/domain/campanas/entity"
	"leal-backend/src/domain/campanas/service"
	"log"
	"net/http"
	"strconv"
	"time"

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

// @Summary Crea una camapaña, para un comercio y sucursal.
// @Description Crea una campaña, relacionada a una sucursal y a una campaña, con algun tipo de los beneficios definidos en el sistema.
// @Tags Campañas
// @Param campana body CampanaDTO true "Datos de la campaña."
// @Accept json
// @Produce json
// @Success 200
// @Router /campanas [post]
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

// @Summary Consultar camapañas de un comercio y sucursal.
// @Description Consulta todas las campañas relacionadas a un comercio y una sucursal, por mediod el id de cada uno.
// @Tags Campañas
// @Accept json
// @Produce json
// @Success 200 {object} CampanaDTO "Información de las campañas registradas."
// @Router /campanas/:comercioID/:sucursalID [get]
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

// Objeto para mostrar en los models del swagger.
type CampanaDTO struct {
	SucursalID      int16     `json:"sucursalID" validate:"required"`
	ComercioID      int16     `json:"comercioID" validate:"required"`
	TipoBeneficioID int16     `json:"tipoBeneficioID" validate:"required"`
	FechaInicio     time.Time `json:"fechaInicio" validate:"required"`
	FechaFin        time.Time `json:"fechaFin" validate:"required,gtfield=FechaInicio"`
	Estado          bool      `json:"estado" validate:"required"`
}
