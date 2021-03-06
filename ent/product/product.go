// Code generated by ent, DO NOT EDIT.

package product

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeFeaturedParameters holds the string denoting the featured_parameters edge name in mutations.
	EdgeFeaturedParameters = "featured_parameters"
	// EdgeFeatures holds the string denoting the features edge name in mutations.
	EdgeFeatures = "features"
	// Table holds the table name of the product in the database.
	Table = "products"
	// FeaturedParametersTable is the table that holds the featured_parameters relation/edge. The primary key declared below.
	FeaturedParametersTable = "features"
	// FeaturedParametersInverseTable is the table name for the Parameter entity.
	// It exists in this package in order to avoid circular dependency with the "parameter" package.
	FeaturedParametersInverseTable = "parameters"
	// FeaturesTable is the table that holds the features relation/edge.
	FeaturesTable = "features"
	// FeaturesInverseTable is the table name for the Feature entity.
	// It exists in this package in order to avoid circular dependency with the "feature" package.
	FeaturesInverseTable = "features"
	// FeaturesColumn is the table column denoting the features relation/edge.
	FeaturesColumn = "product_id"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// FeaturedParametersPrimaryKey and FeaturedParametersColumn2 are the table columns denoting the
	// primary key for the featured_parameters relation (M2M).
	FeaturedParametersPrimaryKey = []string{"product_id", "parameter_id"}
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
)
