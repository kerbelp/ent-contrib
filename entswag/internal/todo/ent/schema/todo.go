package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent-contrib/entswag"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Annotations of the User.
func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entswag.Annotation{Model: true},
	}
}

// Fields returns todo fields.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			),
		field.Int("priority").
			Default(0),
		field.Text("text").
			NotEmpty(),
	}
}

// Edges returns todo edges.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Todo.Type).
			From("parent").
			Unique(),
	}
}
