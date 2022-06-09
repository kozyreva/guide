package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Segment holds the schema definition for the Segment entity.
type Segment struct {
	ent.Schema
}

// Fields of the Segment.
func (Segment) Fields() []ent.Field {
	return []ent.Field{
		field.Float("number").
			Positive(),
		field.Text("text"),
	}
}

// Edges of the Segment.
func (Segment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Article.Type).
			Ref("segments").
			Unique().
			Required(),
		edge.To("features", Feature.Type),
	}
}
