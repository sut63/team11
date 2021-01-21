// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/bookborrow"
	"github.com/team11/app/ent/bookreturn"
	"github.com/team11/app/ent/location"
	"github.com/team11/app/ent/user"
)

// BookreturnCreate is the builder for creating a Bookreturn entity.
type BookreturnCreate struct {
	config
	mutation *BookreturnMutation
	hooks    []Hook
}

// SetRETURNTIME sets the RETURN_TIME field.
func (bc *BookreturnCreate) SetRETURNTIME(t time.Time) *BookreturnCreate {
	bc.mutation.SetRETURNTIME(t)
	return bc
}

// SetDAMAGEDPOINT sets the DAMAGED_POINT field.
func (bc *BookreturnCreate) SetDAMAGEDPOINT(i int) *BookreturnCreate {
	bc.mutation.SetDAMAGEDPOINT(i)
	return bc
}

// SetDAMAGEDPOINTNAME sets the DAMAGED_POINTNAME field.
func (bc *BookreturnCreate) SetDAMAGEDPOINTNAME(s string) *BookreturnCreate {
	bc.mutation.SetDAMAGEDPOINTNAME(s)
	return bc
}

// SetLOST sets the LOST field.
func (bc *BookreturnCreate) SetLOST(s string) *BookreturnCreate {
	bc.mutation.SetLOST(s)
	return bc
}

// SetUserID sets the user edge to User by id.
func (bc *BookreturnCreate) SetUserID(id int) *BookreturnCreate {
	bc.mutation.SetUserID(id)
	return bc
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (bc *BookreturnCreate) SetNillableUserID(id *int) *BookreturnCreate {
	if id != nil {
		bc = bc.SetUserID(*id)
	}
	return bc
}

// SetUser sets the user edge to User.
func (bc *BookreturnCreate) SetUser(u *User) *BookreturnCreate {
	return bc.SetUserID(u.ID)
}

// SetLocationID sets the location edge to Location by id.
func (bc *BookreturnCreate) SetLocationID(id int) *BookreturnCreate {
	bc.mutation.SetLocationID(id)
	return bc
}

// SetNillableLocationID sets the location edge to Location by id if the given value is not nil.
func (bc *BookreturnCreate) SetNillableLocationID(id *int) *BookreturnCreate {
	if id != nil {
		bc = bc.SetLocationID(*id)
	}
	return bc
}

// SetLocation sets the location edge to Location.
func (bc *BookreturnCreate) SetLocation(l *Location) *BookreturnCreate {
	return bc.SetLocationID(l.ID)
}

// SetMustreturnID sets the mustreturn edge to Bookborrow by id.
func (bc *BookreturnCreate) SetMustreturnID(id int) *BookreturnCreate {
	bc.mutation.SetMustreturnID(id)
	return bc
}

// SetNillableMustreturnID sets the mustreturn edge to Bookborrow by id if the given value is not nil.
func (bc *BookreturnCreate) SetNillableMustreturnID(id *int) *BookreturnCreate {
	if id != nil {
		bc = bc.SetMustreturnID(*id)
	}
	return bc
}

// SetMustreturn sets the mustreturn edge to Bookborrow.
func (bc *BookreturnCreate) SetMustreturn(b *Bookborrow) *BookreturnCreate {
	return bc.SetMustreturnID(b.ID)
}

// Mutation returns the BookreturnMutation object of the builder.
func (bc *BookreturnCreate) Mutation() *BookreturnMutation {
	return bc.mutation
}

// Save creates the Bookreturn in the database.
func (bc *BookreturnCreate) Save(ctx context.Context) (*Bookreturn, error) {
	if _, ok := bc.mutation.RETURNTIME(); !ok {
		return nil, &ValidationError{Name: "RETURN_TIME", err: errors.New("ent: missing required field \"RETURN_TIME\"")}
	}
	if _, ok := bc.mutation.DAMAGEDPOINT(); !ok {
		return nil, &ValidationError{Name: "DAMAGED_POINT", err: errors.New("ent: missing required field \"DAMAGED_POINT\"")}
	}
	if v, ok := bc.mutation.DAMAGEDPOINT(); ok {
		if err := bookreturn.DAMAGEDPOINTValidator(v); err != nil {
			return nil, &ValidationError{Name: "DAMAGED_POINT", err: fmt.Errorf("ent: validator failed for field \"DAMAGED_POINT\": %w", err)}
		}
	}
	if _, ok := bc.mutation.DAMAGEDPOINTNAME(); !ok {
		return nil, &ValidationError{Name: "DAMAGED_POINTNAME", err: errors.New("ent: missing required field \"DAMAGED_POINTNAME\"")}
	}
	if v, ok := bc.mutation.DAMAGEDPOINTNAME(); ok {
		if err := bookreturn.DAMAGEDPOINTNAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "DAMAGED_POINTNAME", err: fmt.Errorf("ent: validator failed for field \"DAMAGED_POINTNAME\": %w", err)}
		}
	}
	if _, ok := bc.mutation.LOST(); !ok {
		return nil, &ValidationError{Name: "LOST", err: errors.New("ent: missing required field \"LOST\"")}
	}
	if v, ok := bc.mutation.LOST(); ok {
		if err := bookreturn.LOSTValidator(v); err != nil {
			return nil, &ValidationError{Name: "LOST", err: fmt.Errorf("ent: validator failed for field \"LOST\": %w", err)}
		}
	}
	var (
		err  error
		node *Bookreturn
	)
	if len(bc.hooks) == 0 {
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookreturnMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookreturnCreate) SaveX(ctx context.Context) *Bookreturn {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bc *BookreturnCreate) sqlSave(ctx context.Context) (*Bookreturn, error) {
	b, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	b.ID = int(id)
	return b, nil
}

func (bc *BookreturnCreate) createSpec() (*Bookreturn, *sqlgraph.CreateSpec) {
	var (
		b     = &Bookreturn{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bookreturn.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bookreturn.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.RETURNTIME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bookreturn.FieldRETURNTIME,
		})
		b.RETURNTIME = value
	}
	if value, ok := bc.mutation.DAMAGEDPOINT(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bookreturn.FieldDAMAGEDPOINT,
		})
		b.DAMAGEDPOINT = value
	}
	if value, ok := bc.mutation.DAMAGEDPOINTNAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bookreturn.FieldDAMAGEDPOINTNAME,
		})
		b.DAMAGEDPOINTNAME = value
	}
	if value, ok := bc.mutation.LOST(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bookreturn.FieldLOST,
		})
		b.LOST = value
	}
	if nodes := bc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookreturn.UserTable,
			Columns: []string{bookreturn.UserColumn},
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
	if nodes := bc.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookreturn.LocationTable,
			Columns: []string{bookreturn.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: location.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.MustreturnIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookreturn.MustreturnTable,
			Columns: []string{bookreturn.MustreturnColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return b, _spec
}
