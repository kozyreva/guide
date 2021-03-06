// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"guide/ent/parameter"
	"guide/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ParameterDelete is the builder for deleting a Parameter entity.
type ParameterDelete struct {
	config
	hooks    []Hook
	mutation *ParameterMutation
}

// Where appends a list predicates to the ParameterDelete builder.
func (pd *ParameterDelete) Where(ps ...predicate.Parameter) *ParameterDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *ParameterDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pd.hooks) == 0 {
		affected, err = pd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ParameterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pd.mutation = mutation
			affected, err = pd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pd.hooks) - 1; i >= 0; i-- {
			if pd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *ParameterDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *ParameterDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: parameter.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: parameter.FieldID,
			},
		},
	}
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
}

// ParameterDeleteOne is the builder for deleting a single Parameter entity.
type ParameterDeleteOne struct {
	pd *ParameterDelete
}

// Exec executes the deletion query.
func (pdo *ParameterDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{parameter.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *ParameterDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}
