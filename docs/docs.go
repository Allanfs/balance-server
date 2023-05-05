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
        "/:id": {
            "get": {
                "responses": {}
            },
            "delete": {
                "description": "Remove o lançamento com base no seu ID\nwith user id and username",
                "summary": "Remove um lançamento",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/lancamentos": {
            "post": {
                "description": "Informe o nome, valor (maior que zero) e tipo do lançamento (C, D ou N/A).\nO campo external_info é opcional. Serve para clientes da API armazenarem informações que lhes sejam uteis.",
                "summary": "Cadastra um lançamento",
                "parameters": [
                    {
                        "description": "Objeto para cadastrar um novo lançamento",
                        "name": "req.body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/inbound.NovoLancamentoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/inbound.NovoLancamentoResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "inbound.NovoLancamentoRequest": {
            "type": "object",
            "properties": {
                "external_info": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "tipo": {
                    "$ref": "#/definitions/model.NaturezaFinanceira"
                },
                "valor": {
                    "type": "number"
                }
            }
        },
        "inbound.NovoLancamentoResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.NaturezaFinanceira": {
            "type": "string",
            "enum": [
                "C",
                "D",
                "N/A"
            ],
            "x-enum-varnames": [
                "Credito",
                "Debito",
                "Indefinido"
            ]
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
