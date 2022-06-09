// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"guide/ent/article"
	"guide/ent/document"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Article is the model entity for the Article schema.
type Article struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Number holds the value of the "number" field.
	Number float64 `json:"number,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArticleQuery when eager-loading is set.
	Edges             ArticleEdges `json:"edges"`
	document_articles *int
}

// ArticleEdges holds the relations/edges for other nodes in the graph.
type ArticleEdges struct {
	// Segments holds the value of the segments edge.
	Segments []*Segment `json:"segments,omitempty"`
	// Owner holds the value of the owner edge.
	Owner *Document `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SegmentsOrErr returns the Segments value or an error if the edge
// was not loaded in eager-loading.
func (e ArticleEdges) SegmentsOrErr() ([]*Segment, error) {
	if e.loadedTypes[0] {
		return e.Segments, nil
	}
	return nil, &NotLoadedError{edge: "segments"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArticleEdges) OwnerOrErr() (*Document, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: document.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Article) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case article.FieldNumber:
			values[i] = new(sql.NullFloat64)
		case article.FieldID:
			values[i] = new(sql.NullInt64)
		case article.ForeignKeys[0]: // document_articles
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Article", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Article fields.
func (a *Article) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case article.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case article.FieldNumber:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				a.Number = value.Float64
			}
		case article.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field document_articles", value)
			} else if value.Valid {
				a.document_articles = new(int)
				*a.document_articles = int(value.Int64)
			}
		}
	}
	return nil
}

// QuerySegments queries the "segments" edge of the Article entity.
func (a *Article) QuerySegments() *SegmentQuery {
	return (&ArticleClient{config: a.config}).QuerySegments(a)
}

// QueryOwner queries the "owner" edge of the Article entity.
func (a *Article) QueryOwner() *DocumentQuery {
	return (&ArticleClient{config: a.config}).QueryOwner(a)
}

// Update returns a builder for updating this Article.
// Note that you need to call Article.Unwrap() before calling this method if this Article
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Article) Update() *ArticleUpdateOne {
	return (&ArticleClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Article entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Article) Unwrap() *Article {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Article is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Article) String() string {
	var builder strings.Builder
	builder.WriteString("Article(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("number=")
	builder.WriteString(fmt.Sprintf("%v", a.Number))
	builder.WriteByte(')')
	return builder.String()
}

// Articles is a parsable slice of Article.
type Articles []*Article

func (a Articles) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
