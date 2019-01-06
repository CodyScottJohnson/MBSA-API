package graphql

import (
	"Hush_API/models"
	//"encoding/json"
	//"fmt"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getMessage": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				msg := "Cody"
				return msg, nil
			},
		},
		"recruits": &graphql.Field{
			Type: graphql.NewList(Recruit),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				recruit := []models.Recruit{}
				db.Preload("Person").Preload("Shouter.Emails").Where("archived = ?", false).Find(&recruit)
				return recruit, nil
			},
		},
	},
})
