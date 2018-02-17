package api

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/raunofreiberg/kyrene/server/model"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"todo": &graphql.Field{
			Type:        TodoType,
			Description: "return a todo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// Query via Todo ID and return a Todo's ID, content and is_completed status
				queryID, _ := params.Args["id"].(int)
				queriedTodo, err := QueryTodo(queryID)

				if err != nil {
					return nil, err
				}

				return queriedTodo, nil
			},
		},
		"todoList": &graphql.Field{
			Type:        graphql.NewList(TodoType),
			Description: "return all todos",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				jwt := params.Context.Value("jwt").(string)
				isAuthorized, err := IsAuthorized(jwt)

				if err != nil {
					return nil, err
				}

				if isAuthorized {
					queriedTodos, err := QueryTodos()

					if err != nil {
						return nil, err
					}

					return queriedTodos, nil
				}

				return nil, errors.New("Unauthorized")
			},
		},
		"getUserStatus": &graphql.Field{
			Type:        StatusType,
			Description: "Verify that the token being used is valid",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				jwt := params.Context.Value("jwt").(string)
				isAuthorized, err := IsAuthorized(jwt)

				if err != nil {
					return nil, err
				}

				return model.UserStatus{
					IsAuthorized: isAuthorized,
				}, nil
			},
		},
	},
})
