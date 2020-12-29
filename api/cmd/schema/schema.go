package schema

import (
	"errors"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/jjg-akers/docker-sql-graphql/cmd/auth"
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

					//make our Publication resolver authenticate the incoming quest - access inside context
					token := p.Context.Value("token").(string)

					// valicate token
					isValid, err := auth.ValidateToken(token)
					if err != nil {
						return nil, err
					}

					if !isValid {
						return nil, gqlerrors.FormatError(errors.New("Invalid Token"))
					}

					pub, err := resolvers.GetPublication(p.Args["id"].(int))
					if err != nil {
						return nil, err
					}

					return pub, nil
				},
			},
		},
	})

	Mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			// "createPublication": &graphql.Field{
			// 	Type:        types.PublicationType,
			// 	Description: "create a new publication",
			// 	Args: graphql.FieldConfigArgument{
			// 		"id": &graphql.ArgumentConfig{
			// 			Description: "id of new publication",
			// 			Type:        graphql.NewNonNull(graphql.Int),
			// 		},
			// 		"title": &graphql.ArgumentConfig{
			// 			Description: "the title of the new publication",
			// 			Type:        graphql.NewNonNull(graphql.String),
			// 		},
			// 		"uri": &graphql.ArgumentConfig{
			// 			Description: "the uri of the new publication",
			// 			Type:        graphql.NewNonNull(graphql.String),
			// 		},
			// 		"date_added": &graphql.ArgumentConfig{
			// 			Description: "the date added of the new publication",
			// 			Type:        graphql.NewNonNull(graphql.String),
			// 		},
			// 	},
			// 	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// 		// marshall and cast the arguments
			// 		id, ok := params.Args["id"].(int)
			// 		if !ok {
			// 			return nil, fmt.Errorf("could not convert id to int")
			// 		}
			// 		title, ok := params.Args["title"].(string)
			// 		if !ok {
			// 			return nil, fmt.Errorf("could not convert name to string")
			// 		}
			// 		uri, ok := params.Args["uri"].(string)
			// 		if !ok {
			// 			return nil, fmt.Errorf("could not convert uri to string")
			// 		}
			// 		dateAdded, ok := params.Args["date_added"].(string)
			// 		if !ok {
			// 			return nil, fmt.Errorf("could not convert date_added to string")
			// 		}

			// 		return resolvers.CreatePublication(id, title, uri, dateAdded), nil
			// 	},
			// },
			"createToken": &graphql.Field{
				Type:        graphql.String,
				Description: "creates a new JWT token ",
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Description: "username",
						Type:        graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Description: "password",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},

				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// mashall and cast the argument value
					username, _ := params.Args["username"].(string)
					password, _ := params.Args["password"].(string)

					token, err := auth.CreateToken(username, password)
					if err != nil {
						return nil, err
					}
					return token, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    Query,
		Mutation: Mutation,
	})
	if err != nil {
		log.Panic(err)
	}

	Schema = schema
}
