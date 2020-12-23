package schema

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/jjg-akers/docker-sql-graphql/cmd/resolvers"
	"github.com/jjg-akers/docker-sql-graphql/cmd/types"
)

var Schema graphql.Schema

func init() {
	Query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
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
						Type:        graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// id, err := strconv.Atoi(p.Args["id"].(string))
					// if err != nil {
					// 	return nil, err
					// }
					return resolvers.GetPublication(p.Args["id"].(int)), nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: Query,
	})
	if err != nil {
		log.Panic(err)
	}

	Schema = schema
}
