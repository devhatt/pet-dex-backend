// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "DevHatt",
            "url": "https://github.com/devhatt"
        },
        "license": {
            "name": "MIT license",
            "url": "https://github.com/devhatt/pet-dex-backend?tab=MIT-1-ov-file#readme"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ongs/{ongID}": {
            "patch": {
                "description": "Updates the details of an existing Ong based on the provided Ong ID.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Ong"
                ],
                "summary": "Update Ong By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ong id to be updated",
                        "name": "ongID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Data to update of the Ong",
                        "name": "ongDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.OngUpdateDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/pets/": {
            "get": {
                "description": "Public route for viewing all pets.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pet"
                ],
                "summary": "View list of all pets.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Pet"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Sends the Pet's registration data via the request body for persistence in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pet"
                ],
                "summary": "Create Pet by petDto",
                "parameters": [
                    {
                        "description": "Pet object information for registration",
                        "name": "petDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PetInsertDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/pets/breeds/": {
            "get": {
                "description": "Retrieves list of all pet breeds",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pet"
                ],
                "summary": "View list of all Breeds",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BreedList"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/pets/{petID}": {
            "get": {
                "description": "Retrieves Pet details based on the pet ID provided as a parameter.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pet"
                ],
                "summary": "Find Pet by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the Pet to be retrieved",
                        "name": "petID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Pet"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user/create-account": {
            "post": {
                "description": "Creates user and insert into the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Creates user",
                "parameters": [
                    {
                        "description": "User object information to create",
                        "name": "userDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserInsertDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user/{userID}/my-pets": {
            "get": {
                "description": "List all pets owned by the user corresponding to the provided user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "List pets by user id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the User",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Pet"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user/{userID}/pets/{petID}": {
            "patch": {
                "description": "Update the Pet's registration data via the request body for persistence in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update an Pet existing.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Pet ID",
                        "name": "petID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Pet object information for update of data",
                        "name": "petDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PetUpdateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.BreedList": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "0e0b8399-1bf1-4ed5-a2f4-b5789ddf5df0"
                },
                "img_url": {
                    "type": "string",
                    "example": "https://images.unsplash.com/photo-1530281700549-e82e7bf110d6?q=80\u0026w=1888\u0026auto=format\u0026fit=crop\u0026ixlib=rb-4.0.3\u0026ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
                },
                "name": {
                    "type": "string",
                    "example": "Pastor Alemão"
                }
            }
        },
        "dto.Link": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string",
                    "example": "Facebook da Ong"
                },
                "url": {
                    "type": "string",
                    "example": "https://www.facebook.com/"
                }
            }
        },
        "dto.OngUpdateDto": {
            "type": "object",
            "properties": {
                "adoptionPolicy": {
                    "type": "string",
                    "example": "não pode rato"
                },
                "links": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Link"
                    }
                },
                "openingHours": {
                    "type": "string",
                    "example": "08:00"
                },
                "phone": {
                    "type": "string",
                    "example": "119596995887"
                },
                "user": {
                    "$ref": "#/definitions/dto.UserUpdateDto"
                }
            }
        },
        "dto.PetInsertDto": {
            "type": "object",
            "properties": {
                "adoption_date": {
                    "type": "string",
                    "example": "2008-01-02T15:04:05Z"
                },
                "birthdate": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05Z"
                },
                "breed_id": {
                    "type": "string",
                    "example": "0e0b8399-1bf1-4ed5-a2f4-b5789ddf5df0"
                },
                "name": {
                    "type": "string",
                    "example": "Thor"
                },
                "size": {
                    "type": "string",
                    "example": "medium"
                },
                "user_id": {
                    "type": "string",
                    "example": "fa1b8ae8-5351-11ef-8f02-0242ac130003"
                },
                "weight": {
                    "type": "number",
                    "example": 4.1
                }
            }
        },
        "dto.PetUpdateDto": {
            "type": "object",
            "properties": {
                "adoption_date": {
                    "type": "string",
                    "example": "2008-01-02T00:00:00Z"
                },
                "available_to_adoption": {
                    "type": "boolean",
                    "example": true
                },
                "birthdate": {
                    "type": "string",
                    "example": "2006-01-02T00:00:00Z"
                },
                "breed_id": {
                    "type": "string",
                    "example": "0e0b8399-1bf1-4ed5-a2f4-b5789ddf5df0"
                },
                "castrated": {
                    "type": "boolean",
                    "example": true
                },
                "comorbidity": {
                    "type": "string",
                    "example": "asma"
                },
                "name": {
                    "type": "string",
                    "example": "Spike"
                },
                "size": {
                    "type": "string",
                    "example": "small"
                },
                "special_care": {
                    "$ref": "#/definitions/dto.SpecialCareDto"
                },
                "tags": {
                    "type": "string",
                    "example": "Dog"
                },
                "vaccines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.VaccinesDto"
                    }
                },
                "weight": {
                    "type": "number",
                    "example": 4.8
                },
                "weight_measure": {
                    "type": "string",
                    "example": "kg"
                }
            }
        },
        "dto.SpecialCareDto": {
            "type": "object",
            "properties": {
                "descriptionSpecialCare": {
                    "type": "string",
                    "example": "obesity"
                },
                "neededSpecialCare": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "dto.UserInsertDto": {
            "type": "object",
            "properties": {
                "avatar_url": {
                    "type": "string",
                    "example": "https://example.com/avatar.jpg"
                },
                "birthdate": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05Z"
                },
                "city": {
                    "type": "string",
                    "example": "São Paulo"
                },
                "document": {
                    "type": "string",
                    "example": "12345678900"
                },
                "email": {
                    "type": "string",
                    "example": "claudio@example.com"
                },
                "name": {
                    "type": "string",
                    "example": "Claúdio"
                },
                "pass": {
                    "type": "string",
                    "example": "senhasegura123"
                },
                "phone": {
                    "type": "string",
                    "example": "21912345678"
                },
                "role": {
                    "type": "string",
                    "example": "developer"
                },
                "state": {
                    "type": "string",
                    "example": "São Paulo"
                },
                "type": {
                    "type": "string",
                    "example": "física"
                }
            }
        },
        "dto.UserUpdateDto": {
            "type": "object",
            "properties": {
                "avatar_url": {
                    "type": "string"
                },
                "birthdate": {
                    "type": "string"
                },
                "document": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "pushNotificationsEnabled": {
                    "type": "boolean"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "dto.VaccinesDto": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string",
                    "example": "2007-01-02T00:00:00Z"
                },
                "doctor_crm": {
                    "type": "string",
                    "example": "000000"
                },
                "name": {
                    "type": "string",
                    "example": "PetVax"
                }
            }
        },
        "entity.Pet": {
            "type": "object",
            "properties": {
                "adoption_date": {
                    "type": "string"
                },
                "available_to_adoption": {
                    "type": "boolean"
                },
                "birthdate": {
                    "type": "string"
                },
                "breed_id": {
                    "type": "string"
                },
                "breed_name": {
                    "type": "string"
                },
                "castrated": {
                    "type": "boolean"
                },
                "comorbidity": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "size": {
                    "type": "string"
                },
                "special_care": {
                    "$ref": "#/definitions/entity.SpecialCare"
                },
                "tags": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "vaccines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Vaccines"
                    }
                },
                "weight": {
                    "type": "number"
                },
                "weight_measure": {
                    "type": "string"
                }
            }
        },
        "entity.SpecialCare": {
            "type": "object",
            "properties": {
                "descriptionSpecialCare": {
                    "type": "string"
                },
                "neededSpecialCare": {
                    "type": "boolean"
                }
            }
        },
        "entity.Vaccines": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "doctor_crm": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pet_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000/api",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "PetDex: Documentação API",
	Description:      "Esta página se destina a documentação da API do projeto PetDex Backend",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
