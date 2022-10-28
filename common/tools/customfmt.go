package tools

import (
	"fmt"
	"orginone/common/tools/linq"
	"reflect"
	"strconv"
	"strings"
)

func NumArrayToStrArray(args interface{}) []string {
	var temps []string
	linq.From(args).SelectT(func(a interface{}) string {
		switch t := reflect.TypeOf(a); t.Kind() {
		case reflect.Int64:
			return strconv.FormatInt(a.(int64), 10)
		case reflect.Int32, reflect.Int:
			return strconv.Itoa(a.(int))
		default:
			return fmt.Sprintf("%d", a)
		}
	}).ToSlice(&temps)
	return temps
}

func StrArrayToNumArray(args []string) []int64 {
	var temps []int64
	linq.From(args).SelectT(func(a string) int64 {
		num, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return -1
		}
		return num
	}).ToSlice(&temps)
	return temps
}

func SqlInArgsMap(args interface{}) string {
	var temps []string
	linq.From(args).SelectT(func(a interface{}) string {
		switch t := reflect.TypeOf(a); t.Kind() {
		case reflect.Int64:
			return strconv.FormatInt(a.(int64), 10)
		case reflect.Int32:
			return strconv.Itoa(a.(int))
		case reflect.String:
			return fmt.Sprintf("\"%s\"", a.(string))
		default:
			return ""
		}
	}).ToSlice(&temps)
	return fmt.Sprintf("(%s)", strings.Join(temps, ","))
}
