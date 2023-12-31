// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (c *Car) Owner(ctx context.Context) (*User, error) {
	result, err := c.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryOwner().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Cars(ctx context.Context) (result []*Car, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedCars(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.CarsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryCars().All(ctx)
	}
	return result, err
}
