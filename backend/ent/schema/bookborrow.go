package schema

import (
	"errors"
	"regexp"
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
		field.Int("DAY_OF_BORROW").Range(1,14),
		field.String("PICKUP").NotEmpty(),
		field.String("PHONE_NUMBER").Validate(func(s string) error {
			match, _ := regexp.MatchString("[0]\\d{9}", s)
			if !match {
				return errors.New("รูปแบบเบอร์โทรไม่ถูกต้อง")
			}
			return nil
		}),
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
		edge.From("STATUS", Status.Type).
			Ref("statusbookborrow").
			Unique(),

		edge.To("borrowed", Bookreturn.Type).
			StorageKey(edge.Column("CLIENT_ID")),
	}
}
