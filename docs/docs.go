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
        "/v1/bind_ip/{ip}": {
            "get": {
                "description": "Get node of specefic ip",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IP"
                ],
                "summary": "Show node of specefic ip",
                "parameters": [
                    {
                        "type": "string",
                        "description": "IP Address",
                        "name": "ip",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "object",
                                "additionalProperties": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/v1/bindings": {
            "get": {
                "description": "Get a list of IP addresses",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IP"
                ],
                "summary": "Show IP addresses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "type": "object",
                                    "additionalProperties": {
                                        "type": "string"
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8123",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "IP Mapper API",
	Description:      "This is a simple API to show IP mapping.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
