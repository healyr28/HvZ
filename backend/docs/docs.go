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
        "/cure": {
            "get": {
                "description": "Cures a zombie or infected human",
                "produces": [
                    "text/plain"
                ],
                "summary": "Cure",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Target ID",
                        "name": "target",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "User is not logged in",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "No cures available",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Target does not exist",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Target is not a zombie",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/kill": {
            "get": {
                "description": "Kills a zombie if user is authenticated as a human. Gives zombie target the stunned state",
                "produces": [
                    "text/plain"
                ],
                "summary": "Kill",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Target ID",
                        "name": "target",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "User is not logged in",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Killer is not a human",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Target does not exist",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Target is not a zombie",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Player login",
                        "name": "info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Player"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "403": {
                        "description": "Forbidden"
                    }
                }
            }
        },
        "/me": {
            "get": {
                "description": "Shows user information",
                "produces": [
                    "application/json"
                ],
                "summary": "Me",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Player"
                        }
                    },
                    "401": {
                        "description": "User is not logged in",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/revive": {
            "get": {
                "description": "Revives a stunned zombie",
                "produces": [
                    "text/plain"
                ],
                "summary": "Revive",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Target ID",
                        "name": "target",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "User is not logged in",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "No revives available",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Target does not exist",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Target is not stunned",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tag": {
            "get": {
                "description": "Tag a human as a zombie. Gives target the infected state if user is a zombie or core zombie.",
                "produces": [
                    "text/plain"
                ],
                "summary": "Tag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Target ID",
                        "name": "target",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "User is not logged in",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Tagger is not a zombie",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Target does not exist",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Target is not human",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Login": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "main.Player": {
            "type": "object",
            "properties": {
                "cures": {
                    "type": "integer"
                },
                "extensions": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "kills": {
                    "type": "integer"
                },
                "last_kill": {
                    "type": "string"
                },
                "last_tagged": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "revives": {
                    "type": "integer"
                },
                "state": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "DCU Games Society Humans Vs Zombies",
	Description:      "This is the Swagger documentation for DCU Games Society's 2022 Humans Vs Zombies event.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}