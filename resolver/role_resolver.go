package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/kivutar/chainz/model"
)

type roleResolver struct {
	role *model.Role
}

func (r *roleResolver) ID() graphql.ID {
	return graphql.ID(r.role.ID)
}

func (r *roleResolver) Name() *string {
	return &r.role.Name
}
