package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)


// Author holds the schema definition for the Author entity.
type Author struct {
	ent.Schema
}

// Fields of the Author.
func (Author) Fields() []ent.Field {
	return []ent.Field{
        field.String("Name").
            NotEmpty().
            Unique(),
    }

}

// Edges of the Author.
func (Author) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("owner", Research.Type).StorageKey(edge.Column("OWNER_ID")),
        edge.To("writer", Book.Type).StorageKey(edge.Column("AUTHOR_ID")),
    }

}
