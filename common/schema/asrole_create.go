// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"orginone/common/schema/asjob"
	"orginone/common/schema/asmenu"
	"orginone/common/schema/asrole"
	"orginone/common/schema/astenantattrrole"
	"orginone/common/schema/asuser"
	"orginone/common/tools/date"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsRoleCreate is the builder for creating a AsRole entity.
type AsRoleCreate struct {
	config
	mutation *AsRoleMutation
	hooks    []Hook
}

// SetSort sets the "sort" field.
func (arc *AsRoleCreate) SetSort(i int64) *AsRoleCreate {
	arc.mutation.SetSort(i)
	return arc
}

// SetRoleAlias sets the "role_alias" field.
func (arc *AsRoleCreate) SetRoleAlias(s string) *AsRoleCreate {
	arc.mutation.SetRoleAlias(s)
	return arc
}

// SetNillableRoleAlias sets the "role_alias" field if the given value is not nil.
func (arc *AsRoleCreate) SetNillableRoleAlias(s *string) *AsRoleCreate {
	if s != nil {
		arc.SetRoleAlias(*s)
	}
	return arc
}

// SetRoleName sets the "role_name" field.
func (arc *AsRoleCreate) SetRoleName(s string) *AsRoleCreate {
	arc.mutation.SetRoleName(s)
	return arc
}

// SetNillableRoleName sets the "role_name" field if the given value is not nil.
func (arc *AsRoleCreate) SetNillableRoleName(s *string) *AsRoleCreate {
	if s != nil {
		arc.SetRoleName(*s)
	}
	return arc
}

// SetRoleDescription sets the "role_description" field.
func (arc *AsRoleCreate) SetRoleDescription(s string) *AsRoleCreate {
	arc.mutation.SetRoleDescription(s)
	return arc
}

// SetNillableRoleDescription sets the "role_description" field if the given value is not nil.
func (arc *AsRoleCreate) SetNillableRoleDescription(s *string) *AsRoleCreate {
	if s != nil {
		arc.SetRoleDescription(*s)
	}
	return arc
}

// SetIsDeleted sets the "is_deleted" field.
func (arc *AsRoleCreate) SetIsDeleted(i int64) *AsRoleCreate {
	arc.mutation.SetIsDeleted(i)
	return arc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (arc *AsRoleCreate) SetNillableIsDeleted(i *int64) *AsRoleCreate {
	if i != nil {
		arc.SetIsDeleted(*i)
	}
	return arc
}

// SetStatus sets the "status" field.
func (arc *AsRoleCreate) SetStatus(i int64) *AsRoleCreate {
	arc.mutation.SetStatus(i)
	return arc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (arc *AsRoleCreate) SetNillableStatus(i *int64) *AsRoleCreate {
	if i != nil {
		arc.SetStatus(*i)
	}
	return arc
}

// SetCreateUser sets the "create_user" field.
func (arc *AsRoleCreate) SetCreateUser(i int64) *AsRoleCreate {
	arc.mutation.SetCreateUser(i)
	return arc
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (arc *AsRoleCreate) SetNillableCreateUser(i *int64) *AsRoleCreate {
	if i != nil {
		arc.SetCreateUser(*i)
	}
	return arc
}

// SetUpdateUser sets the "update_user" field.
func (arc *AsRoleCreate) SetUpdateUser(i int64) *AsRoleCreate {
	arc.mutation.SetUpdateUser(i)
	return arc
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (arc *AsRoleCreate) SetNillableUpdateUser(i *int64) *AsRoleCreate {
	if i != nil {
		arc.SetUpdateUser(*i)
	}
	return arc
}

// SetCreateTime sets the "create_time" field.
func (arc *AsRoleCreate) SetCreateTime(dt date.DateTime) *AsRoleCreate {
	arc.mutation.SetCreateTime(dt)
	return arc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (arc *AsRoleCreate) SetNillableCreateTime(dt *date.DateTime) *AsRoleCreate {
	if dt != nil {
		arc.SetCreateTime(*dt)
	}
	return arc
}

// SetUpdateTime sets the "update_time" field.
func (arc *AsRoleCreate) SetUpdateTime(dt date.DateTime) *AsRoleCreate {
	arc.mutation.SetUpdateTime(dt)
	return arc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (arc *AsRoleCreate) SetNillableUpdateTime(dt *date.DateTime) *AsRoleCreate {
	if dt != nil {
		arc.SetUpdateTime(*dt)
	}
	return arc
}

// SetID sets the "id" field.
func (arc *AsRoleCreate) SetID(i int64) *AsRoleCreate {
	arc.mutation.SetID(i)
	return arc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (arc *AsRoleCreate) SetNillableID(i *int64) *AsRoleCreate {
	if i != nil {
		arc.SetID(*i)
	}
	return arc
}

// AddUserIDs adds the "users" edge to the AsUser entity by IDs.
func (arc *AsRoleCreate) AddUserIDs(ids ...int64) *AsRoleCreate {
	arc.mutation.AddUserIDs(ids...)
	return arc
}

// AddUsers adds the "users" edges to the AsUser entity.
func (arc *AsRoleCreate) AddUsers(a ...*AsUser) *AsRoleCreate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return arc.AddUserIDs(ids...)
}

// AddJobIDs adds the "jobs" edge to the AsJob entity by IDs.
func (arc *AsRoleCreate) AddJobIDs(ids ...int64) *AsRoleCreate {
	arc.mutation.AddJobIDs(ids...)
	return arc
}

// AddJobs adds the "jobs" edges to the AsJob entity.
func (arc *AsRoleCreate) AddJobs(a ...*AsJob) *AsRoleCreate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return arc.AddJobIDs(ids...)
}

// AddMenuIDs adds the "menus" edge to the AsMenu entity by IDs.
func (arc *AsRoleCreate) AddMenuIDs(ids ...int64) *AsRoleCreate {
	arc.mutation.AddMenuIDs(ids...)
	return arc
}

// AddMenus adds the "menus" edges to the AsMenu entity.
func (arc *AsRoleCreate) AddMenus(a ...*AsMenu) *AsRoleCreate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return arc.AddMenuIDs(ids...)
}

// AddAttrRoleIDs adds the "attrRoles" edge to the AsTenantAttrRole entity by IDs.
func (arc *AsRoleCreate) AddAttrRoleIDs(ids ...int64) *AsRoleCreate {
	arc.mutation.AddAttrRoleIDs(ids...)
	return arc
}

// AddAttrRoles adds the "attrRoles" edges to the AsTenantAttrRole entity.
func (arc *AsRoleCreate) AddAttrRoles(a ...*AsTenantAttrRole) *AsRoleCreate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return arc.AddAttrRoleIDs(ids...)
}

// Mutation returns the AsRoleMutation object of the builder.
func (arc *AsRoleCreate) Mutation() *AsRoleMutation {
	return arc.mutation
}

// Save creates the AsRole in the database.
func (arc *AsRoleCreate) Save(ctx context.Context) (*AsRole, error) {
	var (
		err  error
		node *AsRole
	)
	arc.defaults()
	if len(arc.hooks) == 0 {
		if err = arc.check(); err != nil {
			return nil, err
		}
		node, err = arc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = arc.check(); err != nil {
				return nil, err
			}
			arc.mutation = mutation
			if node, err = arc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(arc.hooks) - 1; i >= 0; i-- {
			if arc.hooks[i] == nil {
				return nil, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = arc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, arc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (arc *AsRoleCreate) SaveX(ctx context.Context) *AsRole {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arc *AsRoleCreate) Exec(ctx context.Context) error {
	_, err := arc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arc *AsRoleCreate) ExecX(ctx context.Context) {
	if err := arc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (arc *AsRoleCreate) defaults() {
	if _, ok := arc.mutation.IsDeleted(); !ok {
		v := asrole.DefaultIsDeleted
		arc.mutation.SetIsDeleted(v)
	}
	if _, ok := arc.mutation.Status(); !ok {
		v := asrole.DefaultStatus
		arc.mutation.SetStatus(v)
	}
	if _, ok := arc.mutation.CreateTime(); !ok {
		v := asrole.DefaultCreateTime()
		arc.mutation.SetCreateTime(v)
	}
	if _, ok := arc.mutation.UpdateTime(); !ok {
		v := asrole.DefaultUpdateTime()
		arc.mutation.SetUpdateTime(v)
	}
	if _, ok := arc.mutation.ID(); !ok {
		v := asrole.DefaultID()
		arc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arc *AsRoleCreate) check() error {
	if _, ok := arc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`schema: missing required field "AsRole.sort"`)}
	}
	if _, ok := arc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`schema: missing required field "AsRole.is_deleted"`)}
	}
	return nil
}

func (arc *AsRoleCreate) sqlSave(ctx context.Context) (*AsRole, error) {
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
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

func (arc *AsRoleCreate) createSpec() (*AsRole, *sqlgraph.CreateSpec) {
	var (
		_node = &AsRole{config: arc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: asrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asrole.FieldID,
			},
		}
	)
	if id, ok := arc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := arc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asrole.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := arc.mutation.RoleAlias(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asrole.FieldRoleAlias,
		})
		_node.RoleAlias = value
	}
	if value, ok := arc.mutation.RoleName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asrole.FieldRoleName,
		})
		_node.RoleName = value
	}
	if value, ok := arc.mutation.RoleDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asrole.FieldRoleDescription,
		})
		_node.RoleDescription = value
	}
	if value, ok := arc.mutation.IsDeleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asrole.FieldIsDeleted,
		})
		_node.IsDeleted = value
	}
	if value, ok := arc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asrole.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := arc.mutation.CreateUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asrole.FieldCreateUser,
		})
		_node.CreateUser = value
	}
	if value, ok := arc.mutation.UpdateUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asrole.FieldUpdateUser,
		})
		_node.UpdateUser = value
	}
	if value, ok := arc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asrole.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := arc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asrole.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if nodes := arc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asrole.UsersTable,
			Columns: asrole.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := arc.mutation.JobsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   asrole.JobsTable,
			Columns: asrole.JobsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asjob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := arc.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asrole.MenusTable,
			Columns: asrole.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := arc.mutation.AttrRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asrole.AttrRolesTable,
			Columns: []string{asrole.AttrRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: astenantattrrole.FieldID,
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

// AsRoleCreateBulk is the builder for creating many AsRole entities in bulk.
type AsRoleCreateBulk struct {
	config
	builders []*AsRoleCreate
}

// Save creates the AsRole entities in the database.
func (arcb *AsRoleCreateBulk) Save(ctx context.Context) ([]*AsRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*AsRole, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AsRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arcb *AsRoleCreateBulk) SaveX(ctx context.Context) []*AsRole {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arcb *AsRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := arcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arcb *AsRoleCreateBulk) ExecX(ctx context.Context) {
	if err := arcb.Exec(ctx); err != nil {
		panic(err)
	}
}