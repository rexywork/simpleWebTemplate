// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplatev2 = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users": {
            "get": {
                "description": "List users existing.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.UserResponse"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.UserResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfov2 holds exported Swagger Info so clients can modify it
var SwaggerInfov2 = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "/v2",
	Schemes:          []string{"http"},
	Title:            "Gin Swagger Example API",
	Description:      "This is a sample server server.",
	InfoInstanceName: "v2",
	SwaggerTemplate:  docTemplatev2,
}

func init() {
	swag.Register(SwaggerInfov2.InstanceName(), SwaggerInfov2)
}
