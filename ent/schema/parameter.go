package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Parameter holds the schema definition for the Parameter entity.
type Parameter struct {
	ent.Schema
}

// Fields of the Parameter.
func (Parameter) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default(""),
	}
}

// Edges of the Parameter.
func (Parameter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("featured_products", Product.Type).
			Ref("featured_parameters").
			Through("features", Feature.Type),
	}
}
