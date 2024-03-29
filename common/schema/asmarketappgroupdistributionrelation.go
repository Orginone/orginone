// Code generated by entc, DO NOT EDIT.

package schema

import (
	"fmt"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketappgroupdistributionrelation"
	"orginone/common/tools/date"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// AsMarketAppGroupDistributionRelation is the model entity for the AsMarketAppGroupDistributionRelation schema.
type AsMarketAppGroupDistributionRelation struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID int64 `json:"id,string"`
	// AppID holds the value of the "app_id" field.
	// 应用id
	AppID int64 `json:"appId"`
	// RelationID holds the value of the "relation_id" field.
	// 被分配租户id
	RelationID int64 `json:"relationId"`
	// GroupID holds the value of the "group_id" field.
	// 集团id
	GroupID int64 `json:"groupId"`
	// IsDeleted holds the value of the "is_deleted" field.
	// 是否被删除
	IsDeleted int64 `json:"isDeleted"`
	// Status holds the value of the "status" field.
	// 状态
	Status int64 `json:"status"`
	// CreateUser holds the value of the "create_user" field.
	// 创建用户
	CreateUser int64 `json:"createUser"`
	// UpdateUser holds the value of the "update_user" field.
	// 更新用户
	UpdateUser int64 `json:"updateUser"`
	// CreateTime holds the value of the "create_time" field.
	// 创建时间
	CreateTime date.DateTime `json:"createTime"`
	// UpdateTime holds the value of the "update_time" field.
	// 更新时间
	UpdateTime date.DateTime `json:"updateTime"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AsMarketAppGroupDistributionRelationQuery when eager-loading is set.
	Edges AsMarketAppGroupDistributionRelationEdges `json:"edges"`
}

// AsMarketAppGroupDistributionRelationEdges holds the relations/edges for other nodes in the graph.
type AsMarketAppGroupDistributionRelationEdges struct {
	// Appx holds the value of the appx edge.
	Appx *AsMarketApp `json:"appx"`
	// Groupx holds the value of the groupx edge.
	Groupx *AsAllGroup `json:"groupx"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AppxOrErr returns the Appx value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AsMarketAppGroupDistributionRelationEdges) AppxOrErr() (*AsMarketApp, error) {
	if e.loadedTypes[0] {
		if e.Appx == nil {
			// The edge appx was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: asmarketapp.Label}
		}
		return e.Appx, nil
	}
	return nil, &NotLoadedError{edge: "appx"}
}

// GroupxOrErr returns the Groupx value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AsMarketAppGroupDistributionRelationEdges) GroupxOrErr() (*AsAllGroup, error) {
	if e.loadedTypes[1] {
		if e.Groupx == nil {
			// The edge groupx was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: asallgroup.Label}
		}
		return e.Groupx, nil
	}
	return nil, &NotLoadedError{edge: "groupx"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AsMarketAppGroupDistributionRelation) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case asmarketappgroupdistributionrelation.FieldID, asmarketappgroupdistributionrelation.FieldAppID, asmarketappgroupdistributionrelation.FieldRelationID, asmarketappgroupdistributionrelation.FieldGroupID, asmarketappgroupdistributionrelation.FieldIsDeleted, asmarketappgroupdistributionrelation.FieldStatus, asmarketappgroupdistributionrelation.FieldCreateUser, asmarketappgroupdistributionrelation.FieldUpdateUser:
			values[i] = new(sql.NullInt64)
		case asmarketappgroupdistributionrelation.FieldCreateTime, asmarketappgroupdistributionrelation.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AsMarketAppGroupDistributionRelation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AsMarketAppGroupDistributionRelation fields.
func (amagdr *AsMarketAppGroupDistributionRelation) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case asmarketappgroupdistributionrelation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			amagdr.ID = int64(value.Int64)
		case asmarketappgroupdistributionrelation.FieldAppID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				amagdr.AppID = value.Int64
			}
		case asmarketappgroupdistributionrelation.FieldRelationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field relation_id", values[i])
			} else if value.Valid {
				amagdr.RelationID = value.Int64
			}
		case asmarketappgroupdistributionrelation.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				amagdr.GroupID = value.Int64
			}
		case asmarketappgroupdistributionrelation.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				amagdr.IsDeleted = value.Int64
			}
		case asmarketappgroupdistributionrelation.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				amagdr.Status = value.Int64
			}
		case asmarketappgroupdistributionrelation.FieldCreateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_user", values[i])
			} else if value.Valid {
				amagdr.CreateUser = value.Int64
			}
		case asmarketappgroupdistributionrelation.FieldUpdateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_user", values[i])
			} else if value.Valid {
				amagdr.UpdateUser = value.Int64
			}
		case asmarketappgroupdistributionrelation.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				amagdr.CreateTime = date.DateTime(value.Time)
			}
		case asmarketappgroupdistributionrelation.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				amagdr.UpdateTime = date.DateTime(value.Time)
			}
		}
	}
	return nil
}

// QueryAppx queries the "appx" edge of the AsMarketAppGroupDistributionRelation entity.
func (amagdr *AsMarketAppGroupDistributionRelation) QueryAppx() *AsMarketAppQuery {
	return (&AsMarketAppGroupDistributionRelationClient{config: amagdr.config}).QueryAppx(amagdr)
}

// QueryGroupx queries the "groupx" edge of the AsMarketAppGroupDistributionRelation entity.
func (amagdr *AsMarketAppGroupDistributionRelation) QueryGroupx() *AsAllGroupQuery {
	return (&AsMarketAppGroupDistributionRelationClient{config: amagdr.config}).QueryGroupx(amagdr)
}

// Update returns a builder for updating this AsMarketAppGroupDistributionRelation.
// Note that you need to call AsMarketAppGroupDistributionRelation.Unwrap() before calling this method if this AsMarketAppGroupDistributionRelation
// was returned from a transaction, and the transaction was committed or rolled back.
func (amagdr *AsMarketAppGroupDistributionRelation) Update() *AsMarketAppGroupDistributionRelationUpdateOne {
	return (&AsMarketAppGroupDistributionRelationClient{config: amagdr.config}).UpdateOne(amagdr)
}

// Unwrap unwraps the AsMarketAppGroupDistributionRelation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (amagdr *AsMarketAppGroupDistributionRelation) Unwrap() *AsMarketAppGroupDistributionRelation {
	tx, ok := amagdr.config.driver.(*txDriver)
	if !ok {
		panic("schema: AsMarketAppGroupDistributionRelation is not a transactional entity")
	}
	amagdr.config.driver = tx.drv
	return amagdr
}

// String implements the fmt.Stringer.
func (amagdr *AsMarketAppGroupDistributionRelation) String() string {
	var builder strings.Builder
	builder.WriteString("AsMarketAppGroupDistributionRelation(")
	builder.WriteString(fmt.Sprintf("id=%v", amagdr.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", amagdr.AppID))
	builder.WriteString(", relation_id=")
	builder.WriteString(fmt.Sprintf("%v", amagdr.RelationID))
	builder.WriteString(", group_id=")
	builder.WriteString(fmt.Sprintf("%v", amagdr.GroupID))
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", amagdr.IsDeleted))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", amagdr.Status))
	builder.WriteString(", create_user=")
	builder.WriteString(fmt.Sprintf("%v", amagdr.CreateUser))
	builder.WriteString(", update_user=")
	builder.WriteString(fmt.Sprintf("%v", amagdr.UpdateUser))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", amagdr.CreateTime))
	builder.WriteString(", update_time=")
	builder.WriteString(fmt.Sprintf("%v", amagdr.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// AsMarketAppGroupDistributionRelations is a parsable slice of AsMarketAppGroupDistributionRelation.
type AsMarketAppGroupDistributionRelations []*AsMarketAppGroupDistributionRelation

func (amagdr AsMarketAppGroupDistributionRelations) config(cfg config) {
	for _i := range amagdr {
		amagdr[_i].config = cfg
	}
}
