// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"guide/ent/article"
	"guide/ent/feature"
	"guide/ent/predicate"
	"guide/ent/segment"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SegmentUpdate is the builder for updating Segment entities.
type SegmentUpdate struct {
	config
	hooks    []Hook
	mutation *SegmentMutation
}

// Where appends a list predicates to the SegmentUpdate builder.
func (su *SegmentUpdate) Where(ps ...predicate.Segment) *SegmentUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetNumber sets the "number" field.
func (su *SegmentUpdate) SetNumber(f float64) *SegmentUpdate {
	su.mutation.ResetNumber()
	su.mutation.SetNumber(f)
	return su
}

// AddNumber adds f to the "number" field.
func (su *SegmentUpdate) AddNumber(f float64) *SegmentUpdate {
	su.mutation.AddNumber(f)
	return su
}

// SetText sets the "text" field.
func (su *SegmentUpdate) SetText(s string) *SegmentUpdate {
	su.mutation.SetText(s)
	return su
}

// SetOwnerID sets the "owner" edge to the Article entity by ID.
func (su *SegmentUpdate) SetOwnerID(id int) *SegmentUpdate {
	su.mutation.SetOwnerID(id)
	return su
}

// SetOwner sets the "owner" edge to the Article entity.
func (su *SegmentUpdate) SetOwner(a *Article) *SegmentUpdate {
	return su.SetOwnerID(a.ID)
}

// AddFeatureIDs adds the "features" edge to the Feature entity by IDs.
func (su *SegmentUpdate) AddFeatureIDs(ids ...int) *SegmentUpdate {
	su.mutation.AddFeatureIDs(ids...)
	return su
}

// AddFeatures adds the "features" edges to the Feature entity.
func (su *SegmentUpdate) AddFeatures(f ...*Feature) *SegmentUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return su.AddFeatureIDs(ids...)
}

// Mutation returns the SegmentMutation object of the builder.
func (su *SegmentUpdate) Mutation() *SegmentMutation {
	return su.mutation
}

// ClearOwner clears the "owner" edge to the Article entity.
func (su *SegmentUpdate) ClearOwner() *SegmentUpdate {
	su.mutation.ClearOwner()
	return su
}

// ClearFeatures clears all "features" edges to the Feature entity.
func (su *SegmentUpdate) ClearFeatures() *SegmentUpdate {
	su.mutation.ClearFeatures()
	return su
}

// RemoveFeatureIDs removes the "features" edge to Feature entities by IDs.
func (su *SegmentUpdate) RemoveFeatureIDs(ids ...int) *SegmentUpdate {
	su.mutation.RemoveFeatureIDs(ids...)
	return su
}

// RemoveFeatures removes "features" edges to Feature entities.
func (su *SegmentUpdate) RemoveFeatures(f ...*Feature) *SegmentUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return su.RemoveFeatureIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SegmentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SegmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SegmentUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SegmentUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SegmentUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SegmentUpdate) check() error {
	if v, ok := su.mutation.Number(); ok {
		if err := segment.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Segment.number": %w`, err)}
		}
	}
	if _, ok := su.mutation.OwnerID(); su.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Segment.owner"`)
	}
	return nil
}

func (su *SegmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   segment.Table,
			Columns: segment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: segment.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Number(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: segment.FieldNumber,
		})
	}
	if value, ok := su.mutation.AddedNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: segment.FieldNumber,
		})
	}
	if value, ok := su.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: segment.FieldText,
		})
	}
	if su.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   segment.OwnerTable,
			Columns: []string{segment.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   segment.OwnerTable,
			Columns: []string{segment.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.FeaturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   segment.FeaturesTable,
			Columns: segment.FeaturesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: feature.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedFeaturesIDs(); len(nodes) > 0 && !su.mutation.FeaturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   segment.FeaturesTable,
			Columns: segment.FeaturesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: feature.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.FeaturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   segment.FeaturesTable,
			Columns: segment.FeaturesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: feature.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{segment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SegmentUpdateOne is the builder for updating a single Segment entity.
type SegmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SegmentMutation
}

// SetNumber sets the "number" field.
func (suo *SegmentUpdateOne) SetNumber(f float64) *SegmentUpdateOne {
	suo.mutation.ResetNumber()
	suo.mutation.SetNumber(f)
	return suo
}

// AddNumber adds f to the "number" field.
func (suo *SegmentUpdateOne) AddNumber(f float64) *SegmentUpdateOne {
	suo.mutation.AddNumber(f)
	return suo
}

// SetText sets the "text" field.
func (suo *SegmentUpdateOne) SetText(s string) *SegmentUpdateOne {
	suo.mutation.SetText(s)
	return suo
}

// SetOwnerID sets the "owner" edge to the Article entity by ID.
func (suo *SegmentUpdateOne) SetOwnerID(id int) *SegmentUpdateOne {
	suo.mutation.SetOwnerID(id)
	return suo
}

// SetOwner sets the "owner" edge to the Article entity.
func (suo *SegmentUpdateOne) SetOwner(a *Article) *SegmentUpdateOne {
	return suo.SetOwnerID(a.ID)
}

// AddFeatureIDs adds the "features" edge to the Feature entity by IDs.
func (suo *SegmentUpdateOne) AddFeatureIDs(ids ...int) *SegmentUpdateOne {
	suo.mutation.AddFeatureIDs(ids...)
	return suo
}

// AddFeatures adds the "features" edges to the Feature entity.
func (suo *SegmentUpdateOne) AddFeatures(f ...*Feature) *SegmentUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return suo.AddFeatureIDs(ids...)
}

// Mutation returns the SegmentMutation object of the builder.
func (suo *SegmentUpdateOne) Mutation() *SegmentMutation {
	return suo.mutation
}

// ClearOwner clears the "owner" edge to the Article entity.
func (suo *SegmentUpdateOne) ClearOwner() *SegmentUpdateOne {
	suo.mutation.ClearOwner()
	return suo
}

// ClearFeatures clears all "features" edges to the Feature entity.
func (suo *SegmentUpdateOne) ClearFeatures() *SegmentUpdateOne {
	suo.mutation.ClearFeatures()
	return suo
}

// RemoveFeatureIDs removes the "features" edge to Feature entities by IDs.
func (suo *SegmentUpdateOne) RemoveFeatureIDs(ids ...int) *SegmentUpdateOne {
	suo.mutation.RemoveFeatureIDs(ids...)
	return suo
}

// RemoveFeatures removes "features" edges to Feature entities.
func (suo *SegmentUpdateOne) RemoveFeatures(f ...*Feature) *SegmentUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return suo.RemoveFeatureIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SegmentUpdateOne) Select(field string, fields ...string) *SegmentUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Segment entity.
func (suo *SegmentUpdateOne) Save(ctx context.Context) (*Segment, error) {
	var (
		err  error
		node *Segment
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SegmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Segment)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SegmentMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SegmentUpdateOne) SaveX(ctx context.Context) *Segment {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SegmentUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SegmentUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SegmentUpdateOne) check() error {
	if v, ok := suo.mutation.Number(); ok {
		if err := segment.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Segment.number": %w`, err)}
		}
	}
	if _, ok := suo.mutation.OwnerID(); suo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Segment.owner"`)
	}
	return nil
}

func (suo *SegmentUpdateOne) sqlSave(ctx context.Context) (_node *Segment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   segment.Table,
			Columns: segment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: segment.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Segment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, segment.FieldID)
		for _, f := range fields {
			if !segment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != segment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Number(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: segment.FieldNumber,
		})
	}
	if value, ok := suo.mutation.AddedNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: segment.FieldNumber,
		})
	}
	if value, ok := suo.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: segment.FieldText,
		})
	}
	if suo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   segment.OwnerTable,
			Columns: []string{segment.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   segment.OwnerTable,
			Columns: []string{segment.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.FeaturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   segment.FeaturesTable,
			Columns: segment.FeaturesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: feature.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedFeaturesIDs(); len(nodes) > 0 && !suo.mutation.FeaturesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   segment.FeaturesTable,
			Columns: segment.FeaturesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: feature.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.FeaturesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   segment.FeaturesTable,
			Columns: segment.FeaturesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: feature.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Segment{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{segment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
