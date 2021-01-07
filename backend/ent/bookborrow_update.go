// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/book"
	"github.com/team11/app/ent/bookborrow"
	"github.com/team11/app/ent/predicate"
	"github.com/team11/app/ent/servicepoint"
	"github.com/team11/app/ent/user"
)

// BookborrowUpdate is the builder for updating Bookborrow entities.
type BookborrowUpdate struct {
	config
	hooks      []Hook
	mutation   *BookborrowMutation
	predicates []predicate.Bookborrow
}

// Where adds a new predicate for the builder.
func (bu *BookborrowUpdate) Where(ps ...predicate.Bookborrow) *BookborrowUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// SetBORROWDATE sets the BORROW_DATE field.
func (bu *BookborrowUpdate) SetBORROWDATE(t time.Time) *BookborrowUpdate {
	bu.mutation.SetBORROWDATE(t)
	return bu
}

// SetNillableBORROWDATE sets the BORROW_DATE field if the given value is not nil.
func (bu *BookborrowUpdate) SetNillableBORROWDATE(t *time.Time) *BookborrowUpdate {
	if t != nil {
		bu.SetBORROWDATE(*t)
	}
	return bu
}

// SetUSERID sets the USER edge to User by id.
func (bu *BookborrowUpdate) SetUSERID(id int) *BookborrowUpdate {
	bu.mutation.SetUSERID(id)
	return bu
}

// SetNillableUSERID sets the USER edge to User by id if the given value is not nil.
func (bu *BookborrowUpdate) SetNillableUSERID(id *int) *BookborrowUpdate {
	if id != nil {
		bu = bu.SetUSERID(*id)
	}
	return bu
}

// SetUSER sets the USER edge to User.
func (bu *BookborrowUpdate) SetUSER(u *User) *BookborrowUpdate {
	return bu.SetUSERID(u.ID)
}

// SetBOOKID sets the BOOK edge to Book by id.
func (bu *BookborrowUpdate) SetBOOKID(id int) *BookborrowUpdate {
	bu.mutation.SetBOOKID(id)
	return bu
}

// SetNillableBOOKID sets the BOOK edge to Book by id if the given value is not nil.
func (bu *BookborrowUpdate) SetNillableBOOKID(id *int) *BookborrowUpdate {
	if id != nil {
		bu = bu.SetBOOKID(*id)
	}
	return bu
}

// SetBOOK sets the BOOK edge to Book.
func (bu *BookborrowUpdate) SetBOOK(b *Book) *BookborrowUpdate {
	return bu.SetBOOKID(b.ID)
}

// SetSERVICEPOINTID sets the SERVICEPOINT edge to ServicePoint by id.
func (bu *BookborrowUpdate) SetSERVICEPOINTID(id int) *BookborrowUpdate {
	bu.mutation.SetSERVICEPOINTID(id)
	return bu
}

// SetNillableSERVICEPOINTID sets the SERVICEPOINT edge to ServicePoint by id if the given value is not nil.
func (bu *BookborrowUpdate) SetNillableSERVICEPOINTID(id *int) *BookborrowUpdate {
	if id != nil {
		bu = bu.SetSERVICEPOINTID(*id)
	}
	return bu
}

// SetSERVICEPOINT sets the SERVICEPOINT edge to ServicePoint.
func (bu *BookborrowUpdate) SetSERVICEPOINT(s *ServicePoint) *BookborrowUpdate {
	return bu.SetSERVICEPOINTID(s.ID)
}

// Mutation returns the BookborrowMutation object of the builder.
func (bu *BookborrowUpdate) Mutation() *BookborrowMutation {
	return bu.mutation
}

// ClearUSER clears the USER edge to User.
func (bu *BookborrowUpdate) ClearUSER() *BookborrowUpdate {
	bu.mutation.ClearUSER()
	return bu
}

// ClearBOOK clears the BOOK edge to Book.
func (bu *BookborrowUpdate) ClearBOOK() *BookborrowUpdate {
	bu.mutation.ClearBOOK()
	return bu
}

// ClearSERVICEPOINT clears the SERVICEPOINT edge to ServicePoint.
func (bu *BookborrowUpdate) ClearSERVICEPOINT() *BookborrowUpdate {
	bu.mutation.ClearSERVICEPOINT()
	return bu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BookborrowUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookborrowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookborrowUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookborrowUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookborrowUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BookborrowUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bookborrow.Table,
			Columns: bookborrow.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bookborrow.FieldID,
			},
		},
	}
	if ps := bu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.BORROWDATE(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bookborrow.FieldBORROWDATE,
		})
	}
	if bu.mutation.USERCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.USERTable,
			Columns: []string{bookborrow.USERColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.USERIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.USERTable,
			Columns: []string{bookborrow.USERColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.BOOKCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.BOOKTable,
			Columns: []string{bookborrow.BOOKColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: book.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.BOOKIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.BOOKTable,
			Columns: []string{bookborrow.BOOKColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: book.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.SERVICEPOINTCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.SERVICEPOINTTable,
			Columns: []string{bookborrow.SERVICEPOINTColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: servicepoint.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.SERVICEPOINTIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.SERVICEPOINTTable,
			Columns: []string{bookborrow.SERVICEPOINTColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: servicepoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookborrow.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BookborrowUpdateOne is the builder for updating a single Bookborrow entity.
type BookborrowUpdateOne struct {
	config
	hooks    []Hook
	mutation *BookborrowMutation
}

// SetBORROWDATE sets the BORROW_DATE field.
func (buo *BookborrowUpdateOne) SetBORROWDATE(t time.Time) *BookborrowUpdateOne {
	buo.mutation.SetBORROWDATE(t)
	return buo
}

// SetNillableBORROWDATE sets the BORROW_DATE field if the given value is not nil.
func (buo *BookborrowUpdateOne) SetNillableBORROWDATE(t *time.Time) *BookborrowUpdateOne {
	if t != nil {
		buo.SetBORROWDATE(*t)
	}
	return buo
}

// SetUSERID sets the USER edge to User by id.
func (buo *BookborrowUpdateOne) SetUSERID(id int) *BookborrowUpdateOne {
	buo.mutation.SetUSERID(id)
	return buo
}

// SetNillableUSERID sets the USER edge to User by id if the given value is not nil.
func (buo *BookborrowUpdateOne) SetNillableUSERID(id *int) *BookborrowUpdateOne {
	if id != nil {
		buo = buo.SetUSERID(*id)
	}
	return buo
}

// SetUSER sets the USER edge to User.
func (buo *BookborrowUpdateOne) SetUSER(u *User) *BookborrowUpdateOne {
	return buo.SetUSERID(u.ID)
}

// SetBOOKID sets the BOOK edge to Book by id.
func (buo *BookborrowUpdateOne) SetBOOKID(id int) *BookborrowUpdateOne {
	buo.mutation.SetBOOKID(id)
	return buo
}

// SetNillableBOOKID sets the BOOK edge to Book by id if the given value is not nil.
func (buo *BookborrowUpdateOne) SetNillableBOOKID(id *int) *BookborrowUpdateOne {
	if id != nil {
		buo = buo.SetBOOKID(*id)
	}
	return buo
}

// SetBOOK sets the BOOK edge to Book.
func (buo *BookborrowUpdateOne) SetBOOK(b *Book) *BookborrowUpdateOne {
	return buo.SetBOOKID(b.ID)
}

// SetSERVICEPOINTID sets the SERVICEPOINT edge to ServicePoint by id.
func (buo *BookborrowUpdateOne) SetSERVICEPOINTID(id int) *BookborrowUpdateOne {
	buo.mutation.SetSERVICEPOINTID(id)
	return buo
}

// SetNillableSERVICEPOINTID sets the SERVICEPOINT edge to ServicePoint by id if the given value is not nil.
func (buo *BookborrowUpdateOne) SetNillableSERVICEPOINTID(id *int) *BookborrowUpdateOne {
	if id != nil {
		buo = buo.SetSERVICEPOINTID(*id)
	}
	return buo
}

// SetSERVICEPOINT sets the SERVICEPOINT edge to ServicePoint.
func (buo *BookborrowUpdateOne) SetSERVICEPOINT(s *ServicePoint) *BookborrowUpdateOne {
	return buo.SetSERVICEPOINTID(s.ID)
}

// Mutation returns the BookborrowMutation object of the builder.
func (buo *BookborrowUpdateOne) Mutation() *BookborrowMutation {
	return buo.mutation
}

// ClearUSER clears the USER edge to User.
func (buo *BookborrowUpdateOne) ClearUSER() *BookborrowUpdateOne {
	buo.mutation.ClearUSER()
	return buo
}

// ClearBOOK clears the BOOK edge to Book.
func (buo *BookborrowUpdateOne) ClearBOOK() *BookborrowUpdateOne {
	buo.mutation.ClearBOOK()
	return buo
}

// ClearSERVICEPOINT clears the SERVICEPOINT edge to ServicePoint.
func (buo *BookborrowUpdateOne) ClearSERVICEPOINT() *BookborrowUpdateOne {
	buo.mutation.ClearSERVICEPOINT()
	return buo
}

// Save executes the query and returns the updated entity.
func (buo *BookborrowUpdateOne) Save(ctx context.Context) (*Bookborrow, error) {

	var (
		err  error
		node *Bookborrow
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookborrowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookborrowUpdateOne) SaveX(ctx context.Context) *Bookborrow {
	b, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// Exec executes the query on the entity.
func (buo *BookborrowUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookborrowUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BookborrowUpdateOne) sqlSave(ctx context.Context) (b *Bookborrow, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bookborrow.Table,
			Columns: bookborrow.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bookborrow.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Bookborrow.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := buo.mutation.BORROWDATE(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bookborrow.FieldBORROWDATE,
		})
	}
	if buo.mutation.USERCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.USERTable,
			Columns: []string{bookborrow.USERColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.USERIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.USERTable,
			Columns: []string{bookborrow.USERColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.BOOKCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.BOOKTable,
			Columns: []string{bookborrow.BOOKColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: book.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.BOOKIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.BOOKTable,
			Columns: []string{bookborrow.BOOKColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: book.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.SERVICEPOINTCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.SERVICEPOINTTable,
			Columns: []string{bookborrow.SERVICEPOINTColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: servicepoint.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.SERVICEPOINTIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.SERVICEPOINTTable,
			Columns: []string{bookborrow.SERVICEPOINTColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: servicepoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	b = &Bookborrow{config: buo.config}
	_spec.Assign = b.assignValues
	_spec.ScanValues = b.scanValues()
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookborrow.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return b, nil
}