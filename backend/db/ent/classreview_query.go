// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
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

// ClassReviewQuery is the builder for querying ClassReview entities.
type ClassReviewQuery struct {
	config
	ctx                  *QueryContext
	order                []classreview.OrderOption
	inters               []Interceptor
	predicates           []predicate.ClassReview
	withUsers            *UserQuery
	withClassReviewLikes *ClassReviewLikeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ClassReviewQuery builder.
func (crq *ClassReviewQuery) Where(ps ...predicate.ClassReview) *ClassReviewQuery {
	crq.predicates = append(crq.predicates, ps...)
	return crq
}

// Limit the number of records to be returned by this query.
func (crq *ClassReviewQuery) Limit(limit int) *ClassReviewQuery {
	crq.ctx.Limit = &limit
	return crq
}

// Offset to start from.
func (crq *ClassReviewQuery) Offset(offset int) *ClassReviewQuery {
	crq.ctx.Offset = &offset
	return crq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (crq *ClassReviewQuery) Unique(unique bool) *ClassReviewQuery {
	crq.ctx.Unique = &unique
	return crq
}

// Order specifies how the records should be ordered.
func (crq *ClassReviewQuery) Order(o ...classreview.OrderOption) *ClassReviewQuery {
	crq.order = append(crq.order, o...)
	return crq
}

// QueryUsers chains the current query on the "users" edge.
func (crq *ClassReviewQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: crq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := crq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(classreview.Table, classreview.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, classreview.UsersTable, classreview.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(crq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryClassReviewLikes chains the current query on the "class_review_likes" edge.
func (crq *ClassReviewQuery) QueryClassReviewLikes() *ClassReviewLikeQuery {
	query := (&ClassReviewLikeClient{config: crq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := crq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := crq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(classreview.Table, classreview.FieldID, selector),
			sqlgraph.To(classreviewlike.Table, classreviewlike.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, classreview.ClassReviewLikesTable, classreview.ClassReviewLikesColumn),
		)
		fromU = sqlgraph.SetNeighbors(crq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ClassReview entity from the query.
// Returns a *NotFoundError when no ClassReview was found.
func (crq *ClassReviewQuery) First(ctx context.Context) (*ClassReview, error) {
	nodes, err := crq.Limit(1).All(setContextOp(ctx, crq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{classreview.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (crq *ClassReviewQuery) FirstX(ctx context.Context) *ClassReview {
	node, err := crq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ClassReview ID from the query.
// Returns a *NotFoundError when no ClassReview ID was found.
func (crq *ClassReviewQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = crq.Limit(1).IDs(setContextOp(ctx, crq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{classreview.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (crq *ClassReviewQuery) FirstIDX(ctx context.Context) string {
	id, err := crq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ClassReview entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ClassReview entity is found.
// Returns a *NotFoundError when no ClassReview entities are found.
func (crq *ClassReviewQuery) Only(ctx context.Context) (*ClassReview, error) {
	nodes, err := crq.Limit(2).All(setContextOp(ctx, crq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{classreview.Label}
	default:
		return nil, &NotSingularError{classreview.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (crq *ClassReviewQuery) OnlyX(ctx context.Context) *ClassReview {
	node, err := crq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ClassReview ID in the query.
// Returns a *NotSingularError when more than one ClassReview ID is found.
// Returns a *NotFoundError when no entities are found.
func (crq *ClassReviewQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = crq.Limit(2).IDs(setContextOp(ctx, crq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{classreview.Label}
	default:
		err = &NotSingularError{classreview.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (crq *ClassReviewQuery) OnlyIDX(ctx context.Context) string {
	id, err := crq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ClassReviews.
func (crq *ClassReviewQuery) All(ctx context.Context) ([]*ClassReview, error) {
	ctx = setContextOp(ctx, crq.ctx, "All")
	if err := crq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ClassReview, *ClassReviewQuery]()
	return withInterceptors[[]*ClassReview](ctx, crq, qr, crq.inters)
}

// AllX is like All, but panics if an error occurs.
func (crq *ClassReviewQuery) AllX(ctx context.Context) []*ClassReview {
	nodes, err := crq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ClassReview IDs.
func (crq *ClassReviewQuery) IDs(ctx context.Context) (ids []string, err error) {
	if crq.ctx.Unique == nil && crq.path != nil {
		crq.Unique(true)
	}
	ctx = setContextOp(ctx, crq.ctx, "IDs")
	if err = crq.Select(classreview.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (crq *ClassReviewQuery) IDsX(ctx context.Context) []string {
	ids, err := crq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (crq *ClassReviewQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, crq.ctx, "Count")
	if err := crq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, crq, querierCount[*ClassReviewQuery](), crq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (crq *ClassReviewQuery) CountX(ctx context.Context) int {
	count, err := crq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (crq *ClassReviewQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, crq.ctx, "Exist")
	switch _, err := crq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (crq *ClassReviewQuery) ExistX(ctx context.Context) bool {
	exist, err := crq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ClassReviewQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (crq *ClassReviewQuery) Clone() *ClassReviewQuery {
	if crq == nil {
		return nil
	}
	return &ClassReviewQuery{
		config:               crq.config,
		ctx:                  crq.ctx.Clone(),
		order:                append([]classreview.OrderOption{}, crq.order...),
		inters:               append([]Interceptor{}, crq.inters...),
		predicates:           append([]predicate.ClassReview{}, crq.predicates...),
		withUsers:            crq.withUsers.Clone(),
		withClassReviewLikes: crq.withClassReviewLikes.Clone(),
		// clone intermediate query.
		sql:  crq.sql.Clone(),
		path: crq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (crq *ClassReviewQuery) WithUsers(opts ...func(*UserQuery)) *ClassReviewQuery {
	query := (&UserClient{config: crq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	crq.withUsers = query
	return crq
}

// WithClassReviewLikes tells the query-builder to eager-load the nodes that are connected to
// the "class_review_likes" edge. The optional arguments are used to configure the query builder of the edge.
func (crq *ClassReviewQuery) WithClassReviewLikes(opts ...func(*ClassReviewLikeQuery)) *ClassReviewQuery {
	query := (&ClassReviewLikeClient{config: crq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	crq.withClassReviewLikes = query
	return crq
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
//	client.ClassReview.Query().
//		GroupBy(classreview.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (crq *ClassReviewQuery) GroupBy(field string, fields ...string) *ClassReviewGroupBy {
	crq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ClassReviewGroupBy{build: crq}
	grbuild.flds = &crq.ctx.Fields
	grbuild.label = classreview.Label
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
//	client.ClassReview.Query().
//		Select(classreview.FieldUserID).
//		Scan(ctx, &v)
func (crq *ClassReviewQuery) Select(fields ...string) *ClassReviewSelect {
	crq.ctx.Fields = append(crq.ctx.Fields, fields...)
	sbuild := &ClassReviewSelect{ClassReviewQuery: crq}
	sbuild.label = classreview.Label
	sbuild.flds, sbuild.scan = &crq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ClassReviewSelect configured with the given aggregations.
func (crq *ClassReviewQuery) Aggregate(fns ...AggregateFunc) *ClassReviewSelect {
	return crq.Select().Aggregate(fns...)
}

func (crq *ClassReviewQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range crq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, crq); err != nil {
				return err
			}
		}
	}
	for _, f := range crq.ctx.Fields {
		if !classreview.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if crq.path != nil {
		prev, err := crq.path(ctx)
		if err != nil {
			return err
		}
		crq.sql = prev
	}
	return nil
}

func (crq *ClassReviewQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ClassReview, error) {
	var (
		nodes       = []*ClassReview{}
		_spec       = crq.querySpec()
		loadedTypes = [2]bool{
			crq.withUsers != nil,
			crq.withClassReviewLikes != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ClassReview).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ClassReview{config: crq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, crq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := crq.withUsers; query != nil {
		if err := crq.loadUsers(ctx, query, nodes, nil,
			func(n *ClassReview, e *User) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	if query := crq.withClassReviewLikes; query != nil {
		if err := crq.loadClassReviewLikes(ctx, query, nodes,
			func(n *ClassReview) { n.Edges.ClassReviewLikes = []*ClassReviewLike{} },
			func(n *ClassReview, e *ClassReviewLike) {
				n.Edges.ClassReviewLikes = append(n.Edges.ClassReviewLikes, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (crq *ClassReviewQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*ClassReview, init func(*ClassReview), assign func(*ClassReview, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ClassReview)
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
func (crq *ClassReviewQuery) loadClassReviewLikes(ctx context.Context, query *ClassReviewLikeQuery, nodes []*ClassReview, init func(*ClassReview), assign func(*ClassReview, *ClassReviewLike)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*ClassReview)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(classreviewlike.FieldClassReviewID)
	}
	query.Where(predicate.ClassReviewLike(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(classreview.ClassReviewLikesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ClassReviewID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "class_review_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (crq *ClassReviewQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := crq.querySpec()
	_spec.Node.Columns = crq.ctx.Fields
	if len(crq.ctx.Fields) > 0 {
		_spec.Unique = crq.ctx.Unique != nil && *crq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, crq.driver, _spec)
}

func (crq *ClassReviewQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(classreview.Table, classreview.Columns, sqlgraph.NewFieldSpec(classreview.FieldID, field.TypeString))
	_spec.From = crq.sql
	if unique := crq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if crq.path != nil {
		_spec.Unique = true
	}
	if fields := crq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, classreview.FieldID)
		for i := range fields {
			if fields[i] != classreview.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if crq.withUsers != nil {
			_spec.Node.AddColumnOnce(classreview.FieldUserID)
		}
	}
	if ps := crq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := crq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := crq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := crq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (crq *ClassReviewQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(crq.driver.Dialect())
	t1 := builder.Table(classreview.Table)
	columns := crq.ctx.Fields
	if len(columns) == 0 {
		columns = classreview.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if crq.sql != nil {
		selector = crq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if crq.ctx.Unique != nil && *crq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range crq.predicates {
		p(selector)
	}
	for _, p := range crq.order {
		p(selector)
	}
	if offset := crq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := crq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ClassReviewGroupBy is the group-by builder for ClassReview entities.
type ClassReviewGroupBy struct {
	selector
	build *ClassReviewQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (crgb *ClassReviewGroupBy) Aggregate(fns ...AggregateFunc) *ClassReviewGroupBy {
	crgb.fns = append(crgb.fns, fns...)
	return crgb
}

// Scan applies the selector query and scans the result into the given value.
func (crgb *ClassReviewGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, crgb.build.ctx, "GroupBy")
	if err := crgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClassReviewQuery, *ClassReviewGroupBy](ctx, crgb.build, crgb, crgb.build.inters, v)
}

func (crgb *ClassReviewGroupBy) sqlScan(ctx context.Context, root *ClassReviewQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(crgb.fns))
	for _, fn := range crgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*crgb.flds)+len(crgb.fns))
		for _, f := range *crgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*crgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ClassReviewSelect is the builder for selecting fields of ClassReview entities.
type ClassReviewSelect struct {
	*ClassReviewQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (crs *ClassReviewSelect) Aggregate(fns ...AggregateFunc) *ClassReviewSelect {
	crs.fns = append(crs.fns, fns...)
	return crs
}

// Scan applies the selector query and scans the result into the given value.
func (crs *ClassReviewSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, crs.ctx, "Select")
	if err := crs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClassReviewQuery, *ClassReviewSelect](ctx, crs.ClassReviewQuery, crs, crs.inters, v)
}

func (crs *ClassReviewSelect) sqlScan(ctx context.Context, root *ClassReviewQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(crs.fns))
	for _, fn := range crs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*crs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := crs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
