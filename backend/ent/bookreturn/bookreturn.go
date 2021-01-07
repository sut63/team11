// Code generated by entc, DO NOT EDIT.

package bookreturn

const (
	// Label holds the string label denoting the bookreturn type in the database.
	Label = "bookreturn"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBookName holds the string denoting the book_name field in the database.
	FieldBookName = "book_name"

	// Table holds the table name of the bookreturn in the database.
	Table = "bookreturns"
)

// Columns holds all SQL columns for bookreturn fields.
var Columns = []string{
	FieldID,
	FieldBookName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Bookreturn type.
var ForeignKeys = []string{
	"location_id",
}

var (
	// BookNameValidator is a validator for the "book_name" field. It is called by the builders before save.
	BookNameValidator func(string) error
)
