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
        "/auth": {
            "post": {
                "description": "Authenticates a user and returns a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Auth Endpoint",
                "parameters": [
                    {
                        "description": "Credentials credentials",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/driver/nearest": {
            "post": {
                "description": "GetNearestDriver  and returns the driver location and distance of the driver",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetNearestDriver"
                ],
                "summary": "GetNearestDriver Endpoint",
                "parameters": [
                    {
                        "description": "models.GeoJSON credentials",
                        "name": "userLocation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GeoJSON"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.DriverLocation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/import": {
            "post": {
                "description": "ImportDrivers get filepath and import the csv to mongodb",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ImportDrivers"
                ],
                "summary": "ImportDrivers Endpoint",
                "parameters": [
                    {
                        "description": "models.FilePath credentials",
                        "name": "filePath",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FilePath"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.FilePath"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/location": {
            "post": {
                "description": "CreateLocation get location info and create new driver location in db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CreateLocation"
                ],
                "summary": "CreateLocation Endpoint",
                "parameters": [
                    {
                        "description": "models.DriverLocation credentials",
                        "name": "location",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DriverLocation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.DriverLocation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.User": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.DriverLocation": {
            "type": "object",
            "properties": {
                "distance": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "$ref": "#/definitions/models.GeoJSON"
                }
            }
        },
        "models.FilePath": {
            "type": "object",
            "properties": {
                "path": {
                    "type": "string"
                }
            }
        },
        "models.GeoJSON": {
            "type": "object",
            "properties": {
                "coordinates": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "type": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
