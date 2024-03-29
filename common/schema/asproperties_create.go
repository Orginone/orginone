// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"orginone/common/schema/asproperties"
	"orginone/common/schema/aspropertiesdistribution"
	"orginone/common/tools/date"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsPropertiesCreate is the builder for creating a AsProperties entity.
type AsPropertiesCreate struct {
	config
	mutation *AsPropertiesMutation
	hooks    []Hook
}

// SetPropertiesName sets the "properties_name" field.
func (apc *AsPropertiesCreate) SetPropertiesName(s string) *AsPropertiesCreate {
	apc.mutation.SetPropertiesName(s)
	return apc
}

// SetGroupID sets the "group_id" field.
func (apc *AsPropertiesCreate) SetGroupID(i int64) *AsPropertiesCreate {
	apc.mutation.SetGroupID(i)
	return apc
}

// SetIsDeleted sets the "is_deleted" field.
func (apc *AsPropertiesCreate) SetIsDeleted(i int64) *AsPropertiesCreate {
	apc.mutation.SetIsDeleted(i)
	return apc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (apc *AsPropertiesCreate) SetNillableIsDeleted(i *int64) *AsPropertiesCreate {
	if i != nil {
		apc.SetIsDeleted(*i)
	}
	return apc
}

// SetStatus sets the "status" field.
func (apc *AsPropertiesCreate) SetStatus(i int64) *AsPropertiesCreate {
	apc.mutation.SetStatus(i)
	return apc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (apc *AsPropertiesCreate) SetNillableStatus(i *int64) *AsPropertiesCreate {
	if i != nil {
		apc.SetStatus(*i)
	}
	return apc
}

// SetCreateUser sets the "create_user" field.
func (apc *AsPropertiesCreate) SetCreateUser(i int64) *AsPropertiesCreate {
	apc.mutation.SetCreateUser(i)
	return apc
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (apc *AsPropertiesCreate) SetNillableCreateUser(i *int64) *AsPropertiesCreate {
	if i != nil {
		apc.SetCreateUser(*i)
	}
	return apc
}

// SetUpdateUser sets the "update_user" field.
func (apc *AsPropertiesCreate) SetUpdateUser(i int64) *AsPropertiesCreate {
	apc.mutation.SetUpdateUser(i)
	return apc
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (apc *AsPropertiesCreate) SetNillableUpdateUser(i *int64) *AsPropertiesCreate {
	if i != nil {
		apc.SetUpdateUser(*i)
	}
	return apc
}

// SetCreateTime sets the "create_time" field.
func (apc *AsPropertiesCreate) SetCreateTime(dt date.DateTime) *AsPropertiesCreate {
	apc.mutation.SetCreateTime(dt)
	return apc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (apc *AsPropertiesCreate) SetNillableCreateTime(dt *date.DateTime) *AsPropertiesCreate {
	if dt != nil {
		apc.SetCreateTime(*dt)
	}
	return apc
}

// SetUpdateTime sets the "update_time" field.
func (apc *AsPropertiesCreate) SetUpdateTime(dt date.DateTime) *AsPropertiesCreate {
	apc.mutation.SetUpdateTime(dt)
	return apc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (apc *AsPropertiesCreate) SetNillableUpdateTime(dt *date.DateTime) *AsPropertiesCreate {
	if dt != nil {
		apc.SetUpdateTime(*dt)
	}
	return apc
}

// SetID sets the "id" field.
func (apc *AsPropertiesCreate) SetID(i int64) *AsPropertiesCreate {
	apc.mutation.SetID(i)
	return apc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (apc *AsPropertiesCreate) SetNillableID(i *int64) *AsPropertiesCreate {
	if i != nil {
		apc.SetID(*i)
	}
	return apc
}

// AddAllTenantIDs adds the "allTenants" edge to the AsPropertiesDistribution entity by IDs.
func (apc *AsPropertiesCreate) AddAllTenantIDs(ids ...int64) *AsPropertiesCreate {
	apc.mutation.AddAllTenantIDs(ids...)
	return apc
}

// AddAllTenants adds the "allTenants" edges to the AsPropertiesDistribution entity.
func (apc *AsPropertiesCreate) AddAllTenants(a ...*AsPropertiesDistribution) *AsPropertiesCreate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apc.AddAllTenantIDs(ids...)
}

// Mutation returns the AsPropertiesMutation object of the builder.
func (apc *AsPropertiesCreate) Mutation() *AsPropertiesMutation {
	return apc.mutation
}

// Save creates the AsProperties in the database.
func (apc *AsPropertiesCreate) Save(ctx context.Context) (*AsProperties, error) {
	var (
		err  error
		node *AsProperties
	)
	apc.defaults()
	if len(apc.hooks) == 0 {
		if err = apc.check(); err != nil {
			return nil, err
		}
		node, err = apc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsPropertiesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = apc.check(); err != nil {
				return nil, err
			}
			apc.mutation = mutation
			if node, err = apc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(apc.hooks) - 1; i >= 0; i-- {
			if apc.hooks[i] == nil {
				return nil, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = apc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, apc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (apc *AsPropertiesCreate) SaveX(ctx context.Context) *AsProperties {
	v, err := apc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apc *AsPropertiesCreate) Exec(ctx context.Context) error {
	_, err := apc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apc *AsPropertiesCreate) ExecX(ctx context.Context) {
	if err := apc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apc *AsPropertiesCreate) defaults() {
	if _, ok := apc.mutation.IsDeleted(); !ok {
		v := asproperties.DefaultIsDeleted
		apc.mutation.SetIsDeleted(v)
	}
	if _, ok := apc.mutation.Status(); !ok {
		v := asproperties.DefaultStatus
		apc.mutation.SetStatus(v)
	}
	if _, ok := apc.mutation.CreateTime(); !ok {
		v := asproperties.DefaultCreateTime()
		apc.mutation.SetCreateTime(v)
	}
	if _, ok := apc.mutation.UpdateTime(); !ok {
		v := asproperties.DefaultUpdateTime()
		apc.mutation.SetUpdateTime(v)
	}
	if _, ok := apc.mutation.ID(); !ok {
		v := asproperties.DefaultID()
		apc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apc *AsPropertiesCreate) check() error {
	if _, ok := apc.mutation.PropertiesName(); !ok {
		return &ValidationError{Name: "properties_name", err: errors.New(`schema: missing required field "AsProperties.properties_name"`)}
	}
	if _, ok := apc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`schema: missing required field "AsProperties.group_id"`)}
	}
	if _, ok := apc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`schema: missing required field "AsProperties.is_deleted"`)}
	}
	return nil
}

func (apc *AsPropertiesCreate) sqlSave(ctx context.Context) (*AsProperties, error) {
	_node, _spec := apc.createSpec()
	if err := sqlgraph.CreateNode(ctx, apc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (apc *AsPropertiesCreate) createSpec() (*AsProperties, *sqlgraph.CreateSpec) {
	var (
		_node = &AsProperties{config: apc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: asproperties.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asproperties.FieldID,
			},
		}
	)
	if id, ok := apc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := apc.mutation.PropertiesName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asproperties.FieldPropertiesName,
		})
		_node.PropertiesName = value
	}
	if value, ok := apc.mutation.GroupID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asproperties.FieldGroupID,
		})
		_node.GroupID = value
	}
	if value, ok := apc.mutation.IsDeleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asproperties.FieldIsDeleted,
		})
		_node.IsDeleted = value
	}
	if value, ok := apc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asproperties.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := apc.mutation.CreateUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asproperties.FieldCreateUser,
		})
		_node.CreateUser = value
	}
	if value, ok := apc.mutation.UpdateUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asproperties.FieldUpdateUser,
		})
		_node.UpdateUser = value
	}
	if value, ok := apc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asproperties.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := apc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asproperties.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if nodes := apc.mutation.AllTenantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asproperties.AllTenantsTable,
			Columns: []string{asproperties.AllTenantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: aspropertiesdistribution.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AsPropertiesCreateBulk is the builder for creating many AsProperties entities in bulk.
type AsPropertiesCreateBulk struct {
	config
	builders []*AsPropertiesCreate
}

// Save creates the AsProperties entities in the database.
func (apcb *AsPropertiesCreateBulk) Save(ctx context.Context) ([]*AsProperties, error) {
	specs := make([]*sqlgraph.CreateSpec, len(apcb.builders))
	nodes := make([]*AsProperties, len(apcb.builders))
	mutators := make([]Mutator, len(apcb.builders))
	for i := range apcb.builders {
		func(i int, root context.Context) {
			builder := apcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AsPropertiesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, apcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, apcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, apcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (apcb *AsPropertiesCreateBulk) SaveX(ctx context.Context) []*AsProperties {
	v, err := apcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apcb *AsPropertiesCreateBulk) Exec(ctx context.Context) error {
	_, err := apcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apcb *AsPropertiesCreateBulk) ExecX(ctx context.Context) {
	if err := apcb.Exec(ctx); err != nil {
		panic(err)
	}
}
