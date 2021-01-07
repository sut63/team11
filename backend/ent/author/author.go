// Code generated by entc, DO NOT EDIT.

package author

const (
	// Label holds the string label denoting the author type in the database.
	Label = "author"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPosition holds the string denoting the position field in the database.
	FieldPosition = "position"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeWriter holds the string denoting the writer edge name in mutations.
	EdgeWriter = "writer"

	// Table holds the table name of the author in the database.
	Table = "authors"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "researches"
	// OwnerInverseTable is the table name for the Research entity.
	// It exists in this package in order to avoid circular dependency with the "research" package.
	OwnerInverseTable = "researches"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "OWNER_ID"
	// WriterTable is the table the holds the writer relation/edge.
	WriterTable = "books"
	// WriterInverseTable is the table name for the Book entity.
	// It exists in this package in order to avoid circular dependency with the "book" package.
	WriterInverseTable = "books"
	// WriterColumn is the table column denoting the writer relation/edge.
	WriterColumn = "AUTHOR_ID"
)

// Columns holds all SQL columns for author fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPosition,
}

var (
	// NameValidator is a validator for the "Name" field. It is called by the builders before save.
	NameValidator func(string) error
	// PositionValidator is a validator for the "Position" field. It is called by the builders before save.
	PositionValidator func(string) error
)