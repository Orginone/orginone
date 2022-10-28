// Code generated by entc, DO NOT EDIT.

package asmarketappcomponenttemplate

import (
	"orginone/common/tools/date"
)

const (
	// Label holds the string label denoting the asmarketappcomponenttemplate type in the database.
	Label = "as_market_app_component_template"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"
	// FieldIsDefault holds the string denoting the is_default field in the database.
	FieldIsDefault = "is_default"
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
	// EdgeAppUserTemplates holds the string denoting the appusertemplates edge name in mutations.
	EdgeAppUserTemplates = "appUserTemplates"
	// Table holds the table name of the asmarketappcomponenttemplate in the database.
	Table = "as_market_app_component_template"
	// AppUserTemplatesTable is the table that holds the appUserTemplates relation/edge.
	AppUserTemplatesTable = "as_market_app_user_template"
	// AppUserTemplatesInverseTable is the table name for the AsMarketAppUserTemplate entity.
	// It exists in this package in order to avoid circular dependency with the "asmarketappusertemplate" package.
	AppUserTemplatesInverseTable = "as_market_app_user_template"
	// AppUserTemplatesColumn is the table column denoting the appUserTemplates relation/edge.
	AppUserTemplatesColumn = "template_id"
)

// Columns holds all SQL columns for asmarketappcomponenttemplate fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldConfig,
	FieldIsDefault,
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
