// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team11/app/ent/bookborrow"
	"github.com/team11/app/ent/bookreturn"
	"github.com/team11/app/ent/location"
	"github.com/team11/app/ent/user"
)

// Bookreturn is the model entity for the Bookreturn schema.
type Bookreturn struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RETURNTIME holds the value of the "RETURN_TIME" field.
	RETURNTIME time.Time `json:"RETURN_TIME,omitempty"`
	// DAMAGEDPOINT holds the value of the "DAMAGED_POINT" field.
	DAMAGEDPOINT int `json:"DAMAGED_POINT,omitempty"`
	// DAMAGEDPOINTNAME holds the value of the "DAMAGED_POINTNAME" field.
	DAMAGEDPOINTNAME string `json:"DAMAGED_POINTNAME,omitempty"`
	// LOST holds the value of the "LOST" field.
	LOST string `json:"LOST,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookreturnQuery when eager-loading is set.
	Edges       BookreturnEdges `json:"edges"`
	CLIENT_ID   *int
	LOCATION_ID *int
	USER_ID     *int
}

// BookreturnEdges holds the relations/edges for other nodes in the graph.
type BookreturnEdges struct {
	// User holds the value of the user edge.
	User *User
	// Location holds the value of the location edge.
	Location *Location
	// Mustreturn holds the value of the mustreturn edge.
	Mustreturn *Bookborrow
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookreturnEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// LocationOrErr returns the Location value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookreturnEdges) LocationOrErr() (*Location, error) {
	if e.loadedTypes[1] {
		if e.Location == nil {
			// The edge location was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: location.Label}
		}
		return e.Location, nil
	}
	return nil, &NotLoadedError{edge: "location"}
}

// MustreturnOrErr returns the Mustreturn value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookreturnEdges) MustreturnOrErr() (*Bookborrow, error) {
	if e.loadedTypes[2] {
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
		&sql.NullTime{},   // RETURN_TIME
		&sql.NullInt64{},  // DAMAGED_POINT
		&sql.NullString{}, // DAMAGED_POINTNAME
		&sql.NullString{}, // LOST
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Bookreturn) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // CLIENT_ID
		&sql.NullInt64{}, // LOCATION_ID
		&sql.NullInt64{}, // USER_ID
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
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field RETURN_TIME", values[0])
	} else if value.Valid {
		b.RETURNTIME = value.Time
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field DAMAGED_POINT", values[1])
	} else if value.Valid {
		b.DAMAGEDPOINT = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field DAMAGED_POINTNAME", values[2])
	} else if value.Valid {
		b.DAMAGEDPOINTNAME = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field LOST", values[3])
	} else if value.Valid {
		b.LOST = value.String
	}
	values = values[4:]
	if len(values) == len(bookreturn.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field CLIENT_ID", value)
		} else if value.Valid {
			b.CLIENT_ID = new(int)
			*b.CLIENT_ID = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field LOCATION_ID", value)
		} else if value.Valid {
			b.LOCATION_ID = new(int)
			*b.LOCATION_ID = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field USER_ID", value)
		} else if value.Valid {
			b.USER_ID = new(int)
			*b.USER_ID = int(value.Int64)
		}
	}
	return nil
}

// QueryUser queries the user edge of the Bookreturn.
func (b *Bookreturn) QueryUser() *UserQuery {
	return (&BookreturnClient{config: b.config}).QueryUser(b)
}

// QueryLocation queries the location edge of the Bookreturn.
func (b *Bookreturn) QueryLocation() *LocationQuery {
	return (&BookreturnClient{config: b.config}).QueryLocation(b)
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
	builder.WriteString(", RETURN_TIME=")
	builder.WriteString(b.RETURNTIME.Format(time.ANSIC))
	builder.WriteString(", DAMAGED_POINT=")
	builder.WriteString(fmt.Sprintf("%v", b.DAMAGEDPOINT))
	builder.WriteString(", DAMAGED_POINTNAME=")
	builder.WriteString(b.DAMAGEDPOINTNAME)
	builder.WriteString(", LOST=")
	builder.WriteString(b.LOST)
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
