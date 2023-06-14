package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			NotEmpty(),
		field.Text("email").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}

}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
