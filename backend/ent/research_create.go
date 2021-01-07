// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/author"
	"github.com/team11/app/ent/research"
	"github.com/team11/app/ent/researchtype"
	"github.com/team11/app/ent/user"
)

// ResearchCreate is the builder for creating a Research entity.
type ResearchCreate struct {
	config
	mutation *ResearchMutation
	hooks    []Hook
}

// SetDOCNAME sets the DOC_NAME field.
func (rc *ResearchCreate) SetDOCNAME(s string) *ResearchCreate {
	rc.mutation.SetDOCNAME(s)
	return rc
}

// SetDATE sets the DATE field.
func (rc *ResearchCreate) SetDATE(t time.Time) *ResearchCreate {
	rc.mutation.SetDATE(t)
	return rc
}

// SetNillableDATE sets the DATE field if the given value is not nil.
func (rc *ResearchCreate) SetNillableDATE(t *time.Time) *ResearchCreate {
	if t != nil {
		rc.SetDATE(*t)
	}
	return rc
}

// SetRegisterID sets the register edge to User by id.
func (rc *ResearchCreate) SetRegisterID(id int) *ResearchCreate {
	rc.mutation.SetRegisterID(id)
	return rc
}

// SetNillableRegisterID sets the register edge to User by id if the given value is not nil.
func (rc *ResearchCreate) SetNillableRegisterID(id *int) *ResearchCreate {
	if id != nil {
		rc = rc.SetRegisterID(*id)
	}
	return rc
}

// SetRegister sets the register edge to User.
func (rc *ResearchCreate) SetRegister(u *User) *ResearchCreate {
	return rc.SetRegisterID(u.ID)
}

// SetMyDocID sets the myDoc edge to Author by id.
func (rc *ResearchCreate) SetMyDocID(id int) *ResearchCreate {
	rc.mutation.SetMyDocID(id)
	return rc
}

// SetNillableMyDocID sets the myDoc edge to Author by id if the given value is not nil.
func (rc *ResearchCreate) SetNillableMyDocID(id *int) *ResearchCreate {
	if id != nil {
		rc = rc.SetMyDocID(*id)
	}
	return rc
}

// SetMyDoc sets the myDoc edge to Author.
func (rc *ResearchCreate) SetMyDoc(a *Author) *ResearchCreate {
	return rc.SetMyDocID(a.ID)
}

// SetDocTypeID sets the docType edge to Researchtype by id.
func (rc *ResearchCreate) SetDocTypeID(id int) *ResearchCreate {
	rc.mutation.SetDocTypeID(id)
	return rc
}

// SetNillableDocTypeID sets the docType edge to Researchtype by id if the given value is not nil.
func (rc *ResearchCreate) SetNillableDocTypeID(id *int) *ResearchCreate {
	if id != nil {
		rc = rc.SetDocTypeID(*id)
	}
	return rc
}

// SetDocType sets the docType edge to Researchtype.
func (rc *ResearchCreate) SetDocType(r *Researchtype) *ResearchCreate {
	return rc.SetDocTypeID(r.ID)
}

// Mutation returns the ResearchMutation object of the builder.
func (rc *ResearchCreate) Mutation() *ResearchMutation {
	return rc.mutation
}

// Save creates the Research in the database.
func (rc *ResearchCreate) Save(ctx context.Context) (*Research, error) {
	if _, ok := rc.mutation.DOCNAME(); !ok {
		return nil, &ValidationError{Name: "DOC_NAME", err: errors.New("ent: missing required field \"DOC_NAME\"")}
	}
	if v, ok := rc.mutation.DOCNAME(); ok {
		if err := research.DOCNAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "DOC_NAME", err: fmt.Errorf("ent: validator failed for field \"DOC_NAME\": %w", err)}
		}
	}
	if _, ok := rc.mutation.DATE(); !ok {
		v := research.DefaultDATE()
		rc.mutation.SetDATE(v)
	}
	var (
		err  error
		node *Research
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResearchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ResearchCreate) SaveX(ctx context.Context) *Research {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *ResearchCreate) sqlSave(ctx context.Context) (*Research, error) {
	r, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	r.ID = int(id)
	return r, nil
}

func (rc *ResearchCreate) createSpec() (*Research, *sqlgraph.CreateSpec) {
	var (
		r     = &Research{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: research.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: research.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.DOCNAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: research.FieldDOCNAME,
		})
		r.DOCNAME = value
	}
	if value, ok := rc.mutation.DATE(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: research.FieldDATE,
		})
		r.DATE = value
	}
	if nodes := rc.mutation.RegisterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.RegisterTable,
			Columns: []string{research.RegisterColumn},
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
	if nodes := rc.mutation.MyDocIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.MyDocTable,
			Columns: []string{research.MyDocColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: author.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.DocTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   research.DocTypeTable,
			Columns: []string{research.DocTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: researchtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return r, _spec
}
