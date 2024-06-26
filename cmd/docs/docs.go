// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Lebedev A.S.",
            "email": "9112441775@mail.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/info/": {
            "get": {
                "description": "get Cars info from the database with a search filter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "methods"
                ],
                "summary": "get Cars info from the database",
                "parameters": [
                    {
                        "type": "string",
                        "description": "regNum",
                        "name": "regNum",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "mark",
                        "name": "mark",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "year",
                        "name": "year",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "owner.name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "owner.surname",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "owner.patronymic",
                        "name": "patronymic",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "read records",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.Car"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "400"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "500"
                        }
                    }
                }
            },
            "post": {
                "description": "add car info to database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "methods"
                ],
                "summary": "add car info to to database",
                "parameters": [
                    {
                        "description": "slice reg numbers",
                        "name": "regNums",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.RegNums"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "added records",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.Car"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "400"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "500"
                        }
                    }
                }
            }
        },
        "/info/:id": {
            "get": {
                "description": "get car info by id from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "methods"
                ],
                "summary": "get car info by id from the database",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "read record",
                        "schema": {
                            "$ref": "#/definitions/service.Car"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "400"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "500"
                        }
                    }
                }
            },
            "put": {
                "description": "update car info in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "methods"
                ],
                "summary": "update car info in the database",
                "parameters": [
                    {
                        "description": "parameters of the record being updated",
                        "name": "car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.Car"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "updated record",
                        "schema": {
                            "$ref": "#/definitions/service.Car"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "400"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "500"
                        }
                    }
                }
            },
            "delete": {
                "description": "del car info from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "methods"
                ],
                "summary": "del car info from the database",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID delete car",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "deleted record",
                        "schema": {
                            "$ref": "#/definitions/service.Car"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "400"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "500"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "service.Car": {
            "description": "Car info information with regNum",
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "mark": {
                    "type": "string",
                    "example": "Lada"
                },
                "model": {
                    "type": "string",
                    "example": "Vesta"
                },
                "owner": {
                    "$ref": "#/definitions/service.People"
                },
                "regnum": {
                    "type": "string",
                    "example": "X123XX150"
                },
                "year": {
                    "type": "integer",
                    "example": 2020
                }
            }
        },
        "service.People": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "name"
                },
                "patronymic": {
                    "type": "string",
                    "example": "patronymic"
                },
                "surname": {
                    "type": "string",
                    "example": "surname"
                }
            }
        },
        "service.RegNums": {
            "description": "regNum list",
            "type": "object",
            "properties": {
                "regNums": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "X123XX150"
                    ]
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "http://127.0.0.1:80",
	BasePath:         "/info/",
	Schemes:          []string{},
	Title:            "Car Info",
	Description:      "для получения данных из внешного источника необходимо изменить значение  переменной URL_GETCARINFO в .env",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
