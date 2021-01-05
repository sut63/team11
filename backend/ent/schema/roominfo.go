package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Preemption holds the schema definition for the Preemption entity.
type Preemption struct {
	ent.Schema
}

// Fields of the Preemption.
func (Preemption) Fields() []ent.Field {
	return []ent.Field{

		
		field.Time("PreemptTime").
			Default(time.Now),
	}
}

// Edges of the Preemption.
func (Preemption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("UserID", User.Type).
			Ref("preemption").Unique(),
		edge.From("PurposeID", Purpose.Type).
			Ref("preemption").Unique(),
		// edge.From("RoomID", Roominfo.Type).
		// 	Ref("preemption").Unique(),
	}
}

