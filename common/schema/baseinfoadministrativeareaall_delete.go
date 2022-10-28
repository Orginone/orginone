// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"fmt"
	"orginone/common/schema/baseinfoadministrativeareaall"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BaseinfoadministrativeareaallDelete is the builder for deleting a Baseinfoadministrativeareaall entity.
type BaseinfoadministrativeareaallDelete struct {
	config
	hooks    []Hook
	mutation *BaseinfoadministrativeareaallMutation
}

// Where appends a list predicates to the BaseinfoadministrativeareaallDelete builder.
func (bd *BaseinfoadministrativeareaallDelete) Where(ps ...predicate.Baseinfoadministrativeareaall) *BaseinfoadministrativeareaallDelete {
	bd.mutation.Where(ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BaseinfoadministrativeareaallDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bd.hooks) == 0 {
		affected, err = bd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaseinfoadministrativeareaallMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bd.mutation = mutation
			affected, err = bd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bd.hooks) - 1; i >= 0; i-- {
			if bd.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = bd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BaseinfoadministrativeareaallDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BaseinfoadministrativeareaallDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: baseinfoadministrativeareaall.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: baseinfoadministrativeareaall.FieldID,
			},
		},
	}
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
}

// BaseinfoadministrativeareaallDeleteOne is the builder for deleting a single Baseinfoadministrativeareaall entity.
type BaseinfoadministrativeareaallDeleteOne struct {
	bd *BaseinfoadministrativeareaallDelete
}

// Exec executes the deletion query.
func (bdo *BaseinfoadministrativeareaallDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{baseinfoadministrativeareaall.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BaseinfoadministrativeareaallDeleteOne) ExecX(ctx context.Context) {
	bdo.bd.ExecX(ctx)
}
