package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	lealHttp "leal-backend/src/adapters/in/http"
	"leal-backend/src/adapters/out/database"

	campanaIn "leal-backend/src/adapters/campanas/campanas-in"
	campanasDatabase "leal-backend/src/adapters/campanas/campanas-out"
	"leal-backend/src/domain/campanas/service"

	comercioin "leal-backend/src/adapters/comercio/comercio-in"
	comercioout "leal-backend/src/adapters/comercio/comercio-out"
	comercio "leal-backend/src/domain/comercios/service"

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

	// Modulo Campañas.
	campanaRepository := campanasDatabase.ConstructorCampanasPostgresRepository(*database)
	campanaService := service.ConstructorCampanaService(campanaRepository)
	campaanaHandler := campanaIn.ConstructorCampanaHandler(campanaService)
	campaanaHandler.RegisterCampanasRoutes(server.Router)

	// Modulo Comercios.
	comercioRepository := comercioout.ConstructorComercioPostgresRepository(*database)
	comercioService := comercio.ConstructorComercioService(comercioRepository)
	comercioHandler := comercioin.ConstructorComercioHandler(comercioService)
	comercioHandler.RegisterComercioRoutes(server.Router)

	server.Router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!. Welcome to leal api♥",
		})
	})
	server.StartServer()
}
