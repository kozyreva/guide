package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Feature holds the schema definition for the Feature entity.
type Feature struct {
	ent.Schema
}

/*
func (Feature) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("product_id", "parameter_id"),
	}
}
*/

// Fields of the Feature.
func (Feature) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Int("product_id"),
		field.Int("parameter_id"),
	}
}

// Edges of the Feature.
func (Feature) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type).
			Required().
			Unique().
			Field("product_id"),
		edge.To("parameter", Parameter.Type).
			Required().
			Unique().
			Field("parameter_id"),
		edge.From("segments", Segment.Type).
			Ref("features"),
	}
}
