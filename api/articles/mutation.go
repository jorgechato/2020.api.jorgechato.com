package articles

import (
	"time"

	"github.com/graphql-go/graphql"
)

var arguments = graphql.FieldConfigArgument{
	"title": {
		Type: graphql.NewNonNull(graphql.String),
	},
	"content": {
		Type: graphql.NewNonNull(graphql.String),
	},
	"thumbnail": {
		Type: graphql.String,
	},
	"public": {
		Type:         graphql.Boolean,
		DefaultValue: false,
	},
	"tags": {
		Type: graphql.NewList(graphql.String),
	},
	"published_at": {
		Type:         graphql.DateTime,
		DefaultValue: time.Now(),
	},
	"id": {
		Type:         graphql.ID,
		DefaultValue: "",
	},
}

// CreateArticle create an article
func CreateArticle() (field *graphql.Field) {
	field = &graphql.Field{
		Type: ArticleType,
		Args: arguments,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return create(p)
		},
	}

	return
}

// UpdateArticle update an article
func UpdateArticle() (field *graphql.Field) {
	field = &graphql.Field{
		Type: ArticleType,
		Args: arguments,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return update(p)
		},
	}

	return
}

// DeleteArticle delete an article
func DeleteArticle() (field *graphql.Field) {
	field = &graphql.Field{
		Type: ArticleType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return delete(p)
		},
	}

	return
}
