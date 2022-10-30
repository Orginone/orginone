// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"math"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketappalert"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsMarketAppAlertQuery is the builder for querying AsMarketAppAlert entities.
type AsMarketAppAlertQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AsMarketAppAlert
	// eager-loading edges.
	withAppx *AsMarketAppQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AsMarketAppAlertQuery builder.
func (amaaq *AsMarketAppAlertQuery) Where(ps ...predicate.AsMarketAppAlert) *AsMarketAppAlertQuery {
	amaaq.predicates = append(amaaq.predicates, ps...)
	return amaaq
}

// Limit adds a limit step to the query.
func (amaaq *AsMarketAppAlertQuery) Limit(limit int) *AsMarketAppAlertQuery {
	amaaq.limit = &limit
	return amaaq
}

// Offset adds an offset step to the query.
func (amaaq *AsMarketAppAlertQuery) Offset(offset int) *AsMarketAppAlertQuery {
	amaaq.offset = &offset
	return amaaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (amaaq *AsMarketAppAlertQuery) Unique(unique bool) *AsMarketAppAlertQuery {
	amaaq.unique = &unique
	return amaaq
}

// Order adds an order step to the query.
func (amaaq *AsMarketAppAlertQuery) Order(o ...OrderFunc) *AsMarketAppAlertQuery {
	amaaq.order = append(amaaq.order, o...)
	return amaaq
}

// QueryAppx chains the current query on the "appx" edge.
func (amaaq *AsMarketAppAlertQuery) QueryAppx() *AsMarketAppQuery {
	query := &AsMarketAppQuery{config: amaaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := amaaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := amaaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asmarketappalert.Table, asmarketappalert.FieldID, selector),
			sqlgraph.To(asmarketapp.Table, asmarketapp.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, asmarketappalert.AppxTable, asmarketappalert.AppxColumn),
		)
		fromU = sqlgraph.SetNeighbors(amaaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AsMarketAppAlert entity from the query.
// Returns a *NotFoundError when no AsMarketAppAlert was found.
func (amaaq *AsMarketAppAlertQuery) First(ctx context.Context) (*AsMarketAppAlert, error) {
	nodes, err := amaaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{asmarketappalert.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (amaaq *AsMarketAppAlertQuery) FirstX(ctx context.Context) *AsMarketAppAlert {
	node, err := amaaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AsMarketAppAlert ID from the query.
// Returns a *NotFoundError when no AsMarketAppAlert ID was found.
func (amaaq *AsMarketAppAlertQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = amaaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{asmarketappalert.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (amaaq *AsMarketAppAlertQuery) FirstIDX(ctx context.Context) int64 {
	id, err := amaaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AsMarketAppAlert entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AsMarketAppAlert entity is found.
// Returns a *NotFoundError when no AsMarketAppAlert entities are found.
func (amaaq *AsMarketAppAlertQuery) Only(ctx context.Context) (*AsMarketAppAlert, error) {
	nodes, err := amaaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{asmarketappalert.Label}
	default:
		return nil, &NotSingularError{asmarketappalert.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (amaaq *AsMarketAppAlertQuery) OnlyX(ctx context.Context) *AsMarketAppAlert {
	node, err := amaaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AsMarketAppAlert ID in the query.
// Returns a *NotSingularError when more than one AsMarketAppAlert ID is found.
// Returns a *NotFoundError when no entities are found.
func (amaaq *AsMarketAppAlertQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = amaaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{asmarketappalert.Label}
	default:
		err = &NotSingularError{asmarketappalert.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (amaaq *AsMarketAppAlertQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := amaaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AsMarketAppAlerts.
func (amaaq *AsMarketAppAlertQuery) All(ctx context.Context) ([]*AsMarketAppAlert, error) {
	if err := amaaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return amaaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (amaaq *AsMarketAppAlertQuery) AllX(ctx context.Context) []*AsMarketAppAlert {
	nodes, err := amaaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AsMarketAppAlert IDs.
func (amaaq *AsMarketAppAlertQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := amaaq.Select(asmarketappalert.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (amaaq *AsMarketAppAlertQuery) IDsX(ctx context.Context) []int64 {
	ids, err := amaaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (amaaq *AsMarketAppAlertQuery) Count(ctx context.Context) (int64, error) {
	if err := amaaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return amaaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (amaaq *AsMarketAppAlertQuery) CountX(ctx context.Context) int64 {
	count, err := amaaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (amaaq *AsMarketAppAlertQuery) Exist(ctx context.Context) (bool, error) {
	if err := amaaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return amaaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (amaaq *AsMarketAppAlertQuery) ExistX(ctx context.Context) bool {
	exist, err := amaaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AsMarketAppAlertQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (amaaq *AsMarketAppAlertQuery) Clone() *AsMarketAppAlertQuery {
	if amaaq == nil {
		return nil
	}
	return &AsMarketAppAlertQuery{
		config:     amaaq.config,
		limit:      amaaq.limit,
		offset:     amaaq.offset,
		order:      append([]OrderFunc{}, amaaq.order...),
		predicates: append([]predicate.AsMarketAppAlert{}, amaaq.predicates...),
		withAppx:   amaaq.withAppx.Clone(),
		// clone intermediate query.
		sql:    amaaq.sql.Clone(),
		path:   amaaq.path,
		unique: amaaq.unique,
	}
}

// WithAppx tells the query-builder to eager-load the nodes that are connected to
// the "appx" edge. The optional arguments are used to configure the query builder of the edge.
func (amaaq *AsMarketAppAlertQuery) WithAppx(opts ...func(*AsMarketAppQuery)) *AsMarketAppAlertQuery {
	query := &AsMarketAppQuery{config: amaaq.config}
	for _, opt := range opts {
		opt(query)
	}
	amaaq.withAppx = query
	return amaaq
}

// ThenAppx tells the query-builder to eager-load the nodes that are connected to
// the "appx" edge. The optional arguments are used to configure the query builder of the edge.
func (amaaq *AsMarketAppAlertQuery) ThenAppx(opts ...func(*AsMarketAppQuery)) *AsMarketAppAlertQuery {
	query := &AsMarketAppQuery{config: amaaq.config}
	for _, opt := range opts {
		opt(query.Where(asmarketapp.IsDeleted(0)))
	}
	amaaq.withAppx = query
	return amaaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AlertTitle string `json:"alertTitle"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AsMarketAppAlert.Query().
//		GroupBy(asmarketappalert.FieldAlertTitle).
//		Aggregate(schema.Count()).
//		Scan(ctx, &v)
//
func (amaaq *AsMarketAppAlertQuery) GroupBy(field string, fields ...string) *AsMarketAppAlertGroupBy {
	group := &AsMarketAppAlertGroupBy{config: amaaq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := amaaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return amaaq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AlertTitle string `json:"alertTitle"`
//	}
//
//	client.AsMarketAppAlert.Query().
//		Select(asmarketappalert.FieldAlertTitle).
//		Scan(ctx, &v)
//
func (amaaq *AsMarketAppAlertQuery) Select(fields ...string) *AsMarketAppAlertSelect {
	amaaq.fields = append(amaaq.fields, fields...)
	return &AsMarketAppAlertSelect{AsMarketAppAlertQuery: amaaq}
}

func (amaaq *AsMarketAppAlertQuery) prepareQuery(ctx context.Context) error {
	for _, f := range amaaq.fields {
		if !asmarketappalert.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
		}
	}
	if amaaq.path != nil {
		prev, err := amaaq.path(ctx)
		if err != nil {
			return err
		}
		amaaq.sql = prev
	}
	return nil
}

func (amaaq *AsMarketAppAlertQuery) sqlAll(ctx context.Context) ([]*AsMarketAppAlert, error) {
	var (
		nodes       = []*AsMarketAppAlert{}
		_spec       = amaaq.querySpec()
		loadedTypes = [1]bool{
			amaaq.withAppx != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AsMarketAppAlert{config: amaaq.config}
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
	if err := sqlgraph.QueryNodes(ctx, amaaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := amaaq.withAppx; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*AsMarketAppAlert)
		for i := range nodes {
			fk := nodes[i].AlertReleaseAppID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(asmarketapp.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "alert_release_app_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Appx = n
			}
		}
	}

	return nodes, nil
}

func (amaaq *AsMarketAppAlertQuery) sqlCount(ctx context.Context) (int64, error) {
	_spec := amaaq.querySpec()
	_spec.Node.Columns = amaaq.fields
	if len(amaaq.fields) > 0 {
		_spec.Unique = amaaq.unique != nil && *amaaq.unique
	}
	c, err := sqlgraph.CountNodes(ctx, amaaq.driver, _spec)
	return int64(c), err
}

func (amaaq *AsMarketAppAlertQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := amaaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("schema: check existence: %w", err)
	}
	return n > 0, nil
}

func (amaaq *AsMarketAppAlertQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asmarketappalert.Table,
			Columns: asmarketappalert.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asmarketappalert.FieldID,
			},
		},
		From:   amaaq.sql,
		Unique: true,
	}
	if unique := amaaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := amaaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asmarketappalert.FieldID)
		for i := range fields {
			if fields[i] != asmarketappalert.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := amaaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := amaaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := amaaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := amaaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (amaaq *AsMarketAppAlertQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(amaaq.driver.Dialect())
	t1 := builder.Table(asmarketappalert.Table)
	columns := amaaq.fields
	if len(columns) == 0 {
		columns = asmarketappalert.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if amaaq.sql != nil {
		selector = amaaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if amaaq.unique != nil && *amaaq.unique {
		selector.Distinct()
	}
	for _, p := range amaaq.predicates {
		p(selector)
	}
	for _, p := range amaaq.order {
		p(selector)
	}
	if offset := amaaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := amaaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AsMarketAppAlertGroupBy is the group-by builder for AsMarketAppAlert entities.
type AsMarketAppAlertGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (amaagb *AsMarketAppAlertGroupBy) Aggregate(fns ...AggregateFunc) *AsMarketAppAlertGroupBy {
	amaagb.fns = append(amaagb.fns, fns...)
	return amaagb
}

// Scan applies the group-by query and scans the result into the given value.
func (amaagb *AsMarketAppAlertGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := amaagb.path(ctx)
	if err != nil {
		return err
	}
	amaagb.sql = query
	return amaagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (amaagb *AsMarketAppAlertGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := amaagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (amaagb *AsMarketAppAlertGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(amaagb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppAlertGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := amaagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (amaagb *AsMarketAppAlertGroupBy) StringsX(ctx context.Context) []string {
	v, err := amaagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amaagb *AsMarketAppAlertGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = amaagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappalert.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppAlertGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (amaagb *AsMarketAppAlertGroupBy) StringX(ctx context.Context) string {
	v, err := amaagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (amaagb *AsMarketAppAlertGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(amaagb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppAlertGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := amaagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (amaagb *AsMarketAppAlertGroupBy) IntsX(ctx context.Context) []int {
	v, err := amaagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amaagb *AsMarketAppAlertGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = amaagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappalert.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppAlertGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (amaagb *AsMarketAppAlertGroupBy) IntX(ctx context.Context) int {
	v, err := amaagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (amaagb *AsMarketAppAlertGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(amaagb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppAlertGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := amaagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (amaagb *AsMarketAppAlertGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := amaagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amaagb *AsMarketAppAlertGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = amaagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappalert.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppAlertGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (amaagb *AsMarketAppAlertGroupBy) Float64X(ctx context.Context) float64 {
	v, err := amaagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (amaagb *AsMarketAppAlertGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(amaagb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppAlertGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := amaagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (amaagb *AsMarketAppAlertGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := amaagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amaagb *AsMarketAppAlertGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = amaagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappalert.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppAlertGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (amaagb *AsMarketAppAlertGroupBy) BoolX(ctx context.Context) bool {
	v, err := amaagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (amaagb *AsMarketAppAlertGroupBy) Int64s(ctx context.Context) ([]int64, error) {
	if len(amaagb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppAlertGroupBy.Int64s is not achievable when grouping more than 1 field")
	}
	var v []int64
	if err := amaagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (amaagb *AsMarketAppAlertGroupBy) Int64sX(ctx context.Context) []int64 {
	v, err := amaagb.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amaagb *AsMarketAppAlertGroupBy) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = amaagb.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappalert.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppAlertGroupBy.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (amaagb *AsMarketAppAlertGroupBy) Int64X(ctx context.Context) int64 {
	v, err := amaagb.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (amaagb *AsMarketAppAlertGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range amaagb.fields {
		if !asmarketappalert.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := amaagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := amaagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (amaagb *AsMarketAppAlertGroupBy) sqlQuery() *sql.Selector {
	selector := amaagb.sql.Select()
	aggregation := make([]string, 0, len(amaagb.fns))
	for _, fn := range amaagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(amaagb.fields)+len(amaagb.fns))
		for _, f := range amaagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(amaagb.fields...)...)
}

// AsMarketAppAlertSelect is the builder for selecting fields of AsMarketAppAlert entities.
type AsMarketAppAlertSelect struct {
	*AsMarketAppAlertQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (amaas *AsMarketAppAlertSelect) Scan(ctx context.Context, v interface{}) error {
	if err := amaas.prepareQuery(ctx); err != nil {
		return err
	}
	amaas.sql = amaas.AsMarketAppAlertQuery.sqlQuery(ctx)
	return amaas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (amaas *AsMarketAppAlertSelect) ScanX(ctx context.Context, v interface{}) {
	if err := amaas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (amaas *AsMarketAppAlertSelect) Strings(ctx context.Context) ([]string, error) {
	if len(amaas.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppAlertSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := amaas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (amaas *AsMarketAppAlertSelect) StringsX(ctx context.Context) []string {
	v, err := amaas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (amaas *AsMarketAppAlertSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = amaas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappalert.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppAlertSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (amaas *AsMarketAppAlertSelect) StringX(ctx context.Context) string {
	v, err := amaas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (amaas *AsMarketAppAlertSelect) Ints(ctx context.Context) ([]int, error) {
	if len(amaas.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppAlertSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := amaas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (amaas *AsMarketAppAlertSelect) IntsX(ctx context.Context) []int {
	v, err := amaas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (amaas *AsMarketAppAlertSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = amaas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappalert.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppAlertSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (amaas *AsMarketAppAlertSelect) IntX(ctx context.Context) int {
	v, err := amaas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (amaas *AsMarketAppAlertSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(amaas.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppAlertSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := amaas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (amaas *AsMarketAppAlertSelect) Float64sX(ctx context.Context) []float64 {
	v, err := amaas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (amaas *AsMarketAppAlertSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = amaas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappalert.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppAlertSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (amaas *AsMarketAppAlertSelect) Float64X(ctx context.Context) float64 {
	v, err := amaas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (amaas *AsMarketAppAlertSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(amaas.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppAlertSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := amaas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (amaas *AsMarketAppAlertSelect) BoolsX(ctx context.Context) []bool {
	v, err := amaas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (amaas *AsMarketAppAlertSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = amaas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappalert.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppAlertSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (amaas *AsMarketAppAlertSelect) BoolX(ctx context.Context) bool {
	v, err := amaas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from a selector. It is only allowed when selecting one field.
func (amaas *AsMarketAppAlertSelect) Int64s(ctx context.Context) ([]int64, error) {
	if len(amaas.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppAlertSelect.Int64s is not achievable when selecting more than 1 field")
	}
	var v []int64
	if err := amaas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (amaas *AsMarketAppAlertSelect) Int64sX(ctx context.Context) []int64 {
	v, err := amaas.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a selector. It is only allowed when selecting one field.
func (amaas *AsMarketAppAlertSelect) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = amaas.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappalert.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppAlertSelect.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (amaas *AsMarketAppAlertSelect) Int64X(ctx context.Context) int64 {
	v, err := amaas.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (amaas *AsMarketAppAlertSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := amaas.sql.Query()
	if err := amaas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}