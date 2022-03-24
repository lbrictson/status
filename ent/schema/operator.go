package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Operator holds the schema definition for the Operator entity.
type Operator struct {
	ent.Schema
}

// Fields of the Operator.
func (Operator) Fields() []ent.Field {
	return []ent.Field{
		field.String("Email").Unique(),
		field.String("HashedPassword"),
		field.String("Role"),
	}
}

// Edges of the Operator.
func (Operator) Edges() []ent.Edge {
	return nil
}
