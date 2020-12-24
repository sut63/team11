package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)


// Bookreturn holds the schema definition for the Bookreturn entity.
type Bookreturn struct {
	ent.Schema
}

// Fields of the Bookreturn.
func (Bookreturn) Fields() []ent.Field {
	return []ent.Field{
		field.String("book_name").NotEmpty(),
    }
}

// Edges of the Bookreturn.
func (Bookreturn) Edges() []ent.Edge {
	edge.To("books",Bill.Type).StorageKey(edge.Column("book_id"))
}
