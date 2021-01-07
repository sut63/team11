// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/users": {
            "get": {
                "description": "list user entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List user entities",
                "operationId": "list-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create user",
                "operationId": "create-user",
                "parameters": [
                    {
                        "description": "User entity",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "get user by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a user entity by ID",
                "operationId": "get-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "put": {
                "description": "update user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a user entity by ID",
                "operationId": "update-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User entity",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "delete": {
                "description": "get user by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a user entity by ID",
                "operationId": "delete-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ent.Author": {
            "type": "object",
            "properties": {
                "Name": {
                    "description": "Name holds the value of the \"Name\" field.",
                    "type": "string"
                },
                "Position": {
                    "description": "Position holds the value of the \"Position\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the AuthorQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.AuthorEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.AuthorEdges": {
            "type": "object",
            "properties": {
                "owner": {
                    "description": "Owner holds the value of the owner edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Research"
                    }
                },
                "writer": {
                    "description": "Writer holds the value of the writer edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Book"
                    }
                }
            }
        },
        "ent.Book": {
            "type": "object",
            "properties": {
                "BookName": {
                    "description": "BookName holds the value of the \"BookName\" field.",
                    "type": "string"
                },
                "author_ID": {
                    "type": "integer"
                },
                "category_id": {
                    "type": "integer"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the BookQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.BookEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "status_ID": {
                    "type": "integer"
                },
                "user_ID": {
                    "type": "integer"
                }
            }
        },
        "ent.BookEdges": {
            "type": "object",
            "properties": {
                "author": {
                    "description": "Author holds the value of the author edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Author"
                },
                "booklist": {
                    "description": "Booklist holds the value of the booklist edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Bookborrow"
                    }
                },
                "category": {
                    "description": "Category holds the value of the category edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Category"
                },
                "status": {
                    "description": "Status holds the value of the Status edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Status"
                },
                "user": {
                    "description": "User holds the value of the user edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.User"
                }
            }
        },
        "ent.Bookborrow": {
            "type": "object",
            "properties": {
                "BORROW_DATE": {
                    "description": "BORROWDATE holds the value of the \"BORROW_DATE\" field.",
                    "type": "string"
                },
                "book_ID": {
                    "type": "integer"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the BookborrowQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.BookborrowEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "servicepoint_ID": {
                    "type": "integer"
                },
                "user_ID": {
                    "type": "integer"
                }
            }
        },
        "ent.BookborrowEdges": {
            "type": "object",
            "properties": {
                "book": {
                    "description": "BOOK holds the value of the BOOK edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Book"
                },
                "servicepoint": {
                    "description": "SERVICEPOINT holds the value of the SERVICEPOINT edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.ServicePoint"
                },
                "user": {
                    "description": "USER holds the value of the USER edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.User"
                }
            }
        },
        "ent.Booking": {
            "type": "object",
            "properties": {
                "BOOKING_DATE": {
                    "description": "BOOKINGDATE holds the value of the \"BOOKING_DATE\" field.",
                    "type": "string"
                },
                "TIME_LEFT": {
                    "description": "TIMELEFT holds the value of the \"TIME_LEFT\" field.",
                    "type": "string"
                },
                "client_ID": {
                    "type": "integer"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the BookingQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.BookingEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "servicepoint_ID": {
                    "type": "integer"
                },
                "user_ID": {
                    "type": "integer"
                }
            }
        },
        "ent.BookingEdges": {
            "type": "object",
            "properties": {
                "getservice": {
                    "description": "Getservice holds the value of the getservice edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.ServicePoint"
                },
                "usedby": {
                    "description": "Usedby holds the value of the usedby edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.User"
                },
                "using": {
                    "description": "Using holds the value of the using edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.ClientEntity"
                }
            }
        },
        "ent.Category": {
            "type": "object",
            "properties": {
                "CategoryName": {
                    "description": "CategoryName holds the value of the \"CategoryName\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the CategoryQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.CategoryEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.CategoryEdges": {
            "type": "object",
            "properties": {
                "catof": {
                    "description": "Catof holds the value of the catof edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Book"
                    }
                }
            }
        },
        "ent.ClientEntity": {
            "type": "object",
            "properties": {
                "CLIENT_NAME": {
                    "description": "CLIENTNAME holds the value of the \"CLIENT_NAME\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ClientEntityQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.ClientEntityEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "status_ID": {
                    "type": "integer"
                }
            }
        },
        "ent.ClientEntityEdges": {
            "type": "object",
            "properties": {
                "booked": {
                    "description": "Booked holds the value of the booked edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Booking"
                    }
                },
                "state": {
                    "description": "State holds the value of the state edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Status"
                }
            }
        },
        "ent.Preemption": {
            "type": "object",
            "properties": {
                "PreemptTime": {
                    "description": "PreemptTime holds the value of the \"PreemptTime\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the PreemptionQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.PreemptionEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "purposeID": {
                    "type": "integer"
                },
                "roomID": {
                    "type": "integer"
                },
                "user_ID": {
                    "type": "integer"
                }
            }
        },
        "ent.PreemptionEdges": {
            "type": "object",
            "properties": {
                "purposeID": {
                    "description": "PurposeID holds the value of the PurposeID edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Purpose"
                },
                "roomID": {
                    "description": "RoomID holds the value of the RoomID edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Roominfo"
                },
                "userID": {
                    "description": "UserID holds the value of the User_ID edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.User"
                }
            }
        },
        "ent.Purpose": {
            "type": "object",
            "properties": {
                "PurposeName": {
                    "description": "PurposeName holds the value of the \"PurposeName\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the PurposeQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.PurposeEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.PurposeEdges": {
            "type": "object",
            "properties": {
                "preemption": {
                    "description": "Preemption holds the value of the preemption edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Preemption"
                    }
                }
            }
        },
        "ent.Research": {
            "type": "object",
            "properties": {
                "DATE": {
                    "description": "DATE holds the value of the \"DATE\" field.",
                    "type": "string"
                },
                "DOC_NAME": {
                    "description": "DOCNAME holds the value of the \"DOC_NAME\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ResearchQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.ResearchEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "owner_ID": {
                    "type": "integer"
                },
                "type_ID": {
                    "type": "integer"
                },
                "user_ID": {
                    "type": "integer"
                }
            }
        },
        "ent.ResearchEdges": {
            "type": "object",
            "properties": {
                "docType": {
                    "description": "DocType holds the value of the docType edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Researchtype"
                },
                "myDoc": {
                    "description": "MyDoc holds the value of the myDoc edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Author"
                },
                "register": {
                    "description": "Register holds the value of the register edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.User"
                }
            }
        },
        "ent.Researchtype": {
            "type": "object",
            "properties": {
                "TYPE_NAME": {
                    "description": "TYPENAME holds the value of the \"TYPE_NAME\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ResearchtypeQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.ResearchtypeEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.ResearchtypeEdges": {
            "type": "object",
            "properties": {
                "researchType": {
                    "description": "ResearchType holds the value of the researchType edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Research"
                    }
                }
            }
        },
        "ent.Role": {
            "type": "object",
            "properties": {
                "ROLE_NAME": {
                    "description": "ROLENAME holds the value of the \"ROLE_NAME\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the RoleQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.RoleEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.RoleEdges": {
            "type": "object",
            "properties": {
                "role": {
                    "description": "Role holds the value of the role edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.User"
                    }
                }
            }
        },
        "ent.Roominfo": {
            "type": "object",
            "properties": {
                "RoomID": {
                    "description": "RoomID holds the value of the \"RoomID\" field.",
                    "type": "string"
                },
                "RoomNo": {
                    "description": "RoomNo holds the value of the \"RoomNo\" field.",
                    "type": "string"
                },
                "RoomStatus": {
                    "description": "RoomStatus holds the value of the \"RoomStatus\" field.",
                    "type": "string"
                },
                "RoomTime": {
                    "description": "RoomTime holds the value of the \"RoomTime\" field.",
                    "type": "string"
                },
                "RoomType": {
                    "description": "RoomType holds the value of the \"RoomType\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the RoominfoQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.RoominfoEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.RoominfoEdges": {
            "type": "object",
            "properties": {
                "preemption": {
                    "description": "Preemption holds the value of the preemption edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Preemption"
                    }
                }
            }
        },
        "ent.ServicePoint": {
            "type": "object",
            "properties": {
                "BUILDING_NAME": {
                    "description": "BUILDINGNAME holds the value of the \"BUILDING_NAME\" field.",
                    "type": "string"
                },
                "COUNTER_NUMBER": {
                    "description": "COUNTERNUMBER holds the value of the \"COUNTER_NUMBER\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the ServicePointQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.ServicePointEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.ServicePointEdges": {
            "type": "object",
            "properties": {
                "from": {
                    "description": "From holds the value of the from edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Bookborrow"
                    }
                },
                "servicepoint": {
                    "description": "Servicepoint holds the value of the servicepoint edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Booking"
                    }
                }
            }
        },
        "ent.Status": {
            "type": "object",
            "properties": {
                "STATUS_NAME": {
                    "description": "STATUSNAME holds the value of the \"STATUS_NAME\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the StatusQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.StatusEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.StatusEdges": {
            "type": "object",
            "properties": {
                "status": {
                    "description": "Status holds the value of the status edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.ClientEntity"
                    }
                },
                "statusofbook": {
                    "description": "Statusofbook holds the value of the statusofbook edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Book"
                    }
                }
            }
        },
        "ent.User": {
            "type": "object",
            "properties": {
                "PASSWORD": {
                    "description": "PASSWORD holds the value of the \"PASSWORD\" field.",
                    "type": "string"
                },
                "USER_EMAIL": {
                    "description": "USEREMAIL holds the value of the \"USER_EMAIL\" field.",
                    "type": "string"
                },
                "USER_NAME": {
                    "description": "USERNAME holds the value of the \"USER_NAME\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the UserQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.UserEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "role_ID": {
                    "type": "integer"
                }
            }
        },
        "ent.UserEdges": {
            "type": "object",
            "properties": {
                "addby": {
                    "description": "Addby holds the value of the addby edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Book"
                    }
                },
                "booking": {
                    "description": "Booking holds the value of the booking edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Booking"
                    }
                },
                "borrow": {
                    "description": "Borrow holds the value of the borrow edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Bookborrow"
                    }
                },
                "position": {
                    "description": "Position holds the value of the position edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Role"
                },
                "preemption": {
                    "description": "Preemption holds the value of the preemption edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Preemption"
                    }
                },
                "record": {
                    "description": "Record holds the value of the record edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Research"
                    }
                }
            }
        },
        "gin.H": {
            "type": "object",
            "additionalProperties": true
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        },
        "OAuth2AccessCode": {
            "type": "oauth2",
            "flow": "accessCode",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information"
            }
        },
        "OAuth2Application": {
            "type": "oauth2",
            "flow": "application",
            "authorizationUrl": "",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Implicit": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "authorizationUrl": "",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "read": " Grants read access",
                "write": " Grants write access"
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "SUT SA Example API Playlist Vidoe",
	Description: "This is a sample server for SUT SE 2563",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
