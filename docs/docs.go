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
        "/api/v1/auth/signup": {
            "post": {
                "description": "Register a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "registration is successful, please check your email for email verification",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "invalid payload\" or \"error validating payload\" or \"registration failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.TRegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "Email of the user (required, should be a valid email address)\nExample: john.doe@example.com\nRequired: true\nFormat: email",
                    "type": "string"
                },
                "name": {
                    "description": "Name of the user (required)\nExample: John Doe\nRequired: true",
                    "type": "string"
                },
                "password": {
                    "description": "Password of the user (required, minimum length 8)\nExample: mySecurePassword\nRequired: true\nMinLength: 8",
                    "type": "string",
                    "minLength": 8
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/swagger/index.html",
	Schemes:          []string{"http"},
	Title:            "CodeNexus API",
	Description:      "Happy integration",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}