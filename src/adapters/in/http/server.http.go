package http

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ServerHttp struct {
	Port   string
	Router *gin.Engine
}

func ConstructorServerHttp(port string) (*ServerHttp, error) {
	fmt.Println("puerto recibido: ", port)
	if port == "" {
		return nil, errors.New("el servidor necesita un puerto para inicializar la applicacion")
	}

	router := gin.Default()

	return &ServerHttp{
		Port:   port,
		Router: router,
	}, nil
}

func (s *ServerHttp) StartServer() {
	fmt.Println("Servidor inicializado en el puerto: ", s.Port)
	s.Router.Run(":" + s.Port)
}
