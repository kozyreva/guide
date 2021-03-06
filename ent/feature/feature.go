// Code generated by ent, DO NOT EDIT.

package feature

import (
	"time"
)

const (
	// Label holds the string label denoting the feature type in the database.
	Label = "feature"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldProductID holds the string denoting the product_id field in the database.
	FieldProductID = "product_id"
	// FieldParameterID holds the string denoting the parameter_id field in the database.
	FieldParameterID = "parameter_id"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// EdgeParameter holds the string denoting the parameter edge name in mutations.
	EdgeParameter = "parameter"
	// EdgeSegments holds the string denoting the segments edge name in mutations.
	EdgeSegments = "segments"
	// Table holds the table name of the feature in the database.
	Table = "features"
	// ProductTable is the table that holds the product relation/edge.
	ProductTable = "features"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "product_id"
	// ParameterTable is the table that holds the parameter relation/edge.
	ParameterTable = "features"
	// ParameterInverseTable is the table name for the Parameter entity.
	// It exists in this package in order to avoid circular dependency with the "parameter" package.
	ParameterInverseTable = "parameters"
	// ParameterColumn is the table column denoting the parameter relation/edge.
	ParameterColumn = "parameter_id"
	// SegmentsTable is the table that holds the segments relation/edge. The primary key declared below.
	SegmentsTable = "segment_features"
	// SegmentsInverseTable is the table name for the Segment entity.
	// It exists in this package in order to avoid circular dependency with the "segment" package.
	SegmentsInverseTable = "segments"
)

// Columns holds all SQL columns for feature fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldProductID,
	FieldParameterID,
}

var (
	// SegmentsPrimaryKey and SegmentsColumn2 are the table columns denoting the
	// primary key for the segments relation (M2M).
	SegmentsPrimaryKey = []string{"segment_id", "feature_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
