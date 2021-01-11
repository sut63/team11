package schema

import (
	"time"
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
		field.Time("RETURN_DATE"),
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
		edge.From("SERVICEPOINT", ServicePoint.Type).
			Ref("from").
			Unique(),
		
		edge.To("borrowed", Bookreturn.Type).
			StorageKey(edge.Column("CLIENT_ID")),

	}
}
