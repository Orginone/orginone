// Code generated by entc, DO NOT EDIT.

package schema

import (
	"fmt"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketappcomponent"
	"orginone/common/tools/date"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// AsMarketAppComponent is the model entity for the AsMarketAppComponent schema.
type AsMarketAppComponent struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID int64 `json:"id,string"`
	// AppID holds the value of the "app_id" field.
	// app主键
	AppID int64 `json:"appId"`
	// Code holds the value of the "code" field.
	// 编码
	Code string `json:"code"`
	// Name holds the value of the "name" field.
	// 名称
	Name string `json:"name"`
	// URL holds the value of the "url" field.
	// url
	URL string `json:"url"`
	// Type holds the value of the "type" field.
	// 类型
	Type int64 `json:"type"`
	// PreviewPic holds the value of the "preview_pic" field.
	// 预览图片
	PreviewPic string `json:"previewPic"`
	// LayoutType holds the value of the "layout_type" field.
	// 布局类型
	LayoutType string `json:"layoutType"`
	// LayoutConfig holds the value of the "layout_config" field.
	// 布局配置
	LayoutConfig string `json:"layoutConfig"`
	// TenantCode holds the value of the "tenant_code" field.
	// 租户编码
	TenantCode string `json:"tenantCode"`
	// Source holds the value of the "source" field.
	// 平台还剩应用
	Source string `json:"source"`
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
	// The values are being populated by the AsMarketAppComponentQuery when eager-loading is set.
	Edges AsMarketAppComponentEdges `json:"edges"`
}

// AsMarketAppComponentEdges holds the relations/edges for other nodes in the graph.
type AsMarketAppComponentEdges struct {
	// Appx holds the value of the appx edge.
	Appx *AsMarketApp `json:"appx"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AppxOrErr returns the Appx value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AsMarketAppComponentEdges) AppxOrErr() (*AsMarketApp, error) {
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

// scanValues returns the types for scanning values from sql.Rows.
func (*AsMarketAppComponent) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case asmarketappcomponent.FieldID, asmarketappcomponent.FieldAppID, asmarketappcomponent.FieldType, asmarketappcomponent.FieldIsDeleted, asmarketappcomponent.FieldStatus, asmarketappcomponent.FieldCreateUser, asmarketappcomponent.FieldUpdateUser:
			values[i] = new(sql.NullInt64)
		case asmarketappcomponent.FieldCode, asmarketappcomponent.FieldName, asmarketappcomponent.FieldURL, asmarketappcomponent.FieldPreviewPic, asmarketappcomponent.FieldLayoutType, asmarketappcomponent.FieldLayoutConfig, asmarketappcomponent.FieldTenantCode, asmarketappcomponent.FieldSource:
			values[i] = new(sql.NullString)
		case asmarketappcomponent.FieldCreateTime, asmarketappcomponent.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AsMarketAppComponent", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AsMarketAppComponent fields.
func (amac *AsMarketAppComponent) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case asmarketappcomponent.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			amac.ID = int64(value.Int64)
		case asmarketappcomponent.FieldAppID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				amac.AppID = value.Int64
			}
		case asmarketappcomponent.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				amac.Code = value.String
			}
		case asmarketappcomponent.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				amac.Name = value.String
			}
		case asmarketappcomponent.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				amac.URL = value.String
			}
		case asmarketappcomponent.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				amac.Type = value.Int64
			}
		case asmarketappcomponent.FieldPreviewPic:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field preview_pic", values[i])
			} else if value.Valid {
				amac.PreviewPic = value.String
			}
		case asmarketappcomponent.FieldLayoutType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field layout_type", values[i])
			} else if value.Valid {
				amac.LayoutType = value.String
			}
		case asmarketappcomponent.FieldLayoutConfig:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field layout_config", values[i])
			} else if value.Valid {
				amac.LayoutConfig = value.String
			}
		case asmarketappcomponent.FieldTenantCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_code", values[i])
			} else if value.Valid {
				amac.TenantCode = value.String
			}
		case asmarketappcomponent.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				amac.Source = value.String
			}
		case asmarketappcomponent.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				amac.IsDeleted = value.Int64
			}
		case asmarketappcomponent.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				amac.Status = value.Int64
			}
		case asmarketappcomponent.FieldCreateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_user", values[i])
			} else if value.Valid {
				amac.CreateUser = value.Int64
			}
		case asmarketappcomponent.FieldUpdateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_user", values[i])
			} else if value.Valid {
				amac.UpdateUser = value.Int64
			}
		case asmarketappcomponent.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				amac.CreateTime = date.DateTime(value.Time)
			}
		case asmarketappcomponent.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				amac.UpdateTime = date.DateTime(value.Time)
			}
		}
	}
	return nil
}

// QueryAppx queries the "appx" edge of the AsMarketAppComponent entity.
func (amac *AsMarketAppComponent) QueryAppx() *AsMarketAppQuery {
	return (&AsMarketAppComponentClient{config: amac.config}).QueryAppx(amac)
}

// Update returns a builder for updating this AsMarketAppComponent.
// Note that you need to call AsMarketAppComponent.Unwrap() before calling this method if this AsMarketAppComponent
// was returned from a transaction, and the transaction was committed or rolled back.
func (amac *AsMarketAppComponent) Update() *AsMarketAppComponentUpdateOne {
	return (&AsMarketAppComponentClient{config: amac.config}).UpdateOne(amac)
}

// Unwrap unwraps the AsMarketAppComponent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (amac *AsMarketAppComponent) Unwrap() *AsMarketAppComponent {
	tx, ok := amac.config.driver.(*txDriver)
	if !ok {
		panic("schema: AsMarketAppComponent is not a transactional entity")
	}
	amac.config.driver = tx.drv
	return amac
}

// String implements the fmt.Stringer.
func (amac *AsMarketAppComponent) String() string {
	var builder strings.Builder
	builder.WriteString("AsMarketAppComponent(")
	builder.WriteString(fmt.Sprintf("id=%v", amac.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", amac.AppID))
	builder.WriteString(", code=")
	builder.WriteString(amac.Code)
	builder.WriteString(", name=")
	builder.WriteString(amac.Name)
	builder.WriteString(", url=")
	builder.WriteString(amac.URL)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", amac.Type))
	builder.WriteString(", preview_pic=")
	builder.WriteString(amac.PreviewPic)
	builder.WriteString(", layout_type=")
	builder.WriteString(amac.LayoutType)
	builder.WriteString(", layout_config=")
	builder.WriteString(amac.LayoutConfig)
	builder.WriteString(", tenant_code=")
	builder.WriteString(amac.TenantCode)
	builder.WriteString(", source=")
	builder.WriteString(amac.Source)
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", amac.IsDeleted))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", amac.Status))
	builder.WriteString(", create_user=")
	builder.WriteString(fmt.Sprintf("%v", amac.CreateUser))
	builder.WriteString(", update_user=")
	builder.WriteString(fmt.Sprintf("%v", amac.UpdateUser))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", amac.CreateTime))
	builder.WriteString(", update_time=")
	builder.WriteString(fmt.Sprintf("%v", amac.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// AsMarketAppComponents is a parsable slice of AsMarketAppComponent.
type AsMarketAppComponents []*AsMarketAppComponent

func (amac AsMarketAppComponents) config(cfg config) {
	for _i := range amac {
		amac[_i].config = cfg
	}
}