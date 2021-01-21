// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team11/app/ent/author"
	"github.com/team11/app/ent/research"
	"github.com/team11/app/ent/researchtype"
	"github.com/team11/app/ent/user"
)

// Research is the model entity for the Research schema.
type Research struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DOCNAME holds the value of the "DOC_NAME" field.
	DOCNAME string `json:"DOC_NAME,omitempty"`
	// DATE holds the value of the "DATE" field.
	DATE time.Time `json:"DATE,omitempty"`
	// PAGENUMBER holds the value of the "PAGE_NUMBER" field.
	PAGENUMBER int `json:"PAGE_NUMBER,omitempty"`
	// YEARNUMBER holds the value of the "YEAR_NUMBER" field.
	YEARNUMBER int `json:"YEAR_NUMBER,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResearchQuery when eager-loading is set.
	Edges    ResearchEdges `json:"edges"`
	OWNER_ID *int
	TYPE_ID  *int
	USER_ID  *int
}

// ResearchEdges holds the relations/edges for other nodes in the graph.
type ResearchEdges struct {
	// Register holds the value of the register edge.
	Register *User
	// MyDoc holds the value of the myDoc edge.
	MyDoc *Author
	// DocType holds the value of the docType edge.
	DocType *Researchtype
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// RegisterOrErr returns the Register value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResearchEdges) RegisterOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Register == nil {
			// The edge register was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Register, nil
	}
	return nil, &NotLoadedError{edge: "register"}
}

// MyDocOrErr returns the MyDoc value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResearchEdges) MyDocOrErr() (*Author, error) {
	if e.loadedTypes[1] {
		if e.MyDoc == nil {
			// The edge myDoc was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: author.Label}
		}
		return e.MyDoc, nil
	}
	return nil, &NotLoadedError{edge: "myDoc"}
}

// DocTypeOrErr returns the DocType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResearchEdges) DocTypeOrErr() (*Researchtype, error) {
	if e.loadedTypes[2] {
		if e.DocType == nil {
			// The edge docType was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: researchtype.Label}
		}
		return e.DocType, nil
	}
	return nil, &NotLoadedError{edge: "docType"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Research) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // DOC_NAME
		&sql.NullTime{},   // DATE
		&sql.NullInt64{},  // PAGE_NUMBER
		&sql.NullInt64{},  // YEAR_NUMBER
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Research) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // OWNER_ID
		&sql.NullInt64{}, // TYPE_ID
		&sql.NullInt64{}, // USER_ID
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Research fields.
func (r *Research) assignValues(values ...interface{}) error {
	if m, n := len(values), len(research.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field DOC_NAME", values[0])
	} else if value.Valid {
		r.DOCNAME = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field DATE", values[1])
	} else if value.Valid {
		r.DATE = value.Time
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field PAGE_NUMBER", values[2])
	} else if value.Valid {
		r.PAGENUMBER = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field YEAR_NUMBER", values[3])
	} else if value.Valid {
		r.YEARNUMBER = int(value.Int64)
	}
	values = values[4:]
	if len(values) == len(research.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field OWNER_ID", value)
		} else if value.Valid {
			r.OWNER_ID = new(int)
			*r.OWNER_ID = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field TYPE_ID", value)
		} else if value.Valid {
			r.TYPE_ID = new(int)
			*r.TYPE_ID = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field USER_ID", value)
		} else if value.Valid {
			r.USER_ID = new(int)
			*r.USER_ID = int(value.Int64)
		}
	}
	return nil
}

// QueryRegister queries the register edge of the Research.
func (r *Research) QueryRegister() *UserQuery {
	return (&ResearchClient{config: r.config}).QueryRegister(r)
}

// QueryMyDoc queries the myDoc edge of the Research.
func (r *Research) QueryMyDoc() *AuthorQuery {
	return (&ResearchClient{config: r.config}).QueryMyDoc(r)
}

// QueryDocType queries the docType edge of the Research.
func (r *Research) QueryDocType() *ResearchtypeQuery {
	return (&ResearchClient{config: r.config}).QueryDocType(r)
}

// Update returns a builder for updating this Research.
// Note that, you need to call Research.Unwrap() before calling this method, if this Research
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Research) Update() *ResearchUpdateOne {
	return (&ResearchClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Research) Unwrap() *Research {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Research is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Research) String() string {
	var builder strings.Builder
	builder.WriteString("Research(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", DOC_NAME=")
	builder.WriteString(r.DOCNAME)
	builder.WriteString(", DATE=")
	builder.WriteString(r.DATE.Format(time.ANSIC))
	builder.WriteString(", PAGE_NUMBER=")
	builder.WriteString(fmt.Sprintf("%v", r.PAGENUMBER))
	builder.WriteString(", YEAR_NUMBER=")
	builder.WriteString(fmt.Sprintf("%v", r.YEARNUMBER))
	builder.WriteByte(')')
	return builder.String()
}

// Researches is a parsable slice of Research.
type Researches []*Research

func (r Researches) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
