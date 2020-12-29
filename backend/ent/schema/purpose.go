package schema

import (
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent"
)
// Purpose holds the schema definition for the Purpose entity.
type Purpose struct {
	ent.Schema
}

// Fields of the Purpose.
func (Purpose) Fields() []ent.Field {
	return []ent.Field{
		field.String("PurposeName").
			Unique(),
			
	}
}

// Edges of the Purpose.
func (Purpose) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("preemption", Preemption.Type).StorageKey(edge.Column("PurposeID")),
	}
}

