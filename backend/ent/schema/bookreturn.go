package schema

import (
	"time"
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
<<<<<<< HEAD
		field.Time("return_deadline"),
    }
=======
		field.String("book_name").NotEmpty(),
	}
>>>>>>> c9cada08bc7f55755261c1e47704a41c6656a7a2
}

// Edges of the Bookreturn.
func (Bookreturn) Edges() []ent.Edge {
	return []ent.Edge{
<<<<<<< HEAD
		edge.From("user",User.Type).Ref("return").Unique(),
		edge.From("location",Location.Type).Ref("returnfrom").Unique(),
		edge.From("mustreturn", Bookborrow.Type).Ref("borrowed").Unique(),
=======
		edge.From("mustreturn", Bookborrow.Type).
			Ref("borrowed").
			Unique(),
>>>>>>> c9cada08bc7f55755261c1e47704a41c6656a7a2
	}
}
