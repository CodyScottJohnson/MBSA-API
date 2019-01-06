package graphql

import (
	"github.com/graphql-go/graphql"

	"Hush_API/models"
)

var Person = graphql.NewObject(graphql.ObjectConfig{
	Name: "Person",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(models.Shouter).ID, nil
			},
		},
		"CreatedAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(models.Shouter).CreatedAt, nil
			},
		},
		"UpdatedAt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source.(models.Shouter).UpdatedAt, nil
			},
		},
		"FirstName": &graphql.Field{
			Type: graphql.String,
		},
		"LastName": &graphql.Field{
			Type: graphql.String,
		},
		"Primary_Email": &graphql.Field{
			Type: graphql.String,
		},
		"Primary_Address": &graphql.Field{
			Type: graphql.String,
		},
		"Primary_Phone": &graphql.Field{
			Type: graphql.String,
		},
		"Emails": &graphql.Field{
			Type: graphql.String,
		},
		"Addresses": &graphql.Field{
			Type: graphql.String,
		},
	},
})
