// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/ent/file"
	"github.com/facebook/ent/entc/integration/ent/filetype"
	"github.com/facebook/ent/entc/integration/ent/predicate"
	"github.com/facebook/ent/schema/field"
)

// FileTypeQuery is the builder for querying FileType entities.
type FileTypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.FileType
	// eager-loading edges.
	withFiles *FileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FileTypeQuery builder.
func (ftq *FileTypeQuery) Where(ps ...predicate.FileType) *FileTypeQuery {
	ftq.predicates = append(ftq.predicates, ps...)
	return ftq
}

// Limit adds a limit step to the query.
func (ftq *FileTypeQuery) Limit(limit int) *FileTypeQuery {
	ftq.limit = &limit
	return ftq
}

// Offset adds an offset step to the query.
func (ftq *FileTypeQuery) Offset(offset int) *FileTypeQuery {
	ftq.offset = &offset
	return ftq
}

// Order adds an order step to the query.
func (ftq *FileTypeQuery) Order(o ...OrderFunc) *FileTypeQuery {
	ftq.order = append(ftq.order, o...)
	return ftq
}

// QueryFiles chains the current query on the "files" edge.
func (ftq *FileTypeQuery) QueryFiles() *FileQuery {
	query := &FileQuery{config: ftq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ftq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ftq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(filetype.Table, filetype.FieldID, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, filetype.FilesTable, filetype.FilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(ftq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FileType entity from the query.
// Returns a *NotFoundError when no FileType was found.
func (ftq *FileTypeQuery) First(ctx context.Context) (*FileType, error) {
	nodes, err := ftq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{filetype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ftq *FileTypeQuery) FirstX(ctx context.Context) *FileType {
	node, err := ftq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FileType ID from the query.
// Returns a *NotFoundError when no FileType ID was found.
func (ftq *FileTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ftq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{filetype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ftq *FileTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := ftq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FileType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one FileType entity is not found.
// Returns a *NotFoundError when no FileType entities are found.
func (ftq *FileTypeQuery) Only(ctx context.Context) (*FileType, error) {
	nodes, err := ftq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{filetype.Label}
	default:
		return nil, &NotSingularError{filetype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ftq *FileTypeQuery) OnlyX(ctx context.Context) *FileType {
	node, err := ftq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FileType ID in the query.
// Returns a *NotSingularError when exactly one FileType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ftq *FileTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ftq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{filetype.Label}
	default:
		err = &NotSingularError{filetype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ftq *FileTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := ftq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FileTypes.
func (ftq *FileTypeQuery) All(ctx context.Context) ([]*FileType, error) {
	if err := ftq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ftq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ftq *FileTypeQuery) AllX(ctx context.Context) []*FileType {
	nodes, err := ftq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FileType IDs.
func (ftq *FileTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ftq.Select(filetype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ftq *FileTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := ftq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ftq *FileTypeQuery) Count(ctx context.Context) (int, error) {
	if err := ftq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ftq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ftq *FileTypeQuery) CountX(ctx context.Context) int {
	count, err := ftq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ftq *FileTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := ftq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ftq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ftq *FileTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ftq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FileTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ftq *FileTypeQuery) Clone() *FileTypeQuery {
	if ftq == nil {
		return nil
	}
	return &FileTypeQuery{
		config:     ftq.config,
		limit:      ftq.limit,
		offset:     ftq.offset,
		order:      append([]OrderFunc{}, ftq.order...),
		predicates: append([]predicate.FileType{}, ftq.predicates...),
		withFiles:  ftq.withFiles.Clone(),
		// clone intermediate query.
		sql:  ftq.sql.Clone(),
		path: ftq.path,
	}
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (ftq *FileTypeQuery) WithFiles(opts ...func(*FileQuery)) *FileTypeQuery {
	query := &FileQuery{config: ftq.config}
	for _, opt := range opts {
		opt(query)
	}
	ftq.withFiles = query
	return ftq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FileType.Query().
//		GroupBy(filetype.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ftq *FileTypeQuery) GroupBy(field string, fields ...string) *FileTypeGroupBy {
	group := &FileTypeGroupBy{config: ftq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ftq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ftq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.FileType.Query().
//		Select(filetype.FieldName).
//		Scan(ctx, &v)
//
func (ftq *FileTypeQuery) Select(field string, fields ...string) *FileTypeSelect {
	ftq.fields = append([]string{field}, fields...)
	return &FileTypeSelect{FileTypeQuery: ftq}
}

func (ftq *FileTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ftq.fields {
		if !filetype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ftq.path != nil {
		prev, err := ftq.path(ctx)
		if err != nil {
			return err
		}
		ftq.sql = prev
	}
	return nil
}

func (ftq *FileTypeQuery) sqlAll(ctx context.Context) ([]*FileType, error) {
	var (
		nodes       = []*FileType{}
		_spec       = ftq.querySpec()
		loadedTypes = [1]bool{
			ftq.withFiles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &FileType{config: ftq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ftq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ftq.withFiles; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*FileType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Files = []*File{}
		}
		query.withFKs = true
		query.Where(predicate.File(func(s *sql.Selector) {
			s.Where(sql.InValues(filetype.FilesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.file_type_files
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "file_type_files" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "file_type_files" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Files = append(node.Edges.Files, n)
		}
	}

	return nodes, nil
}

func (ftq *FileTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ftq.querySpec()
	return sqlgraph.CountNodes(ctx, ftq.driver, _spec)
}

func (ftq *FileTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ftq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ftq *FileTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   filetype.Table,
			Columns: filetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: filetype.FieldID,
			},
		},
		From:   ftq.sql,
		Unique: true,
	}
	if fields := ftq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, filetype.FieldID)
		for i := range fields {
			if fields[i] != filetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ftq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ftq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ftq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ftq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, filetype.ValidColumn)
			}
		}
	}
	return _spec
}

func (ftq *FileTypeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ftq.driver.Dialect())
	t1 := builder.Table(filetype.Table)
	selector := builder.Select(t1.Columns(filetype.Columns...)...).From(t1)
	if ftq.sql != nil {
		selector = ftq.sql
		selector.Select(selector.Columns(filetype.Columns...)...)
	}
	for _, p := range ftq.predicates {
		p(selector)
	}
	for _, p := range ftq.order {
		p(selector, filetype.ValidColumn)
	}
	if offset := ftq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ftq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FileTypeGroupBy is the group-by builder for FileType entities.
type FileTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ftgb *FileTypeGroupBy) Aggregate(fns ...AggregateFunc) *FileTypeGroupBy {
	ftgb.fns = append(ftgb.fns, fns...)
	return ftgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ftgb *FileTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ftgb.path(ctx)
	if err != nil {
		return err
	}
	ftgb.sql = query
	return ftgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ftgb *FileTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ftgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ftgb *FileTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FileTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ftgb *FileTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := ftgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ftgb *FileTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ftgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{filetype.Label}
	default:
		err = fmt.Errorf("ent: FileTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ftgb *FileTypeGroupBy) StringX(ctx context.Context) string {
	v, err := ftgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ftgb *FileTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FileTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ftgb *FileTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := ftgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ftgb *FileTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ftgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{filetype.Label}
	default:
		err = fmt.Errorf("ent: FileTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ftgb *FileTypeGroupBy) IntX(ctx context.Context) int {
	v, err := ftgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ftgb *FileTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FileTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ftgb *FileTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ftgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ftgb *FileTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ftgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{filetype.Label}
	default:
		err = fmt.Errorf("ent: FileTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ftgb *FileTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ftgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ftgb *FileTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FileTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ftgb *FileTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ftgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ftgb *FileTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ftgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{filetype.Label}
	default:
		err = fmt.Errorf("ent: FileTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ftgb *FileTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := ftgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ftgb *FileTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ftgb.fields {
		if !filetype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ftgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ftgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ftgb *FileTypeGroupBy) sqlQuery() *sql.Selector {
	selector := ftgb.sql
	columns := make([]string, 0, len(ftgb.fields)+len(ftgb.fns))
	columns = append(columns, ftgb.fields...)
	for _, fn := range ftgb.fns {
		columns = append(columns, fn(selector, filetype.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(ftgb.fields...)
}

// FileTypeSelect is the builder for selecting fields of FileType entities.
type FileTypeSelect struct {
	*FileTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fts *FileTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fts.prepareQuery(ctx); err != nil {
		return err
	}
	fts.sql = fts.FileTypeQuery.sqlQuery()
	return fts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fts *FileTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (fts *FileTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FileTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fts *FileTypeSelect) StringsX(ctx context.Context) []string {
	v, err := fts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (fts *FileTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{filetype.Label}
	default:
		err = fmt.Errorf("ent: FileTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fts *FileTypeSelect) StringX(ctx context.Context) string {
	v, err := fts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (fts *FileTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FileTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fts *FileTypeSelect) IntsX(ctx context.Context) []int {
	v, err := fts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (fts *FileTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{filetype.Label}
	default:
		err = fmt.Errorf("ent: FileTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fts *FileTypeSelect) IntX(ctx context.Context) int {
	v, err := fts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (fts *FileTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FileTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fts *FileTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (fts *FileTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{filetype.Label}
	default:
		err = fmt.Errorf("ent: FileTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fts *FileTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := fts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (fts *FileTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FileTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fts *FileTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := fts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (fts *FileTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{filetype.Label}
	default:
		err = fmt.Errorf("ent: FileTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fts *FileTypeSelect) BoolX(ctx context.Context) bool {
	v, err := fts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fts *FileTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fts.sqlQuery().Query()
	if err := fts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fts *FileTypeSelect) sqlQuery() sql.Querier {
	selector := fts.sql
	selector.Select(selector.Columns(fts.fields...)...)
	return selector
}
