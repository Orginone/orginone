// Code generated by entc, DO NOT EDIT.

package baseinfoadministrativeareaall

import (
	"orginone/common/tools/date"
)

const (
	// Label holds the string label denoting the baseinfoadministrativeareaall type in the database.
	Label = "baseinfoadministrativeareaall"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPid holds the string denoting the pid field in the database.
	FieldPid = "pid"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldProvince holds the string denoting the province field in the database.
	FieldProvince = "province"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldArea holds the string denoting the area field in the database.
	FieldArea = "area"
	// FieldTown holds the string denoting the town field in the database.
	FieldTown = "town"
	// FieldAllName holds the string denoting the all_name field in the database.
	FieldAllName = "all_name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldTsVersion holds the string denoting the ts_version field in the database.
	FieldTsVersion = "ts_version"
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
	// EdgeParentx holds the string denoting the parentx edge name in mutations.
	EdgeParentx = "parentx"
	// EdgeChildrens holds the string denoting the childrens edge name in mutations.
	EdgeChildrens = "childrens"
	// Table holds the table name of the baseinfoadministrativeareaall in the database.
	Table = "baseinfo_administrative_area_all"
	// ParentxTable is the table that holds the parentx relation/edge.
	ParentxTable = "baseinfo_administrative_area_all"
	// ParentxColumn is the table column denoting the parentx relation/edge.
	ParentxColumn = "pid"
	// ChildrensTable is the table that holds the childrens relation/edge.
	ChildrensTable = "baseinfo_administrative_area_all"
	// ChildrensColumn is the table column denoting the childrens relation/edge.
	ChildrensColumn = "pid"
)

// Columns holds all SQL columns for baseinfoadministrativeareaall fields.
var Columns = []string{
	FieldID,
	FieldPid,
	FieldCode,
	FieldName,
	FieldProvince,
	FieldCity,
	FieldArea,
	FieldTown,
	FieldAllName,
	FieldType,
	FieldTsVersion,
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
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType int32
	// DefaultTsVersion holds the default value on creation for the "ts_version" field.
	DefaultTsVersion int32
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
