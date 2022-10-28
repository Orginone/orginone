// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"orginone/common/schema/astenantattr"
	"orginone/common/schema/astenantattrrole"
	"orginone/common/schema/predicate"
	"orginone/common/tools/date"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsTenantAttrUpdate is the builder for updating AsTenantAttr entities.
type AsTenantAttrUpdate struct {
	config
	hooks    []Hook
	mutation *AsTenantAttrMutation
}

// Where appends a list predicates to the AsTenantAttrUpdate builder.
func (atau *AsTenantAttrUpdate) Where(ps ...predicate.AsTenantAttr) *AsTenantAttrUpdate {
	atau.mutation.Where(ps...)
	return atau
}

// SetParentID sets the "parent_id" field.
func (atau *AsTenantAttrUpdate) SetParentID(i int64) *AsTenantAttrUpdate {
	atau.mutation.SetParentID(i)
	return atau
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (atau *AsTenantAttrUpdate) SetNillableParentID(i *int64) *AsTenantAttrUpdate {
	if i != nil {
		atau.SetParentID(*i)
	}
	return atau
}

// ClearParentID clears the value of the "parent_id" field.
func (atau *AsTenantAttrUpdate) ClearParentID() *AsTenantAttrUpdate {
	atau.mutation.ClearParentID()
	return atau
}

// SetAttrName sets the "attr_name" field.
func (atau *AsTenantAttrUpdate) SetAttrName(s string) *AsTenantAttrUpdate {
	atau.mutation.SetAttrName(s)
	return atau
}

// SetAttrRemark sets the "attr_remark" field.
func (atau *AsTenantAttrUpdate) SetAttrRemark(s string) *AsTenantAttrUpdate {
	atau.mutation.SetAttrRemark(s)
	return atau
}

// SetNillableAttrRemark sets the "attr_remark" field if the given value is not nil.
func (atau *AsTenantAttrUpdate) SetNillableAttrRemark(s *string) *AsTenantAttrUpdate {
	if s != nil {
		atau.SetAttrRemark(*s)
	}
	return atau
}

// ClearAttrRemark clears the value of the "attr_remark" field.
func (atau *AsTenantAttrUpdate) ClearAttrRemark() *AsTenantAttrUpdate {
	atau.mutation.ClearAttrRemark()
	return atau
}

// SetIsDeleted sets the "is_deleted" field.
func (atau *AsTenantAttrUpdate) SetIsDeleted(i int64) *AsTenantAttrUpdate {
	atau.mutation.ResetIsDeleted()
	atau.mutation.SetIsDeleted(i)
	return atau
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (atau *AsTenantAttrUpdate) SetNillableIsDeleted(i *int64) *AsTenantAttrUpdate {
	if i != nil {
		atau.SetIsDeleted(*i)
	}
	return atau
}

// AddIsDeleted adds i to the "is_deleted" field.
func (atau *AsTenantAttrUpdate) AddIsDeleted(i int64) *AsTenantAttrUpdate {
	atau.mutation.AddIsDeleted(i)
	return atau
}

// SetStatus sets the "status" field.
func (atau *AsTenantAttrUpdate) SetStatus(i int64) *AsTenantAttrUpdate {
	atau.mutation.ResetStatus()
	atau.mutation.SetStatus(i)
	return atau
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (atau *AsTenantAttrUpdate) SetNillableStatus(i *int64) *AsTenantAttrUpdate {
	if i != nil {
		atau.SetStatus(*i)
	}
	return atau
}

// AddStatus adds i to the "status" field.
func (atau *AsTenantAttrUpdate) AddStatus(i int64) *AsTenantAttrUpdate {
	atau.mutation.AddStatus(i)
	return atau
}

// ClearStatus clears the value of the "status" field.
func (atau *AsTenantAttrUpdate) ClearStatus() *AsTenantAttrUpdate {
	atau.mutation.ClearStatus()
	return atau
}

// SetCreateUser sets the "create_user" field.
func (atau *AsTenantAttrUpdate) SetCreateUser(i int64) *AsTenantAttrUpdate {
	atau.mutation.ResetCreateUser()
	atau.mutation.SetCreateUser(i)
	return atau
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (atau *AsTenantAttrUpdate) SetNillableCreateUser(i *int64) *AsTenantAttrUpdate {
	if i != nil {
		atau.SetCreateUser(*i)
	}
	return atau
}

// AddCreateUser adds i to the "create_user" field.
func (atau *AsTenantAttrUpdate) AddCreateUser(i int64) *AsTenantAttrUpdate {
	atau.mutation.AddCreateUser(i)
	return atau
}

// ClearCreateUser clears the value of the "create_user" field.
func (atau *AsTenantAttrUpdate) ClearCreateUser() *AsTenantAttrUpdate {
	atau.mutation.ClearCreateUser()
	return atau
}

// SetUpdateUser sets the "update_user" field.
func (atau *AsTenantAttrUpdate) SetUpdateUser(i int64) *AsTenantAttrUpdate {
	atau.mutation.ResetUpdateUser()
	atau.mutation.SetUpdateUser(i)
	return atau
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (atau *AsTenantAttrUpdate) SetNillableUpdateUser(i *int64) *AsTenantAttrUpdate {
	if i != nil {
		atau.SetUpdateUser(*i)
	}
	return atau
}

// AddUpdateUser adds i to the "update_user" field.
func (atau *AsTenantAttrUpdate) AddUpdateUser(i int64) *AsTenantAttrUpdate {
	atau.mutation.AddUpdateUser(i)
	return atau
}

// ClearUpdateUser clears the value of the "update_user" field.
func (atau *AsTenantAttrUpdate) ClearUpdateUser() *AsTenantAttrUpdate {
	atau.mutation.ClearUpdateUser()
	return atau
}

// SetUpdateTime sets the "update_time" field.
func (atau *AsTenantAttrUpdate) SetUpdateTime(dt date.DateTime) *AsTenantAttrUpdate {
	atau.mutation.SetUpdateTime(dt)
	return atau
}

// ClearUpdateTime clears the value of the "update_time" field.
func (atau *AsTenantAttrUpdate) ClearUpdateTime() *AsTenantAttrUpdate {
	atau.mutation.ClearUpdateTime()
	return atau
}

// SetParentxID sets the "parentx" edge to the AsTenantAttr entity by ID.
func (atau *AsTenantAttrUpdate) SetParentxID(id int64) *AsTenantAttrUpdate {
	atau.mutation.SetParentxID(id)
	return atau
}

// SetNillableParentxID sets the "parentx" edge to the AsTenantAttr entity by ID if the given value is not nil.
func (atau *AsTenantAttrUpdate) SetNillableParentxID(id *int64) *AsTenantAttrUpdate {
	if id != nil {
		atau = atau.SetParentxID(*id)
	}
	return atau
}

// SetParentx sets the "parentx" edge to the AsTenantAttr entity.
func (atau *AsTenantAttrUpdate) SetParentx(a *AsTenantAttr) *AsTenantAttrUpdate {
	return atau.SetParentxID(a.ID)
}

// AddChildrenIDs adds the "childrens" edge to the AsTenantAttr entity by IDs.
func (atau *AsTenantAttrUpdate) AddChildrenIDs(ids ...int64) *AsTenantAttrUpdate {
	atau.mutation.AddChildrenIDs(ids...)
	return atau
}

// AddChildrens adds the "childrens" edges to the AsTenantAttr entity.
func (atau *AsTenantAttrUpdate) AddChildrens(a ...*AsTenantAttr) *AsTenantAttrUpdate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atau.AddChildrenIDs(ids...)
}

// AddAttrRoleIDs adds the "attrRoles" edge to the AsTenantAttrRole entity by IDs.
func (atau *AsTenantAttrUpdate) AddAttrRoleIDs(ids ...int64) *AsTenantAttrUpdate {
	atau.mutation.AddAttrRoleIDs(ids...)
	return atau
}

// AddAttrRoles adds the "attrRoles" edges to the AsTenantAttrRole entity.
func (atau *AsTenantAttrUpdate) AddAttrRoles(a ...*AsTenantAttrRole) *AsTenantAttrUpdate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atau.AddAttrRoleIDs(ids...)
}

// Mutation returns the AsTenantAttrMutation object of the builder.
func (atau *AsTenantAttrUpdate) Mutation() *AsTenantAttrMutation {
	return atau.mutation
}

// ClearParentx clears the "parentx" edge to the AsTenantAttr entity.
func (atau *AsTenantAttrUpdate) ClearParentx() *AsTenantAttrUpdate {
	atau.mutation.ClearParentx()
	return atau
}

// ClearChildrens clears all "childrens" edges to the AsTenantAttr entity.
func (atau *AsTenantAttrUpdate) ClearChildrens() *AsTenantAttrUpdate {
	atau.mutation.ClearChildrens()
	return atau
}

// RemoveChildrenIDs removes the "childrens" edge to AsTenantAttr entities by IDs.
func (atau *AsTenantAttrUpdate) RemoveChildrenIDs(ids ...int64) *AsTenantAttrUpdate {
	atau.mutation.RemoveChildrenIDs(ids...)
	return atau
}

// RemoveChildrens removes "childrens" edges to AsTenantAttr entities.
func (atau *AsTenantAttrUpdate) RemoveChildrens(a ...*AsTenantAttr) *AsTenantAttrUpdate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atau.RemoveChildrenIDs(ids...)
}

// ClearAttrRoles clears all "attrRoles" edges to the AsTenantAttrRole entity.
func (atau *AsTenantAttrUpdate) ClearAttrRoles() *AsTenantAttrUpdate {
	atau.mutation.ClearAttrRoles()
	return atau
}

// RemoveAttrRoleIDs removes the "attrRoles" edge to AsTenantAttrRole entities by IDs.
func (atau *AsTenantAttrUpdate) RemoveAttrRoleIDs(ids ...int64) *AsTenantAttrUpdate {
	atau.mutation.RemoveAttrRoleIDs(ids...)
	return atau
}

// RemoveAttrRoles removes "attrRoles" edges to AsTenantAttrRole entities.
func (atau *AsTenantAttrUpdate) RemoveAttrRoles(a ...*AsTenantAttrRole) *AsTenantAttrUpdate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atau.RemoveAttrRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atau *AsTenantAttrUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	atau.defaults()
	if len(atau.hooks) == 0 {
		affected, err = atau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsTenantAttrMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atau.mutation = mutation
			affected, err = atau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atau.hooks) - 1; i >= 0; i-- {
			if atau.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = atau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (atau *AsTenantAttrUpdate) SaveX(ctx context.Context) int {
	affected, err := atau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atau *AsTenantAttrUpdate) Exec(ctx context.Context) error {
	_, err := atau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atau *AsTenantAttrUpdate) ExecX(ctx context.Context) {
	if err := atau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atau *AsTenantAttrUpdate) defaults() {
	if _, ok := atau.mutation.UpdateTime(); !ok && !atau.mutation.UpdateTimeCleared() {
		v := astenantattr.UpdateDefaultUpdateTime()
		atau.mutation.SetUpdateTime(v)
	}
}

func (atau *AsTenantAttrUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   astenantattr.Table,
			Columns: astenantattr.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: astenantattr.FieldID,
			},
		},
	}
	if ps := atau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atau.mutation.AttrName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: astenantattr.FieldAttrName,
		})
	}
	if value, ok := atau.mutation.AttrRemark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: astenantattr.FieldAttrRemark,
		})
	}
	if atau.mutation.AttrRemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: astenantattr.FieldAttrRemark,
		})
	}
	if value, ok := atau.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldIsDeleted,
		})
	}
	if value, ok := atau.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldIsDeleted,
		})
	}
	if value, ok := atau.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldStatus,
		})
	}
	if value, ok := atau.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldStatus,
		})
	}
	if atau.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: astenantattr.FieldStatus,
		})
	}
	if value, ok := atau.mutation.CreateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldCreateUser,
		})
	}
	if value, ok := atau.mutation.AddedCreateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldCreateUser,
		})
	}
	if atau.mutation.CreateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: astenantattr.FieldCreateUser,
		})
	}
	if value, ok := atau.mutation.UpdateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldUpdateUser,
		})
	}
	if value, ok := atau.mutation.AddedUpdateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldUpdateUser,
		})
	}
	if atau.mutation.UpdateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: astenantattr.FieldUpdateUser,
		})
	}
	if atau.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: astenantattr.FieldCreateTime,
		})
	}
	if value, ok := atau.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: astenantattr.FieldUpdateTime,
		})
	}
	if atau.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: astenantattr.FieldUpdateTime,
		})
	}
	if atau.mutation.ParentxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   astenantattr.ParentxTable,
			Columns: []string{astenantattr.ParentxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: astenantattr.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atau.mutation.ParentxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   astenantattr.ParentxTable,
			Columns: []string{astenantattr.ParentxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: astenantattr.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atau.mutation.ChildrensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   astenantattr.ChildrensTable,
			Columns: []string{astenantattr.ChildrensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: astenantattr.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atau.mutation.RemovedChildrensIDs(); len(nodes) > 0 && !atau.mutation.ChildrensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   astenantattr.ChildrensTable,
			Columns: []string{astenantattr.ChildrensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: astenantattr.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atau.mutation.ChildrensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   astenantattr.ChildrensTable,
			Columns: []string{astenantattr.ChildrensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: astenantattr.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atau.mutation.AttrRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   astenantattr.AttrRolesTable,
			Columns: []string{astenantattr.AttrRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: astenantattrrole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atau.mutation.RemovedAttrRolesIDs(); len(nodes) > 0 && !atau.mutation.AttrRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   astenantattr.AttrRolesTable,
			Columns: []string{astenantattr.AttrRolesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atau.mutation.AttrRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   astenantattr.AttrRolesTable,
			Columns: []string{astenantattr.AttrRolesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{astenantattr.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AsTenantAttrUpdateOne is the builder for updating a single AsTenantAttr entity.
type AsTenantAttrUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AsTenantAttrMutation
}

// SetParentID sets the "parent_id" field.
func (atauo *AsTenantAttrUpdateOne) SetParentID(i int64) *AsTenantAttrUpdateOne {
	atauo.mutation.SetParentID(i)
	return atauo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (atauo *AsTenantAttrUpdateOne) SetNillableParentID(i *int64) *AsTenantAttrUpdateOne {
	if i != nil {
		atauo.SetParentID(*i)
	}
	return atauo
}

// ClearParentID clears the value of the "parent_id" field.
func (atauo *AsTenantAttrUpdateOne) ClearParentID() *AsTenantAttrUpdateOne {
	atauo.mutation.ClearParentID()
	return atauo
}

// SetAttrName sets the "attr_name" field.
func (atauo *AsTenantAttrUpdateOne) SetAttrName(s string) *AsTenantAttrUpdateOne {
	atauo.mutation.SetAttrName(s)
	return atauo
}

// SetAttrRemark sets the "attr_remark" field.
func (atauo *AsTenantAttrUpdateOne) SetAttrRemark(s string) *AsTenantAttrUpdateOne {
	atauo.mutation.SetAttrRemark(s)
	return atauo
}

// SetNillableAttrRemark sets the "attr_remark" field if the given value is not nil.
func (atauo *AsTenantAttrUpdateOne) SetNillableAttrRemark(s *string) *AsTenantAttrUpdateOne {
	if s != nil {
		atauo.SetAttrRemark(*s)
	}
	return atauo
}

// ClearAttrRemark clears the value of the "attr_remark" field.
func (atauo *AsTenantAttrUpdateOne) ClearAttrRemark() *AsTenantAttrUpdateOne {
	atauo.mutation.ClearAttrRemark()
	return atauo
}

// SetIsDeleted sets the "is_deleted" field.
func (atauo *AsTenantAttrUpdateOne) SetIsDeleted(i int64) *AsTenantAttrUpdateOne {
	atauo.mutation.ResetIsDeleted()
	atauo.mutation.SetIsDeleted(i)
	return atauo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (atauo *AsTenantAttrUpdateOne) SetNillableIsDeleted(i *int64) *AsTenantAttrUpdateOne {
	if i != nil {
		atauo.SetIsDeleted(*i)
	}
	return atauo
}

// AddIsDeleted adds i to the "is_deleted" field.
func (atauo *AsTenantAttrUpdateOne) AddIsDeleted(i int64) *AsTenantAttrUpdateOne {
	atauo.mutation.AddIsDeleted(i)
	return atauo
}

// SetStatus sets the "status" field.
func (atauo *AsTenantAttrUpdateOne) SetStatus(i int64) *AsTenantAttrUpdateOne {
	atauo.mutation.ResetStatus()
	atauo.mutation.SetStatus(i)
	return atauo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (atauo *AsTenantAttrUpdateOne) SetNillableStatus(i *int64) *AsTenantAttrUpdateOne {
	if i != nil {
		atauo.SetStatus(*i)
	}
	return atauo
}

// AddStatus adds i to the "status" field.
func (atauo *AsTenantAttrUpdateOne) AddStatus(i int64) *AsTenantAttrUpdateOne {
	atauo.mutation.AddStatus(i)
	return atauo
}

// ClearStatus clears the value of the "status" field.
func (atauo *AsTenantAttrUpdateOne) ClearStatus() *AsTenantAttrUpdateOne {
	atauo.mutation.ClearStatus()
	return atauo
}

// SetCreateUser sets the "create_user" field.
func (atauo *AsTenantAttrUpdateOne) SetCreateUser(i int64) *AsTenantAttrUpdateOne {
	atauo.mutation.ResetCreateUser()
	atauo.mutation.SetCreateUser(i)
	return atauo
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (atauo *AsTenantAttrUpdateOne) SetNillableCreateUser(i *int64) *AsTenantAttrUpdateOne {
	if i != nil {
		atauo.SetCreateUser(*i)
	}
	return atauo
}

// AddCreateUser adds i to the "create_user" field.
func (atauo *AsTenantAttrUpdateOne) AddCreateUser(i int64) *AsTenantAttrUpdateOne {
	atauo.mutation.AddCreateUser(i)
	return atauo
}

// ClearCreateUser clears the value of the "create_user" field.
func (atauo *AsTenantAttrUpdateOne) ClearCreateUser() *AsTenantAttrUpdateOne {
	atauo.mutation.ClearCreateUser()
	return atauo
}

// SetUpdateUser sets the "update_user" field.
func (atauo *AsTenantAttrUpdateOne) SetUpdateUser(i int64) *AsTenantAttrUpdateOne {
	atauo.mutation.ResetUpdateUser()
	atauo.mutation.SetUpdateUser(i)
	return atauo
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (atauo *AsTenantAttrUpdateOne) SetNillableUpdateUser(i *int64) *AsTenantAttrUpdateOne {
	if i != nil {
		atauo.SetUpdateUser(*i)
	}
	return atauo
}

// AddUpdateUser adds i to the "update_user" field.
func (atauo *AsTenantAttrUpdateOne) AddUpdateUser(i int64) *AsTenantAttrUpdateOne {
	atauo.mutation.AddUpdateUser(i)
	return atauo
}

// ClearUpdateUser clears the value of the "update_user" field.
func (atauo *AsTenantAttrUpdateOne) ClearUpdateUser() *AsTenantAttrUpdateOne {
	atauo.mutation.ClearUpdateUser()
	return atauo
}

// SetUpdateTime sets the "update_time" field.
func (atauo *AsTenantAttrUpdateOne) SetUpdateTime(dt date.DateTime) *AsTenantAttrUpdateOne {
	atauo.mutation.SetUpdateTime(dt)
	return atauo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (atauo *AsTenantAttrUpdateOne) ClearUpdateTime() *AsTenantAttrUpdateOne {
	atauo.mutation.ClearUpdateTime()
	return atauo
}

// SetParentxID sets the "parentx" edge to the AsTenantAttr entity by ID.
func (atauo *AsTenantAttrUpdateOne) SetParentxID(id int64) *AsTenantAttrUpdateOne {
	atauo.mutation.SetParentxID(id)
	return atauo
}

// SetNillableParentxID sets the "parentx" edge to the AsTenantAttr entity by ID if the given value is not nil.
func (atauo *AsTenantAttrUpdateOne) SetNillableParentxID(id *int64) *AsTenantAttrUpdateOne {
	if id != nil {
		atauo = atauo.SetParentxID(*id)
	}
	return atauo
}

// SetParentx sets the "parentx" edge to the AsTenantAttr entity.
func (atauo *AsTenantAttrUpdateOne) SetParentx(a *AsTenantAttr) *AsTenantAttrUpdateOne {
	return atauo.SetParentxID(a.ID)
}

// AddChildrenIDs adds the "childrens" edge to the AsTenantAttr entity by IDs.
func (atauo *AsTenantAttrUpdateOne) AddChildrenIDs(ids ...int64) *AsTenantAttrUpdateOne {
	atauo.mutation.AddChildrenIDs(ids...)
	return atauo
}

// AddChildrens adds the "childrens" edges to the AsTenantAttr entity.
func (atauo *AsTenantAttrUpdateOne) AddChildrens(a ...*AsTenantAttr) *AsTenantAttrUpdateOne {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atauo.AddChildrenIDs(ids...)
}

// AddAttrRoleIDs adds the "attrRoles" edge to the AsTenantAttrRole entity by IDs.
func (atauo *AsTenantAttrUpdateOne) AddAttrRoleIDs(ids ...int64) *AsTenantAttrUpdateOne {
	atauo.mutation.AddAttrRoleIDs(ids...)
	return atauo
}

// AddAttrRoles adds the "attrRoles" edges to the AsTenantAttrRole entity.
func (atauo *AsTenantAttrUpdateOne) AddAttrRoles(a ...*AsTenantAttrRole) *AsTenantAttrUpdateOne {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atauo.AddAttrRoleIDs(ids...)
}

// Mutation returns the AsTenantAttrMutation object of the builder.
func (atauo *AsTenantAttrUpdateOne) Mutation() *AsTenantAttrMutation {
	return atauo.mutation
}

// ClearParentx clears the "parentx" edge to the AsTenantAttr entity.
func (atauo *AsTenantAttrUpdateOne) ClearParentx() *AsTenantAttrUpdateOne {
	atauo.mutation.ClearParentx()
	return atauo
}

// ClearChildrens clears all "childrens" edges to the AsTenantAttr entity.
func (atauo *AsTenantAttrUpdateOne) ClearChildrens() *AsTenantAttrUpdateOne {
	atauo.mutation.ClearChildrens()
	return atauo
}

// RemoveChildrenIDs removes the "childrens" edge to AsTenantAttr entities by IDs.
func (atauo *AsTenantAttrUpdateOne) RemoveChildrenIDs(ids ...int64) *AsTenantAttrUpdateOne {
	atauo.mutation.RemoveChildrenIDs(ids...)
	return atauo
}

// RemoveChildrens removes "childrens" edges to AsTenantAttr entities.
func (atauo *AsTenantAttrUpdateOne) RemoveChildrens(a ...*AsTenantAttr) *AsTenantAttrUpdateOne {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atauo.RemoveChildrenIDs(ids...)
}

// ClearAttrRoles clears all "attrRoles" edges to the AsTenantAttrRole entity.
func (atauo *AsTenantAttrUpdateOne) ClearAttrRoles() *AsTenantAttrUpdateOne {
	atauo.mutation.ClearAttrRoles()
	return atauo
}

// RemoveAttrRoleIDs removes the "attrRoles" edge to AsTenantAttrRole entities by IDs.
func (atauo *AsTenantAttrUpdateOne) RemoveAttrRoleIDs(ids ...int64) *AsTenantAttrUpdateOne {
	atauo.mutation.RemoveAttrRoleIDs(ids...)
	return atauo
}

// RemoveAttrRoles removes "attrRoles" edges to AsTenantAttrRole entities.
func (atauo *AsTenantAttrUpdateOne) RemoveAttrRoles(a ...*AsTenantAttrRole) *AsTenantAttrUpdateOne {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return atauo.RemoveAttrRoleIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atauo *AsTenantAttrUpdateOne) Select(field string, fields ...string) *AsTenantAttrUpdateOne {
	atauo.fields = append([]string{field}, fields...)
	return atauo
}

// Save executes the query and returns the updated AsTenantAttr entity.
func (atauo *AsTenantAttrUpdateOne) Save(ctx context.Context) (*AsTenantAttr, error) {
	var (
		err  error
		node *AsTenantAttr
	)
	atauo.defaults()
	if len(atauo.hooks) == 0 {
		node, err = atauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsTenantAttrMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atauo.mutation = mutation
			node, err = atauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(atauo.hooks) - 1; i >= 0; i-- {
			if atauo.hooks[i] == nil {
				return nil, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = atauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (atauo *AsTenantAttrUpdateOne) SaveX(ctx context.Context) *AsTenantAttr {
	node, err := atauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atauo *AsTenantAttrUpdateOne) Exec(ctx context.Context) error {
	_, err := atauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atauo *AsTenantAttrUpdateOne) ExecX(ctx context.Context) {
	if err := atauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atauo *AsTenantAttrUpdateOne) defaults() {
	if _, ok := atauo.mutation.UpdateTime(); !ok && !atauo.mutation.UpdateTimeCleared() {
		v := astenantattr.UpdateDefaultUpdateTime()
		atauo.mutation.SetUpdateTime(v)
	}
}

func (atauo *AsTenantAttrUpdateOne) sqlSave(ctx context.Context) (_node *AsTenantAttr, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   astenantattr.Table,
			Columns: astenantattr.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: astenantattr.FieldID,
			},
		},
	}
	id, ok := atauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`schema: missing "AsTenantAttr.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, astenantattr.FieldID)
		for _, f := range fields {
			if !astenantattr.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
			}
			if f != astenantattr.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atauo.mutation.AttrName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: astenantattr.FieldAttrName,
		})
	}
	if value, ok := atauo.mutation.AttrRemark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: astenantattr.FieldAttrRemark,
		})
	}
	if atauo.mutation.AttrRemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: astenantattr.FieldAttrRemark,
		})
	}
	if value, ok := atauo.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldIsDeleted,
		})
	}
	if value, ok := atauo.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldIsDeleted,
		})
	}
	if value, ok := atauo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldStatus,
		})
	}
	if value, ok := atauo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldStatus,
		})
	}
	if atauo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: astenantattr.FieldStatus,
		})
	}
	if value, ok := atauo.mutation.CreateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldCreateUser,
		})
	}
	if value, ok := atauo.mutation.AddedCreateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldCreateUser,
		})
	}
	if atauo.mutation.CreateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: astenantattr.FieldCreateUser,
		})
	}
	if value, ok := atauo.mutation.UpdateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldUpdateUser,
		})
	}
	if value, ok := atauo.mutation.AddedUpdateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: astenantattr.FieldUpdateUser,
		})
	}
	if atauo.mutation.UpdateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: astenantattr.FieldUpdateUser,
		})
	}
	if atauo.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: astenantattr.FieldCreateTime,
		})
	}
	if value, ok := atauo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: astenantattr.FieldUpdateTime,
		})
	}
	if atauo.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: astenantattr.FieldUpdateTime,
		})
	}
	if atauo.mutation.ParentxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   astenantattr.ParentxTable,
			Columns: []string{astenantattr.ParentxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: astenantattr.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atauo.mutation.ParentxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   astenantattr.ParentxTable,
			Columns: []string{astenantattr.ParentxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: astenantattr.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atauo.mutation.ChildrensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   astenantattr.ChildrensTable,
			Columns: []string{astenantattr.ChildrensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: astenantattr.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atauo.mutation.RemovedChildrensIDs(); len(nodes) > 0 && !atauo.mutation.ChildrensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   astenantattr.ChildrensTable,
			Columns: []string{astenantattr.ChildrensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: astenantattr.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atauo.mutation.ChildrensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   astenantattr.ChildrensTable,
			Columns: []string{astenantattr.ChildrensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: astenantattr.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atauo.mutation.AttrRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   astenantattr.AttrRolesTable,
			Columns: []string{astenantattr.AttrRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: astenantattrrole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atauo.mutation.RemovedAttrRolesIDs(); len(nodes) > 0 && !atauo.mutation.AttrRolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   astenantattr.AttrRolesTable,
			Columns: []string{astenantattr.AttrRolesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atauo.mutation.AttrRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   astenantattr.AttrRolesTable,
			Columns: []string{astenantattr.AttrRolesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AsTenantAttr{config: atauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{astenantattr.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
