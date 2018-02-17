package api

import (
	"github.com/graphql-go/graphql"
)

var TodoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"isCompleted": &graphql.Field{
			Type: graphql.Boolean,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var StatusType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Status",
	Fields: graphql.Fields{
		"isAuthorized": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
