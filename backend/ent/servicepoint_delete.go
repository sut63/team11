// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/predicate"
	"github.com/team11/app/ent/servicepoint"
)

// ServicePointDelete is the builder for deleting a ServicePoint entity.
type ServicePointDelete struct {
	config
	hooks      []Hook
	mutation   *ServicePointMutation
	predicates []predicate.ServicePoint
}

// Where adds a new predicate to the delete builder.
func (spd *ServicePointDelete) Where(ps ...predicate.ServicePoint) *ServicePointDelete {
	spd.predicates = append(spd.predicates, ps...)
	return spd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (spd *ServicePointDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(spd.hooks) == 0 {
		affected, err = spd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServicePointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spd.mutation = mutation
			affected, err = spd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(spd.hooks) - 1; i >= 0; i-- {
			mut = spd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (spd *ServicePointDelete) ExecX(ctx context.Context) int {
	n, err := spd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (spd *ServicePointDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: servicepoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: servicepoint.FieldID,
			},
		},
	}
	if ps := spd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, spd.driver, _spec)
}

// ServicePointDeleteOne is the builder for deleting a single ServicePoint entity.
type ServicePointDeleteOne struct {
	spd *ServicePointDelete
}

// Exec executes the deletion query.
func (spdo *ServicePointDeleteOne) Exec(ctx context.Context) error {
	n, err := spdo.spd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{servicepoint.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (spdo *ServicePointDeleteOne) ExecX(ctx context.Context) {
	spdo.spd.ExecX(ctx)
}
