package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Roominfo holds the schema definition for the Roominfo entity.
type Roominfo struct {
	ent.Schema
}

// Fields of the Roominfo.
func (Roominfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("RoomID").
			Unique(),
		field.String("RoomNo").NotEmpty(),
		field.String("RoomType").NotEmpty(),
		field.String("RoomTime").NotEmpty(),
		field.String("RoomStatus").NotEmpty(),
	}
}

// Edges of the Roominfo.
func (Roominfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("preemption", Preemption.Type).StorageKey(edge.Column("RoomID")),
	}
}
