package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/schema/edge"
)

// Hobby holds the schema definition for the Hobby entity.
type Hobby struct {
	ent.Schema
}

// Fields of the Hobby.
func (Hobby) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
	}
}

// Edges of the Hobby.
func (Hobby) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owners", User.Type).
			Ref("hobbies"),
	}
}

func (Hobby) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
