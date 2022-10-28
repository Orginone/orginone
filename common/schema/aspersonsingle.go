// Code generated by entc, DO NOT EDIT.

package schema

import (
	"fmt"
	"orginone/common/schema/aspersonsingle"
	"orginone/common/tools/date"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// AsPersonSingle is the model entity for the AsPersonSingle schema.
type AsPersonSingle struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID int64 `json:"id,string"`
	// RealName holds the value of the "real_name" field.
	// 姓名
	RealName string `json:"realName"`
	// IDCard holds the value of the "id_card" field.
	// 身份证号
	IDCard string `json:"idCard"`
	// Gender holds the value of the "gender" field.
	// 性别;1、男2、女
	Gender int64 `json:"gender"`
	// UserBirthday holds the value of the "user_birthday" field.
	// 出生日期
	UserBirthday date.DateTime `json:"userBirthday"`
	// UserEmail holds the value of the "user_email" field.
	// 邮箱号码
	UserEmail string `json:"userEmail"`
	// UserPhoto holds the value of the "user_photo" field.
	// 照片
	UserPhoto string `json:"userPhoto"`
	// PhoneNumber holds the value of the "phone_number" field.
	// 手机号码
	PhoneNumber string `json:"phoneNumber"`
	// Province holds the value of the "province" field.
	// 省
	Province string `json:"province"`
	// City holds the value of the "city" field.
	// 市
	City string `json:"city"`
	// StreetAddress holds the value of the "street_address" field.
	// 地区
	StreetAddress string `json:"streetAddress"`
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
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AsPersonSingle) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case aspersonsingle.FieldID, aspersonsingle.FieldGender, aspersonsingle.FieldIsDeleted, aspersonsingle.FieldStatus, aspersonsingle.FieldCreateUser, aspersonsingle.FieldUpdateUser:
			values[i] = new(sql.NullInt64)
		case aspersonsingle.FieldRealName, aspersonsingle.FieldIDCard, aspersonsingle.FieldUserEmail, aspersonsingle.FieldUserPhoto, aspersonsingle.FieldPhoneNumber, aspersonsingle.FieldProvince, aspersonsingle.FieldCity, aspersonsingle.FieldStreetAddress:
			values[i] = new(sql.NullString)
		case aspersonsingle.FieldUserBirthday, aspersonsingle.FieldCreateTime, aspersonsingle.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AsPersonSingle", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AsPersonSingle fields.
func (aps *AsPersonSingle) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case aspersonsingle.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			aps.ID = int64(value.Int64)
		case aspersonsingle.FieldRealName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field real_name", values[i])
			} else if value.Valid {
				aps.RealName = value.String
			}
		case aspersonsingle.FieldIDCard:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id_card", values[i])
			} else if value.Valid {
				aps.IDCard = value.String
			}
		case aspersonsingle.FieldGender:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				aps.Gender = value.Int64
			}
		case aspersonsingle.FieldUserBirthday:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field user_birthday", values[i])
			} else if value.Valid {
				aps.UserBirthday = date.DateTime(value.Time)
			}
		case aspersonsingle.FieldUserEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_email", values[i])
			} else if value.Valid {
				aps.UserEmail = value.String
			}
		case aspersonsingle.FieldUserPhoto:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_photo", values[i])
			} else if value.Valid {
				aps.UserPhoto = value.String
			}
		case aspersonsingle.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				aps.PhoneNumber = value.String
			}
		case aspersonsingle.FieldProvince:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field province", values[i])
			} else if value.Valid {
				aps.Province = value.String
			}
		case aspersonsingle.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				aps.City = value.String
			}
		case aspersonsingle.FieldStreetAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field street_address", values[i])
			} else if value.Valid {
				aps.StreetAddress = value.String
			}
		case aspersonsingle.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				aps.IsDeleted = value.Int64
			}
		case aspersonsingle.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				aps.Status = value.Int64
			}
		case aspersonsingle.FieldCreateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_user", values[i])
			} else if value.Valid {
				aps.CreateUser = value.Int64
			}
		case aspersonsingle.FieldUpdateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_user", values[i])
			} else if value.Valid {
				aps.UpdateUser = value.Int64
			}
		case aspersonsingle.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				aps.CreateTime = date.DateTime(value.Time)
			}
		case aspersonsingle.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				aps.UpdateTime = date.DateTime(value.Time)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AsPersonSingle.
// Note that you need to call AsPersonSingle.Unwrap() before calling this method if this AsPersonSingle
// was returned from a transaction, and the transaction was committed or rolled back.
func (aps *AsPersonSingle) Update() *AsPersonSingleUpdateOne {
	return (&AsPersonSingleClient{config: aps.config}).UpdateOne(aps)
}

// Unwrap unwraps the AsPersonSingle entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aps *AsPersonSingle) Unwrap() *AsPersonSingle {
	tx, ok := aps.config.driver.(*txDriver)
	if !ok {
		panic("schema: AsPersonSingle is not a transactional entity")
	}
	aps.config.driver = tx.drv
	return aps
}

// String implements the fmt.Stringer.
func (aps *AsPersonSingle) String() string {
	var builder strings.Builder
	builder.WriteString("AsPersonSingle(")
	builder.WriteString(fmt.Sprintf("id=%v", aps.ID))
	builder.WriteString(", real_name=")
	builder.WriteString(aps.RealName)
	builder.WriteString(", id_card=")
	builder.WriteString(aps.IDCard)
	builder.WriteString(", gender=")
	builder.WriteString(fmt.Sprintf("%v", aps.Gender))
	builder.WriteString(", user_birthday=")
	builder.WriteString(fmt.Sprintf("%v", aps.UserBirthday))
	builder.WriteString(", user_email=")
	builder.WriteString(aps.UserEmail)
	builder.WriteString(", user_photo=")
	builder.WriteString(aps.UserPhoto)
	builder.WriteString(", phone_number=")
	builder.WriteString(aps.PhoneNumber)
	builder.WriteString(", province=")
	builder.WriteString(aps.Province)
	builder.WriteString(", city=")
	builder.WriteString(aps.City)
	builder.WriteString(", street_address=")
	builder.WriteString(aps.StreetAddress)
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", aps.IsDeleted))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", aps.Status))
	builder.WriteString(", create_user=")
	builder.WriteString(fmt.Sprintf("%v", aps.CreateUser))
	builder.WriteString(", update_user=")
	builder.WriteString(fmt.Sprintf("%v", aps.UpdateUser))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", aps.CreateTime))
	builder.WriteString(", update_time=")
	builder.WriteString(fmt.Sprintf("%v", aps.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// AsPersonSingles is a parsable slice of AsPersonSingle.
type AsPersonSingles []*AsPersonSingle

func (aps AsPersonSingles) config(cfg config) {
	for _i := range aps {
		aps[_i].config = cfg
	}
}
