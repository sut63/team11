// Code generated by entc, DO NOT EDIT.

package status

const (
	// Label holds the string label denoting the status type in the database.
	Label = "status"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSTATUSNAME holds the string denoting the status_name field in the database.
	FieldSTATUSNAME = "status_name"

	// EdgeStatus holds the string denoting the status edge name in mutations.
	EdgeStatus = "status"
	// EdgeStatusofbook holds the string denoting the statusofbook edge name in mutations.
	EdgeStatusofbook = "statusofbook"
	// EdgeStatusbookborrow holds the string denoting the statusbookborrow edge name in mutations.
	EdgeStatusbookborrow = "statusbookborrow"

	// Table holds the table name of the status in the database.
	Table = "status"
	// StatusTable is the table the holds the status relation/edge.
	StatusTable = "client_entities"
	// StatusInverseTable is the table name for the ClientEntity entity.
	// It exists in this package in order to avoid circular dependency with the "cliententity" package.
	StatusInverseTable = "client_entities"
	// StatusColumn is the table column denoting the status relation/edge.
	StatusColumn = "STATUS_ID"
	// StatusofbookTable is the table the holds the statusofbook relation/edge.
	StatusofbookTable = "books"
	// StatusofbookInverseTable is the table name for the Book entity.
	// It exists in this package in order to avoid circular dependency with the "book" package.
	StatusofbookInverseTable = "books"
	// StatusofbookColumn is the table column denoting the statusofbook relation/edge.
	StatusofbookColumn = "STATUS_ID"
	// StatusbookborrowTable is the table the holds the statusbookborrow relation/edge.
	StatusbookborrowTable = "bookborrows"
	// StatusbookborrowInverseTable is the table name for the Bookborrow entity.
	// It exists in this package in order to avoid circular dependency with the "bookborrow" package.
	StatusbookborrowInverseTable = "bookborrows"
	// StatusbookborrowColumn is the table column denoting the statusbookborrow relation/edge.
	StatusbookborrowColumn = "STATUS_ID"
)

// Columns holds all SQL columns for status fields.
var Columns = []string{
	FieldID,
	FieldSTATUSNAME,
}

var (
	// STATUSNAMEValidator is a validator for the "STATUS_NAME" field. It is called by the builders before save.
	STATUSNAMEValidator func(string) error
)
