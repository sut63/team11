package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Bookborrow holds the schema definition for the Bookborrow entity.
type Bookborrow struct {
	ent.Schema
}

// Fields of the Bookborrow.
func (Bookborrow) Fields() []ent.Field {
	return []ent.Field{
		field.Time("BORROW_DATE").Default(time.Now),
    }
}

// Edges of the Bookborrow.
func (Bookborrow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("USER", User.Type).
			Ref("borrow").
			Unique(),
		edge.From("BOOK", Book.Type).
			Ref("booklist").
			Unique(),
		edge.From("SERVICEPOINT", Servicepoint.Type).
			Ref("from").
			Unique(),
	}
}
