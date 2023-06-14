package gqlCursor

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"gqlCursor/ent"
)

// CreateCar is the resolver for the createCar field.
func (r *carMutationResolver) CreateCar(ctx context.Context, obj *ent.CarMutation, input ent.CreateCarInput) (*ent.Car, error) {
	return ent.FromContext(ctx).Car.Create().SetInput(input).Save(ctx)
}

// UpdateCar is the resolver for the updateCar field.
func (r *carMutationResolver) UpdateCar(ctx context.Context, obj *ent.CarMutation, id int, input ent.UpdateCarInput) (*ent.Car, error) {
	return ent.FromContext(ctx).Car.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteCar is the resolver for the deleteCar field.
func (r *carMutationResolver) DeleteCar(ctx context.Context, obj *ent.CarMutation, id int) (*ent.Car, error) {
	return nil, r.client.Car.DeleteOneID(id).Exec(ctx)
}

// CarMutation returns CarMutationResolver implementation.
func (r *Resolver) CarMutation() CarMutationResolver { return &carMutationResolver{r} }

type carMutationResolver struct{ *Resolver }
