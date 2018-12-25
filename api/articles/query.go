package articles

import (
	"github.com/graphql-go/graphql"
)

// GetArticles get all articles
func GetArticles() (field *graphql.Field) {
	field = &graphql.Field{
		Type: graphql.NewList(ArticleType),
		Args: graphql.FieldConfigArgument{
			"first": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: -1,
			},
			"offset": &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: -1,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return getList(p)
		},
	}

	return
}

// GetArticleByID get article by ID
func GetArticleByID() (field *graphql.Field) {
	field = &graphql.Field{
		Type: ArticleType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return get(p)
		},
	}

	return
}
