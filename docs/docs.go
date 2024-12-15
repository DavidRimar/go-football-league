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
        "/api/fixtures/{fixtureId}": {
            "put": {
                "description": "Update the fixture's status and scores by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fixtures"
                ],
                "summary": "Update Fixture",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Fixture ID",
                        "name": "fixtureId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Fixture Update Payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.UpdateFixtureDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Fixture updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request body or missing fixture ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Fixture not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/fixtures/{gameweekId}": {
            "get": {
                "description": "Get Fixtures by a Gameweek Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fixtures"
                ],
                "summary": "Get Fixtures by Gameweek",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Gameweek",
                        "name": "gameweekId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Fixture"
                            }
                        }
                    }
                }
            }
        },
        "/api/standings": {
            "get": {
                "description": "Retrieve team statistics for all teams",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Standings"
                ],
                "summary": "Get standings",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.GetTeamStatisticsDTO"
                            }
                        }
                    }
                }
            }
        },
        "/api/teams": {
            "get": {
                "description": "Retrieve details of all teams",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Get all teams",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Team"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.GetTeamStatisticsDTO": {
            "type": "object",
            "properties": {
                "draws": {
                    "type": "integer"
                },
                "gamesPlayed": {
                    "type": "integer"
                },
                "goalDifference": {
                    "type": "integer"
                },
                "goalsConceded": {
                    "type": "integer"
                },
                "goalsScored": {
                    "type": "integer"
                },
                "losses": {
                    "type": "integer"
                },
                "points": {
                    "type": "integer"
                },
                "team": {
                    "type": "string"
                },
                "wins": {
                    "type": "integer"
                }
            }
        },
        "dtos.UpdateFixtureDTO": {
            "type": "object",
            "properties": {
                "awayScore": {
                    "type": "integer"
                },
                "homeScore": {
                    "type": "integer"
                },
                "status": {
                    "$ref": "#/definitions/models.FixtureStatus"
                }
            }
        },
        "models.Fixture": {
            "type": "object",
            "properties": {
                "awayScore": {
                    "type": "integer"
                },
                "awayTeam": {
                    "type": "string"
                },
                "awayTeamId": {
                    "description": "References Team.ID",
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "gameweekId": {
                    "type": "integer"
                },
                "homeScore": {
                    "type": "integer"
                },
                "homeTeam": {
                    "type": "string"
                },
                "homeTeamId": {
                    "description": "References Team.ID",
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/models.FixtureStatus"
                }
            }
        },
        "models.FixtureStatus": {
            "type": "string",
            "enum": [
                "Played",
                "Live",
                "Upcoming"
            ],
            "x-enum-varnames": [
                "StatusPlayed",
                "StatusLive",
                "StatusUpcoming"
            ]
        },
        "models.Team": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "stadium": {
                    "type": "string"
                },
                "stadiumCapacity": {
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Football League API",
	Description:      "This is the Football League API documentation for the Football League service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
