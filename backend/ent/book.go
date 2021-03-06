// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/team11/app/ent/author"
	"github.com/team11/app/ent/book"
	"github.com/team11/app/ent/category"
	"github.com/team11/app/ent/status"
	"github.com/team11/app/ent/user"
)

// Book is the model entity for the Book schema.
type Book struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BookName holds the value of the "BookName" field.
	BookName string `json:"BookName,omitempty"`
	// Barcode holds the value of the "Barcode" field.
	Barcode string `json:"Barcode,omitempty"`
	// BookPage holds the value of the "BookPage" field.
	BookPage int `json:"BookPage,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookQuery when eager-loading is set.
	Edges       BookEdges `json:"edges"`
	AUTHOR_ID   *int
	Category_id *int
	STATUS_ID   *int
	USER_ID     *int
}

// BookEdges holds the relations/edges for other nodes in the graph.
type BookEdges struct {
	// Category holds the value of the category edge.
	Category *Category
	// Author holds the value of the author edge.
	Author *Author
	// User holds the value of the user edge.
	User *User
	// Status holds the value of the status edge.
	Status *Status
	// Booklist holds the value of the booklist edge.
	Booklist []*Bookborrow
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookEdges) CategoryOrErr() (*Category, error) {
	if e.loadedTypes[0] {
		if e.Category == nil {
			// The edge category was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookEdges) AuthorOrErr() (*Author, error) {
	if e.loadedTypes[1] {
		if e.Author == nil {
			// The edge author was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: author.Label}
		}
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// StatusOrErr returns the Status value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookEdges) StatusOrErr() (*Status, error) {
	if e.loadedTypes[3] {
		if e.Status == nil {
			// The edge status was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: status.Label}
		}
		return e.Status, nil
	}
	return nil, &NotLoadedError{edge: "status"}
}

// BooklistOrErr returns the Booklist value or an error if the edge
// was not loaded in eager-loading.
func (e BookEdges) BooklistOrErr() ([]*Bookborrow, error) {
	if e.loadedTypes[4] {
		return e.Booklist, nil
	}
	return nil, &NotLoadedError{edge: "booklist"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Book) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // BookName
		&sql.NullString{}, // Barcode
		&sql.NullInt64{},  // BookPage
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Book) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // AUTHOR_ID
		&sql.NullInt64{}, // Category_id
		&sql.NullInt64{}, // STATUS_ID
		&sql.NullInt64{}, // USER_ID
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Book fields.
func (b *Book) assignValues(values ...interface{}) error {
	if m, n := len(values), len(book.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	b.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field BookName", values[0])
	} else if value.Valid {
		b.BookName = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Barcode", values[1])
	} else if value.Valid {
		b.Barcode = value.String
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field BookPage", values[2])
	} else if value.Valid {
		b.BookPage = int(value.Int64)
	}
	values = values[3:]
	if len(values) == len(book.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field AUTHOR_ID", value)
		} else if value.Valid {
			b.AUTHOR_ID = new(int)
			*b.AUTHOR_ID = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field Category_id", value)
		} else if value.Valid {
			b.Category_id = new(int)
			*b.Category_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field STATUS_ID", value)
		} else if value.Valid {
			b.STATUS_ID = new(int)
			*b.STATUS_ID = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field USER_ID", value)
		} else if value.Valid {
			b.USER_ID = new(int)
			*b.USER_ID = int(value.Int64)
		}
	}
	return nil
}

// QueryCategory queries the category edge of the Book.
func (b *Book) QueryCategory() *CategoryQuery {
	return (&BookClient{config: b.config}).QueryCategory(b)
}

// QueryAuthor queries the author edge of the Book.
func (b *Book) QueryAuthor() *AuthorQuery {
	return (&BookClient{config: b.config}).QueryAuthor(b)
}

// QueryUser queries the user edge of the Book.
func (b *Book) QueryUser() *UserQuery {
	return (&BookClient{config: b.config}).QueryUser(b)
}

// QueryStatus queries the status edge of the Book.
func (b *Book) QueryStatus() *StatusQuery {
	return (&BookClient{config: b.config}).QueryStatus(b)
}

// QueryBooklist queries the booklist edge of the Book.
func (b *Book) QueryBooklist() *BookborrowQuery {
	return (&BookClient{config: b.config}).QueryBooklist(b)
}

// Update returns a builder for updating this Book.
// Note that, you need to call Book.Unwrap() before calling this method, if this Book
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Book) Update() *BookUpdateOne {
	return (&BookClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Book) Unwrap() *Book {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Book is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Book) String() string {
	var builder strings.Builder
	builder.WriteString("Book(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", BookName=")
	builder.WriteString(b.BookName)
	builder.WriteString(", Barcode=")
	builder.WriteString(b.Barcode)
	builder.WriteString(", BookPage=")
	builder.WriteString(fmt.Sprintf("%v", b.BookPage))
	builder.WriteByte(')')
	return builder.String()
}

// Books is a parsable slice of Book.
type Books []*Book

func (b Books) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
