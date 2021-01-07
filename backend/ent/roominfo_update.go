// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/predicate"
	"github.com/team11/app/ent/preemption"
	"github.com/team11/app/ent/roominfo"
)

// RoominfoUpdate is the builder for updating Roominfo entities.
type RoominfoUpdate struct {
	config
	hooks      []Hook
	mutation   *RoominfoMutation
	predicates []predicate.Roominfo
}

// Where adds a new predicate for the builder.
func (ru *RoominfoUpdate) Where(ps ...predicate.Roominfo) *RoominfoUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetRoomID sets the RoomID field.
func (ru *RoominfoUpdate) SetRoomID(s string) *RoominfoUpdate {
	ru.mutation.SetRoomID(s)
	return ru
}

// SetRoomNo sets the RoomNo field.
func (ru *RoominfoUpdate) SetRoomNo(s string) *RoominfoUpdate {
	ru.mutation.SetRoomNo(s)
	return ru
}

// SetRoomType sets the RoomType field.
func (ru *RoominfoUpdate) SetRoomType(s string) *RoominfoUpdate {
	ru.mutation.SetRoomType(s)
	return ru
}

// SetRoomTime sets the RoomTime field.
func (ru *RoominfoUpdate) SetRoomTime(s string) *RoominfoUpdate {
	ru.mutation.SetRoomTime(s)
	return ru
}

// SetRoomStatus sets the RoomStatus field.
func (ru *RoominfoUpdate) SetRoomStatus(s string) *RoominfoUpdate {
	ru.mutation.SetRoomStatus(s)
	return ru
}

// AddPreemptionIDs adds the preemption edge to Preemption by ids.
func (ru *RoominfoUpdate) AddPreemptionIDs(ids ...int) *RoominfoUpdate {
	ru.mutation.AddPreemptionIDs(ids...)
	return ru
}

// AddPreemption adds the preemption edges to Preemption.
func (ru *RoominfoUpdate) AddPreemption(p ...*Preemption) *RoominfoUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddPreemptionIDs(ids...)
}

// Mutation returns the RoominfoMutation object of the builder.
func (ru *RoominfoUpdate) Mutation() *RoominfoMutation {
	return ru.mutation
}

// RemovePreemptionIDs removes the preemption edge to Preemption by ids.
func (ru *RoominfoUpdate) RemovePreemptionIDs(ids ...int) *RoominfoUpdate {
	ru.mutation.RemovePreemptionIDs(ids...)
	return ru
}

// RemovePreemption removes preemption edges to Preemption.
func (ru *RoominfoUpdate) RemovePreemption(p ...*Preemption) *RoominfoUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemovePreemptionIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *RoominfoUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := ru.mutation.RoomNo(); ok {
		if err := roominfo.RoomNoValidator(v); err != nil {
			return 0, &ValidationError{Name: "RoomNo", err: fmt.Errorf("ent: validator failed for field \"RoomNo\": %w", err)}
		}
	}
	if v, ok := ru.mutation.RoomType(); ok {
		if err := roominfo.RoomTypeValidator(v); err != nil {
			return 0, &ValidationError{Name: "RoomType", err: fmt.Errorf("ent: validator failed for field \"RoomType\": %w", err)}
		}
	}
	if v, ok := ru.mutation.RoomTime(); ok {
		if err := roominfo.RoomTimeValidator(v); err != nil {
			return 0, &ValidationError{Name: "RoomTime", err: fmt.Errorf("ent: validator failed for field \"RoomTime\": %w", err)}
		}
	}
	if v, ok := ru.mutation.RoomStatus(); ok {
		if err := roominfo.RoomStatusValidator(v); err != nil {
			return 0, &ValidationError{Name: "RoomStatus", err: fmt.Errorf("ent: validator failed for field \"RoomStatus\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoominfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoominfoUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoominfoUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoominfoUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RoominfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roominfo.Table,
			Columns: roominfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roominfo.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.RoomID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roominfo.FieldRoomID,
		})
	}
	if value, ok := ru.mutation.RoomNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roominfo.FieldRoomNo,
		})
	}
	if value, ok := ru.mutation.RoomType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roominfo.FieldRoomType,
		})
	}
	if value, ok := ru.mutation.RoomTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roominfo.FieldRoomTime,
		})
	}
	if value, ok := ru.mutation.RoomStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roominfo.FieldRoomStatus,
		})
	}
	if nodes := ru.mutation.RemovedPreemptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roominfo.PreemptionTable,
			Columns: []string{roominfo.PreemptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: preemption.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.PreemptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roominfo.PreemptionTable,
			Columns: []string{roominfo.PreemptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: preemption.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roominfo.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RoominfoUpdateOne is the builder for updating a single Roominfo entity.
type RoominfoUpdateOne struct {
	config
	hooks    []Hook
	mutation *RoominfoMutation
}

// SetRoomID sets the RoomID field.
func (ruo *RoominfoUpdateOne) SetRoomID(s string) *RoominfoUpdateOne {
	ruo.mutation.SetRoomID(s)
	return ruo
}

// SetRoomNo sets the RoomNo field.
func (ruo *RoominfoUpdateOne) SetRoomNo(s string) *RoominfoUpdateOne {
	ruo.mutation.SetRoomNo(s)
	return ruo
}

// SetRoomType sets the RoomType field.
func (ruo *RoominfoUpdateOne) SetRoomType(s string) *RoominfoUpdateOne {
	ruo.mutation.SetRoomType(s)
	return ruo
}

// SetRoomTime sets the RoomTime field.
func (ruo *RoominfoUpdateOne) SetRoomTime(s string) *RoominfoUpdateOne {
	ruo.mutation.SetRoomTime(s)
	return ruo
}

// SetRoomStatus sets the RoomStatus field.
func (ruo *RoominfoUpdateOne) SetRoomStatus(s string) *RoominfoUpdateOne {
	ruo.mutation.SetRoomStatus(s)
	return ruo
}

// AddPreemptionIDs adds the preemption edge to Preemption by ids.
func (ruo *RoominfoUpdateOne) AddPreemptionIDs(ids ...int) *RoominfoUpdateOne {
	ruo.mutation.AddPreemptionIDs(ids...)
	return ruo
}

// AddPreemption adds the preemption edges to Preemption.
func (ruo *RoominfoUpdateOne) AddPreemption(p ...*Preemption) *RoominfoUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddPreemptionIDs(ids...)
}

// Mutation returns the RoominfoMutation object of the builder.
func (ruo *RoominfoUpdateOne) Mutation() *RoominfoMutation {
	return ruo.mutation
}

// RemovePreemptionIDs removes the preemption edge to Preemption by ids.
func (ruo *RoominfoUpdateOne) RemovePreemptionIDs(ids ...int) *RoominfoUpdateOne {
	ruo.mutation.RemovePreemptionIDs(ids...)
	return ruo
}

// RemovePreemption removes preemption edges to Preemption.
func (ruo *RoominfoUpdateOne) RemovePreemption(p ...*Preemption) *RoominfoUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemovePreemptionIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ruo *RoominfoUpdateOne) Save(ctx context.Context) (*Roominfo, error) {
	if v, ok := ruo.mutation.RoomNo(); ok {
		if err := roominfo.RoomNoValidator(v); err != nil {
			return nil, &ValidationError{Name: "RoomNo", err: fmt.Errorf("ent: validator failed for field \"RoomNo\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.RoomType(); ok {
		if err := roominfo.RoomTypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "RoomType", err: fmt.Errorf("ent: validator failed for field \"RoomType\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.RoomTime(); ok {
		if err := roominfo.RoomTimeValidator(v); err != nil {
			return nil, &ValidationError{Name: "RoomTime", err: fmt.Errorf("ent: validator failed for field \"RoomTime\": %w", err)}
		}
	}
	if v, ok := ruo.mutation.RoomStatus(); ok {
		if err := roominfo.RoomStatusValidator(v); err != nil {
			return nil, &ValidationError{Name: "RoomStatus", err: fmt.Errorf("ent: validator failed for field \"RoomStatus\": %w", err)}
		}
	}

	var (
		err  error
		node *Roominfo
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoominfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoominfoUpdateOne) SaveX(ctx context.Context) *Roominfo {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *RoominfoUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoominfoUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RoominfoUpdateOne) sqlSave(ctx context.Context) (r *Roominfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roominfo.Table,
			Columns: roominfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roominfo.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Roominfo.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.RoomID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roominfo.FieldRoomID,
		})
	}
	if value, ok := ruo.mutation.RoomNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roominfo.FieldRoomNo,
		})
	}
	if value, ok := ruo.mutation.RoomType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roominfo.FieldRoomType,
		})
	}
	if value, ok := ruo.mutation.RoomTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roominfo.FieldRoomTime,
		})
	}
	if value, ok := ruo.mutation.RoomStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roominfo.FieldRoomStatus,
		})
	}
	if nodes := ruo.mutation.RemovedPreemptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roominfo.PreemptionTable,
			Columns: []string{roominfo.PreemptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: preemption.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.PreemptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roominfo.PreemptionTable,
			Columns: []string{roominfo.PreemptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: preemption.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	r = &Roominfo{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roominfo.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}