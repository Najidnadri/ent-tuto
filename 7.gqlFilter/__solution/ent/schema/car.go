package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.Text("model").
			NotEmpty().
			Annotations(
				entgql.OrderField("MODEL"),
			),
		field.Text("plate_number").
			Unique().
			NotEmpty().
			Annotations(
				entgql.OrderField("PLATE_NUMBER"),
			),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			).
			Immutable(),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("cars").
			Unique(),
	}
}

func (Car) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
