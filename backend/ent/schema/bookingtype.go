package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Bookingtype holds the schema definition for the Bookingtype entity.
type Bookingtype struct {
	ent.Schema
}

// Fields of the Bookingtype.
func (Bookingtype) Fields() []ent.Field {
	return []ent.Field{
		field.String("BOOKTYPE_NAME").
			Unique().
			NotEmpty(),
	}
}

// Edges of the Bookingtype.
func (Bookingtype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("booktype", Booking.Type).
			StorageKey(edge.Column("BOOKINGTYPE_ID")),
	}
}
