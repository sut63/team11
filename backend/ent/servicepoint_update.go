// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/bookborrow"
	"github.com/team11/app/ent/booking"
	"github.com/team11/app/ent/predicate"
	"github.com/team11/app/ent/servicepoint"
)

// ServicePointUpdate is the builder for updating ServicePoint entities.
type ServicePointUpdate struct {
	config
	hooks      []Hook
	mutation   *ServicePointMutation
	predicates []predicate.ServicePoint
}

// Where adds a new predicate for the builder.
func (spu *ServicePointUpdate) Where(ps ...predicate.ServicePoint) *ServicePointUpdate {
	spu.predicates = append(spu.predicates, ps...)
	return spu
}

// SetBUILDINGNAME sets the BUILDING_NAME field.
func (spu *ServicePointUpdate) SetBUILDINGNAME(s string) *ServicePointUpdate {
	spu.mutation.SetBUILDINGNAME(s)
	return spu
}

// SetCOUNTERNUMBER sets the COUNTER_NUMBER field.
func (spu *ServicePointUpdate) SetCOUNTERNUMBER(s string) *ServicePointUpdate {
	spu.mutation.SetCOUNTERNUMBER(s)
	return spu
}

// AddFromIDs adds the from edge to Bookborrow by ids.
func (spu *ServicePointUpdate) AddFromIDs(ids ...int) *ServicePointUpdate {
	spu.mutation.AddFromIDs(ids...)
	return spu
}

// AddFrom adds the from edges to Bookborrow.
func (spu *ServicePointUpdate) AddFrom(b ...*Bookborrow) *ServicePointUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return spu.AddFromIDs(ids...)
}

// AddServicepointIDs adds the servicepoint edge to Booking by ids.
func (spu *ServicePointUpdate) AddServicepointIDs(ids ...int) *ServicePointUpdate {
	spu.mutation.AddServicepointIDs(ids...)
	return spu
}

// AddServicepoint adds the servicepoint edges to Booking.
func (spu *ServicePointUpdate) AddServicepoint(b ...*Booking) *ServicePointUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return spu.AddServicepointIDs(ids...)
}

// Mutation returns the ServicePointMutation object of the builder.
func (spu *ServicePointUpdate) Mutation() *ServicePointMutation {
	return spu.mutation
}

// RemoveFromIDs removes the from edge to Bookborrow by ids.
func (spu *ServicePointUpdate) RemoveFromIDs(ids ...int) *ServicePointUpdate {
	spu.mutation.RemoveFromIDs(ids...)
	return spu
}

// RemoveFrom removes from edges to Bookborrow.
func (spu *ServicePointUpdate) RemoveFrom(b ...*Bookborrow) *ServicePointUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return spu.RemoveFromIDs(ids...)
}

// RemoveServicepointIDs removes the servicepoint edge to Booking by ids.
func (spu *ServicePointUpdate) RemoveServicepointIDs(ids ...int) *ServicePointUpdate {
	spu.mutation.RemoveServicepointIDs(ids...)
	return spu
}

// RemoveServicepoint removes servicepoint edges to Booking.
func (spu *ServicePointUpdate) RemoveServicepoint(b ...*Booking) *ServicePointUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return spu.RemoveServicepointIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (spu *ServicePointUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := spu.mutation.BUILDINGNAME(); ok {
		if err := servicepoint.BUILDINGNAMEValidator(v); err != nil {
			return 0, &ValidationError{Name: "BUILDING_NAME", err: fmt.Errorf("ent: validator failed for field \"BUILDING_NAME\": %w", err)}
		}
	}
	if v, ok := spu.mutation.COUNTERNUMBER(); ok {
		if err := servicepoint.COUNTERNUMBERValidator(v); err != nil {
			return 0, &ValidationError{Name: "COUNTER_NUMBER", err: fmt.Errorf("ent: validator failed for field \"COUNTER_NUMBER\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(spu.hooks) == 0 {
		affected, err = spu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServicePointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spu.mutation = mutation
			affected, err = spu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(spu.hooks) - 1; i >= 0; i-- {
			mut = spu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (spu *ServicePointUpdate) SaveX(ctx context.Context) int {
	affected, err := spu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (spu *ServicePointUpdate) Exec(ctx context.Context) error {
	_, err := spu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spu *ServicePointUpdate) ExecX(ctx context.Context) {
	if err := spu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (spu *ServicePointUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   servicepoint.Table,
			Columns: servicepoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: servicepoint.FieldID,
			},
		},
	}
	if ps := spu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spu.mutation.BUILDINGNAME(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: servicepoint.FieldBUILDINGNAME,
		})
	}
	if value, ok := spu.mutation.COUNTERNUMBER(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: servicepoint.FieldCOUNTERNUMBER,
		})
	}
	if nodes := spu.mutation.RemovedFromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   servicepoint.FromTable,
			Columns: []string{servicepoint.FromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bookborrow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spu.mutation.FromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   servicepoint.FromTable,
			Columns: []string{servicepoint.FromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bookborrow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := spu.mutation.RemovedServicepointIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   servicepoint.ServicepointTable,
			Columns: []string{servicepoint.ServicepointColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spu.mutation.ServicepointIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   servicepoint.ServicepointTable,
			Columns: []string{servicepoint.ServicepointColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, spu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servicepoint.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ServicePointUpdateOne is the builder for updating a single ServicePoint entity.
type ServicePointUpdateOne struct {
	config
	hooks    []Hook
	mutation *ServicePointMutation
}

// SetBUILDINGNAME sets the BUILDING_NAME field.
func (spuo *ServicePointUpdateOne) SetBUILDINGNAME(s string) *ServicePointUpdateOne {
	spuo.mutation.SetBUILDINGNAME(s)
	return spuo
}

// SetCOUNTERNUMBER sets the COUNTER_NUMBER field.
func (spuo *ServicePointUpdateOne) SetCOUNTERNUMBER(s string) *ServicePointUpdateOne {
	spuo.mutation.SetCOUNTERNUMBER(s)
	return spuo
}

// AddFromIDs adds the from edge to Bookborrow by ids.
func (spuo *ServicePointUpdateOne) AddFromIDs(ids ...int) *ServicePointUpdateOne {
	spuo.mutation.AddFromIDs(ids...)
	return spuo
}

// AddFrom adds the from edges to Bookborrow.
func (spuo *ServicePointUpdateOne) AddFrom(b ...*Bookborrow) *ServicePointUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return spuo.AddFromIDs(ids...)
}

// AddServicepointIDs adds the servicepoint edge to Booking by ids.
func (spuo *ServicePointUpdateOne) AddServicepointIDs(ids ...int) *ServicePointUpdateOne {
	spuo.mutation.AddServicepointIDs(ids...)
	return spuo
}

// AddServicepoint adds the servicepoint edges to Booking.
func (spuo *ServicePointUpdateOne) AddServicepoint(b ...*Booking) *ServicePointUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return spuo.AddServicepointIDs(ids...)
}

// Mutation returns the ServicePointMutation object of the builder.
func (spuo *ServicePointUpdateOne) Mutation() *ServicePointMutation {
	return spuo.mutation
}

// RemoveFromIDs removes the from edge to Bookborrow by ids.
func (spuo *ServicePointUpdateOne) RemoveFromIDs(ids ...int) *ServicePointUpdateOne {
	spuo.mutation.RemoveFromIDs(ids...)
	return spuo
}

// RemoveFrom removes from edges to Bookborrow.
func (spuo *ServicePointUpdateOne) RemoveFrom(b ...*Bookborrow) *ServicePointUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return spuo.RemoveFromIDs(ids...)
}

// RemoveServicepointIDs removes the servicepoint edge to Booking by ids.
func (spuo *ServicePointUpdateOne) RemoveServicepointIDs(ids ...int) *ServicePointUpdateOne {
	spuo.mutation.RemoveServicepointIDs(ids...)
	return spuo
}

// RemoveServicepoint removes servicepoint edges to Booking.
func (spuo *ServicePointUpdateOne) RemoveServicepoint(b ...*Booking) *ServicePointUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return spuo.RemoveServicepointIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (spuo *ServicePointUpdateOne) Save(ctx context.Context) (*ServicePoint, error) {
	if v, ok := spuo.mutation.BUILDINGNAME(); ok {
		if err := servicepoint.BUILDINGNAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "BUILDING_NAME", err: fmt.Errorf("ent: validator failed for field \"BUILDING_NAME\": %w", err)}
		}
	}
	if v, ok := spuo.mutation.COUNTERNUMBER(); ok {
		if err := servicepoint.COUNTERNUMBERValidator(v); err != nil {
			return nil, &ValidationError{Name: "COUNTER_NUMBER", err: fmt.Errorf("ent: validator failed for field \"COUNTER_NUMBER\": %w", err)}
		}
	}

	var (
		err  error
		node *ServicePoint
	)
	if len(spuo.hooks) == 0 {
		node, err = spuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServicePointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spuo.mutation = mutation
			node, err = spuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(spuo.hooks) - 1; i >= 0; i-- {
			mut = spuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (spuo *ServicePointUpdateOne) SaveX(ctx context.Context) *ServicePoint {
	sp, err := spuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return sp
}

// Exec executes the query on the entity.
func (spuo *ServicePointUpdateOne) Exec(ctx context.Context) error {
	_, err := spuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spuo *ServicePointUpdateOne) ExecX(ctx context.Context) {
	if err := spuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (spuo *ServicePointUpdateOne) sqlSave(ctx context.Context) (sp *ServicePoint, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   servicepoint.Table,
			Columns: servicepoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: servicepoint.FieldID,
			},
		},
	}
	id, ok := spuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ServicePoint.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := spuo.mutation.BUILDINGNAME(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: servicepoint.FieldBUILDINGNAME,
		})
	}
	if value, ok := spuo.mutation.COUNTERNUMBER(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: servicepoint.FieldCOUNTERNUMBER,
		})
	}
	if nodes := spuo.mutation.RemovedFromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   servicepoint.FromTable,
			Columns: []string{servicepoint.FromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bookborrow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spuo.mutation.FromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   servicepoint.FromTable,
			Columns: []string{servicepoint.FromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bookborrow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := spuo.mutation.RemovedServicepointIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   servicepoint.ServicepointTable,
			Columns: []string{servicepoint.ServicepointColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spuo.mutation.ServicepointIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   servicepoint.ServicepointTable,
			Columns: []string{servicepoint.ServicepointColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	sp = &ServicePoint{config: spuo.config}
	_spec.Assign = sp.assignValues
	_spec.ScanValues = sp.scanValues()
	if err = sqlgraph.UpdateNode(ctx, spuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servicepoint.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return sp, nil
}
