package schema

import (
	"errors"
	"regexp"
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Research holds the schema definition for the Research entity.
type Research struct {
	ent.Schema
}

// Fields of the Research.
func (Research) Fields() []ent.Field {
	return []ent.Field{
		field.String("DOC_NAME").
			Validate(func(s string) error {
				match, _ := regexp.MatchString("^[0-9a-zA-Zก-๙]+$", s)
				if !match {
					return errors.New("มีอักษรพิเศษ")
				}
				return nil
			}).
			NotEmpty(),
		field.Time("DATE").Default(time.Now),
		field.Int("PAGE_NUMBER").Min(1),
		field.Int("YEAR_NUMBER").Range(1000, 2999),
	}
}

// Edges of the Research.
func (Research) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("register", User.Type).Ref("record").Unique(),
		edge.From("myDoc", Author.Type).Ref("owner").Unique(),
		edge.From("docType", Researchtype.Type).Ref("researchType").
			Unique(),
	}
}
