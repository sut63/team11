package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.String("location_name").NotEmpty(),
   }
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("locations",Bookreturn.Type).StorageKey(edge.Column("location_id")),
    }
}
