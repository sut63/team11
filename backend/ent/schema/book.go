package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
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
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).Ref("catof").Unique(),
		edge.From("author", Author.Type).Ref("writer").Unique(),
		edge.From("user", User.Type).Ref("addby").Unique(),

		edge.To("booklist", Bookborrow.Type).
			StorageKey(edge.Column("BOOK_ID")),
	}
}
