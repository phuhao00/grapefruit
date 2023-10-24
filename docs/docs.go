// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/login": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "login success response",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/do.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/do.LoginRsp"
                        }
                    }
                }
            }
        },
        "/api/register": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "user register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "register success response",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/do.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/do.RegisterRsp"
                        }
                    }
                }
            }
        },
        "/dev/api/login": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "login success response",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/do.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/do.LoginRsp"
                        }
                    }
                }
            }
        },
        "/dev/api/register": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "user register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "register success response",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/do.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/do.RegisterRsp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "do.LoginReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "pwd": {
                    "type": "string"
                }
            }
        },
        "do.LoginRsp": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "do.RegisterReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "pwd": {
                    "type": "string"
                }
            }
        },
        "do.RegisterRsp": {
            "type": "object"
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "grapefruit",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}