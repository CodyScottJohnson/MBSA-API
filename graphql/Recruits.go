package graphql

import (
	"fmt"
	"github.com/graphql-go/graphql"

	"Hush_API/models"
)

var Recruit = graphql.NewObject(graphql.ObjectConfig{
	Name: "Recruit",
	Fields: graphql.Fields{
		"ID": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Println("%#v\n",  p.Source.(models.Shouter))
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
		"Person": &graphql.Field{
			Type: Person,
		},
	},
})
