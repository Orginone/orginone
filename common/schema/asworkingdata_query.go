// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"math"
	"orginone/common/schema/asuser"
	"orginone/common/schema/asworkingdata"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsWorkingDataQuery is the builder for querying AsWorkingData entities.
type AsWorkingDataQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AsWorkingData
	// eager-loading edges.
	withUser *AsUserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AsWorkingDataQuery builder.
func (awdq *AsWorkingDataQuery) Where(ps ...predicate.AsWorkingData) *AsWorkingDataQuery {
	awdq.predicates = append(awdq.predicates, ps...)
	return awdq
}

// Limit adds a limit step to the query.
func (awdq *AsWorkingDataQuery) Limit(limit int) *AsWorkingDataQuery {
	awdq.limit = &limit
	return awdq
}

// Offset adds an offset step to the query.
func (awdq *AsWorkingDataQuery) Offset(offset int) *AsWorkingDataQuery {
	awdq.offset = &offset
	return awdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (awdq *AsWorkingDataQuery) Unique(unique bool) *AsWorkingDataQuery {
	awdq.unique = &unique
	return awdq
}

// Order adds an order step to the query.
func (awdq *AsWorkingDataQuery) Order(o ...OrderFunc) *AsWorkingDataQuery {
	awdq.order = append(awdq.order, o...)
	return awdq
}

// QueryUser chains the current query on the "user" edge.
func (awdq *AsWorkingDataQuery) QueryUser() *AsUserQuery {
	query := &AsUserQuery{config: awdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := awdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := awdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asworkingdata.Table, asworkingdata.FieldID, selector),
			sqlgraph.To(asuser.Table, asuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, asworkingdata.UserTable, asworkingdata.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(awdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AsWorkingData entity from the query.
// Returns a *NotFoundError when no AsWorkingData was found.
func (awdq *AsWorkingDataQuery) First(ctx context.Context) (*AsWorkingData, error) {
	nodes, err := awdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{asworkingdata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (awdq *AsWorkingDataQuery) FirstX(ctx context.Context) *AsWorkingData {
	node, err := awdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AsWorkingData ID from the query.
// Returns a *NotFoundError when no AsWorkingData ID was found.
func (awdq *AsWorkingDataQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = awdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{asworkingdata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (awdq *AsWorkingDataQuery) FirstIDX(ctx context.Context) int64 {
	id, err := awdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AsWorkingData entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AsWorkingData entity is found.
// Returns a *NotFoundError when no AsWorkingData entities are found.
func (awdq *AsWorkingDataQuery) Only(ctx context.Context) (*AsWorkingData, error) {
	nodes, err := awdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{asworkingdata.Label}
	default:
		return nil, &NotSingularError{asworkingdata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (awdq *AsWorkingDataQuery) OnlyX(ctx context.Context) *AsWorkingData {
	node, err := awdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AsWorkingData ID in the query.
// Returns a *NotSingularError when more than one AsWorkingData ID is found.
// Returns a *NotFoundError when no entities are found.
func (awdq *AsWorkingDataQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = awdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{asworkingdata.Label}
	default:
		err = &NotSingularError{asworkingdata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (awdq *AsWorkingDataQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := awdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AsWorkingDataSlice.
func (awdq *AsWorkingDataQuery) All(ctx context.Context) ([]*AsWorkingData, error) {
	if err := awdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return awdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (awdq *AsWorkingDataQuery) AllX(ctx context.Context) []*AsWorkingData {
	nodes, err := awdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AsWorkingData IDs.
func (awdq *AsWorkingDataQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := awdq.Select(asworkingdata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (awdq *AsWorkingDataQuery) IDsX(ctx context.Context) []int64 {
	ids, err := awdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (awdq *AsWorkingDataQuery) Count(ctx context.Context) (int64, error) {
	if err := awdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return awdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (awdq *AsWorkingDataQuery) CountX(ctx context.Context) int64 {
	count, err := awdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (awdq *AsWorkingDataQuery) Exist(ctx context.Context) (bool, error) {
	if err := awdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return awdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (awdq *AsWorkingDataQuery) ExistX(ctx context.Context) bool {
	exist, err := awdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AsWorkingDataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (awdq *AsWorkingDataQuery) Clone() *AsWorkingDataQuery {
	if awdq == nil {
		return nil
	}
	return &AsWorkingDataQuery{
		config:     awdq.config,
		limit:      awdq.limit,
		offset:     awdq.offset,
		order:      append([]OrderFunc{}, awdq.order...),
		predicates: append([]predicate.AsWorkingData{}, awdq.predicates...),
		withUser:   awdq.withUser.Clone(),
		// clone intermediate query.
		sql:    awdq.sql.Clone(),
		path:   awdq.path,
		unique: awdq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (awdq *AsWorkingDataQuery) WithUser(opts ...func(*AsUserQuery)) *AsWorkingDataQuery {
	query := &AsUserQuery{config: awdq.config}
	for _, opt := range opts {
		opt(query)
	}
	awdq.withUser = query
	return awdq
}

// ThenUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (awdq *AsWorkingDataQuery) ThenUser(opts ...func(*AsUserQuery)) *AsWorkingDataQuery {
	query := &AsUserQuery{config: awdq.config}
	for _, opt := range opts {
		opt(query.Where(asuser.IsDeleted(0)))
	}
	awdq.withUser = query
	return awdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AppID int64 `json:"appId"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AsWorkingData.Query().
//		GroupBy(asworkingdata.FieldAppID).
//		Aggregate(schema.Count()).
//		Scan(ctx, &v)
//
func (awdq *AsWorkingDataQuery) GroupBy(field string, fields ...string) *AsWorkingDataGroupBy {
	group := &AsWorkingDataGroupBy{config: awdq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := awdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return awdq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AppID int64 `json:"appId"`
//	}
//
//	client.AsWorkingData.Query().
//		Select(asworkingdata.FieldAppID).
//		Scan(ctx, &v)
//
func (awdq *AsWorkingDataQuery) Select(fields ...string) *AsWorkingDataSelect {
	awdq.fields = append(awdq.fields, fields...)
	return &AsWorkingDataSelect{AsWorkingDataQuery: awdq}
}

func (awdq *AsWorkingDataQuery) prepareQuery(ctx context.Context) error {
	for _, f := range awdq.fields {
		if !asworkingdata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
		}
	}
	if awdq.path != nil {
		prev, err := awdq.path(ctx)
		if err != nil {
			return err
		}
		awdq.sql = prev
	}
	return nil
}

func (awdq *AsWorkingDataQuery) sqlAll(ctx context.Context) ([]*AsWorkingData, error) {
	var (
		nodes       = []*AsWorkingData{}
		_spec       = awdq.querySpec()
		loadedTypes = [1]bool{
			awdq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AsWorkingData{config: awdq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("schema: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, awdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := awdq.withUser; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*AsWorkingData)
		for i := range nodes {
			fk := nodes[i].UserID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(asuser.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (awdq *AsWorkingDataQuery) sqlCount(ctx context.Context) (int64, error) {
	_spec := awdq.querySpec()
	_spec.Node.Columns = awdq.fields
	if len(awdq.fields) > 0 {
		_spec.Unique = awdq.unique != nil && *awdq.unique
	}
	c, err := sqlgraph.CountNodes(ctx, awdq.driver, _spec)
	return int64(c), err
}

func (awdq *AsWorkingDataQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := awdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("schema: check existence: %w", err)
	}
	return n > 0, nil
}

func (awdq *AsWorkingDataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asworkingdata.Table,
			Columns: asworkingdata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asworkingdata.FieldID,
			},
		},
		From:   awdq.sql,
		Unique: true,
	}
	if unique := awdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := awdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asworkingdata.FieldID)
		for i := range fields {
			if fields[i] != asworkingdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := awdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := awdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := awdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := awdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (awdq *AsWorkingDataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(awdq.driver.Dialect())
	t1 := builder.Table(asworkingdata.Table)
	columns := awdq.fields
	if len(columns) == 0 {
		columns = asworkingdata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if awdq.sql != nil {
		selector = awdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if awdq.unique != nil && *awdq.unique {
		selector.Distinct()
	}
	for _, p := range awdq.predicates {
		p(selector)
	}
	for _, p := range awdq.order {
		p(selector)
	}
	if offset := awdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := awdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AsWorkingDataGroupBy is the group-by builder for AsWorkingData entities.
type AsWorkingDataGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (awdgb *AsWorkingDataGroupBy) Aggregate(fns ...AggregateFunc) *AsWorkingDataGroupBy {
	awdgb.fns = append(awdgb.fns, fns...)
	return awdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (awdgb *AsWorkingDataGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := awdgb.path(ctx)
	if err != nil {
		return err
	}
	awdgb.sql = query
	return awdgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (awdgb *AsWorkingDataGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := awdgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (awdgb *AsWorkingDataGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(awdgb.fields) > 1 {
		return nil, errors.New("schema: AsWorkingDataGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := awdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (awdgb *AsWorkingDataGroupBy) StringsX(ctx context.Context) []string {
	v, err := awdgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (awdgb *AsWorkingDataGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = awdgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asworkingdata.Label}
	default:
		err = fmt.Errorf("schema: AsWorkingDataGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (awdgb *AsWorkingDataGroupBy) StringX(ctx context.Context) string {
	v, err := awdgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (awdgb *AsWorkingDataGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(awdgb.fields) > 1 {
		return nil, errors.New("schema: AsWorkingDataGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := awdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (awdgb *AsWorkingDataGroupBy) IntsX(ctx context.Context) []int {
	v, err := awdgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (awdgb *AsWorkingDataGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = awdgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asworkingdata.Label}
	default:
		err = fmt.Errorf("schema: AsWorkingDataGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (awdgb *AsWorkingDataGroupBy) IntX(ctx context.Context) int {
	v, err := awdgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (awdgb *AsWorkingDataGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(awdgb.fields) > 1 {
		return nil, errors.New("schema: AsWorkingDataGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := awdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (awdgb *AsWorkingDataGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := awdgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (awdgb *AsWorkingDataGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = awdgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asworkingdata.Label}
	default:
		err = fmt.Errorf("schema: AsWorkingDataGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (awdgb *AsWorkingDataGroupBy) Float64X(ctx context.Context) float64 {
	v, err := awdgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (awdgb *AsWorkingDataGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(awdgb.fields) > 1 {
		return nil, errors.New("schema: AsWorkingDataGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := awdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (awdgb *AsWorkingDataGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := awdgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (awdgb *AsWorkingDataGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = awdgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asworkingdata.Label}
	default:
		err = fmt.Errorf("schema: AsWorkingDataGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (awdgb *AsWorkingDataGroupBy) BoolX(ctx context.Context) bool {
	v, err := awdgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (awdgb *AsWorkingDataGroupBy) Int64s(ctx context.Context) ([]int64, error) {
	if len(awdgb.fields) > 1 {
		return nil, errors.New("schema: AsWorkingDataGroupBy.Int64s is not achievable when grouping more than 1 field")
	}
	var v []int64
	if err := awdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (awdgb *AsWorkingDataGroupBy) Int64sX(ctx context.Context) []int64 {
	v, err := awdgb.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (awdgb *AsWorkingDataGroupBy) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = awdgb.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asworkingdata.Label}
	default:
		err = fmt.Errorf("schema: AsWorkingDataGroupBy.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (awdgb *AsWorkingDataGroupBy) Int64X(ctx context.Context) int64 {
	v, err := awdgb.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (awdgb *AsWorkingDataGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range awdgb.fields {
		if !asworkingdata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := awdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := awdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (awdgb *AsWorkingDataGroupBy) sqlQuery() *sql.Selector {
	selector := awdgb.sql.Select()
	aggregation := make([]string, 0, len(awdgb.fns))
	for _, fn := range awdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(awdgb.fields)+len(awdgb.fns))
		for _, f := range awdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(awdgb.fields...)...)
}

// AsWorkingDataSelect is the builder for selecting fields of AsWorkingData entities.
type AsWorkingDataSelect struct {
	*AsWorkingDataQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (awds *AsWorkingDataSelect) Scan(ctx context.Context, v interface{}) error {
	if err := awds.prepareQuery(ctx); err != nil {
		return err
	}
	awds.sql = awds.AsWorkingDataQuery.sqlQuery(ctx)
	return awds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (awds *AsWorkingDataSelect) ScanX(ctx context.Context, v interface{}) {
	if err := awds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (awds *AsWorkingDataSelect) Strings(ctx context.Context) ([]string, error) {
	if len(awds.fields) > 1 {
		return nil, errors.New("schema: AsWorkingDataSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := awds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (awds *AsWorkingDataSelect) StringsX(ctx context.Context) []string {
	v, err := awds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (awds *AsWorkingDataSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = awds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asworkingdata.Label}
	default:
		err = fmt.Errorf("schema: AsWorkingDataSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (awds *AsWorkingDataSelect) StringX(ctx context.Context) string {
	v, err := awds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (awds *AsWorkingDataSelect) Ints(ctx context.Context) ([]int, error) {
	if len(awds.fields) > 1 {
		return nil, errors.New("schema: AsWorkingDataSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := awds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (awds *AsWorkingDataSelect) IntsX(ctx context.Context) []int {
	v, err := awds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (awds *AsWorkingDataSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = awds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asworkingdata.Label}
	default:
		err = fmt.Errorf("schema: AsWorkingDataSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (awds *AsWorkingDataSelect) IntX(ctx context.Context) int {
	v, err := awds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (awds *AsWorkingDataSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(awds.fields) > 1 {
		return nil, errors.New("schema: AsWorkingDataSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := awds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (awds *AsWorkingDataSelect) Float64sX(ctx context.Context) []float64 {
	v, err := awds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (awds *AsWorkingDataSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = awds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asworkingdata.Label}
	default:
		err = fmt.Errorf("schema: AsWorkingDataSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (awds *AsWorkingDataSelect) Float64X(ctx context.Context) float64 {
	v, err := awds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (awds *AsWorkingDataSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(awds.fields) > 1 {
		return nil, errors.New("schema: AsWorkingDataSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := awds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (awds *AsWorkingDataSelect) BoolsX(ctx context.Context) []bool {
	v, err := awds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (awds *AsWorkingDataSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = awds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asworkingdata.Label}
	default:
		err = fmt.Errorf("schema: AsWorkingDataSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (awds *AsWorkingDataSelect) BoolX(ctx context.Context) bool {
	v, err := awds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from a selector. It is only allowed when selecting one field.
func (awds *AsWorkingDataSelect) Int64s(ctx context.Context) ([]int64, error) {
	if len(awds.fields) > 1 {
		return nil, errors.New("schema: AsWorkingDataSelect.Int64s is not achievable when selecting more than 1 field")
	}
	var v []int64
	if err := awds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (awds *AsWorkingDataSelect) Int64sX(ctx context.Context) []int64 {
	v, err := awds.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a selector. It is only allowed when selecting one field.
func (awds *AsWorkingDataSelect) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = awds.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asworkingdata.Label}
	default:
		err = fmt.Errorf("schema: AsWorkingDataSelect.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (awds *AsWorkingDataSelect) Int64X(ctx context.Context) int64 {
	v, err := awds.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (awds *AsWorkingDataSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := awds.sql.Query()
	if err := awds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}