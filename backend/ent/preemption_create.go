// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/preemption"
	"github.com/team11/app/ent/purpose"
	"github.com/team11/app/ent/roominfo"
	"github.com/team11/app/ent/user"
)

// PreemptionCreate is the builder for creating a Preemption entity.
type PreemptionCreate struct {
	config
	mutation *PreemptionMutation
	hooks    []Hook
}

// SetPreemptTime sets the PreemptTime field.
func (pc *PreemptionCreate) SetPreemptTime(t time.Time) *PreemptionCreate {
	pc.mutation.SetPreemptTime(t)
	return pc
}

// SetNillablePreemptTime sets the PreemptTime field if the given value is not nil.
func (pc *PreemptionCreate) SetNillablePreemptTime(t *time.Time) *PreemptionCreate {
	if t != nil {
		pc.SetPreemptTime(*t)
	}
	return pc
}

// SetPhonenumber sets the Phonenumber field.
func (pc *PreemptionCreate) SetPhonenumber(s string) *PreemptionCreate {
	pc.mutation.SetPhonenumber(s)
	return pc
}

// SetSurrogateid sets the Surrogateid field.
func (pc *PreemptionCreate) SetSurrogateid(s string) *PreemptionCreate {
	pc.mutation.SetSurrogateid(s)
	return pc
}

// SetSurrogatephone sets the Surrogatephone field.
func (pc *PreemptionCreate) SetSurrogatephone(s string) *PreemptionCreate {
	pc.mutation.SetSurrogatephone(s)
	return pc
}

// SetUserIDID sets the User_ID edge to User by id.
func (pc *PreemptionCreate) SetUserIDID(id int) *PreemptionCreate {
	pc.mutation.SetUserIDID(id)
	return pc
}

// SetNillableUserIDID sets the User_ID edge to User by id if the given value is not nil.
func (pc *PreemptionCreate) SetNillableUserIDID(id *int) *PreemptionCreate {
	if id != nil {
		pc = pc.SetUserIDID(*id)
	}
	return pc
}

// SetUserID sets the User_ID edge to User.
func (pc *PreemptionCreate) SetUserID(u *User) *PreemptionCreate {
	return pc.SetUserIDID(u.ID)
}

// SetPurposeIDID sets the PurposeID edge to Purpose by id.
func (pc *PreemptionCreate) SetPurposeIDID(id int) *PreemptionCreate {
	pc.mutation.SetPurposeIDID(id)
	return pc
}

// SetNillablePurposeIDID sets the PurposeID edge to Purpose by id if the given value is not nil.
func (pc *PreemptionCreate) SetNillablePurposeIDID(id *int) *PreemptionCreate {
	if id != nil {
		pc = pc.SetPurposeIDID(*id)
	}
	return pc
}

// SetPurposeID sets the PurposeID edge to Purpose.
func (pc *PreemptionCreate) SetPurposeID(p *Purpose) *PreemptionCreate {
	return pc.SetPurposeIDID(p.ID)
}

// SetRoomIDID sets the RoomID edge to Roominfo by id.
func (pc *PreemptionCreate) SetRoomIDID(id int) *PreemptionCreate {
	pc.mutation.SetRoomIDID(id)
	return pc
}

// SetNillableRoomIDID sets the RoomID edge to Roominfo by id if the given value is not nil.
func (pc *PreemptionCreate) SetNillableRoomIDID(id *int) *PreemptionCreate {
	if id != nil {
		pc = pc.SetRoomIDID(*id)
	}
	return pc
}

// SetRoomID sets the RoomID edge to Roominfo.
func (pc *PreemptionCreate) SetRoomID(r *Roominfo) *PreemptionCreate {
	return pc.SetRoomIDID(r.ID)
}

// Mutation returns the PreemptionMutation object of the builder.
func (pc *PreemptionCreate) Mutation() *PreemptionMutation {
	return pc.mutation
}

// Save creates the Preemption in the database.
func (pc *PreemptionCreate) Save(ctx context.Context) (*Preemption, error) {
	if _, ok := pc.mutation.PreemptTime(); !ok {
		v := preemption.DefaultPreemptTime()
		pc.mutation.SetPreemptTime(v)
	}
	if _, ok := pc.mutation.Phonenumber(); !ok {
		return nil, &ValidationError{Name: "Phonenumber", err: errors.New("ent: missing required field \"Phonenumber\"")}
	}
	if v, ok := pc.mutation.Phonenumber(); ok {
		if err := preemption.PhonenumberValidator(v); err != nil {
			return nil, &ValidationError{Name: "Phonenumber", err: fmt.Errorf("ent: validator failed for field \"Phonenumber\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Surrogateid(); !ok {
		return nil, &ValidationError{Name: "Surrogateid", err: errors.New("ent: missing required field \"Surrogateid\"")}
	}
	if v, ok := pc.mutation.Surrogateid(); ok {
		if err := preemption.SurrogateidValidator(v); err != nil {
			return nil, &ValidationError{Name: "Surrogateid", err: fmt.Errorf("ent: validator failed for field \"Surrogateid\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Surrogatephone(); !ok {
		return nil, &ValidationError{Name: "Surrogatephone", err: errors.New("ent: missing required field \"Surrogatephone\"")}
	}
	if v, ok := pc.mutation.Surrogatephone(); ok {
		if err := preemption.SurrogatephoneValidator(v); err != nil {
			return nil, &ValidationError{Name: "Surrogatephone", err: fmt.Errorf("ent: validator failed for field \"Surrogatephone\": %w", err)}
		}
	}
	var (
		err  error
		node *Preemption
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PreemptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PreemptionCreate) SaveX(ctx context.Context) *Preemption {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PreemptionCreate) sqlSave(ctx context.Context) (*Preemption, error) {
	pr, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pr.ID = int(id)
	return pr, nil
}

func (pc *PreemptionCreate) createSpec() (*Preemption, *sqlgraph.CreateSpec) {
	var (
		pr    = &Preemption{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: preemption.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: preemption.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.PreemptTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: preemption.FieldPreemptTime,
		})
		pr.PreemptTime = value
	}
	if value, ok := pc.mutation.Phonenumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: preemption.FieldPhonenumber,
		})
		pr.Phonenumber = value
	}
	if value, ok := pc.mutation.Surrogateid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: preemption.FieldSurrogateid,
		})
		pr.Surrogateid = value
	}
	if value, ok := pc.mutation.Surrogatephone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: preemption.FieldSurrogatephone,
		})
		pr.Surrogatephone = value
	}
	if nodes := pc.mutation.UserIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   preemption.UserIDTable,
			Columns: []string{preemption.UserIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PurposeIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   preemption.PurposeIDTable,
			Columns: []string{preemption.PurposeIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: purpose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.RoomIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   preemption.RoomIDTable,
			Columns: []string{preemption.RoomIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roominfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pr, _spec
}
