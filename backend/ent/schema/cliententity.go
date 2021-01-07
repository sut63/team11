package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// ClientEntity holds the schema definition for the ClientEntity entity.
type ClientEntity struct {
	ent.Schema
}

// Fields of the ClientEntity.
func (ClientEntity) Fields() []ent.Field {
	return []ent.Field{
		field.String("CLIENT_NAME").
			NotEmpty().
			Unique(),
	}
}

// Edges of the ClientEntity.
func (ClientEntity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("booked", Booking.Type).
			StorageKey(edge.Column("CLIENT_ID")),
		edge.From("state", Status.Type).
			Ref("status").
			Unique(),
	}
}
