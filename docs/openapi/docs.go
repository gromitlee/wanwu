// Package openapi Code generated by swaggo/swag. DO NOT EDIT
package openapi

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
        "/agent/chat": {
            "post": {
                "description": "智能体对话OpenAPI",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "openapi"
                ],
                "summary": "智能体对话OpenAPI",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.OpenAPIAgentChatRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.OpenAPIAgentChatResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/agent/conversation": {
            "post": {
                "description": "智能体创建对话OpenAPI",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "openapi"
                ],
                "summary": "智能体创建对话OpenAPI",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.OpenAPIAgentCreateConversationRequest"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.OpenAPIAgentCreateConversationResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/rag/chat": {
            "post": {
                "description": "文本问答OpenAPI",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "openapi"
                ],
                "summary": "文本问答OpenAPI",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.OpenAPIRagChatRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.OpenAPIRagChatResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.OpenAPIAgentChatRequest": {
            "type": "object",
            "required": [
                "conversation_id",
                "query"
            ],
            "properties": {
                "conversation_id": {
                    "type": "string"
                },
                "query": {
                    "type": "string"
                },
                "stream": {
                    "type": "boolean"
                }
            }
        },
        "request.OpenAPIAgentCreateConversationRequest": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "request.OpenAPIRagChatRequest": {
            "type": "object",
            "required": [
                "query"
            ],
            "properties": {
                "query": {
                    "type": "string"
                },
                "stream": {
                    "type": "boolean"
                }
            }
        },
        "response.OpenAIChatHistory": {
            "type": "object",
            "properties": {
                "query": {
                    "type": "string"
                },
                "response": {
                    "type": "string"
                }
            }
        },
        "response.OpenAIChatSearch": {
            "type": "object",
            "properties": {
                "kb_name": {
                    "type": "string"
                },
                "snippet": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "response.OpenAPIAgentChatFile": {
            "type": "object",
            "properties": {
                "output_file_url": {
                    "type": "string"
                }
            }
        },
        "response.OpenAPIAgentChatResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "finish": {
                    "type": "integer"
                },
                "gen_file_url_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.OpenAPIAgentChatFile"
                    }
                },
                "history": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.OpenAIChatHistory"
                    }
                },
                "message": {
                    "type": "string"
                },
                "response": {
                    "type": "string"
                },
                "search_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.OpenAIChatSearch"
                    }
                },
                "usage": {
                    "$ref": "#/definitions/response.OpenAPIAgentChatUsage"
                }
            }
        },
        "response.OpenAPIAgentChatUsage": {
            "type": "object",
            "properties": {
                "completion_tokens": {
                    "type": "integer"
                },
                "prompt_tokens": {
                    "type": "integer"
                },
                "total_tokens": {
                    "type": "integer"
                }
            }
        },
        "response.OpenAPIAgentCreateConversationResponse": {
            "type": "object",
            "properties": {
                "conversation_id": {
                    "type": "string"
                }
            }
        },
        "response.OpenAPIRagChatData": {
            "type": "object",
            "properties": {
                "output": {
                    "type": "string"
                },
                "searchList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.OpenAIChatSearch"
                    }
                }
            }
        },
        "response.OpenAPIRagChatResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/response.OpenAPIRagChatData"
                },
                "finish": {
                    "type": "integer"
                },
                "history": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.OpenAIChatHistory"
                    }
                },
                "message": {
                    "type": "string"
                },
                "msg_id": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v0.0.1",
	Host:             "",
	BasePath:         "/openapi/v1",
	Schemes:          []string{},
	Title:            "AI Agent Productivity Platform - Open API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
