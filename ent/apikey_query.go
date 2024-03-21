// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Encedeus/panel/ent/apikey"
	"github.com/Encedeus/panel/ent/predicate"
	"github.com/Encedeus/panel/ent/user"
	"github.com/google/uuid"
)

// ApiKeyQuery is the builder for querying ApiKey entities.
type ApiKeyQuery struct {
	config
	ctx        *QueryContext
	order      []apikey.OrderOption
	inters     []Interceptor
	predicates []predicate.ApiKey
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ApiKeyQuery builder.
func (akq *ApiKeyQuery) Where(ps ...predicate.ApiKey) *ApiKeyQuery {
	akq.predicates = append(akq.predicates, ps...)
	return akq
}

// Limit the number of records to be returned by this query.
func (akq *ApiKeyQuery) Limit(limit int) *ApiKeyQuery {
	akq.ctx.Limit = &limit
	return akq
}

// Offset to start from.
func (akq *ApiKeyQuery) Offset(offset int) *ApiKeyQuery {
	akq.ctx.Offset = &offset
	return akq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (akq *ApiKeyQuery) Unique(unique bool) *ApiKeyQuery {
	akq.ctx.Unique = &unique
	return akq
}

// Order specifies how the records should be ordered.
func (akq *ApiKeyQuery) Order(o ...apikey.OrderOption) *ApiKeyQuery {
	akq.order = append(akq.order, o...)
	return akq
}

// QueryUser chains the current query on the "user" edge.
func (akq *ApiKeyQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: akq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := akq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := akq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(apikey.Table, apikey.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, apikey.UserTable, apikey.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(akq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ApiKey entity from the query.
// Returns a *NotFoundError when no ApiKey was found.
func (akq *ApiKeyQuery) First(ctx context.Context) (*ApiKey, error) {
	nodes, err := akq.Limit(1).All(setContextOp(ctx, akq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{apikey.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (akq *ApiKeyQuery) FirstX(ctx context.Context) *ApiKey {
	node, err := akq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ApiKey ID from the query.
// Returns a *NotFoundError when no ApiKey ID was found.
func (akq *ApiKeyQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = akq.Limit(1).IDs(setContextOp(ctx, akq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{apikey.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (akq *ApiKeyQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := akq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ApiKey entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ApiKey entity is found.
// Returns a *NotFoundError when no ApiKey entities are found.
func (akq *ApiKeyQuery) Only(ctx context.Context) (*ApiKey, error) {
	nodes, err := akq.Limit(2).All(setContextOp(ctx, akq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{apikey.Label}
	default:
		return nil, &NotSingularError{apikey.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (akq *ApiKeyQuery) OnlyX(ctx context.Context) *ApiKey {
	node, err := akq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ApiKey ID in the query.
// Returns a *NotSingularError when more than one ApiKey ID is found.
// Returns a *NotFoundError when no entities are found.
func (akq *ApiKeyQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = akq.Limit(2).IDs(setContextOp(ctx, akq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{apikey.Label}
	default:
		err = &NotSingularError{apikey.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (akq *ApiKeyQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := akq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ApiKeys.
func (akq *ApiKeyQuery) All(ctx context.Context) ([]*ApiKey, error) {
	ctx = setContextOp(ctx, akq.ctx, "All")
	if err := akq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ApiKey, *ApiKeyQuery]()
	return withInterceptors[[]*ApiKey](ctx, akq, qr, akq.inters)
}

// AllX is like All, but panics if an error occurs.
func (akq *ApiKeyQuery) AllX(ctx context.Context) []*ApiKey {
	nodes, err := akq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ApiKey IDs.
func (akq *ApiKeyQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if akq.ctx.Unique == nil && akq.path != nil {
		akq.Unique(true)
	}
	ctx = setContextOp(ctx, akq.ctx, "IDs")
	if err = akq.Select(apikey.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (akq *ApiKeyQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := akq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (akq *ApiKeyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, akq.ctx, "Count")
	if err := akq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, akq, querierCount[*ApiKeyQuery](), akq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (akq *ApiKeyQuery) CountX(ctx context.Context) int {
	count, err := akq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (akq *ApiKeyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, akq.ctx, "Exist")
	switch _, err := akq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (akq *ApiKeyQuery) ExistX(ctx context.Context) bool {
	exist, err := akq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ApiKeyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (akq *ApiKeyQuery) Clone() *ApiKeyQuery {
	if akq == nil {
		return nil
	}
	return &ApiKeyQuery{
		config:     akq.config,
		ctx:        akq.ctx.Clone(),
		order:      append([]apikey.OrderOption{}, akq.order...),
		inters:     append([]Interceptor{}, akq.inters...),
		predicates: append([]predicate.ApiKey{}, akq.predicates...),
		withUser:   akq.withUser.Clone(),
		// clone intermediate query.
		sql:  akq.sql.Clone(),
		path: akq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (akq *ApiKeyQuery) WithUser(opts ...func(*UserQuery)) *ApiKeyQuery {
	query := (&UserClient{config: akq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	akq.withUser = query
	return akq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ApiKey.Query().
//		GroupBy(apikey.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (akq *ApiKeyQuery) GroupBy(field string, fields ...string) *ApiKeyGroupBy {
	akq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ApiKeyGroupBy{build: akq}
	grbuild.flds = &akq.ctx.Fields
	grbuild.label = apikey.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.ApiKey.Query().
//		Select(apikey.FieldCreatedAt).
//		Scan(ctx, &v)
func (akq *ApiKeyQuery) Select(fields ...string) *ApiKeySelect {
	akq.ctx.Fields = append(akq.ctx.Fields, fields...)
	sbuild := &ApiKeySelect{ApiKeyQuery: akq}
	sbuild.label = apikey.Label
	sbuild.flds, sbuild.scan = &akq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ApiKeySelect configured with the given aggregations.
func (akq *ApiKeyQuery) Aggregate(fns ...AggregateFunc) *ApiKeySelect {
	return akq.Select().Aggregate(fns...)
}

func (akq *ApiKeyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range akq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, akq); err != nil {
				return err
			}
		}
	}
	for _, f := range akq.ctx.Fields {
		if !apikey.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if akq.path != nil {
		prev, err := akq.path(ctx)
		if err != nil {
			return err
		}
		akq.sql = prev
	}
	return nil
}

func (akq *ApiKeyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ApiKey, error) {
	var (
		nodes       = []*ApiKey{}
		_spec       = akq.querySpec()
		loadedTypes = [1]bool{
			akq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ApiKey).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ApiKey{config: akq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, akq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := akq.withUser; query != nil {
		if err := akq.loadUser(ctx, query, nodes, nil,
			func(n *ApiKey, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (akq *ApiKeyQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*ApiKey, init func(*ApiKey), assign func(*ApiKey, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ApiKey)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (akq *ApiKeyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := akq.querySpec()
	_spec.Node.Columns = akq.ctx.Fields
	if len(akq.ctx.Fields) > 0 {
		_spec.Unique = akq.ctx.Unique != nil && *akq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, akq.driver, _spec)
}

func (akq *ApiKeyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(apikey.Table, apikey.Columns, sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeUUID))
	_spec.From = akq.sql
	if unique := akq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if akq.path != nil {
		_spec.Unique = true
	}
	if fields := akq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apikey.FieldID)
		for i := range fields {
			if fields[i] != apikey.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if akq.withUser != nil {
			_spec.Node.AddColumnOnce(apikey.FieldUserID)
		}
	}
	if ps := akq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := akq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := akq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := akq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (akq *ApiKeyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(akq.driver.Dialect())
	t1 := builder.Table(apikey.Table)
	columns := akq.ctx.Fields
	if len(columns) == 0 {
		columns = apikey.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if akq.sql != nil {
		selector = akq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if akq.ctx.Unique != nil && *akq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range akq.predicates {
		p(selector)
	}
	for _, p := range akq.order {
		p(selector)
	}
	if offset := akq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := akq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ApiKeyGroupBy is the group-by builder for ApiKey entities.
type ApiKeyGroupBy struct {
	selector
	build *ApiKeyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (akgb *ApiKeyGroupBy) Aggregate(fns ...AggregateFunc) *ApiKeyGroupBy {
	akgb.fns = append(akgb.fns, fns...)
	return akgb
}

// Scan applies the selector query and scans the result into the given value.
func (akgb *ApiKeyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, akgb.build.ctx, "GroupBy")
	if err := akgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ApiKeyQuery, *ApiKeyGroupBy](ctx, akgb.build, akgb, akgb.build.inters, v)
}

func (akgb *ApiKeyGroupBy) sqlScan(ctx context.Context, root *ApiKeyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(akgb.fns))
	for _, fn := range akgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*akgb.flds)+len(akgb.fns))
		for _, f := range *akgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*akgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := akgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ApiKeySelect is the builder for selecting fields of ApiKey entities.
type ApiKeySelect struct {
	*ApiKeyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aks *ApiKeySelect) Aggregate(fns ...AggregateFunc) *ApiKeySelect {
	aks.fns = append(aks.fns, fns...)
	return aks
}

// Scan applies the selector query and scans the result into the given value.
func (aks *ApiKeySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aks.ctx, "Select")
	if err := aks.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ApiKeyQuery, *ApiKeySelect](ctx, aks.ApiKeyQuery, aks, aks.inters, v)
}

func (aks *ApiKeySelect) sqlScan(ctx context.Context, root *ApiKeyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aks.fns))
	for _, fn := range aks.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aks.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
