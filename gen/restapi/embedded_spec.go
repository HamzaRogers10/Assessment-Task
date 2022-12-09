// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Assessment Task Wancloud",
    "version": "0.1.0"
  },
  "host": "localhost:3000",
  "basePath": "/v1",
  "paths": {
    "/login": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Users"
        ],
        "summary": "logging the user",
        "operationId": "loginUser",
        "parameters": [
          {
            "description": "This is how the body of the login user request body will look like.",
            "name": "LoginUserBody",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginUserDefinition"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Sucessfull Registeration",
            "schema": {
              "$ref": "#/definitions/LoginSuccessResponseDefinition"
            }
          },
          "400": {
            "description": "Invalid username/password supplied"
          }
        }
      }
    },
    "/register": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Users"
        ],
        "summary": "This endpoint entertain the registering user functionality",
        "operationId": "RegisterUser",
        "parameters": [
          {
            "description": "This is how the body of the register user request body will look like.",
            "name": "RegisterUserBody",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RegisterUserDefinition"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Sucessfull Registeration",
            "schema": {
              "$ref": "#/definitions/SuccessResponseDefinition"
            }
          },
          "400": {
            "description": "server could not understand the request due to invalid syntax"
          }
        }
      }
    }
  },
  "definitions": {
    "LoginSuccessResponseDefinition": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "LoginUserDefinition": {
      "type": "object",
      "required": [
        "Email",
        "Password"
      ],
      "properties": {
        "Email": {
          "type": "string"
        },
        "Password": {
          "type": "string"
        }
      }
    },
    "RegisterUserDefinition": {
      "type": "object",
      "required": [
        "Name",
        "Email",
        "Password"
      ],
      "properties": {
        "Email": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Password": {
          "type": "string"
        }
      }
    },
    "SuccessResponseDefinition": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Assessment Task Wancloud",
    "version": "0.1.0"
  },
  "host": "localhost:3000",
  "basePath": "/v1",
  "paths": {
    "/login": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Users"
        ],
        "summary": "logging the user",
        "operationId": "loginUser",
        "parameters": [
          {
            "description": "This is how the body of the login user request body will look like.",
            "name": "LoginUserBody",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginUserDefinition"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Sucessfull Registeration",
            "schema": {
              "$ref": "#/definitions/LoginSuccessResponseDefinition"
            }
          },
          "400": {
            "description": "Invalid username/password supplied"
          }
        }
      }
    },
    "/register": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Users"
        ],
        "summary": "This endpoint entertain the registering user functionality",
        "operationId": "RegisterUser",
        "parameters": [
          {
            "description": "This is how the body of the register user request body will look like.",
            "name": "RegisterUserBody",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RegisterUserDefinition"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Sucessfull Registeration",
            "schema": {
              "$ref": "#/definitions/SuccessResponseDefinition"
            }
          },
          "400": {
            "description": "server could not understand the request due to invalid syntax"
          }
        }
      }
    }
  },
  "definitions": {
    "LoginSuccessResponseDefinition": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "LoginUserDefinition": {
      "type": "object",
      "required": [
        "Email",
        "Password"
      ],
      "properties": {
        "Email": {
          "type": "string"
        },
        "Password": {
          "type": "string"
        }
      }
    },
    "RegisterUserDefinition": {
      "type": "object",
      "required": [
        "Name",
        "Email",
        "Password"
      ],
      "properties": {
        "Email": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Password": {
          "type": "string"
        }
      }
    },
    "SuccessResponseDefinition": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
}
