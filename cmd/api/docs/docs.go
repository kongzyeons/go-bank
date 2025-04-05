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
        "/api/v1/account/edit/{accountID}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Edit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account API"
                ],
                "summary": "Edit",
                "operationId": "AccountEdit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "accountID",
                        "name": "accountID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body",
                        "name": "AccountEditReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AccountEditReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/account/getlist": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GetList",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HomePage API"
                ],
                "summary": "GetList",
                "operationId": "AccountGetList",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "AccountGetListReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AccountGetListReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/auth/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Logout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth API"
                ],
                "summary": "Logout",
                "operationId": "AuthLogout",
                "responses": {}
            }
        },
        "/api/v1/auth/ping": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth API"
                ],
                "summary": "Ping",
                "operationId": "AuthPing",
                "responses": {}
            }
        },
        "/api/v1/auth/refresh": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Refresh",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth API"
                ],
                "summary": "Refresh",
                "operationId": "AuthRefresh",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "AuthRefreshReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AuthRefreshReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/banner/getlist": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GetList",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HomePage API"
                ],
                "summary": "GetList",
                "operationId": "BannerGetList",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "BannerGetListReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BannerGetListReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/debitCard/getlist": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GetList",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HomePage API"
                ],
                "summary": "GetList",
                "operationId": "debitCardGetList",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "DebitCardGetListReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DebitCardGetListReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth API"
                ],
                "summary": "Login",
                "operationId": "AuthLogin",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "AuthLoginReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AuthLoginReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/register": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth API"
                ],
                "summary": "Register",
                "operationId": "AuthRegister",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "AuthRegisterReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AuthRegisterReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/user/greeting": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "GetGeeting",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HomePage API"
                ],
                "summary": "GetGeeting",
                "operationId": "UserGetGeeting",
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.AccountEditReq": {
            "type": "object",
            "required": [
                "color",
                "name"
            ],
            "properties": {
                "color": {
                    "type": "string",
                    "maxLength": 10,
                    "example": "color"
                },
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "example": "name"
                }
            }
        },
        "models.AccountGetListReq": {
            "type": "object",
            "properties": {
                "isManinAccount": {
                    "type": "boolean",
                    "example": true
                },
                "page": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 1
                },
                "perPage": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 10
                },
                "searchText": {
                    "type": "string",
                    "example": "search by type or name"
                },
                "sortBy": {
                    "type": "object",
                    "properties": {
                        "field": {
                            "type": "string",
                            "example": "updatedDate"
                        },
                        "mode": {
                            "type": "string",
                            "example": "desc"
                        }
                    }
                }
            }
        },
        "models.AuthLoginReq": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 6,
                    "example": "123456"
                },
                "username": {
                    "type": "string",
                    "maxLength": 100,
                    "example": "admin"
                }
            }
        },
        "models.AuthRefreshReq": {
            "type": "object",
            "properties": {
                "refToken": {
                    "type": "string"
                }
            }
        },
        "models.AuthRegisterReq": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 6,
                    "example": "123456"
                },
                "username": {
                    "type": "string",
                    "maxLength": 100,
                    "example": "admin"
                }
            }
        },
        "models.BannerGetListReq": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 1
                },
                "perPage": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 10
                },
                "searchText": {
                    "type": "string",
                    "example": "search by title"
                },
                "sortBy": {
                    "type": "object",
                    "properties": {
                        "field": {
                            "type": "string",
                            "example": "updatedDate"
                        },
                        "mode": {
                            "type": "string",
                            "example": "desc"
                        }
                    }
                }
            }
        },
        "models.DebitCardGetListReq": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 1
                },
                "perPage": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 10
                },
                "searchText": {
                    "type": "string",
                    "example": "search by name"
                },
                "sortBy": {
                    "type": "object",
                    "properties": {
                        "field": {
                            "type": "string",
                            "example": "updatedDate"
                        },
                        "mode": {
                            "type": "string",
                            "example": "desc"
                        }
                    }
                },
                "status": {
                    "type": "string",
                    "example": ""
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "API-BANK",
	Description:      "This is a sample server for a github.com/kongzyeons/go-bank",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
