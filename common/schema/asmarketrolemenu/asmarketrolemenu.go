// Code generated by entc, DO NOT EDIT.

package asmarketrolemenu

import (
	"orginone/common/tools/date"
)

const (
	// Label holds the string label denoting the asmarketrolemenu type in the database.
	Label = "as_market_role_menu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldMenuID holds the string denoting the menu_id field in the database.
	FieldMenuID = "menu_id"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreateUser holds the string denoting the create_user field in the database.
	FieldCreateUser = "create_user"
	// FieldUpdateUser holds the string denoting the update_user field in the database.
	FieldUpdateUser = "update_user"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// EdgeMenux holds the string denoting the menux edge name in mutations.
	EdgeMenux = "menux"
	// EdgeRolex holds the string denoting the rolex edge name in mutations.
	EdgeRolex = "rolex"
	// Table holds the table name of the asmarketrolemenu in the database.
	Table = "as_market_role_menu"
	// MenuxTable is the table that holds the menux relation/edge.
	MenuxTable = "as_market_role_menu"
	// MenuxInverseTable is the table name for the AsMarketMenu entity.
	// It exists in this package in order to avoid circular dependency with the "asmarketmenu" package.
	MenuxInverseTable = "as_market_menu"
	// MenuxColumn is the table column denoting the menux relation/edge.
	MenuxColumn = "menu_id"
	// RolexTable is the table that holds the rolex relation/edge.
	RolexTable = "as_market_role_menu"
	// RolexInverseTable is the table name for the AsMarketAppRole entity.
	// It exists in this package in order to avoid circular dependency with the "asmarketapprole" package.
	RolexInverseTable = "as_market_app_role"
	// RolexColumn is the table column denoting the rolex relation/edge.
	RolexColumn = "role_id"
)

// Columns holds all SQL columns for asmarketrolemenu fields.
var Columns = []string{
	FieldID,
	FieldRoleID,
	FieldMenuID,
	FieldIsDeleted,
	FieldStatus,
	FieldCreateUser,
	FieldUpdateUser,
	FieldCreateTime,
	FieldUpdateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted int64
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int64
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() date.DateTime
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() date.DateTime
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() date.DateTime
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)
