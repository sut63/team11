// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team11/app/ent/bookborrow"
	"github.com/team11/app/ent/bookreturn"
)

// Bookreturn is the model entity for the Bookreturn schema.
type Bookreturn struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BookName holds the value of the "book_name" field.
	BookName string `json:"book_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookreturnQuery when eager-loading is set.
	Edges       BookreturnEdges `json:"edges"`
	CLIENT_ID   *int
	location_id *int
}

// BookreturnEdges holds the relations/edges for other nodes in the graph.
type BookreturnEdges struct {
	// Mustreturn holds the value of the mustreturn edge.
	Mustreturn *Bookborrow
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MustreturnOrErr returns the Mustreturn value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookreturnEdges) MustreturnOrErr() (*Bookborrow, error) {
	if e.loadedTypes[0] {
		if e.Mustreturn == nil {
			// The edge mustreturn was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: bookborrow.Label}
		}
		return e.Mustreturn, nil
	}
	return nil, &NotLoadedError{edge: "mustreturn"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bookreturn) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // book_name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Bookreturn) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // CLIENT_ID
		&sql.NullInt64{}, // location_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bookreturn fields.
func (b *Bookreturn) assignValues(values ...interface{}) error {
	if m, n := len(values), len(bookreturn.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	b.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field book_name", values[0])
	} else if value.Valid {
		b.BookName = value.String
	}
	values = values[1:]
	if len(values) == len(bookreturn.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field CLIENT_ID", value)
		} else if value.Valid {
			b.CLIENT_ID = new(int)
			*b.CLIENT_ID = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field location_id", value)
		} else if value.Valid {
			b.location_id = new(int)
			*b.location_id = int(value.Int64)
		}
	}
	return nil
}

// QueryMustreturn queries the mustreturn edge of the Bookreturn.
func (b *Bookreturn) QueryMustreturn() *BookborrowQuery {
	return (&BookreturnClient{config: b.config}).QueryMustreturn(b)
}

// Update returns a builder for updating this Bookreturn.
// Note that, you need to call Bookreturn.Unwrap() before calling this method, if this Bookreturn
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bookreturn) Update() *BookreturnUpdateOne {
	return (&BookreturnClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Bookreturn) Unwrap() *Bookreturn {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bookreturn is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bookreturn) String() string {
	var builder strings.Builder
	builder.WriteString("Bookreturn(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", book_name=")
	builder.WriteString(b.BookName)
	builder.WriteByte(')')
	return builder.String()
}

// Bookreturns is a parsable slice of Bookreturn.
type Bookreturns []*Bookreturn

func (b Bookreturns) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
