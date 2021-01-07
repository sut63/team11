// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUSEREMAIL holds the string denoting the user_email field in the database.
	FieldUSEREMAIL = "user_email"
	// FieldUSERNAME holds the string denoting the user_name field in the database.
	FieldUSERNAME = "user_name"
	// FieldPASSWORD holds the string denoting the password field in the database.
	FieldPASSWORD = "password"

	// EdgePosition holds the string denoting the position edge name in mutations.
	EdgePosition = "position"
	// EdgeBooking holds the string denoting the booking edge name in mutations.
	EdgeBooking = "booking"
	// EdgeAddby holds the string denoting the addby edge name in mutations.
	EdgeAddby = "addby"
	// EdgeBorrow holds the string denoting the borrow edge name in mutations.
	EdgeBorrow = "borrow"
	// EdgePreemption holds the string denoting the preemption edge name in mutations.
	EdgePreemption = "preemption"
	// EdgeRecord holds the string denoting the record edge name in mutations.
	EdgeRecord = "record"

	// Table holds the table name of the user in the database.
	Table = "users"
	// PositionTable is the table the holds the position relation/edge.
	PositionTable = "users"
	// PositionInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	PositionInverseTable = "roles"
	// PositionColumn is the table column denoting the position relation/edge.
	PositionColumn = "ROLE_ID"
	// BookingTable is the table the holds the booking relation/edge.
	BookingTable = "bookings"
	// BookingInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	BookingInverseTable = "bookings"
	// BookingColumn is the table column denoting the booking relation/edge.
	BookingColumn = "USER_ID"
	// AddbyTable is the table the holds the addby relation/edge.
	AddbyTable = "books"
	// AddbyInverseTable is the table name for the Book entity.
	// It exists in this package in order to avoid circular dependency with the "book" package.
	AddbyInverseTable = "books"
	// AddbyColumn is the table column denoting the addby relation/edge.
	AddbyColumn = "USER_ID"
	// BorrowTable is the table the holds the borrow relation/edge.
	BorrowTable = "bookborrows"
	// BorrowInverseTable is the table name for the Bookborrow entity.
	// It exists in this package in order to avoid circular dependency with the "bookborrow" package.
	BorrowInverseTable = "bookborrows"
	// BorrowColumn is the table column denoting the borrow relation/edge.
	BorrowColumn = "USER_ID"
	// PreemptionTable is the table the holds the preemption relation/edge.
	PreemptionTable = "preemptions"
	// PreemptionInverseTable is the table name for the Preemption entity.
	// It exists in this package in order to avoid circular dependency with the "preemption" package.
	PreemptionInverseTable = "preemptions"
	// PreemptionColumn is the table column denoting the preemption relation/edge.
	PreemptionColumn = "USER_ID"
	// RecordTable is the table the holds the record relation/edge.
	RecordTable = "researches"
	// RecordInverseTable is the table name for the Research entity.
	// It exists in this package in order to avoid circular dependency with the "research" package.
	RecordInverseTable = "researches"
	// RecordColumn is the table column denoting the record relation/edge.
	RecordColumn = "USER_ID"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUSEREMAIL,
	FieldUSERNAME,
	FieldPASSWORD,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the User type.
var ForeignKeys = []string{
	"ROLE_ID",
}

var (
	// USEREMAILValidator is a validator for the "USER_EMAIL" field. It is called by the builders before save.
	USEREMAILValidator func(string) error
	// USERNAMEValidator is a validator for the "USER_NAME" field. It is called by the builders before save.
	USERNAMEValidator func(string) error
	// PASSWORDValidator is a validator for the "PASSWORD" field. It is called by the builders before save.
	PASSWORDValidator func(string) error
)
