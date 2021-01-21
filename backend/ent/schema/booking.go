package schema

import (
	"errors"
	"regexp"
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.Time("BOOKING_DATE").
			Default(time.Now),
		field.Time("TIME_LEFT"),
		field.Int("USER_NUMBER").Range(1, 6),
		field.Int("BORROW_ITEM").Range(1, 6),
		field.String("PHONE_NUMBER").Validate(func(s string) error {
			match, _ := regexp.MatchString("[0][689]\\d{8}", s)
			if !match {
				return errors.New("รูปแบบของเบอร์มือถือไม่ถูกต้อง")
			}
			return nil
		}),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("usedby", User.Type).
			Ref("booking").
			Unique(),
		edge.From("getservice", ServicePoint.Type).
			Ref("servicepoint").
			Unique(),
		edge.From("using", ClientEntity.Type).
			Ref("booked").
			Unique(),
	}
}
