// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/virtualmachine"
	"go.infratographer.com/x/gidx"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[gidx.PrefixedID]
	PageInfo       = entgql.PageInfo[gidx.PrefixedID]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// VirtualMachineEdge is the edge representation of VirtualMachine.
type VirtualMachineEdge struct {
	Node   *VirtualMachine `json:"node"`
	Cursor Cursor          `json:"cursor"`
}

// VirtualMachineConnection is the connection containing edges to VirtualMachine.
type VirtualMachineConnection struct {
	Edges      []*VirtualMachineEdge `json:"edges"`
	PageInfo   PageInfo              `json:"pageInfo"`
	TotalCount int                   `json:"totalCount"`
}

func (c *VirtualMachineConnection) build(nodes []*VirtualMachine, pager *virtualmachinePager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *VirtualMachine
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *VirtualMachine {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *VirtualMachine {
			return nodes[i]
		}
	}
	c.Edges = make([]*VirtualMachineEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &VirtualMachineEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// VirtualMachinePaginateOption enables pagination customization.
type VirtualMachinePaginateOption func(*virtualmachinePager) error

// WithVirtualMachineOrder configures pagination ordering.
func WithVirtualMachineOrder(order *VirtualMachineOrder) VirtualMachinePaginateOption {
	if order == nil {
		order = DefaultVirtualMachineOrder
	}
	o := *order
	return func(pager *virtualmachinePager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultVirtualMachineOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithVirtualMachineFilter configures pagination filter.
func WithVirtualMachineFilter(filter func(*VirtualMachineQuery) (*VirtualMachineQuery, error)) VirtualMachinePaginateOption {
	return func(pager *virtualmachinePager) error {
		if filter == nil {
			return errors.New("VirtualMachineQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type virtualmachinePager struct {
	reverse bool
	order   *VirtualMachineOrder
	filter  func(*VirtualMachineQuery) (*VirtualMachineQuery, error)
}

func newVirtualMachinePager(opts []VirtualMachinePaginateOption, reverse bool) (*virtualmachinePager, error) {
	pager := &virtualmachinePager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultVirtualMachineOrder
	}
	return pager, nil
}

func (p *virtualmachinePager) applyFilter(query *VirtualMachineQuery) (*VirtualMachineQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *virtualmachinePager) toCursor(vm *VirtualMachine) Cursor {
	return p.order.Field.toCursor(vm)
}

func (p *virtualmachinePager) applyCursors(query *VirtualMachineQuery, after, before *Cursor) (*VirtualMachineQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultVirtualMachineOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *virtualmachinePager) applyOrder(query *VirtualMachineQuery) *VirtualMachineQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultVirtualMachineOrder.Field {
		query = query.Order(DefaultVirtualMachineOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *virtualmachinePager) orderExpr(query *VirtualMachineQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultVirtualMachineOrder.Field {
			b.Comma().Ident(DefaultVirtualMachineOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to VirtualMachine.
func (vm *VirtualMachineQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...VirtualMachinePaginateOption,
) (*VirtualMachineConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newVirtualMachinePager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if vm, err = pager.applyFilter(vm); err != nil {
		return nil, err
	}
	conn := &VirtualMachineConnection{Edges: []*VirtualMachineEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = vm.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if vm, err = pager.applyCursors(vm, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		vm.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := vm.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	vm = pager.applyOrder(vm)
	nodes, err := vm.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// VirtualMachineOrderFieldID orders VirtualMachine by id.
	VirtualMachineOrderFieldID = &VirtualMachineOrderField{
		Value: func(vm *VirtualMachine) (ent.Value, error) {
			return vm.ID, nil
		},
		column: virtualmachine.FieldID,
		toTerm: virtualmachine.ByID,
		toCursor: func(vm *VirtualMachine) Cursor {
			return Cursor{
				ID:    vm.ID,
				Value: vm.ID,
			}
		},
	}
	// VirtualMachineOrderFieldCreatedAt orders VirtualMachine by created_at.
	VirtualMachineOrderFieldCreatedAt = &VirtualMachineOrderField{
		Value: func(vm *VirtualMachine) (ent.Value, error) {
			return vm.CreatedAt, nil
		},
		column: virtualmachine.FieldCreatedAt,
		toTerm: virtualmachine.ByCreatedAt,
		toCursor: func(vm *VirtualMachine) Cursor {
			return Cursor{
				ID:    vm.ID,
				Value: vm.CreatedAt,
			}
		},
	}
	// VirtualMachineOrderFieldUpdatedAt orders VirtualMachine by updated_at.
	VirtualMachineOrderFieldUpdatedAt = &VirtualMachineOrderField{
		Value: func(vm *VirtualMachine) (ent.Value, error) {
			return vm.UpdatedAt, nil
		},
		column: virtualmachine.FieldUpdatedAt,
		toTerm: virtualmachine.ByUpdatedAt,
		toCursor: func(vm *VirtualMachine) Cursor {
			return Cursor{
				ID:    vm.ID,
				Value: vm.UpdatedAt,
			}
		},
	}
	// VirtualMachineOrderFieldName orders VirtualMachine by name.
	VirtualMachineOrderFieldName = &VirtualMachineOrderField{
		Value: func(vm *VirtualMachine) (ent.Value, error) {
			return vm.Name, nil
		},
		column: virtualmachine.FieldName,
		toTerm: virtualmachine.ByName,
		toCursor: func(vm *VirtualMachine) Cursor {
			return Cursor{
				ID:    vm.ID,
				Value: vm.Name,
			}
		},
	}
	// VirtualMachineOrderFieldOwnerID orders VirtualMachine by owner_id.
	VirtualMachineOrderFieldOwnerID = &VirtualMachineOrderField{
		Value: func(vm *VirtualMachine) (ent.Value, error) {
			return vm.OwnerID, nil
		},
		column: virtualmachine.FieldOwnerID,
		toTerm: virtualmachine.ByOwnerID,
		toCursor: func(vm *VirtualMachine) Cursor {
			return Cursor{
				ID:    vm.ID,
				Value: vm.OwnerID,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f VirtualMachineOrderField) String() string {
	var str string
	switch f.column {
	case VirtualMachineOrderFieldID.column:
		str = "ID"
	case VirtualMachineOrderFieldCreatedAt.column:
		str = "CREATED_AT"
	case VirtualMachineOrderFieldUpdatedAt.column:
		str = "UPDATED_AT"
	case VirtualMachineOrderFieldName.column:
		str = "NAME"
	case VirtualMachineOrderFieldOwnerID.column:
		str = "OWNER"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f VirtualMachineOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *VirtualMachineOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("VirtualMachineOrderField %T must be a string", v)
	}
	switch str {
	case "ID":
		*f = *VirtualMachineOrderFieldID
	case "CREATED_AT":
		*f = *VirtualMachineOrderFieldCreatedAt
	case "UPDATED_AT":
		*f = *VirtualMachineOrderFieldUpdatedAt
	case "NAME":
		*f = *VirtualMachineOrderFieldName
	case "OWNER":
		*f = *VirtualMachineOrderFieldOwnerID
	default:
		return fmt.Errorf("%s is not a valid VirtualMachineOrderField", str)
	}
	return nil
}

// VirtualMachineOrderField defines the ordering field of VirtualMachine.
type VirtualMachineOrderField struct {
	// Value extracts the ordering value from the given VirtualMachine.
	Value    func(*VirtualMachine) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) virtualmachine.OrderOption
	toCursor func(*VirtualMachine) Cursor
}

// VirtualMachineOrder defines the ordering of VirtualMachine.
type VirtualMachineOrder struct {
	Direction OrderDirection            `json:"direction"`
	Field     *VirtualMachineOrderField `json:"field"`
}

// DefaultVirtualMachineOrder is the default ordering of VirtualMachine.
var DefaultVirtualMachineOrder = &VirtualMachineOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &VirtualMachineOrderField{
		Value: func(vm *VirtualMachine) (ent.Value, error) {
			return vm.ID, nil
		},
		column: virtualmachine.FieldID,
		toTerm: virtualmachine.ByID,
		toCursor: func(vm *VirtualMachine) Cursor {
			return Cursor{ID: vm.ID}
		},
	},
}

// ToEdge converts VirtualMachine into VirtualMachineEdge.
func (vm *VirtualMachine) ToEdge(order *VirtualMachineOrder) *VirtualMachineEdge {
	if order == nil {
		order = DefaultVirtualMachineOrder
	}
	return &VirtualMachineEdge{
		Node:   vm,
		Cursor: order.Field.toCursor(vm),
	}
}
