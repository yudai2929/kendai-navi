// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yudai2929/kendai-navi/backend/db/ent/classreview"
	"github.com/yudai2929/kendai-navi/backend/db/ent/classreviewlike"
	"github.com/yudai2929/kendai-navi/backend/db/ent/predicate"
	"github.com/yudai2929/kendai-navi/backend/db/ent/user"
)

// ClassReviewLikeQuery is the builder for querying ClassReviewLike entities.
type ClassReviewLikeQuery struct {
	config
	ctx              *QueryContext
	order            []classreviewlike.OrderOption
	inters           []Interceptor
	predicates       []predicate.ClassReviewLike
	withUsers        *UserQuery
	withClassReviews *ClassReviewQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ClassReviewLikeQuery builder.
func (crlq *ClassReviewLikeQuery) Where(ps ...predicate.ClassReviewLike) *ClassReviewLikeQuery {
	crlq.predicates = append(crlq.predicates, ps...)
	return crlq
}

// Limit the number of records to be returned by this query.
func (crlq *ClassReviewLikeQuery) Limit(limit int) *ClassReviewLikeQuery {
	crlq.ctx.Limit = &limit
	return crlq
}

// Offset to start from.
func (crlq *ClassReviewLikeQuery) Offset(offset int) *ClassReviewLikeQuery {
	crlq.ctx.Offset = &offset
	return crlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (crlq *ClassReviewLikeQuery) Unique(unique bool) *ClassReviewLikeQuery {
	crlq.ctx.Unique = &unique
	return crlq
}

// Order specifies how the records should be ordered.
func (crlq *ClassReviewLikeQuery) Order(o ...classreviewlike.OrderOption) *ClassReviewLikeQuery {
	crlq.order = append(crlq.order, o...)
	return crlq
}

// QueryUsers chains the current query on the "users" edge.
func (crlq *ClassReviewLikeQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: crlq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := crlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := crlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(classreviewlike.Table, classreviewlike.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, classreviewlike.UsersTable, classreviewlike.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(crlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryClassReviews chains the current query on the "class_reviews" edge.
func (crlq *ClassReviewLikeQuery) QueryClassReviews() *ClassReviewQuery {
	query := (&ClassReviewClient{config: crlq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := crlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := crlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(classreviewlike.Table, classreviewlike.FieldID, selector),
			sqlgraph.To(classreview.Table, classreview.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, classreviewlike.ClassReviewsTable, classreviewlike.ClassReviewsColumn),
		)
		fromU = sqlgraph.SetNeighbors(crlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ClassReviewLike entity from the query.
// Returns a *NotFoundError when no ClassReviewLike was found.
func (crlq *ClassReviewLikeQuery) First(ctx context.Context) (*ClassReviewLike, error) {
	nodes, err := crlq.Limit(1).All(setContextOp(ctx, crlq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{classreviewlike.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (crlq *ClassReviewLikeQuery) FirstX(ctx context.Context) *ClassReviewLike {
	node, err := crlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ClassReviewLike ID from the query.
// Returns a *NotFoundError when no ClassReviewLike ID was found.
func (crlq *ClassReviewLikeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crlq.Limit(1).IDs(setContextOp(ctx, crlq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{classreviewlike.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (crlq *ClassReviewLikeQuery) FirstIDX(ctx context.Context) int {
	id, err := crlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ClassReviewLike entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ClassReviewLike entity is found.
// Returns a *NotFoundError when no ClassReviewLike entities are found.
func (crlq *ClassReviewLikeQuery) Only(ctx context.Context) (*ClassReviewLike, error) {
	nodes, err := crlq.Limit(2).All(setContextOp(ctx, crlq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{classreviewlike.Label}
	default:
		return nil, &NotSingularError{classreviewlike.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (crlq *ClassReviewLikeQuery) OnlyX(ctx context.Context) *ClassReviewLike {
	node, err := crlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ClassReviewLike ID in the query.
// Returns a *NotSingularError when more than one ClassReviewLike ID is found.
// Returns a *NotFoundError when no entities are found.
func (crlq *ClassReviewLikeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = crlq.Limit(2).IDs(setContextOp(ctx, crlq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{classreviewlike.Label}
	default:
		err = &NotSingularError{classreviewlike.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (crlq *ClassReviewLikeQuery) OnlyIDX(ctx context.Context) int {
	id, err := crlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ClassReviewLikes.
func (crlq *ClassReviewLikeQuery) All(ctx context.Context) ([]*ClassReviewLike, error) {
	ctx = setContextOp(ctx, crlq.ctx, "All")
	if err := crlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ClassReviewLike, *ClassReviewLikeQuery]()
	return withInterceptors[[]*ClassReviewLike](ctx, crlq, qr, crlq.inters)
}

// AllX is like All, but panics if an error occurs.
func (crlq *ClassReviewLikeQuery) AllX(ctx context.Context) []*ClassReviewLike {
	nodes, err := crlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ClassReviewLike IDs.
func (crlq *ClassReviewLikeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if crlq.ctx.Unique == nil && crlq.path != nil {
		crlq.Unique(true)
	}
	ctx = setContextOp(ctx, crlq.ctx, "IDs")
	if err = crlq.Select(classreviewlike.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (crlq *ClassReviewLikeQuery) IDsX(ctx context.Context) []int {
	ids, err := crlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (crlq *ClassReviewLikeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, crlq.ctx, "Count")
	if err := crlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, crlq, querierCount[*ClassReviewLikeQuery](), crlq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (crlq *ClassReviewLikeQuery) CountX(ctx context.Context) int {
	count, err := crlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (crlq *ClassReviewLikeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, crlq.ctx, "Exist")
	switch _, err := crlq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (crlq *ClassReviewLikeQuery) ExistX(ctx context.Context) bool {
	exist, err := crlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ClassReviewLikeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (crlq *ClassReviewLikeQuery) Clone() *ClassReviewLikeQuery {
	if crlq == nil {
		return nil
	}
	return &ClassReviewLikeQuery{
		config:           crlq.config,
		ctx:              crlq.ctx.Clone(),
		order:            append([]classreviewlike.OrderOption{}, crlq.order...),
		inters:           append([]Interceptor{}, crlq.inters...),
		predicates:       append([]predicate.ClassReviewLike{}, crlq.predicates...),
		withUsers:        crlq.withUsers.Clone(),
		withClassReviews: crlq.withClassReviews.Clone(),
		// clone intermediate query.
		sql:  crlq.sql.Clone(),
		path: crlq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (crlq *ClassReviewLikeQuery) WithUsers(opts ...func(*UserQuery)) *ClassReviewLikeQuery {
	query := (&UserClient{config: crlq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	crlq.withUsers = query
	return crlq
}

// WithClassReviews tells the query-builder to eager-load the nodes that are connected to
// the "class_reviews" edge. The optional arguments are used to configure the query builder of the edge.
func (crlq *ClassReviewLikeQuery) WithClassReviews(opts ...func(*ClassReviewQuery)) *ClassReviewLikeQuery {
	query := (&ClassReviewClient{config: crlq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	crlq.withClassReviews = query
	return crlq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ClassReviewLike.Query().
//		GroupBy(classreviewlike.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (crlq *ClassReviewLikeQuery) GroupBy(field string, fields ...string) *ClassReviewLikeGroupBy {
	crlq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ClassReviewLikeGroupBy{build: crlq}
	grbuild.flds = &crlq.ctx.Fields
	grbuild.label = classreviewlike.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//	}
//
//	client.ClassReviewLike.Query().
//		Select(classreviewlike.FieldUserID).
//		Scan(ctx, &v)
func (crlq *ClassReviewLikeQuery) Select(fields ...string) *ClassReviewLikeSelect {
	crlq.ctx.Fields = append(crlq.ctx.Fields, fields...)
	sbuild := &ClassReviewLikeSelect{ClassReviewLikeQuery: crlq}
	sbuild.label = classreviewlike.Label
	sbuild.flds, sbuild.scan = &crlq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ClassReviewLikeSelect configured with the given aggregations.
func (crlq *ClassReviewLikeQuery) Aggregate(fns ...AggregateFunc) *ClassReviewLikeSelect {
	return crlq.Select().Aggregate(fns...)
}

func (crlq *ClassReviewLikeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range crlq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, crlq); err != nil {
				return err
			}
		}
	}
	for _, f := range crlq.ctx.Fields {
		if !classreviewlike.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if crlq.path != nil {
		prev, err := crlq.path(ctx)
		if err != nil {
			return err
		}
		crlq.sql = prev
	}
	return nil
}

func (crlq *ClassReviewLikeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ClassReviewLike, error) {
	var (
		nodes       = []*ClassReviewLike{}
		_spec       = crlq.querySpec()
		loadedTypes = [2]bool{
			crlq.withUsers != nil,
			crlq.withClassReviews != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ClassReviewLike).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ClassReviewLike{config: crlq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, crlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := crlq.withUsers; query != nil {
		if err := crlq.loadUsers(ctx, query, nodes, nil,
			func(n *ClassReviewLike, e *User) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	if query := crlq.withClassReviews; query != nil {
		if err := crlq.loadClassReviews(ctx, query, nodes, nil,
			func(n *ClassReviewLike, e *ClassReview) { n.Edges.ClassReviews = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (crlq *ClassReviewLikeQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*ClassReviewLike, init func(*ClassReviewLike), assign func(*ClassReviewLike, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ClassReviewLike)
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
func (crlq *ClassReviewLikeQuery) loadClassReviews(ctx context.Context, query *ClassReviewQuery, nodes []*ClassReviewLike, init func(*ClassReviewLike), assign func(*ClassReviewLike, *ClassReview)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ClassReviewLike)
	for i := range nodes {
		fk := nodes[i].ClassReviewID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(classreview.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "class_review_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (crlq *ClassReviewLikeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := crlq.querySpec()
	_spec.Node.Columns = crlq.ctx.Fields
	if len(crlq.ctx.Fields) > 0 {
		_spec.Unique = crlq.ctx.Unique != nil && *crlq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, crlq.driver, _spec)
}

func (crlq *ClassReviewLikeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(classreviewlike.Table, classreviewlike.Columns, sqlgraph.NewFieldSpec(classreviewlike.FieldID, field.TypeInt))
	_spec.From = crlq.sql
	if unique := crlq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if crlq.path != nil {
		_spec.Unique = true
	}
	if fields := crlq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, classreviewlike.FieldID)
		for i := range fields {
			if fields[i] != classreviewlike.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if crlq.withUsers != nil {
			_spec.Node.AddColumnOnce(classreviewlike.FieldUserID)
		}
		if crlq.withClassReviews != nil {
			_spec.Node.AddColumnOnce(classreviewlike.FieldClassReviewID)
		}
	}
	if ps := crlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := crlq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := crlq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := crlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (crlq *ClassReviewLikeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(crlq.driver.Dialect())
	t1 := builder.Table(classreviewlike.Table)
	columns := crlq.ctx.Fields
	if len(columns) == 0 {
		columns = classreviewlike.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if crlq.sql != nil {
		selector = crlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if crlq.ctx.Unique != nil && *crlq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range crlq.predicates {
		p(selector)
	}
	for _, p := range crlq.order {
		p(selector)
	}
	if offset := crlq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := crlq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ClassReviewLikeGroupBy is the group-by builder for ClassReviewLike entities.
type ClassReviewLikeGroupBy struct {
	selector
	build *ClassReviewLikeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (crlgb *ClassReviewLikeGroupBy) Aggregate(fns ...AggregateFunc) *ClassReviewLikeGroupBy {
	crlgb.fns = append(crlgb.fns, fns...)
	return crlgb
}

// Scan applies the selector query and scans the result into the given value.
func (crlgb *ClassReviewLikeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, crlgb.build.ctx, "GroupBy")
	if err := crlgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClassReviewLikeQuery, *ClassReviewLikeGroupBy](ctx, crlgb.build, crlgb, crlgb.build.inters, v)
}

func (crlgb *ClassReviewLikeGroupBy) sqlScan(ctx context.Context, root *ClassReviewLikeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(crlgb.fns))
	for _, fn := range crlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*crlgb.flds)+len(crlgb.fns))
		for _, f := range *crlgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*crlgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crlgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ClassReviewLikeSelect is the builder for selecting fields of ClassReviewLike entities.
type ClassReviewLikeSelect struct {
	*ClassReviewLikeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (crls *ClassReviewLikeSelect) Aggregate(fns ...AggregateFunc) *ClassReviewLikeSelect {
	crls.fns = append(crls.fns, fns...)
	return crls
}

// Scan applies the selector query and scans the result into the given value.
func (crls *ClassReviewLikeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, crls.ctx, "Select")
	if err := crls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClassReviewLikeQuery, *ClassReviewLikeSelect](ctx, crls.ClassReviewLikeQuery, crls, crls.inters, v)
}

func (crls *ClassReviewLikeSelect) sqlScan(ctx context.Context, root *ClassReviewLikeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(crls.fns))
	for _, fn := range crls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*crls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}