// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://www.swagger.io/terms",
        "contact": {
            "name": "Allan Cordeiro",
            "url": "http://www.allancordeiro.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/retrieve/{cep}": {
            "get": {
                "description": "Find CEP through different providers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cep retriever"
                ],
                "summary": "HandleGet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cep code",
                        "name": "cep",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/webserver.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/webserver.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "webserver.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "CEP Retriever",
	Description:      "CEP retriever document. Fetch values like address, district and city through cep code.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
