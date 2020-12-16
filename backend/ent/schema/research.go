package schema

import (
    "time"

    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

type Research struct {
    ent.Schema
}

func (Research) Fields() []ent.Field {
    return []ent.Field{
        field.String("DOC_NAME").NotEmpty(),
        field.Time("DATE").Default(time.Now),
    }
}

func (Research) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("register", User.Type).Ref("record").Unique(),
        edge.From("myDoc", Author.Type).Ref("owner").Unique(),
        edge.From("docType", ResearchType.Type).Ref("researchType").
            Unique(),
    }
}
