package schema

import (
	"errors"
	"regexp"

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
		field.Time("RETURN_TIME"),
		field.Int("DAMAGED_POINT").Range(-1, 10),
		
		field.String("DAMAGED_POINTNAME").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[a-zA-Z]+$", s)
			if !match {
				return errors.New("จุดที่เสียหายเป็นภาษาอังกฤษเท่านั้น เช่น TopFront,BottomBack")
			}
			return nil
		}),
		field.String("LOST").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[a-zA-Z]+$", s)
			if !match {
				return errors.New("ถ้าหายให้พิมพ์ lost ถ้าไม่พิมพ์ no")
			}
			return nil
		}).MaxLen(5),
	}
}

// Edges of the Bookreturn.
func (Bookreturn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("return").Unique(),
		edge.From("location", Location.Type).Ref("returnfrom").Unique(),
		edge.From("mustreturn", Bookborrow.Type).Ref("borrowed").Unique(),
	}
}
