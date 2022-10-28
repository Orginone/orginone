// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"orginone/common/schema/asinputdata"
	"orginone/common/tools/date"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsInputDataCreate is the builder for creating a AsInputData entity.
type AsInputDataCreate struct {
	config
	mutation *AsInputDataMutation
	hooks    []Hook
}

// SetFileID sets the "file_id" field.
func (aidc *AsInputDataCreate) SetFileID(i int64) *AsInputDataCreate {
	aidc.mutation.SetFileID(i)
	return aidc
}

// SetFileName sets the "file_name" field.
func (aidc *AsInputDataCreate) SetFileName(s string) *AsInputDataCreate {
	aidc.mutation.SetFileName(s)
	return aidc
}

// SetTableName sets the "table_name" field.
func (aidc *AsInputDataCreate) SetTableName(s string) *AsInputDataCreate {
	aidc.mutation.SetTableName(s)
	return aidc
}

// SetType sets the "type" field.
func (aidc *AsInputDataCreate) SetType(i int64) *AsInputDataCreate {
	aidc.mutation.SetType(i)
	return aidc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableType(i *int64) *AsInputDataCreate {
	if i != nil {
		aidc.SetType(*i)
	}
	return aidc
}

// SetTCount sets the "t_count" field.
func (aidc *AsInputDataCreate) SetTCount(i int64) *AsInputDataCreate {
	aidc.mutation.SetTCount(i)
	return aidc
}

// SetNillableTCount sets the "t_count" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableTCount(i *int64) *AsInputDataCreate {
	if i != nil {
		aidc.SetTCount(*i)
	}
	return aidc
}

// SetFCount sets the "f_count" field.
func (aidc *AsInputDataCreate) SetFCount(i int64) *AsInputDataCreate {
	aidc.mutation.SetFCount(i)
	return aidc
}

// SetNillableFCount sets the "f_count" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableFCount(i *int64) *AsInputDataCreate {
	if i != nil {
		aidc.SetFCount(*i)
	}
	return aidc
}

// SetContext sets the "context" field.
func (aidc *AsInputDataCreate) SetContext(s string) *AsInputDataCreate {
	aidc.mutation.SetContext(s)
	return aidc
}

// SetNillableContext sets the "context" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableContext(s *string) *AsInputDataCreate {
	if s != nil {
		aidc.SetContext(*s)
	}
	return aidc
}

// SetEndTime sets the "end_time" field.
func (aidc *AsInputDataCreate) SetEndTime(dt date.DateTime) *AsInputDataCreate {
	aidc.mutation.SetEndTime(dt)
	return aidc
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableEndTime(dt *date.DateTime) *AsInputDataCreate {
	if dt != nil {
		aidc.SetEndTime(*dt)
	}
	return aidc
}

// SetTotalTime sets the "total_time" field.
func (aidc *AsInputDataCreate) SetTotalTime(i int64) *AsInputDataCreate {
	aidc.mutation.SetTotalTime(i)
	return aidc
}

// SetNillableTotalTime sets the "total_time" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableTotalTime(i *int64) *AsInputDataCreate {
	if i != nil {
		aidc.SetTotalTime(*i)
	}
	return aidc
}

// SetTenantCode sets the "tenant_code" field.
func (aidc *AsInputDataCreate) SetTenantCode(s string) *AsInputDataCreate {
	aidc.mutation.SetTenantCode(s)
	return aidc
}

// SetNillableTenantCode sets the "tenant_code" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableTenantCode(s *string) *AsInputDataCreate {
	if s != nil {
		aidc.SetTenantCode(*s)
	}
	return aidc
}

// SetImportType sets the "import_type" field.
func (aidc *AsInputDataCreate) SetImportType(i int64) *AsInputDataCreate {
	aidc.mutation.SetImportType(i)
	return aidc
}

// SetNillableImportType sets the "import_type" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableImportType(i *int64) *AsInputDataCreate {
	if i != nil {
		aidc.SetImportType(*i)
	}
	return aidc
}

// SetIsDeleted sets the "is_deleted" field.
func (aidc *AsInputDataCreate) SetIsDeleted(i int64) *AsInputDataCreate {
	aidc.mutation.SetIsDeleted(i)
	return aidc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableIsDeleted(i *int64) *AsInputDataCreate {
	if i != nil {
		aidc.SetIsDeleted(*i)
	}
	return aidc
}

// SetStatus sets the "status" field.
func (aidc *AsInputDataCreate) SetStatus(i int64) *AsInputDataCreate {
	aidc.mutation.SetStatus(i)
	return aidc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableStatus(i *int64) *AsInputDataCreate {
	if i != nil {
		aidc.SetStatus(*i)
	}
	return aidc
}

// SetCreateUser sets the "create_user" field.
func (aidc *AsInputDataCreate) SetCreateUser(i int64) *AsInputDataCreate {
	aidc.mutation.SetCreateUser(i)
	return aidc
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableCreateUser(i *int64) *AsInputDataCreate {
	if i != nil {
		aidc.SetCreateUser(*i)
	}
	return aidc
}

// SetUpdateUser sets the "update_user" field.
func (aidc *AsInputDataCreate) SetUpdateUser(i int64) *AsInputDataCreate {
	aidc.mutation.SetUpdateUser(i)
	return aidc
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableUpdateUser(i *int64) *AsInputDataCreate {
	if i != nil {
		aidc.SetUpdateUser(*i)
	}
	return aidc
}

// SetCreateTime sets the "create_time" field.
func (aidc *AsInputDataCreate) SetCreateTime(dt date.DateTime) *AsInputDataCreate {
	aidc.mutation.SetCreateTime(dt)
	return aidc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableCreateTime(dt *date.DateTime) *AsInputDataCreate {
	if dt != nil {
		aidc.SetCreateTime(*dt)
	}
	return aidc
}

// SetUpdateTime sets the "update_time" field.
func (aidc *AsInputDataCreate) SetUpdateTime(dt date.DateTime) *AsInputDataCreate {
	aidc.mutation.SetUpdateTime(dt)
	return aidc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableUpdateTime(dt *date.DateTime) *AsInputDataCreate {
	if dt != nil {
		aidc.SetUpdateTime(*dt)
	}
	return aidc
}

// SetID sets the "id" field.
func (aidc *AsInputDataCreate) SetID(i int64) *AsInputDataCreate {
	aidc.mutation.SetID(i)
	return aidc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (aidc *AsInputDataCreate) SetNillableID(i *int64) *AsInputDataCreate {
	if i != nil {
		aidc.SetID(*i)
	}
	return aidc
}

// Mutation returns the AsInputDataMutation object of the builder.
func (aidc *AsInputDataCreate) Mutation() *AsInputDataMutation {
	return aidc.mutation
}

// Save creates the AsInputData in the database.
func (aidc *AsInputDataCreate) Save(ctx context.Context) (*AsInputData, error) {
	var (
		err  error
		node *AsInputData
	)
	aidc.defaults()
	if len(aidc.hooks) == 0 {
		if err = aidc.check(); err != nil {
			return nil, err
		}
		node, err = aidc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsInputDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aidc.check(); err != nil {
				return nil, err
			}
			aidc.mutation = mutation
			if node, err = aidc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(aidc.hooks) - 1; i >= 0; i-- {
			if aidc.hooks[i] == nil {
				return nil, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = aidc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aidc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (aidc *AsInputDataCreate) SaveX(ctx context.Context) *AsInputData {
	v, err := aidc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aidc *AsInputDataCreate) Exec(ctx context.Context) error {
	_, err := aidc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aidc *AsInputDataCreate) ExecX(ctx context.Context) {
	if err := aidc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aidc *AsInputDataCreate) defaults() {
	if _, ok := aidc.mutation.IsDeleted(); !ok {
		v := asinputdata.DefaultIsDeleted
		aidc.mutation.SetIsDeleted(v)
	}
	if _, ok := aidc.mutation.Status(); !ok {
		v := asinputdata.DefaultStatus
		aidc.mutation.SetStatus(v)
	}
	if _, ok := aidc.mutation.CreateTime(); !ok {
		v := asinputdata.DefaultCreateTime()
		aidc.mutation.SetCreateTime(v)
	}
	if _, ok := aidc.mutation.UpdateTime(); !ok {
		v := asinputdata.DefaultUpdateTime()
		aidc.mutation.SetUpdateTime(v)
	}
	if _, ok := aidc.mutation.ID(); !ok {
		v := asinputdata.DefaultID()
		aidc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aidc *AsInputDataCreate) check() error {
	if _, ok := aidc.mutation.FileID(); !ok {
		return &ValidationError{Name: "file_id", err: errors.New(`schema: missing required field "AsInputData.file_id"`)}
	}
	if _, ok := aidc.mutation.FileName(); !ok {
		return &ValidationError{Name: "file_name", err: errors.New(`schema: missing required field "AsInputData.file_name"`)}
	}
	if _, ok := aidc.mutation.TableName(); !ok {
		return &ValidationError{Name: "table_name", err: errors.New(`schema: missing required field "AsInputData.table_name"`)}
	}
	if _, ok := aidc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`schema: missing required field "AsInputData.is_deleted"`)}
	}
	return nil
}

func (aidc *AsInputDataCreate) sqlSave(ctx context.Context) (*AsInputData, error) {
	_node, _spec := aidc.createSpec()
	if err := sqlgraph.CreateNode(ctx, aidc.driver, _spec); err != nil {
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

func (aidc *AsInputDataCreate) createSpec() (*AsInputData, *sqlgraph.CreateSpec) {
	var (
		_node = &AsInputData{config: aidc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: asinputdata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asinputdata.FieldID,
			},
		}
	)
	if id, ok := aidc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := aidc.mutation.FileID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asinputdata.FieldFileID,
		})
		_node.FileID = value
	}
	if value, ok := aidc.mutation.FileName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asinputdata.FieldFileName,
		})
		_node.FileName = value
	}
	if value, ok := aidc.mutation.TableName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asinputdata.FieldTableName,
		})
		_node.TableName = value
	}
	if value, ok := aidc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asinputdata.FieldType,
		})
		_node.Type = value
	}
	if value, ok := aidc.mutation.TCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asinputdata.FieldTCount,
		})
		_node.TCount = value
	}
	if value, ok := aidc.mutation.FCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asinputdata.FieldFCount,
		})
		_node.FCount = value
	}
	if value, ok := aidc.mutation.Context(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asinputdata.FieldContext,
		})
		_node.Context = value
	}
	if value, ok := aidc.mutation.EndTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asinputdata.FieldEndTime,
		})
		_node.EndTime = value
	}
	if value, ok := aidc.mutation.TotalTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asinputdata.FieldTotalTime,
		})
		_node.TotalTime = value
	}
	if value, ok := aidc.mutation.TenantCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asinputdata.FieldTenantCode,
		})
		_node.TenantCode = value
	}
	if value, ok := aidc.mutation.ImportType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asinputdata.FieldImportType,
		})
		_node.ImportType = value
	}
	if value, ok := aidc.mutation.IsDeleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asinputdata.FieldIsDeleted,
		})
		_node.IsDeleted = value
	}
	if value, ok := aidc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asinputdata.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := aidc.mutation.CreateUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asinputdata.FieldCreateUser,
		})
		_node.CreateUser = value
	}
	if value, ok := aidc.mutation.UpdateUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asinputdata.FieldUpdateUser,
		})
		_node.UpdateUser = value
	}
	if value, ok := aidc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asinputdata.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := aidc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asinputdata.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	return _node, _spec
}

// AsInputDataCreateBulk is the builder for creating many AsInputData entities in bulk.
type AsInputDataCreateBulk struct {
	config
	builders []*AsInputDataCreate
}

// Save creates the AsInputData entities in the database.
func (aidcb *AsInputDataCreateBulk) Save(ctx context.Context) ([]*AsInputData, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aidcb.builders))
	nodes := make([]*AsInputData, len(aidcb.builders))
	mutators := make([]Mutator, len(aidcb.builders))
	for i := range aidcb.builders {
		func(i int, root context.Context) {
			builder := aidcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AsInputDataMutation)
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
					_, err = mutators[i+1].Mutate(root, aidcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aidcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, aidcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aidcb *AsInputDataCreateBulk) SaveX(ctx context.Context) []*AsInputData {
	v, err := aidcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aidcb *AsInputDataCreateBulk) Exec(ctx context.Context) error {
	_, err := aidcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aidcb *AsInputDataCreateBulk) ExecX(ctx context.Context) {
	if err := aidcb.Exec(ctx); err != nil {
		panic(err)
	}
}
