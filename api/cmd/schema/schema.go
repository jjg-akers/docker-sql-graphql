package schema

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/jjg-akers/go-docker-sql/resolvers"
	"github.com/jjg-akers/go-docker-sql/types"
)

var Schema graphql.Schema

func init() {
	Query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: &graphql.Fields{
			"publications": &graphql.Field{
				Type: graphql.NewList(types.PublicationType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolvers.GetPublications(), nil
				},
			},
			"publication": &graphql.Field{
				Type: types.PublicationType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the publication",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, err := strconv.Atoi(p.Args["id"].(string))
					if err != nil {
						return nil, err
					}
					return resolvers.GetPublication(id), nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: Query,
	})
	if err != nil {
		panic(err)
	}

	Schema = schema
}
