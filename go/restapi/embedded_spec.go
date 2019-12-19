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
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "api for Bellese fullstack challenge",
    "title": "Bellese challenge",
    "version": "0.1.0"
  },
  "host": "localhost:59828",
  "paths": {
    "/item": {
      "get": {
        "responses": {
          "200": {
            "description": "list all items",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/item"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "parameters": [
          {
            "description": "the item to create",
            "name": "item",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "create new item",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/item/{id}": {
      "get": {
        "parameters": [
          {
            "type": "number",
            "description": "id of the item to fetch",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "get an item",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "parameters": [
          {
            "type": "number",
            "description": "id of the item to fetch",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "202": {
            "description": "remove an item"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "parameters": [
          {
            "type": "number",
            "description": "id of the item to fetch",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "the item to create",
            "name": "item",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "update an item",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "item": {
      "required": [
        "name",
        "description",
        "picture"
      ],
      "properties": {
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "keyValueFields": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/keyValue"
          }
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "picture": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "keyValue": {
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "value": {
          "type": "string",
          "minLength": 1
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
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "api for Bellese fullstack challenge",
    "title": "Bellese challenge",
    "version": "0.1.0"
  },
  "host": "localhost:59828",
  "paths": {
    "/item": {
      "get": {
        "responses": {
          "200": {
            "description": "list all items",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/item"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "parameters": [
          {
            "description": "the item to create",
            "name": "item",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "create new item",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/item/{id}": {
      "get": {
        "parameters": [
          {
            "type": "number",
            "description": "id of the item to fetch",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "get an item",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "parameters": [
          {
            "type": "number",
            "description": "id of the item to fetch",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "202": {
            "description": "remove an item"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "parameters": [
          {
            "type": "number",
            "description": "id of the item to fetch",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "the item to create",
            "name": "item",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "update an item",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "item": {
      "required": [
        "name",
        "description",
        "picture"
      ],
      "properties": {
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "keyValueFields": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/keyValue"
          }
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "picture": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "keyValue": {
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "minLength": 1
        },
        "value": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
}`))
}
