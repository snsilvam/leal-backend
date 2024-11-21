package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	campanaIn "leal-backend/src/adapters/campanas/campanas-in"
	campanasDatabase "leal-backend/src/adapters/campanas/campanas-out"
	lealHttp "leal-backend/src/adapters/in/http"
	"leal-backend/src/adapters/out/database"
	"leal-backend/src/domain/campanas/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DNS := os.Getenv("DNS")
	database, err := database.ConstructorDatabase(DNS)
	if err != nil {
		log.Fatal("error conectando a la base de datos", err)
	}
	fmt.Println(database)

	PORT := os.Getenv("PORT")
	server, err := lealHttp.ConstructorServerHttp(PORT)
	if err != nil {
		log.Fatal("error inicializando el servidor ", err)
	}

	campanaRepository := campanasDatabase.ConstructorCampanasPostgresRepository(*database)
	campanaService := service.ConstructorCampanaService(campanaRepository)
	campaanaHandler := campanaIn.ConstructorCampanaHandler(campanaService)
	campaanaHandler.RegisterCampanasRoutes(server.Router)

	server.Router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!. Welcome to leal api♥",
		})
	})
	server.StartServer()
}
