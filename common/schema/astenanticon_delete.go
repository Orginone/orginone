// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"fmt"
	"orginone/common/schema/astenanticon"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsTenantIconDelete is the builder for deleting a AsTenantIcon entity.
type AsTenantIconDelete struct {
	config
	hooks    []Hook
	mutation *AsTenantIconMutation
}

// Where appends a list predicates to the AsTenantIconDelete builder.
func (atid *AsTenantIconDelete) Where(ps ...predicate.AsTenantIcon) *AsTenantIconDelete {
	atid.mutation.Where(ps...)
	return atid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atid *AsTenantIconDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atid.hooks) == 0 {
		affected, err = atid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsTenantIconMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atid.mutation = mutation
			affected, err = atid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atid.hooks) - 1; i >= 0; i-- {
			if atid.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = atid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (atid *AsTenantIconDelete) ExecX(ctx context.Context) int {
	n, err := atid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atid *AsTenantIconDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: astenanticon.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: astenanticon.FieldID,
			},
		},
	}
	if ps := atid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, atid.driver, _spec)
}

// AsTenantIconDeleteOne is the builder for deleting a single AsTenantIcon entity.
type AsTenantIconDeleteOne struct {
	atid *AsTenantIconDelete
}

// Exec executes the deletion query.
func (atido *AsTenantIconDeleteOne) Exec(ctx context.Context) error {
	n, err := atido.atid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{astenanticon.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atido *AsTenantIconDeleteOne) ExecX(ctx context.Context) {
	atido.atid.ExecX(ctx)
}