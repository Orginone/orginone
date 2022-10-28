// Code generated by entc, DO NOT EDIT.

package schema

import (
	"fmt"
	"orginone/common/schema/baseinfoadministrativeareaall"
	"orginone/common/tools/date"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Baseinfoadministrativeareaall is the model entity for the Baseinfoadministrativeareaall schema.
type Baseinfoadministrativeareaall struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID int64 `json:"id,string"`
	// Pid holds the value of the "pid" field.
	// 上级行政区域
	Pid int64 `json:"pid"`
	// Code holds the value of the "code" field.
	// 行政区域编码
	Code string `json:"code"`
	// Name holds the value of the "name" field.
	// 行政区域名称
	Name string `json:"name"`
	// Province holds the value of the "province" field.
	// 省/直辖市
	Province string `json:"province"`
	// City holds the value of the "city" field.
	// 市
	City string `json:"city"`
	// Area holds the value of the "area" field.
	// 区
	Area string `json:"area"`
	// Town holds the value of the "town" field.
	// 城镇地区
	Town string `json:"town"`
	// AllName holds the value of the "all_name" field.
	// 全名
	AllName string `json:"allName"`
	// Type holds the value of the "type" field.
	// 行政区域级别
	Type int32 `json:"type"`
	// TsVersion holds the value of the "ts_version" field.
	// 乐观锁字段
	TsVersion int32 `json:"tsVersion"`
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
	// The values are being populated by the BaseinfoadministrativeareaallQuery when eager-loading is set.
	Edges BaseinfoadministrativeareaallEdges `json:"edges"`
}

// BaseinfoadministrativeareaallEdges holds the relations/edges for other nodes in the graph.
type BaseinfoadministrativeareaallEdges struct {
	// Parentx holds the value of the parentx edge.
	Parentx *Baseinfoadministrativeareaall `json:"parentx"`
	// Childrens holds the value of the childrens edge.
	Childrens []*Baseinfoadministrativeareaall `json:"childrens"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentxOrErr returns the Parentx value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BaseinfoadministrativeareaallEdges) ParentxOrErr() (*Baseinfoadministrativeareaall, error) {
	if e.loadedTypes[0] {
		if e.Parentx == nil {
			// The edge parentx was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: baseinfoadministrativeareaall.Label}
		}
		return e.Parentx, nil
	}
	return nil, &NotLoadedError{edge: "parentx"}
}

// ChildrensOrErr returns the Childrens value or an error if the edge
// was not loaded in eager-loading.
func (e BaseinfoadministrativeareaallEdges) ChildrensOrErr() ([]*Baseinfoadministrativeareaall, error) {
	if e.loadedTypes[1] {
		return e.Childrens, nil
	}
	return nil, &NotLoadedError{edge: "childrens"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Baseinfoadministrativeareaall) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case baseinfoadministrativeareaall.FieldID, baseinfoadministrativeareaall.FieldPid, baseinfoadministrativeareaall.FieldType, baseinfoadministrativeareaall.FieldTsVersion, baseinfoadministrativeareaall.FieldIsDeleted, baseinfoadministrativeareaall.FieldStatus, baseinfoadministrativeareaall.FieldCreateUser, baseinfoadministrativeareaall.FieldUpdateUser:
			values[i] = new(sql.NullInt64)
		case baseinfoadministrativeareaall.FieldCode, baseinfoadministrativeareaall.FieldName, baseinfoadministrativeareaall.FieldProvince, baseinfoadministrativeareaall.FieldCity, baseinfoadministrativeareaall.FieldArea, baseinfoadministrativeareaall.FieldTown, baseinfoadministrativeareaall.FieldAllName:
			values[i] = new(sql.NullString)
		case baseinfoadministrativeareaall.FieldCreateTime, baseinfoadministrativeareaall.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Baseinfoadministrativeareaall", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Baseinfoadministrativeareaall fields.
func (b *Baseinfoadministrativeareaall) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case baseinfoadministrativeareaall.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int64(value.Int64)
		case baseinfoadministrativeareaall.FieldPid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pid", values[i])
			} else if value.Valid {
				b.Pid = value.Int64
			}
		case baseinfoadministrativeareaall.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				b.Code = value.String
			}
		case baseinfoadministrativeareaall.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		case baseinfoadministrativeareaall.FieldProvince:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field province", values[i])
			} else if value.Valid {
				b.Province = value.String
			}
		case baseinfoadministrativeareaall.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				b.City = value.String
			}
		case baseinfoadministrativeareaall.FieldArea:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field area", values[i])
			} else if value.Valid {
				b.Area = value.String
			}
		case baseinfoadministrativeareaall.FieldTown:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field town", values[i])
			} else if value.Valid {
				b.Town = value.String
			}
		case baseinfoadministrativeareaall.FieldAllName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field all_name", values[i])
			} else if value.Valid {
				b.AllName = value.String
			}
		case baseinfoadministrativeareaall.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				b.Type = int32(value.Int64)
			}
		case baseinfoadministrativeareaall.FieldTsVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ts_version", values[i])
			} else if value.Valid {
				b.TsVersion = int32(value.Int64)
			}
		case baseinfoadministrativeareaall.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				b.IsDeleted = value.Int64
			}
		case baseinfoadministrativeareaall.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = value.Int64
			}
		case baseinfoadministrativeareaall.FieldCreateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_user", values[i])
			} else if value.Valid {
				b.CreateUser = value.Int64
			}
		case baseinfoadministrativeareaall.FieldUpdateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_user", values[i])
			} else if value.Valid {
				b.UpdateUser = value.Int64
			}
		case baseinfoadministrativeareaall.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				b.CreateTime = date.DateTime(value.Time)
			}
		case baseinfoadministrativeareaall.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				b.UpdateTime = date.DateTime(value.Time)
			}
		}
	}
	return nil
}

// QueryParentx queries the "parentx" edge of the Baseinfoadministrativeareaall entity.
func (b *Baseinfoadministrativeareaall) QueryParentx() *BaseinfoadministrativeareaallQuery {
	return (&BaseinfoadministrativeareaallClient{config: b.config}).QueryParentx(b)
}

// QueryChildrens queries the "childrens" edge of the Baseinfoadministrativeareaall entity.
func (b *Baseinfoadministrativeareaall) QueryChildrens() *BaseinfoadministrativeareaallQuery {
	return (&BaseinfoadministrativeareaallClient{config: b.config}).QueryChildrens(b)
}

// Update returns a builder for updating this Baseinfoadministrativeareaall.
// Note that you need to call Baseinfoadministrativeareaall.Unwrap() before calling this method if this Baseinfoadministrativeareaall
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Baseinfoadministrativeareaall) Update() *BaseinfoadministrativeareaallUpdateOne {
	return (&BaseinfoadministrativeareaallClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Baseinfoadministrativeareaall entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Baseinfoadministrativeareaall) Unwrap() *Baseinfoadministrativeareaall {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("schema: Baseinfoadministrativeareaall is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Baseinfoadministrativeareaall) String() string {
	var builder strings.Builder
	builder.WriteString("Baseinfoadministrativeareaall(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", pid=")
	builder.WriteString(fmt.Sprintf("%v", b.Pid))
	builder.WriteString(", code=")
	builder.WriteString(b.Code)
	builder.WriteString(", name=")
	builder.WriteString(b.Name)
	builder.WriteString(", province=")
	builder.WriteString(b.Province)
	builder.WriteString(", city=")
	builder.WriteString(b.City)
	builder.WriteString(", area=")
	builder.WriteString(b.Area)
	builder.WriteString(", town=")
	builder.WriteString(b.Town)
	builder.WriteString(", all_name=")
	builder.WriteString(b.AllName)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", b.Type))
	builder.WriteString(", ts_version=")
	builder.WriteString(fmt.Sprintf("%v", b.TsVersion))
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", b.IsDeleted))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", b.Status))
	builder.WriteString(", create_user=")
	builder.WriteString(fmt.Sprintf("%v", b.CreateUser))
	builder.WriteString(", update_user=")
	builder.WriteString(fmt.Sprintf("%v", b.UpdateUser))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", b.CreateTime))
	builder.WriteString(", update_time=")
	builder.WriteString(fmt.Sprintf("%v", b.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// Baseinfoadministrativeareaalls is a parsable slice of Baseinfoadministrativeareaall.
type Baseinfoadministrativeareaalls []*Baseinfoadministrativeareaall

func (b Baseinfoadministrativeareaalls) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
