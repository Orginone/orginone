package tools

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

var (
	errNotExported = errors.New("field is not exported")
	errNotSettable = errors.New("field is not settable")
	DefaultTagName = "json"
)

type Struct struct {
	raw     interface{}
	value   reflect.Value
	TagName string
	JTag    bool
}

func NewStruct(s interface{}) *Struct {
	return &Struct{
		raw:     s,
		JTag:    true,
		value:   strctVal(s),
		TagName: DefaultTagName,
	}
}

func CallStruct(s interface{}) *Struct {
	return &Struct{
		raw:     s,
		JTag:    false,
		value:   strctVal(s),
		TagName: DefaultTagName,
	}
}

func (s *Struct) Map() map[string]interface{} {
	out := make(map[string]interface{})
	s.FillMap(out)
	return out
}

func (s *Struct) FillMap(out map[string]interface{}) {
	if out == nil {
		return
	}

	fields := s.structFields()

	for _, field := range fields {
		name := field.Name
		if strings.ToLower(name) == "edges" {
			continue
		}
		val := s.value.FieldByName(name)
		isSubStruct := false
		var finalVal interface{}

		tagName, tagOpts := parseTag(field.Tag.Get(s.TagName))
		if tagName != "" {
			name = tagName
		} else {
			name = string(unicode.ToLower(rune(field.Name[0]))) + field.Name[1:]
		}

		if tagOpts.Has("omitempty") {
			zero := reflect.Zero(val.Type()).Interface()
			current := val.Interface()

			if reflect.DeepEqual(current, zero) {
				continue
			}
		}

		if !tagOpts.Has("omitnested") {
			if fromValuer, ok := val.Addr().Interface().(driver.Valuer); ok {
				v, err := fromValuer.Value()
				if err == nil {
					finalVal = v
				}
			} else {
				finalVal = s.nested(val)
			}

			v := reflect.ValueOf(val.Interface())
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}

			switch v.Kind() {
			case reflect.Map, reflect.Struct:
				isSubStruct = true
			}
		} else {
			finalVal = val.Interface()
		}

		if s.JTag && tagOpts.Has("string") {
			switch val.Type().Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				out[name] = fmt.Sprintf("%d", val.Interface())
				break
			}
			continue
		}

		if isSubStruct && (tagOpts.Has("flatten")) {
			for k := range finalVal.(map[string]interface{}) {
				out[k] = finalVal.(map[string]interface{})[k]
			}
		} else {
			out[name] = finalVal
		}
	}
}

func (s *Struct) Values() []interface{} {
	fields := s.structFields()

	var t []interface{}

	for _, field := range fields {
		val := s.value.FieldByName(field.Name)

		_, tagOpts := parseTag(field.Tag.Get(s.TagName))

		if tagOpts.Has("omitempty") {
			zero := reflect.Zero(val.Type()).Interface()
			current := val.Interface()

			if reflect.DeepEqual(current, zero) {
				continue
			}
		}

		if tagOpts.Has("string") {
			s, ok := val.Interface().(fmt.Stringer)
			if ok {
				t = append(t, s.String())
			}
			continue
		}

		if IsStruct(val.Interface()) && !tagOpts.Has("omitnested") {
			t = append(t, Values(val.Interface())...)
		} else {
			t = append(t, val.Interface())
		}
	}

	return t
}

func (s *Struct) Fields() []*Field {
	return getFields(s.value, s.TagName)
}

func (s *Struct) Names() []string {
	fields := getFields(s.value, s.TagName)

	names := make([]string, len(fields))

	for i, field := range fields {
		names[i] = field.Name()
	}

	return names
}

func getFields(v reflect.Value, tagName string) []*Field {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	var fields []*Field

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if tag := field.Tag.Get(tagName); tag == "-" {
			continue
		}

		f := &Field{
			field: field,
			value: v.FieldByName(field.Name),
		}

		fields = append(fields, f)

	}

	return fields
}

func (s *Struct) Field(name string) *Field {
	f, ok := s.FieldOk(name)
	if !ok {
		panic("field not found")
	}

	return f
}

func (s *Struct) FieldOk(name string) (*Field, bool) {
	t := s.value.Type()

	field, ok := t.FieldByName(name)
	if !ok {
		return nil, false
	}

	return &Field{
		field:      field,
		value:      s.value.FieldByName(name),
		defaultTag: s.TagName,
	}, true
}

func (s *Struct) IsZero() bool {
	fields := s.structFields()

	for _, field := range fields {
		val := s.value.FieldByName(field.Name)

		_, tagOpts := parseTag(field.Tag.Get(s.TagName))

		if IsStruct(val.Interface()) && !tagOpts.Has("omitnested") {
			ok := IsZero(val.Interface())
			if !ok {
				return false
			}

			continue
		}
		zero := reflect.Zero(val.Type()).Interface()
		current := val.Interface()

		if !reflect.DeepEqual(current, zero) {
			return false
		}
	}

	return true
}

func (s *Struct) HasZero() bool {
	fields := s.structFields()

	for _, field := range fields {
		val := s.value.FieldByName(field.Name)

		_, tagOpts := parseTag(field.Tag.Get(s.TagName))

		if IsStruct(val.Interface()) && !tagOpts.Has("omitnested") {
			ok := HasZero(val.Interface())
			if ok {
				return true
			}

			continue
		}
		zero := reflect.Zero(val.Type()).Interface()
		current := val.Interface()

		if reflect.DeepEqual(current, zero) {
			return true
		}
	}

	return false
}

func (s *Struct) Name() string {
	return s.value.Type().Name()
}

func (s *Struct) structFields() []reflect.StructField {
	t := s.value.Type()

	var f []reflect.StructField

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue
		}
		if tag := field.Tag.Get(s.TagName); tag == "-" {
			continue
		}

		f = append(f, field)
	}

	return f
}

func strctVal(s interface{}) reflect.Value {
	v := reflect.ValueOf(s)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("not struct")
	}

	return v
}

func Map(s interface{}) map[string]interface{} {
	return NewStruct(s).Map()
}

func FillMap(s interface{}, out map[string]interface{}) {
	NewStruct(s).FillMap(out)
}

func Values(s interface{}) []interface{} {
	return NewStruct(s).Values()
}

func Fields(s interface{}) []*Field {
	return NewStruct(s).Fields()
}

func Names(s interface{}) []string {
	return NewStruct(s).Names()
}

func IsZero(s interface{}) bool {
	return NewStruct(s).IsZero()
}

func HasZero(s interface{}) bool {
	return NewStruct(s).HasZero()
}

func IsStruct(s interface{}) bool {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Invalid {
		return false
	}

	return v.Kind() == reflect.Struct
}

func Name(s interface{}) string {
	return NewStruct(s).Name()
}

func (s *Struct) nested(val reflect.Value) interface{} {
	var finalVal interface{}

	v := reflect.ValueOf(val.Interface())
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		n := NewStruct(val.Interface())
		n.TagName = s.TagName
		m := n.Map()
		if len(m) == 0 {
			finalVal = val.Interface()
		} else {
			finalVal = m
		}
	case reflect.Map:
		mapElem := val.Type()
		switch val.Type().Kind() {
		case reflect.Ptr, reflect.Array, reflect.Map,
			reflect.Slice, reflect.Chan:
			mapElem = val.Type().Elem()
			if mapElem.Kind() == reflect.Ptr {
				mapElem = mapElem.Elem()
			}
		}

		if mapElem.Kind() == reflect.Struct ||
			(mapElem.Kind() == reflect.Slice &&
				mapElem.Elem().Kind() == reflect.Struct) {
			m := make(map[string]interface{}, val.Len())
			for _, k := range val.MapKeys() {
				m[k.String()] = s.nested(val.MapIndex(k))
			}
			finalVal = m
			break
		}

		finalVal = val.Interface()
	case reflect.Slice, reflect.Array:
		if val.Type().Kind() == reflect.Interface {
			finalVal = val.Interface()
			break
		}

		if val.Type().Elem().Kind() != reflect.Struct &&
			!(val.Type().Elem().Kind() == reflect.Ptr &&
				val.Type().Elem().Elem().Kind() == reflect.Struct) {
			finalVal = val.Interface()
			break
		}

		slices := make([]interface{}, val.Len())
		for x := 0; x < val.Len(); x++ {
			slices[x] = s.nested(val.Index(x))
		}
		finalVal = slices
	default:
		finalVal = val.Interface()
	}

	return finalVal
}

type tagOptions []string

func (t tagOptions) Has(opt string) bool {
	for _, tagOpt := range t {
		if tagOpt == opt {
			return true
		}
	}
	return false
}

func parseTag(tag string) (string, tagOptions) {
	res := strings.Split(tag, ",")
	return res[0], res[1:]
}

type Field struct {
	value      reflect.Value
	field      reflect.StructField
	defaultTag string
}

func (f *Field) Tag(key string) string {
	return f.field.Tag.Get(key)
}

func (f *Field) Value() interface{} {
	return f.value.Interface()
}

func (f *Field) IsEmbedded() bool {
	return f.field.Anonymous
}

func (f *Field) IsExported() bool {
	return f.field.PkgPath == ""
}

func (f *Field) IsZero() bool {
	zero := reflect.Zero(f.value.Type()).Interface()
	current := f.Value()

	return reflect.DeepEqual(current, zero)
}

func (f *Field) Name() string {
	return f.field.Name
}

func (f *Field) Kind() reflect.Kind {
	return f.value.Kind()
}

func (f *Field) Set(val interface{}) error {
	if !f.IsExported() {
		return errNotExported
	}

	if !f.value.CanSet() {
		return errNotSettable
	}

	given := reflect.ValueOf(val)

	if f.value.Kind() != given.Kind() {
		return fmt.Errorf("wrong kind. got: %s want: %s", given.Kind(), f.value.Kind())
	}

	f.value.Set(given)
	return nil
}

func (f *Field) Zero() error {
	zero := reflect.Zero(f.value.Type()).Interface()
	return f.Set(zero)
}

func (f *Field) Fields() []*Field {
	return getFields(f.value, f.defaultTag)
}

func (f *Field) Field(name string) *Field {
	field, ok := f.FieldOk(name)
	if !ok {
		panic("field not found")
	}

	return field
}

func (f *Field) FieldOk(name string) (*Field, bool) {
	value := &f.value
	if f.value.Kind() != reflect.Ptr {
		a := f.value.Addr()
		value = &a
	}
	v := strctVal(value.Interface())
	t := v.Type()

	field, ok := t.FieldByName(name)
	if !ok {
		return nil, false
	}

	return &Field{
		field: field,
		value: v.FieldByName(name),
	}, true
}
