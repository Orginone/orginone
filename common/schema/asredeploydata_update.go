// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asredeploydata"
	"orginone/common/schema/predicate"
	"orginone/common/tools/date"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsRedeployDataUpdate is the builder for updating AsRedeployData entities.
type AsRedeployDataUpdate struct {
	config
	hooks    []Hook
	mutation *AsRedeployDataMutation
}

// Where appends a list predicates to the AsRedeployDataUpdate builder.
func (ardu *AsRedeployDataUpdate) Where(ps ...predicate.AsRedeployData) *AsRedeployDataUpdate {
	ardu.mutation.Where(ps...)
	return ardu
}

// SetAppID sets the "app_id" field.
func (ardu *AsRedeployDataUpdate) SetAppID(i int64) *AsRedeployDataUpdate {
	ardu.mutation.SetAppID(i)
	return ardu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ardu *AsRedeployDataUpdate) SetNillableAppID(i *int64) *AsRedeployDataUpdate {
	if i != nil {
		ardu.SetAppID(*i)
	}
	return ardu
}

// ClearAppID clears the value of the "app_id" field.
func (ardu *AsRedeployDataUpdate) ClearAppID() *AsRedeployDataUpdate {
	ardu.mutation.ClearAppID()
	return ardu
}

// SetApplication sets the "application" field.
func (ardu *AsRedeployDataUpdate) SetApplication(s string) *AsRedeployDataUpdate {
	ardu.mutation.SetApplication(s)
	return ardu
}

// SetNillableApplication sets the "application" field if the given value is not nil.
func (ardu *AsRedeployDataUpdate) SetNillableApplication(s *string) *AsRedeployDataUpdate {
	if s != nil {
		ardu.SetApplication(*s)
	}
	return ardu
}

// ClearApplication clears the value of the "application" field.
func (ardu *AsRedeployDataUpdate) ClearApplication() *AsRedeployDataUpdate {
	ardu.mutation.ClearApplication()
	return ardu
}

// SetIsDeleted sets the "is_deleted" field.
func (ardu *AsRedeployDataUpdate) SetIsDeleted(i int64) *AsRedeployDataUpdate {
	ardu.mutation.ResetIsDeleted()
	ardu.mutation.SetIsDeleted(i)
	return ardu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (ardu *AsRedeployDataUpdate) SetNillableIsDeleted(i *int64) *AsRedeployDataUpdate {
	if i != nil {
		ardu.SetIsDeleted(*i)
	}
	return ardu
}

// AddIsDeleted adds i to the "is_deleted" field.
func (ardu *AsRedeployDataUpdate) AddIsDeleted(i int64) *AsRedeployDataUpdate {
	ardu.mutation.AddIsDeleted(i)
	return ardu
}

// SetStatus sets the "status" field.
func (ardu *AsRedeployDataUpdate) SetStatus(i int64) *AsRedeployDataUpdate {
	ardu.mutation.ResetStatus()
	ardu.mutation.SetStatus(i)
	return ardu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ardu *AsRedeployDataUpdate) SetNillableStatus(i *int64) *AsRedeployDataUpdate {
	if i != nil {
		ardu.SetStatus(*i)
	}
	return ardu
}

// AddStatus adds i to the "status" field.
func (ardu *AsRedeployDataUpdate) AddStatus(i int64) *AsRedeployDataUpdate {
	ardu.mutation.AddStatus(i)
	return ardu
}

// ClearStatus clears the value of the "status" field.
func (ardu *AsRedeployDataUpdate) ClearStatus() *AsRedeployDataUpdate {
	ardu.mutation.ClearStatus()
	return ardu
}

// SetCreateUser sets the "create_user" field.
func (ardu *AsRedeployDataUpdate) SetCreateUser(i int64) *AsRedeployDataUpdate {
	ardu.mutation.ResetCreateUser()
	ardu.mutation.SetCreateUser(i)
	return ardu
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (ardu *AsRedeployDataUpdate) SetNillableCreateUser(i *int64) *AsRedeployDataUpdate {
	if i != nil {
		ardu.SetCreateUser(*i)
	}
	return ardu
}

// AddCreateUser adds i to the "create_user" field.
func (ardu *AsRedeployDataUpdate) AddCreateUser(i int64) *AsRedeployDataUpdate {
	ardu.mutation.AddCreateUser(i)
	return ardu
}

// ClearCreateUser clears the value of the "create_user" field.
func (ardu *AsRedeployDataUpdate) ClearCreateUser() *AsRedeployDataUpdate {
	ardu.mutation.ClearCreateUser()
	return ardu
}

// SetUpdateUser sets the "update_user" field.
func (ardu *AsRedeployDataUpdate) SetUpdateUser(i int64) *AsRedeployDataUpdate {
	ardu.mutation.ResetUpdateUser()
	ardu.mutation.SetUpdateUser(i)
	return ardu
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (ardu *AsRedeployDataUpdate) SetNillableUpdateUser(i *int64) *AsRedeployDataUpdate {
	if i != nil {
		ardu.SetUpdateUser(*i)
	}
	return ardu
}

// AddUpdateUser adds i to the "update_user" field.
func (ardu *AsRedeployDataUpdate) AddUpdateUser(i int64) *AsRedeployDataUpdate {
	ardu.mutation.AddUpdateUser(i)
	return ardu
}

// ClearUpdateUser clears the value of the "update_user" field.
func (ardu *AsRedeployDataUpdate) ClearUpdateUser() *AsRedeployDataUpdate {
	ardu.mutation.ClearUpdateUser()
	return ardu
}

// SetUpdateTime sets the "update_time" field.
func (ardu *AsRedeployDataUpdate) SetUpdateTime(dt date.DateTime) *AsRedeployDataUpdate {
	ardu.mutation.SetUpdateTime(dt)
	return ardu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (ardu *AsRedeployDataUpdate) ClearUpdateTime() *AsRedeployDataUpdate {
	ardu.mutation.ClearUpdateTime()
	return ardu
}

// SetAppxID sets the "appx" edge to the AsMarketApp entity by ID.
func (ardu *AsRedeployDataUpdate) SetAppxID(id int64) *AsRedeployDataUpdate {
	ardu.mutation.SetAppxID(id)
	return ardu
}

// SetNillableAppxID sets the "appx" edge to the AsMarketApp entity by ID if the given value is not nil.
func (ardu *AsRedeployDataUpdate) SetNillableAppxID(id *int64) *AsRedeployDataUpdate {
	if id != nil {
		ardu = ardu.SetAppxID(*id)
	}
	return ardu
}

// SetAppx sets the "appx" edge to the AsMarketApp entity.
func (ardu *AsRedeployDataUpdate) SetAppx(a *AsMarketApp) *AsRedeployDataUpdate {
	return ardu.SetAppxID(a.ID)
}

// Mutation returns the AsRedeployDataMutation object of the builder.
func (ardu *AsRedeployDataUpdate) Mutation() *AsRedeployDataMutation {
	return ardu.mutation
}

// ClearAppx clears the "appx" edge to the AsMarketApp entity.
func (ardu *AsRedeployDataUpdate) ClearAppx() *AsRedeployDataUpdate {
	ardu.mutation.ClearAppx()
	return ardu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ardu *AsRedeployDataUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ardu.defaults()
	if len(ardu.hooks) == 0 {
		affected, err = ardu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsRedeployDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ardu.mutation = mutation
			affected, err = ardu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ardu.hooks) - 1; i >= 0; i-- {
			if ardu.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = ardu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ardu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ardu *AsRedeployDataUpdate) SaveX(ctx context.Context) int {
	affected, err := ardu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ardu *AsRedeployDataUpdate) Exec(ctx context.Context) error {
	_, err := ardu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ardu *AsRedeployDataUpdate) ExecX(ctx context.Context) {
	if err := ardu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ardu *AsRedeployDataUpdate) defaults() {
	if _, ok := ardu.mutation.UpdateTime(); !ok && !ardu.mutation.UpdateTimeCleared() {
		v := asredeploydata.UpdateDefaultUpdateTime()
		ardu.mutation.SetUpdateTime(v)
	}
}

func (ardu *AsRedeployDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asredeploydata.Table,
			Columns: asredeploydata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asredeploydata.FieldID,
			},
		},
	}
	if ps := ardu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ardu.mutation.Application(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asredeploydata.FieldApplication,
		})
	}
	if ardu.mutation.ApplicationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: asredeploydata.FieldApplication,
		})
	}
	if value, ok := ardu.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldIsDeleted,
		})
	}
	if value, ok := ardu.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldIsDeleted,
		})
	}
	if value, ok := ardu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldStatus,
		})
	}
	if value, ok := ardu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldStatus,
		})
	}
	if ardu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asredeploydata.FieldStatus,
		})
	}
	if value, ok := ardu.mutation.CreateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldCreateUser,
		})
	}
	if value, ok := ardu.mutation.AddedCreateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldCreateUser,
		})
	}
	if ardu.mutation.CreateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asredeploydata.FieldCreateUser,
		})
	}
	if value, ok := ardu.mutation.UpdateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldUpdateUser,
		})
	}
	if value, ok := ardu.mutation.AddedUpdateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldUpdateUser,
		})
	}
	if ardu.mutation.UpdateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asredeploydata.FieldUpdateUser,
		})
	}
	if ardu.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asredeploydata.FieldCreateTime,
		})
	}
	if value, ok := ardu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asredeploydata.FieldUpdateTime,
		})
	}
	if ardu.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asredeploydata.FieldUpdateTime,
		})
	}
	if ardu.mutation.AppxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asredeploydata.AppxTable,
			Columns: []string{asredeploydata.AppxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketapp.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ardu.mutation.AppxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asredeploydata.AppxTable,
			Columns: []string{asredeploydata.AppxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketapp.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ardu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asredeploydata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AsRedeployDataUpdateOne is the builder for updating a single AsRedeployData entity.
type AsRedeployDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AsRedeployDataMutation
}

// SetAppID sets the "app_id" field.
func (arduo *AsRedeployDataUpdateOne) SetAppID(i int64) *AsRedeployDataUpdateOne {
	arduo.mutation.SetAppID(i)
	return arduo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (arduo *AsRedeployDataUpdateOne) SetNillableAppID(i *int64) *AsRedeployDataUpdateOne {
	if i != nil {
		arduo.SetAppID(*i)
	}
	return arduo
}

// ClearAppID clears the value of the "app_id" field.
func (arduo *AsRedeployDataUpdateOne) ClearAppID() *AsRedeployDataUpdateOne {
	arduo.mutation.ClearAppID()
	return arduo
}

// SetApplication sets the "application" field.
func (arduo *AsRedeployDataUpdateOne) SetApplication(s string) *AsRedeployDataUpdateOne {
	arduo.mutation.SetApplication(s)
	return arduo
}

// SetNillableApplication sets the "application" field if the given value is not nil.
func (arduo *AsRedeployDataUpdateOne) SetNillableApplication(s *string) *AsRedeployDataUpdateOne {
	if s != nil {
		arduo.SetApplication(*s)
	}
	return arduo
}

// ClearApplication clears the value of the "application" field.
func (arduo *AsRedeployDataUpdateOne) ClearApplication() *AsRedeployDataUpdateOne {
	arduo.mutation.ClearApplication()
	return arduo
}

// SetIsDeleted sets the "is_deleted" field.
func (arduo *AsRedeployDataUpdateOne) SetIsDeleted(i int64) *AsRedeployDataUpdateOne {
	arduo.mutation.ResetIsDeleted()
	arduo.mutation.SetIsDeleted(i)
	return arduo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (arduo *AsRedeployDataUpdateOne) SetNillableIsDeleted(i *int64) *AsRedeployDataUpdateOne {
	if i != nil {
		arduo.SetIsDeleted(*i)
	}
	return arduo
}

// AddIsDeleted adds i to the "is_deleted" field.
func (arduo *AsRedeployDataUpdateOne) AddIsDeleted(i int64) *AsRedeployDataUpdateOne {
	arduo.mutation.AddIsDeleted(i)
	return arduo
}

// SetStatus sets the "status" field.
func (arduo *AsRedeployDataUpdateOne) SetStatus(i int64) *AsRedeployDataUpdateOne {
	arduo.mutation.ResetStatus()
	arduo.mutation.SetStatus(i)
	return arduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (arduo *AsRedeployDataUpdateOne) SetNillableStatus(i *int64) *AsRedeployDataUpdateOne {
	if i != nil {
		arduo.SetStatus(*i)
	}
	return arduo
}

// AddStatus adds i to the "status" field.
func (arduo *AsRedeployDataUpdateOne) AddStatus(i int64) *AsRedeployDataUpdateOne {
	arduo.mutation.AddStatus(i)
	return arduo
}

// ClearStatus clears the value of the "status" field.
func (arduo *AsRedeployDataUpdateOne) ClearStatus() *AsRedeployDataUpdateOne {
	arduo.mutation.ClearStatus()
	return arduo
}

// SetCreateUser sets the "create_user" field.
func (arduo *AsRedeployDataUpdateOne) SetCreateUser(i int64) *AsRedeployDataUpdateOne {
	arduo.mutation.ResetCreateUser()
	arduo.mutation.SetCreateUser(i)
	return arduo
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (arduo *AsRedeployDataUpdateOne) SetNillableCreateUser(i *int64) *AsRedeployDataUpdateOne {
	if i != nil {
		arduo.SetCreateUser(*i)
	}
	return arduo
}

// AddCreateUser adds i to the "create_user" field.
func (arduo *AsRedeployDataUpdateOne) AddCreateUser(i int64) *AsRedeployDataUpdateOne {
	arduo.mutation.AddCreateUser(i)
	return arduo
}

// ClearCreateUser clears the value of the "create_user" field.
func (arduo *AsRedeployDataUpdateOne) ClearCreateUser() *AsRedeployDataUpdateOne {
	arduo.mutation.ClearCreateUser()
	return arduo
}

// SetUpdateUser sets the "update_user" field.
func (arduo *AsRedeployDataUpdateOne) SetUpdateUser(i int64) *AsRedeployDataUpdateOne {
	arduo.mutation.ResetUpdateUser()
	arduo.mutation.SetUpdateUser(i)
	return arduo
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (arduo *AsRedeployDataUpdateOne) SetNillableUpdateUser(i *int64) *AsRedeployDataUpdateOne {
	if i != nil {
		arduo.SetUpdateUser(*i)
	}
	return arduo
}

// AddUpdateUser adds i to the "update_user" field.
func (arduo *AsRedeployDataUpdateOne) AddUpdateUser(i int64) *AsRedeployDataUpdateOne {
	arduo.mutation.AddUpdateUser(i)
	return arduo
}

// ClearUpdateUser clears the value of the "update_user" field.
func (arduo *AsRedeployDataUpdateOne) ClearUpdateUser() *AsRedeployDataUpdateOne {
	arduo.mutation.ClearUpdateUser()
	return arduo
}

// SetUpdateTime sets the "update_time" field.
func (arduo *AsRedeployDataUpdateOne) SetUpdateTime(dt date.DateTime) *AsRedeployDataUpdateOne {
	arduo.mutation.SetUpdateTime(dt)
	return arduo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (arduo *AsRedeployDataUpdateOne) ClearUpdateTime() *AsRedeployDataUpdateOne {
	arduo.mutation.ClearUpdateTime()
	return arduo
}

// SetAppxID sets the "appx" edge to the AsMarketApp entity by ID.
func (arduo *AsRedeployDataUpdateOne) SetAppxID(id int64) *AsRedeployDataUpdateOne {
	arduo.mutation.SetAppxID(id)
	return arduo
}

// SetNillableAppxID sets the "appx" edge to the AsMarketApp entity by ID if the given value is not nil.
func (arduo *AsRedeployDataUpdateOne) SetNillableAppxID(id *int64) *AsRedeployDataUpdateOne {
	if id != nil {
		arduo = arduo.SetAppxID(*id)
	}
	return arduo
}

// SetAppx sets the "appx" edge to the AsMarketApp entity.
func (arduo *AsRedeployDataUpdateOne) SetAppx(a *AsMarketApp) *AsRedeployDataUpdateOne {
	return arduo.SetAppxID(a.ID)
}

// Mutation returns the AsRedeployDataMutation object of the builder.
func (arduo *AsRedeployDataUpdateOne) Mutation() *AsRedeployDataMutation {
	return arduo.mutation
}

// ClearAppx clears the "appx" edge to the AsMarketApp entity.
func (arduo *AsRedeployDataUpdateOne) ClearAppx() *AsRedeployDataUpdateOne {
	arduo.mutation.ClearAppx()
	return arduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (arduo *AsRedeployDataUpdateOne) Select(field string, fields ...string) *AsRedeployDataUpdateOne {
	arduo.fields = append([]string{field}, fields...)
	return arduo
}

// Save executes the query and returns the updated AsRedeployData entity.
func (arduo *AsRedeployDataUpdateOne) Save(ctx context.Context) (*AsRedeployData, error) {
	var (
		err  error
		node *AsRedeployData
	)
	arduo.defaults()
	if len(arduo.hooks) == 0 {
		node, err = arduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsRedeployDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			arduo.mutation = mutation
			node, err = arduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(arduo.hooks) - 1; i >= 0; i-- {
			if arduo.hooks[i] == nil {
				return nil, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = arduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, arduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (arduo *AsRedeployDataUpdateOne) SaveX(ctx context.Context) *AsRedeployData {
	node, err := arduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (arduo *AsRedeployDataUpdateOne) Exec(ctx context.Context) error {
	_, err := arduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arduo *AsRedeployDataUpdateOne) ExecX(ctx context.Context) {
	if err := arduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (arduo *AsRedeployDataUpdateOne) defaults() {
	if _, ok := arduo.mutation.UpdateTime(); !ok && !arduo.mutation.UpdateTimeCleared() {
		v := asredeploydata.UpdateDefaultUpdateTime()
		arduo.mutation.SetUpdateTime(v)
	}
}

func (arduo *AsRedeployDataUpdateOne) sqlSave(ctx context.Context) (_node *AsRedeployData, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asredeploydata.Table,
			Columns: asredeploydata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asredeploydata.FieldID,
			},
		},
	}
	id, ok := arduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`schema: missing "AsRedeployData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := arduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asredeploydata.FieldID)
		for _, f := range fields {
			if !asredeploydata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
			}
			if f != asredeploydata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := arduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := arduo.mutation.Application(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asredeploydata.FieldApplication,
		})
	}
	if arduo.mutation.ApplicationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: asredeploydata.FieldApplication,
		})
	}
	if value, ok := arduo.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldIsDeleted,
		})
	}
	if value, ok := arduo.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldIsDeleted,
		})
	}
	if value, ok := arduo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldStatus,
		})
	}
	if value, ok := arduo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldStatus,
		})
	}
	if arduo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asredeploydata.FieldStatus,
		})
	}
	if value, ok := arduo.mutation.CreateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldCreateUser,
		})
	}
	if value, ok := arduo.mutation.AddedCreateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldCreateUser,
		})
	}
	if arduo.mutation.CreateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asredeploydata.FieldCreateUser,
		})
	}
	if value, ok := arduo.mutation.UpdateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldUpdateUser,
		})
	}
	if value, ok := arduo.mutation.AddedUpdateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asredeploydata.FieldUpdateUser,
		})
	}
	if arduo.mutation.UpdateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asredeploydata.FieldUpdateUser,
		})
	}
	if arduo.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asredeploydata.FieldCreateTime,
		})
	}
	if value, ok := arduo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asredeploydata.FieldUpdateTime,
		})
	}
	if arduo.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asredeploydata.FieldUpdateTime,
		})
	}
	if arduo.mutation.AppxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asredeploydata.AppxTable,
			Columns: []string{asredeploydata.AppxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketapp.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := arduo.mutation.AppxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asredeploydata.AppxTable,
			Columns: []string{asredeploydata.AppxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketapp.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AsRedeployData{config: arduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, arduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asredeploydata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}