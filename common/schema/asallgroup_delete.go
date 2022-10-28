// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"fmt"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsAllGroupDelete is the builder for deleting a AsAllGroup entity.
type AsAllGroupDelete struct {
	config
	hooks    []Hook
	mutation *AsAllGroupMutation
}

// Where appends a list predicates to the AsAllGroupDelete builder.
func (aagd *AsAllGroupDelete) Where(ps ...predicate.AsAllGroup) *AsAllGroupDelete {
	aagd.mutation.Where(ps...)
	return aagd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aagd *AsAllGroupDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aagd.hooks) == 0 {
		affected, err = aagd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsAllGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aagd.mutation = mutation
			affected, err = aagd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aagd.hooks) - 1; i >= 0; i-- {
			if aagd.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = aagd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aagd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aagd *AsAllGroupDelete) ExecX(ctx context.Context) int {
	n, err := aagd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aagd *AsAllGroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: asallgroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asallgroup.FieldID,
			},
		},
	}
	if ps := aagd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, aagd.driver, _spec)
}

// AsAllGroupDeleteOne is the builder for deleting a single AsAllGroup entity.
type AsAllGroupDeleteOne struct {
	aagd *AsAllGroupDelete
}

// Exec executes the deletion query.
func (aagdo *AsAllGroupDeleteOne) Exec(ctx context.Context) error {
	n, err := aagdo.aagd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{asallgroup.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aagdo *AsAllGroupDeleteOne) ExecX(ctx context.Context) {
	aagdo.aagd.ExecX(ctx)
}
