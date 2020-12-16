package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("BookName").
			NotEmpty().
			Unique(),
		field.String("Auther").
			NotEmpty(),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return nil
}
