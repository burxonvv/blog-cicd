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
        "/": {
            "get": {
                "description": "Just to ensure that server is running",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ping"
                ],
                "summary": "Ping pong",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/about": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Here about can be Updated.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "About"
                ],
                "summary": "Update about",
                "parameters": [
                    {
                        "description": "post info",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/aboutsh.UpdateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/aboutsh.FullResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Here about can be created.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "About"
                ],
                "summary": "Create about",
                "parameters": [
                    {
                        "description": "post info",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/aboutsh.CreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/aboutsh.FullResponse"
                        }
                    }
                }
            }
        },
        "/about/{id}": {
            "get": {
                "description": "Here about can be got.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "About"
                ],
                "summary": "Get about by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/aboutsh.FullResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Here about can be deleted.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "About"
                ],
                "summary": "Delete about",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/aboutsh.FullResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "aboutsh.CreateReq": {
            "type": "object",
            "properties": {
                "facebook": {
                    "type": "string"
                },
                "intro": {
                    "type": "string"
                },
                "linkedin": {
                    "type": "string"
                },
                "resumelink": {
                    "type": "string"
                },
                "telegram": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "youtube": {
                    "type": "string"
                }
            }
        },
        "aboutsh.FullResponse": {
            "type": "object",
            "properties": {
                "body": {
                    "$ref": "#/definitions/aboutsh.FullResponseBody"
                },
                "error_code": {
                    "type": "integer"
                },
                "error_message": {
                    "type": "string"
                }
            }
        },
        "aboutsh.FullResponseBody": {
            "type": "object",
            "properties": {
                "facebook": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "intro": {
                    "type": "string"
                },
                "linkedin": {
                    "type": "string"
                },
                "resumelink": {
                    "type": "string"
                },
                "telegram": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "youtube": {
                    "type": "string"
                }
            }
        },
        "aboutsh.UpdateReq": {
            "type": "object",
            "properties": {
                "facebook": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "intro": {
                    "type": "string"
                },
                "linkedin": {
                    "type": "string"
                },
                "resumelink": {
                    "type": "string"
                },
                "telegram": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "youtube": {
                    "type": "string"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "error_code": {
                    "type": "integer"
                },
                "error_message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Blog site project API Endpoints",
	Description:      "Here QA can test and frontend or mobile developers can get information of API endpoints.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
