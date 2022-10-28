// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"fmt"
	"orginone/common/schema/astenantattrrole"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsTenantAttrRoleDelete is the builder for deleting a AsTenantAttrRole entity.
type AsTenantAttrRoleDelete struct {
	config
	hooks    []Hook
	mutation *AsTenantAttrRoleMutation
}

// Where appends a list predicates to the AsTenantAttrRoleDelete builder.
func (atard *AsTenantAttrRoleDelete) Where(ps ...predicate.AsTenantAttrRole) *AsTenantAttrRoleDelete {
	atard.mutation.Where(ps...)
	return atard
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atard *AsTenantAttrRoleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atard.hooks) == 0 {
		affected, err = atard.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsTenantAttrRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atard.mutation = mutation
			affected, err = atard.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atard.hooks) - 1; i >= 0; i-- {
			if atard.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = atard.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atard.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (atard *AsTenantAttrRoleDelete) ExecX(ctx context.Context) int {
	n, err := atard.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atard *AsTenantAttrRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: astenantattrrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: astenantattrrole.FieldID,
			},
		},
	}
	if ps := atard.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, atard.driver, _spec)
}

// AsTenantAttrRoleDeleteOne is the builder for deleting a single AsTenantAttrRole entity.
type AsTenantAttrRoleDeleteOne struct {
	atard *AsTenantAttrRoleDelete
}

// Exec executes the deletion query.
func (atardo *AsTenantAttrRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := atardo.atard.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{astenantattrrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atardo *AsTenantAttrRoleDeleteOne) ExecX(ctx context.Context) {
	atardo.atard.ExecX(ctx)
}