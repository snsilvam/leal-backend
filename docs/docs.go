// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/beneficios": {
            "get": {
                "description": "Devuelve todos los tipos de beneficios definidos en el sistema para crear una campaña.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tipos de beneficios definidos para las campañas."
                ],
                "summary": "Obtiene todos los tipos de beneficios definidos en el sistema.",
                "responses": {
                    "200": {
                        "description": "Beneficios registrados.",
                        "schema": {
                            "$ref": "#/definitions/beneficioshandler.TipoBeneficiosDTO"
                        }
                    }
                }
            }
        },
        "/campanas": {
            "post": {
                "description": "Crea una campaña, relacionada a una sucursal y a una campaña, con algun tipo de los beneficios definidos en el sistema.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Campañas"
                ],
                "summary": "Crea una camapaña, para un comercio y sucursal.",
                "parameters": [
                    {
                        "description": "Datos de la campaña.",
                        "name": "campana",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/campanaIn.CampanaDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/campanas/:comercioID/:sucursalID": {
            "get": {
                "description": "Consulta todas las campañas relacionadas a un comercio y una sucursal, por mediod el id de cada uno.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Campañas"
                ],
                "summary": "Consultar camapañas de un comercio y sucursal.",
                "responses": {
                    "200": {
                        "description": "Información de las campañas registradas.",
                        "schema": {
                            "$ref": "#/definitions/campanaIn.CampanaDTO"
                        }
                    }
                }
            }
        },
        "/comercios": {
            "get": {
                "description": "Consultar los comercios disponibles, para crear una sucursal o una campaña.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comercio"
                ],
                "summary": "Consultar los comercios disponibles.",
                "responses": {
                    "200": {
                        "description": "Información del beneficio",
                        "schema": {
                            "$ref": "#/definitions/comercioin.Comercio"
                        }
                    }
                }
            }
        },
        "/compras": {
            "post": {
                "description": "Crea una compra en una sucursal/comercio y si tiene activa una campaña, dispara la acumulacion de puntos leal o cashback.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Compras"
                ],
                "summary": "Crea una compra :D",
                "parameters": [
                    {
                        "description": "Datos para realizar una compra.",
                        "name": "compra",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/compraIn.CreateCompraDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/sucursales/:comercioId": {
            "get": {
                "description": "Consultar todas las sucursales relacionadas a un comercio.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sucursales"
                ],
                "summary": "Consultar las sucursales de un comercio.",
                "responses": {
                    "200": {
                        "description": "Sucursales registradas para un comercio.",
                        "schema": {
                            "$ref": "#/definitions/sucursaleshandler.SucursalesDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "beneficioshandler.TipoBeneficiosDTO": {
            "type": "object",
            "properties": {
                "descripcion": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nombre": {
                    "type": "string"
                },
                "tipoPuntos": {
                    "type": "string"
                }
            }
        },
        "campanaIn.CampanaDTO": {
            "type": "object",
            "required": [
                "comercioID",
                "estado",
                "fechaFin",
                "fechaInicio",
                "sucursalID",
                "tipoBeneficioID"
            ],
            "properties": {
                "comercioID": {
                    "type": "integer"
                },
                "estado": {
                    "type": "boolean"
                },
                "fechaFin": {
                    "type": "string"
                },
                "fechaInicio": {
                    "type": "string"
                },
                "sucursalID": {
                    "type": "integer"
                },
                "tipoBeneficioID": {
                    "type": "integer"
                }
            }
        },
        "comercioin.Comercio": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "NombreComercio": {
                    "type": "string"
                }
            }
        },
        "compraIn.CreateCompraDTO": {
            "type": "object",
            "required": [
                "ValorCompra",
                "comercioID",
                "sucursalID",
                "usuarioID"
            ],
            "properties": {
                "ValorCompra": {
                    "type": "number"
                },
                "campanaID": {
                    "type": "integer"
                },
                "comercioID": {
                    "type": "integer"
                },
                "sucursalID": {
                    "type": "integer"
                },
                "usuarioID": {
                    "type": "integer"
                }
            }
        },
        "sucursaleshandler.SucursalesDTO": {
            "type": "object",
            "properties": {
                "comercioID": {
                    "type": "integer"
                },
                "comprasRegistradas": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "ubicacion": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Prueba tecnica - Leal Backend Go",
	Description:      "Este es un proyecto con arquitectura hexagonal que permite realizar y si aplica acumular puntos para un usuario.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
