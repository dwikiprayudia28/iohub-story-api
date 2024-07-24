// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "IOHub Story API Support",
            "email": "dwikiprayudia123@gmail.com"
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
        "/auth/signin": {
            "post": {
                "description": "Authenticates a user based on provided login and password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign in",
                "parameters": [
                    {
                        "description": "Signin data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignInReqDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully authenticated.",
                        "schema": {
                            "$ref": "#/definitions/jwt.JWTResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized: Invalid credentials.",
                        "schema": {
                            "$ref": "#/definitions/errorhandler.Response"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Registers a new user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Sign Up Info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignupReqDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/jwt.JWTResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorhandler.Response"
                        }
                    }
                }
            }
        },
        "/chat/room": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Allows authenticated users to create a new chat room.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "Create a new chat room",
                "parameters": [
                    {
                        "description": "Create Room",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateRoomReqDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/values.RoomValue"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorhandler.Response"
                        }
                    }
                }
            }
        },
        "/chat/rooms": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves a list of chat rooms based on filters, sorting and pagination.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "Query chat rooms",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number for pagination",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Number of items per page",
                        "name": "per_page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "\"created_at\"",
                        "description": "Field to sort by",
                        "name": "sort_field",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "default": "\"desc\"",
                        "description": "Order of sorting",
                        "name": "sort_order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter rooms where user is in",
                        "name": "user_in",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter rooms where user is not in",
                        "name": "user_not_in",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/values.RoomValue"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorhandler.Response"
                        }
                    }
                }
            }
        },
        "/chat/rooms/{roomID}/messages": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Creates a new message within the specified chat room.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "Create a new message",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the chat room",
                        "name": "roomID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Message details",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateMessageReqDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/values.MessageValue"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorhandler.Response"
                        }
                    }
                }
            }
        },
        "/rooms/{roomID}/messages": {
            "get": {
                "description": "Retrieve messages for a room, with optional pagination and sorting.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "Query messages for a specific room.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the Room",
                        "name": "roomID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page number for pagination. Defaults to 1.",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of messages per page for pagination. Defaults to 10.",
                        "name": "per_page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to sort by. Defaults to created_at.",
                        "name": "sort_field",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Order of sorting (asc/desc). Defaults to desc.",
                        "name": "sort_order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/values.MessageValue"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorhandler.Response"
                        }
                    }
                }
            }
        },
        "/user/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get the authenticated user's own details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User's own details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/values.UserValue"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorhandler.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateMessageReqDTO": {
            "type": "object",
            "required": [
                "content",
                "message_type"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "message_type": {
                    "type": "string"
                }
            }
        },
        "dto.CreateRoomReqDTO": {
            "type": "object",
            "required": [
                "owner"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "is_group": {
                    "description": "Config",
                    "type": "boolean"
                },
                "name": {
                    "description": "Meta",
                    "type": "string"
                },
                "owner": {
                    "description": "Users",
                    "type": "string"
                },
                "user_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.SignInReqDTO": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "description": "The login of the user. It must be alphanumeric and have a length between 3 and 100.\nrequired: true\nexample: john_doe",
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3
                },
                "password": {
                    "description": "The password of the user. It must have a length between 8 and 32.\nrequired: true\nexample: securePassword123",
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8
                }
            }
        },
        "dto.SignupReqDTO": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "description": "The email address of the user. It must be a valid email format.\nrequired: true\nexample: john.doe@example.com",
                    "type": "string"
                },
                "password": {
                    "description": "The desired password for the user. It must have a length between 8 and 32.\nrequired: true\nexample: securePassword123",
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8
                },
                "username": {
                    "description": "The desired username for the user. It must be alphanumeric and have a length between 3 and 20.\nrequired: true\nexample: john_doe",
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                }
            }
        },
        "entity.MessageType": {
            "type": "string",
            "enum": [
                "text",
                "image",
                "video",
                "audio",
                "file",
                "link"
            ],
            "x-enum-varnames": [
                "MessageTypeText",
                "MessageTypeImage",
                "MessageTypeVideo",
                "MessageTypeAudio",
                "MessageTypeFile",
                "MessageTypeLink"
            ]
        },
        "entity.RoomConfig": {
            "type": "object",
            "properties": {
                "room_type": {
                    "$ref": "#/definitions/entity.RoomType"
                }
            }
        },
        "entity.RoomMeta": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.RoomType": {
            "type": "string",
            "enum": [
                "group",
                "private"
            ],
            "x-enum-varnames": [
                "RoomTypeGroup",
                "RoomTypePrivate"
            ]
        },
        "entity.RoomUser": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "Timestamps",
                    "type": "string"
                },
                "is_muted": {
                    "description": "Booleans",
                    "type": "boolean"
                },
                "role": {
                    "$ref": "#/definitions/entity.RoomUserRole"
                },
                "status": {
                    "description": "Status",
                    "allOf": [
                        {
                            "$ref": "#/definitions/entity.RoomUserStatus"
                        }
                    ]
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "description": "User Details",
                    "type": "string"
                }
            }
        },
        "entity.RoomUserRole": {
            "type": "string",
            "enum": [
                "admin",
                "member"
            ],
            "x-enum-varnames": [
                "RoomUserRoleAdmin",
                "RoomUserRoleMember"
            ]
        },
        "entity.RoomUserStatus": {
            "type": "string",
            "enum": [
                "active",
                "left"
            ],
            "x-enum-varnames": [
                "RoomUserStatusActive",
                "RoomUserStatusLeft"
            ]
        },
        "errorhandler.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {},
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "jwt.JWTResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_in": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "values.MessageValue": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "Attributes",
                    "type": "string"
                },
                "created_at": {
                    "description": "Timestamps",
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "message_type": {
                    "$ref": "#/definitions/entity.MessageType"
                },
                "room_id": {
                    "description": "Relations",
                    "type": "string"
                },
                "sender_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "values.RoomValue": {
            "type": "object",
            "properties": {
                "config": {
                    "$ref": "#/definitions/entity.RoomConfig"
                },
                "created_at": {
                    "description": "Timestamps",
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_message_at": {
                    "type": "string"
                },
                "last_message_id": {
                    "description": "Last Message Fields",
                    "type": "string"
                },
                "meta": {
                    "description": "Room Details",
                    "allOf": [
                        {
                            "$ref": "#/definitions/entity.RoomMeta"
                        }
                    ]
                },
                "updated_at": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.RoomUser"
                    }
                }
            }
        },
        "values.UserValue": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "Timestamps",
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "description": "Account information",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1/",
	Schemes:          []string{},
	Title:            "IOHub Story API",
	Description:      "IOHub Story API Documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
