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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user": {
            "get": {
                "tags": [
                    "User"
                ],
                "summary": "mengambil data-data user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Jika tidak memasukkan user ID maka akan get semua",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.RespGetUser"
                        }
                    }
                }
            },
            "put": {
                "tags": [
                    "User"
                ],
                "summary": "mengubah data user berdasarkan ID",
                "parameters": [
                    {
                        "description": "User ID (angka positif)",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.ReqUpdateUserById"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "Membuat user baru",
                "parameters": [
                    {
                        "description": "request body JSON",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.ReqCreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.RespCreateUser"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "delete": {
                "tags": [
                    "User"
                ],
                "summary": "Menghapus data user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID (angka positif)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "user.ReqCreateUser": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "user.ReqUpdateUserById": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "user.RespCreateUser": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "user.RespGetUser": {
            "type": "object",
            "properties": {
                "total": {
                    "type": "integer"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.User"
                    }
                }
            }
        },
        "user.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
