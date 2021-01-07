// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team11/app/ent/researchtype"
)

// Researchtype is the model entity for the Researchtype schema.
type Researchtype struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TYPENAME holds the value of the "TYPE_NAME" field.
	TYPENAME string `json:"TYPE_NAME,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResearchtypeQuery when eager-loading is set.
	Edges ResearchtypeEdges `json:"edges"`
}

// ResearchtypeEdges holds the relations/edges for other nodes in the graph.
type ResearchtypeEdges struct {
	// ResearchType holds the value of the researchType edge.
	ResearchType []*Research
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ResearchTypeOrErr returns the ResearchType value or an error if the edge
// was not loaded in eager-loading.
func (e ResearchtypeEdges) ResearchTypeOrErr() ([]*Research, error) {
	if e.loadedTypes[0] {
		return e.ResearchType, nil
	}
	return nil, &NotLoadedError{edge: "researchType"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Researchtype) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // TYPE_NAME
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Researchtype fields.
func (r *Researchtype) assignValues(values ...interface{}) error {
	if m, n := len(values), len(researchtype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field TYPE_NAME", values[0])
	} else if value.Valid {
		r.TYPENAME = value.String
	}
	return nil
}

// QueryResearchType queries the researchType edge of the Researchtype.
func (r *Researchtype) QueryResearchType() *ResearchQuery {
	return (&ResearchtypeClient{config: r.config}).QueryResearchType(r)
}

// Update returns a builder for updating this Researchtype.
// Note that, you need to call Researchtype.Unwrap() before calling this method, if this Researchtype
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Researchtype) Update() *ResearchtypeUpdateOne {
	return (&ResearchtypeClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Researchtype) Unwrap() *Researchtype {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Researchtype is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Researchtype) String() string {
	var builder strings.Builder
	builder.WriteString("Researchtype(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", TYPE_NAME=")
	builder.WriteString(r.TYPENAME)
	builder.WriteByte(')')
	return builder.String()
}

// Researchtypes is a parsable slice of Researchtype.
type Researchtypes []*Researchtype

func (r Researchtypes) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
