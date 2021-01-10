// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/bookreturn"
	"github.com/team11/app/ent/location"
)

// LocationCreate is the builder for creating a Location entity.
type LocationCreate struct {
	config
	mutation *LocationMutation
	hooks    []Hook
}

// SetLOCATIONNAME sets the LOCATION_NAME field.
func (lc *LocationCreate) SetLOCATIONNAME(s string) *LocationCreate {
	lc.mutation.SetLOCATIONNAME(s)
	return lc
}

// AddReturnfromIDs adds the returnfrom edge to Bookreturn by ids.
func (lc *LocationCreate) AddReturnfromIDs(ids ...int) *LocationCreate {
	lc.mutation.AddReturnfromIDs(ids...)
	return lc
}

// AddReturnfrom adds the returnfrom edges to Bookreturn.
func (lc *LocationCreate) AddReturnfrom(b ...*Bookreturn) *LocationCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return lc.AddReturnfromIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (lc *LocationCreate) Mutation() *LocationMutation {
	return lc.mutation
}

// Save creates the Location in the database.
func (lc *LocationCreate) Save(ctx context.Context) (*Location, error) {
	if _, ok := lc.mutation.LOCATIONNAME(); !ok {
		return nil, &ValidationError{Name: "LOCATION_NAME", err: errors.New("ent: missing required field \"LOCATION_NAME\"")}
	}
	if v, ok := lc.mutation.LOCATIONNAME(); ok {
		if err := location.LOCATIONNAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "LOCATION_NAME", err: fmt.Errorf("ent: validator failed for field \"LOCATION_NAME\": %w", err)}
		}
	}
	var (
		err  error
		node *Location
	)
	if len(lc.hooks) == 0 {
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LocationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lc.mutation = mutation
			node, err = lc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			mut = lc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LocationCreate) SaveX(ctx context.Context) *Location {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lc *LocationCreate) sqlSave(ctx context.Context) (*Location, error) {
	l, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	l.ID = int(id)
	return l, nil
}

func (lc *LocationCreate) createSpec() (*Location, *sqlgraph.CreateSpec) {
	var (
		l     = &Location{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: location.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: location.FieldID,
			},
		}
	)
	if value, ok := lc.mutation.LOCATIONNAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: location.FieldLOCATIONNAME,
		})
		l.LOCATIONNAME = value
	}
	if nodes := lc.mutation.ReturnfromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ReturnfromTable,
			Columns: []string{location.ReturnfromColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bookreturn.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return l, _spec
}
