// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team11/app/ent/location"
)

// Location is the model entity for the Location schema.
type Location struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// LOCATIONNAME holds the value of the "LOCATION_NAME" field.
	LOCATIONNAME string `json:"LOCATION_NAME,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LocationQuery when eager-loading is set.
	Edges LocationEdges `json:"edges"`
}

// LocationEdges holds the relations/edges for other nodes in the graph.
type LocationEdges struct {
	// Returnfrom holds the value of the returnfrom edge.
	Returnfrom []*Bookreturn
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ReturnfromOrErr returns the Returnfrom value or an error if the edge
// was not loaded in eager-loading.
func (e LocationEdges) ReturnfromOrErr() ([]*Bookreturn, error) {
	if e.loadedTypes[0] {
		return e.Returnfrom, nil
	}
	return nil, &NotLoadedError{edge: "returnfrom"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Location) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // LOCATION_NAME
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Location fields.
func (l *Location) assignValues(values ...interface{}) error {
	if m, n := len(values), len(location.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	l.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field LOCATION_NAME", values[0])
	} else if value.Valid {
		l.LOCATIONNAME = value.String
	}
	return nil
}

// QueryReturnfrom queries the returnfrom edge of the Location.
func (l *Location) QueryReturnfrom() *BookreturnQuery {
	return (&LocationClient{config: l.config}).QueryReturnfrom(l)
}

// Update returns a builder for updating this Location.
// Note that, you need to call Location.Unwrap() before calling this method, if this Location
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Location) Update() *LocationUpdateOne {
	return (&LocationClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (l *Location) Unwrap() *Location {
	tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Location is not a transactional entity")
	}
	l.config.driver = tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Location) String() string {
	var builder strings.Builder
	builder.WriteString("Location(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteString(", LOCATION_NAME=")
	builder.WriteString(l.LOCATIONNAME)
	builder.WriteByte(')')
	return builder.String()
}

// Locations is a parsable slice of Location.
type Locations []*Location

func (l Locations) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
