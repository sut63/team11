// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// AuthorsColumns holds the columns for the "authors" table.
	AuthorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// AuthorsTable holds the schema information for the "authors" table.
	AuthorsTable = &schema.Table{
		Name:        "authors",
		Columns:     AuthorsColumns,
		PrimaryKey:  []*schema.Column{AuthorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// BooksColumns holds the columns for the "books" table.
	BooksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "book_name", Type: field.TypeString, Unique: true},
		{Name: "AUTHOR_ID", Type: field.TypeInt, Nullable: true},
		{Name: "Category_id", Type: field.TypeInt, Nullable: true},
		{Name: "STATUS_ID", Type: field.TypeInt, Nullable: true},
		{Name: "USER_ID", Type: field.TypeInt, Nullable: true},
	}
	// BooksTable holds the schema information for the "books" table.
	BooksTable = &schema.Table{
		Name:       "books",
		Columns:    BooksColumns,
		PrimaryKey: []*schema.Column{BooksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "books_authors_writer",
				Columns: []*schema.Column{BooksColumns[2]},

				RefColumns: []*schema.Column{AuthorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "books_categories_catof",
				Columns: []*schema.Column{BooksColumns[3]},

				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "books_status_statusofbook",
				Columns: []*schema.Column{BooksColumns[4]},

				RefColumns: []*schema.Column{StatusColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "books_users_addby",
				Columns: []*schema.Column{BooksColumns[5]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BookborrowsColumns holds the columns for the "bookborrows" table.
	BookborrowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "borrow_date", Type: field.TypeTime},
		{Name: "BOOK_ID", Type: field.TypeInt, Nullable: true},
		{Name: "SERVICEPOINT_ID", Type: field.TypeInt, Nullable: true},
		{Name: "USER_ID", Type: field.TypeInt, Nullable: true},
	}
	// BookborrowsTable holds the schema information for the "bookborrows" table.
	BookborrowsTable = &schema.Table{
		Name:       "bookborrows",
		Columns:    BookborrowsColumns,
		PrimaryKey: []*schema.Column{BookborrowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "bookborrows_books_booklist",
				Columns: []*schema.Column{BookborrowsColumns[2]},

				RefColumns: []*schema.Column{BooksColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bookborrows_service_points_from",
				Columns: []*schema.Column{BookborrowsColumns[3]},

				RefColumns: []*schema.Column{ServicePointsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bookborrows_users_borrow",
				Columns: []*schema.Column{BookborrowsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BookingsColumns holds the columns for the "bookings" table.
	BookingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "booking_date", Type: field.TypeTime},
		{Name: "time_left", Type: field.TypeTime},
		{Name: "CLIENT_ID", Type: field.TypeInt, Nullable: true},
		{Name: "SERVICEPOINT_ID", Type: field.TypeInt, Nullable: true},
		{Name: "USER_ID", Type: field.TypeInt, Nullable: true},
	}
	// BookingsTable holds the schema information for the "bookings" table.
	BookingsTable = &schema.Table{
		Name:       "bookings",
		Columns:    BookingsColumns,
		PrimaryKey: []*schema.Column{BookingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "bookings_client_entities_booked",
				Columns: []*schema.Column{BookingsColumns[3]},

				RefColumns: []*schema.Column{ClientEntitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bookings_service_points_servicepoint",
				Columns: []*schema.Column{BookingsColumns[4]},

				RefColumns: []*schema.Column{ServicePointsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bookings_users_booking",
				Columns: []*schema.Column{BookingsColumns[5]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BookreturnsColumns holds the columns for the "bookreturns" table.
	BookreturnsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "return_time", Type: field.TypeTime},
		{Name: "CLIENT_ID", Type: field.TypeInt, Nullable: true},
		{Name: "LOCATION_ID", Type: field.TypeInt, Nullable: true},
		{Name: "USER_ID", Type: field.TypeInt, Nullable: true},
	}
	// BookreturnsTable holds the schema information for the "bookreturns" table.
	BookreturnsTable = &schema.Table{
		Name:       "bookreturns",
		Columns:    BookreturnsColumns,
		PrimaryKey: []*schema.Column{BookreturnsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "bookreturns_bookborrows_borrowed",
				Columns: []*schema.Column{BookreturnsColumns[2]},

				RefColumns: []*schema.Column{BookborrowsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bookreturns_locations_returnfrom",
				Columns: []*schema.Column{BookreturnsColumns[3]},

				RefColumns: []*schema.Column{LocationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bookreturns_users_return",
				Columns: []*schema.Column{BookreturnsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "category_name", Type: field.TypeString, Unique: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:        "categories",
		Columns:     CategoriesColumns,
		PrimaryKey:  []*schema.Column{CategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ClientEntitiesColumns holds the columns for the "client_entities" table.
	ClientEntitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "client_name", Type: field.TypeString, Unique: true},
		{Name: "STATUS_ID", Type: field.TypeInt, Nullable: true},
	}
	// ClientEntitiesTable holds the schema information for the "client_entities" table.
	ClientEntitiesTable = &schema.Table{
		Name:       "client_entities",
		Columns:    ClientEntitiesColumns,
		PrimaryKey: []*schema.Column{ClientEntitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "client_entities_status_status",
				Columns: []*schema.Column{ClientEntitiesColumns[2]},

				RefColumns: []*schema.Column{StatusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// LocationsColumns holds the columns for the "locations" table.
	LocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "location_name", Type: field.TypeString, Unique: true},
	}
	// LocationsTable holds the schema information for the "locations" table.
	LocationsTable = &schema.Table{
		Name:        "locations",
		Columns:     LocationsColumns,
		PrimaryKey:  []*schema.Column{LocationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PreemptionsColumns holds the columns for the "preemptions" table.
	PreemptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "preempt_time", Type: field.TypeTime},
		{Name: "PurposeID", Type: field.TypeInt, Nullable: true},
		{Name: "RoomID", Type: field.TypeInt, Nullable: true},
		{Name: "USER_ID", Type: field.TypeInt, Nullable: true},
	}
	// PreemptionsTable holds the schema information for the "preemptions" table.
	PreemptionsTable = &schema.Table{
		Name:       "preemptions",
		Columns:    PreemptionsColumns,
		PrimaryKey: []*schema.Column{PreemptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "preemptions_purposes_preemption",
				Columns: []*schema.Column{PreemptionsColumns[2]},

				RefColumns: []*schema.Column{PurposesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "preemptions_roominfos_preemption",
				Columns: []*schema.Column{PreemptionsColumns[3]},

				RefColumns: []*schema.Column{RoominfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "preemptions_users_preemption",
				Columns: []*schema.Column{PreemptionsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PurposesColumns holds the columns for the "purposes" table.
	PurposesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "purpose_name", Type: field.TypeString, Unique: true},
	}
	// PurposesTable holds the schema information for the "purposes" table.
	PurposesTable = &schema.Table{
		Name:        "purposes",
		Columns:     PurposesColumns,
		PrimaryKey:  []*schema.Column{PurposesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ResearchesColumns holds the columns for the "researches" table.
	ResearchesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "doc_name", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime},
		{Name: "OWNER_ID", Type: field.TypeInt, Nullable: true},
		{Name: "TYPE_ID", Type: field.TypeInt, Nullable: true},
		{Name: "USER_ID", Type: field.TypeInt, Nullable: true},
	}
	// ResearchesTable holds the schema information for the "researches" table.
	ResearchesTable = &schema.Table{
		Name:       "researches",
		Columns:    ResearchesColumns,
		PrimaryKey: []*schema.Column{ResearchesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "researches_authors_owner",
				Columns: []*schema.Column{ResearchesColumns[3]},

				RefColumns: []*schema.Column{AuthorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "researches_researchtypes_researchType",
				Columns: []*schema.Column{ResearchesColumns[4]},

				RefColumns: []*schema.Column{ResearchtypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "researches_users_record",
				Columns: []*schema.Column{ResearchesColumns[5]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ResearchtypesColumns holds the columns for the "researchtypes" table.
	ResearchtypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type_name", Type: field.TypeString, Unique: true},
	}
	// ResearchtypesTable holds the schema information for the "researchtypes" table.
	ResearchtypesTable = &schema.Table{
		Name:        "researchtypes",
		Columns:     ResearchtypesColumns,
		PrimaryKey:  []*schema.Column{ResearchtypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role_name", Type: field.TypeString, Unique: true},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:        "roles",
		Columns:     RolesColumns,
		PrimaryKey:  []*schema.Column{RolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RoominfosColumns holds the columns for the "roominfos" table.
	RoominfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "room_id", Type: field.TypeString, Unique: true},
		{Name: "room_no", Type: field.TypeString},
		{Name: "room_type", Type: field.TypeString},
		{Name: "room_time", Type: field.TypeString},
		{Name: "room_status", Type: field.TypeString},
	}
	// RoominfosTable holds the schema information for the "roominfos" table.
	RoominfosTable = &schema.Table{
		Name:        "roominfos",
		Columns:     RoominfosColumns,
		PrimaryKey:  []*schema.Column{RoominfosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ServicePointsColumns holds the columns for the "service_points" table.
	ServicePointsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "counter_number", Type: field.TypeString, Unique: true},
	}
	// ServicePointsTable holds the schema information for the "service_points" table.
	ServicePointsTable = &schema.Table{
		Name:        "service_points",
		Columns:     ServicePointsColumns,
		PrimaryKey:  []*schema.Column{ServicePointsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// StatusColumns holds the columns for the "status" table.
	StatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status_name", Type: field.TypeString, Unique: true},
	}
	// StatusTable holds the schema information for the "status" table.
	StatusTable = &schema.Table{
		Name:        "status",
		Columns:     StatusColumns,
		PrimaryKey:  []*schema.Column{StatusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_email", Type: field.TypeString, Unique: true},
		{Name: "user_name", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "ROLE_ID", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_roles_role",
				Columns: []*schema.Column{UsersColumns[4]},

				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthorsTable,
		BooksTable,
		BookborrowsTable,
		BookingsTable,
		BookreturnsTable,
		CategoriesTable,
		ClientEntitiesTable,
		LocationsTable,
		PreemptionsTable,
		PurposesTable,
		ResearchesTable,
		ResearchtypesTable,
		RolesTable,
		RoominfosTable,
		ServicePointsTable,
		StatusTable,
		UsersTable,
	}
)

func init() {
	BooksTable.ForeignKeys[0].RefTable = AuthorsTable
	BooksTable.ForeignKeys[1].RefTable = CategoriesTable
	BooksTable.ForeignKeys[2].RefTable = StatusTable
	BooksTable.ForeignKeys[3].RefTable = UsersTable
	BookborrowsTable.ForeignKeys[0].RefTable = BooksTable
	BookborrowsTable.ForeignKeys[1].RefTable = ServicePointsTable
	BookborrowsTable.ForeignKeys[2].RefTable = UsersTable
	BookingsTable.ForeignKeys[0].RefTable = ClientEntitiesTable
	BookingsTable.ForeignKeys[1].RefTable = ServicePointsTable
	BookingsTable.ForeignKeys[2].RefTable = UsersTable
	BookreturnsTable.ForeignKeys[0].RefTable = BookborrowsTable
	BookreturnsTable.ForeignKeys[1].RefTable = LocationsTable
	BookreturnsTable.ForeignKeys[2].RefTable = UsersTable
	ClientEntitiesTable.ForeignKeys[0].RefTable = StatusTable
	PreemptionsTable.ForeignKeys[0].RefTable = PurposesTable
	PreemptionsTable.ForeignKeys[1].RefTable = RoominfosTable
	PreemptionsTable.ForeignKeys[2].RefTable = UsersTable
	ResearchesTable.ForeignKeys[0].RefTable = AuthorsTable
	ResearchesTable.ForeignKeys[1].RefTable = ResearchtypesTable
	ResearchesTable.ForeignKeys[2].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = RolesTable
}
