package schema

import (
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
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("usedby", User.Type).
			Ref("booking").
			Unique(),
		edge.From("book", Bookingtype.Type).
			Ref("booktype").
			Unique(),
		edge.From("using", ClientEntity.Type).
			Ref("booked").
			Unique(),
	}
}
