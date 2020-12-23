package types

import (
	"github.com/graphql-go/graphql"
)

var PublicationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Publication",
	Description: "The title of the Website",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The id of the publication",
		},
		"title": &graphql.Field{
			Type:        graphql.String,
			Description: "The title of the publication",
		},
		"uri": &graphql.Field{
			Type:        graphql.String,
			Description: "The uri of the publication",
		},
		"date": &graphql.Field{
			Type:        graphql.String,
			Description: "The date of the publication",
		},
	},
})
