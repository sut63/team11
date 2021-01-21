// Code generated by entc, DO NOT EDIT.

package preemption

import (
	"time"
)

const (
	// Label holds the string label denoting the preemption type in the database.
	Label = "preemption"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPreemptTime holds the string denoting the preempttime field in the database.
	FieldPreemptTime = "preempt_time"
	// FieldPhonenumber holds the string denoting the phonenumber field in the database.
	FieldPhonenumber = "phonenumber"
	// FieldSurrogateid holds the string denoting the surrogateid field in the database.
	FieldSurrogateid = "surrogateid"
	// FieldSurrogatephone holds the string denoting the surrogatephone field in the database.
	FieldSurrogatephone = "surrogatephone"

	// EdgeUserID holds the string denoting the user_id edge name in mutations.
	EdgeUserID = "User_ID"
	// EdgePurposeID holds the string denoting the purposeid edge name in mutations.
	EdgePurposeID = "PurposeID"
	// EdgeRoomID holds the string denoting the roomid edge name in mutations.
	EdgeRoomID = "RoomID"

	// Table holds the table name of the preemption in the database.
	Table = "preemptions"
	// UserIDTable is the table the holds the User_ID relation/edge.
	UserIDTable = "preemptions"
	// UserIDInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserIDInverseTable = "users"
	// UserIDColumn is the table column denoting the User_ID relation/edge.
	UserIDColumn = "USER_ID"
	// PurposeIDTable is the table the holds the PurposeID relation/edge.
	PurposeIDTable = "preemptions"
	// PurposeIDInverseTable is the table name for the Purpose entity.
	// It exists in this package in order to avoid circular dependency with the "purpose" package.
	PurposeIDInverseTable = "purposes"
	// PurposeIDColumn is the table column denoting the PurposeID relation/edge.
	PurposeIDColumn = "PurposeID"
	// RoomIDTable is the table the holds the RoomID relation/edge.
	RoomIDTable = "preemptions"
	// RoomIDInverseTable is the table name for the Roominfo entity.
	// It exists in this package in order to avoid circular dependency with the "roominfo" package.
	RoomIDInverseTable = "roominfos"
	// RoomIDColumn is the table column denoting the RoomID relation/edge.
	RoomIDColumn = "RoomID"
)

// Columns holds all SQL columns for preemption fields.
var Columns = []string{
	FieldID,
	FieldPreemptTime,
	FieldPhonenumber,
	FieldSurrogateid,
	FieldSurrogatephone,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Preemption type.
var ForeignKeys = []string{
	"PurposeID",
	"RoomID",
	"USER_ID",
}

var (
	// DefaultPreemptTime holds the default value on creation for the PreemptTime field.
	DefaultPreemptTime func() time.Time
	// PhonenumberValidator is a validator for the "Phonenumber" field. It is called by the builders before save.
	PhonenumberValidator func(string) error
	// SurrogateidValidator is a validator for the "Surrogateid" field. It is called by the builders before save.
	SurrogateidValidator func(string) error
	// SurrogatephoneValidator is a validator for the "Surrogatephone" field. It is called by the builders before save.
	SurrogatephoneValidator func(string) error
)
