// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (h *Hobby) Owners(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = h.NamedOwners(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = h.Edges.OwnersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = h.QueryOwners().All(ctx)
	}
	return result, err
}

func (u *User) Friends(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy []*UserOrder, where *UserWhereInput,
) (*UserConnection, error) {
	opts := []UserPaginateOption{
		WithUserOrder(orderBy),
		WithUserFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[0][alias]
	if nodes, err := u.NamedFriends(alias); err == nil || hasTotalCount {
		pager, err := newUserPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &UserConnection{Edges: []*UserEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryFriends().Paginate(ctx, after, first, before, last, opts...)
}

func (u *User) Hobbies(ctx context.Context) (result []*Hobby, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedHobbies(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.HobbiesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryHobbies().All(ctx)
	}
	return result, err
}
