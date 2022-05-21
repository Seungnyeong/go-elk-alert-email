// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://confluence.wemakeprice.com/pages/viewpage.action?pageId=206230173",
        "contact": {
            "name": "보안기술실 메일 전송",
            "email": "secutech@wemakeprice.com"
        },
        "license": {
            "name": "위메프 CERT팀 제공",
            "url": "https://stash.wemakeprice.com/projects/SECUTECH/repos/wkms-alert/browse"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/job/delete/{key}": {
            "delete": {
                "description": "key를 입력하세요",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "스케줄"
                ],
                "summary": "Alert Instance 삭제",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Remove Instance",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    }
                }
            }
        },
        "/job/list": {
            "get": {
                "description": "현재 실행되고 있는 잡을 알수있음.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "스케줄"
                ],
                "summary": "Alerting 이 되고 있는 인스턴스 전체 확인.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    }
                }
            }
        },
        "/job/start": {
            "post": {
                "description": "ipv4를 입력하세요",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "스케줄"
                ],
                "summary": "Alert Instance 등록",
                "parameters": [
                    {
                        "description": "Alerting 설정할 인스턴스 ipv4를 입력하세요",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/swagger.Param"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    }
                }
            }
        },
        "/users/list": {
            "get": {
                "description": "WKMS 관리자 전체 조회",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "계정"
                ],
                "summary": "WKMS 관리자 전체 조회",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    }
                }
            }
        },
        "/users/{username}": {
            "get": {
                "description": "username을 입력하세요",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "계정"
                ],
                "summary": "사용자 조회",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get One User",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/swagger.httpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "swagger.Param": {
            "type": "object",
            "properties": {
                "ipv4": {
                    "type": "string"
                }
            }
        },
        "swagger.httpResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "result": {}
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "wkms-alert module",
	Description:      "wkms alert moddule",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
