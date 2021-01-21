package schema

import (
	"errors"
	"regexp"

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
		field.String("Barcode").Validate(func(s string) error {
			match, _ := regexp.MatchString("[0-9]\\d{9}", s)
			if !match {
				return errors.New("รูปแบบของ Barcode ไม่ถูกต้อง")
			}
			return nil
		}),
		field.Int("BookPage").Min(1),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).Ref("catof").Unique(),
		edge.From("author", Author.Type).Ref("writer").Unique(),
		edge.From("user", User.Type).Ref("addby").Unique(),
		edge.From("status", Status.Type).Ref("statusofbook").Unique(),

		edge.To("booklist", Bookborrow.Type).
			StorageKey(edge.Column("BOOK_ID")),
	}
}
