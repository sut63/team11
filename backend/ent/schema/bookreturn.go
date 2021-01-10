package schema

import (
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
		field.Time("DEADLINE"),
    }
}

// Edges of the Bookreturn.
func (Bookreturn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user",User.Type).Ref("return").Unique(),
		edge.From("location",Location.Type).Ref("returnfrom").Unique(),
		edge.From("mustreturn", Bookborrow.Type).Ref("borrowed").Unique(),	
	}
}
