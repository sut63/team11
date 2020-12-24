package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// Researchtype holds the schema definition for the Researchtype entity.
type Researchtype struct {
	ent.Schema
}

// Fields of the Researchtype.
func (Researchtype) Fields() []ent.Field {
	return []ent.Field{
        field.String("TYPE_NAME").
            NotEmpty().
            Unique(),
    }

}

// Edges of the Researchtype.
func (Researchtype) Edges() []ent.Edge {
	edge.To("researchType", Research.Type).StorageKey(edge.Column("TYPE_ID")),

}
