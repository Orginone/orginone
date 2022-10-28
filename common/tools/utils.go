package tools

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"orginone/common/tools/aes"
	"orginone/common/tools/linq"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"entgo.io/ent"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/oauth2.v3/utils/uuid"
)

const (
	MyCost    int    = 10
	numUpChar string = "0123456789ABCDEFGHJKLMNPQRTUWXY"
	chars     string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

//获取元素所在数组下标
func ArrIndexOf(array interface{}, val interface{}) int {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					return i
				}
			}
		}
	}
	return -1
}

//判断数组是否包含指定元素
func ArrIsExist(array interface{}, val interface{}) bool {
	return ArrIndexOf(array, val) != -1
}

//解析前端Aes密码
func AesPwdDecrypt(password string, key string) (string, error) {
	if len(password) < 7 {
		return password, errors.New("password length must be greater than 6")
	}
	if password[0] == '$' {
		return password[1:], nil
	}
	pwdBase64, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return password, err
	}
	des, err := aes.AesECBDecrypt(pwdBase64, []byte(key), aes.PKCS7_PADDING)
	if err != nil {
		return password, err
	}
	return string(des), nil
}

//密码强度检验,不符合要求为true
func CheckPwd(password string) bool {
	if len(password) < 6 {
		return true
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	if b, err := regexp.MatchString(num, password); !b || err != nil {
		return true
	}
	if b, err := regexp.MatchString(a_z, password); !b || err != nil {
		return true
	}
	if b, err := regexp.MatchString(A_Z, password); !b || err != nil {
		return true
	}
	return false
}

//密码加密
func BcryptEncryptPwd(passwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), MyCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}

//校验密码
func ValidateBcryptPwd(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		return false
	}
	return true
}

//生产TenantCode
func GenTenantCode() string {
	timeStr := time.Now().Format("060102150405")
	randStr := GetRandString(6)
	return fmt.Sprintf("ZH%s%s", timeStr, randStr)
}

//获取指定长度的随机字符
func GetRandString(length int32, sources ...string) string {
	sourceStr := chars
	if len(sources) > 0 {
		sourceStr = sources[0]
	}
	b := make([]rune, length)
	for i := range b {
		b[i] = rune(sourceStr[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(sourceStr))])
	}
	return string(b)
}

//社会统一代码验证器
func SocialCreditCodeValidator(code string) bool {
	if pass, _ := regexp.Match("^[A-Z0-9]+$", []byte(code)); IsNilOrEmpty(code) || len(code) != 18 || !pass {
		return false
	}
	wis := []int{1, 3, 9, 27, 19, 26, 16, 17, 20, 29, 25, 13, 8, 24, 10, 30, 28}
	sum := 0
	for i := 0; i < 17; i++ {

		sum += strings.Index(numUpChar, string(code[i])) * wis[i]
	}
	ret := 0
	if sum%31 != 0 {
		ret = 31 - sum%31
	}
	return strings.Index(numUpChar, string(code[17])) == ret
}

//创建社会统一代码
func CreateSocialCreditCode() string {
	timeStr := time.Now().Format("SC060102150405")
	randStr := GetRandString(4, numUpChar)
	return fmt.Sprintf("%s%s", timeStr, randStr)
}

//判断字符串是否为空字符串
func IsNilOrEmpty(str string) bool {
	return len(strings.TrimSpace(str)) <= 0
}
func IsIDCard(str string) bool {
	ok, err := regexp.MatchString("^\\d{15}$)|(^\\d{17}([0-9]|X)$", str)
	return err == nil && ok
}
func IsPhone(str string) bool {
	ok, err := regexp.MatchString("^[1](([3][0-9])|([4][5-9])|([5][0-3,5-9])|([6][5,6])|([7][0-8])|([8][0-9])|([9][1,8,9]))[0-9]{8}$", str)
	return err == nil && ok
}

func IsSocialCreditCode(str string) bool {
	return len(str) >= 17
}

// snake converts the given struct or field name into a snake_case.
//
//	Username => username
//	FullName => full_name
//	HTTPCode => http_code
//
func Snake(s string) string {
	var (
		j int
		b strings.Builder
	)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		// Put '_' if it is not a start or end of a word, current letter is uppercase,
		// and previous is lowercase (cases like: "UserInfo"), or next letter is also
		// a lowercase and previous letter is not "_".
		if i > 0 && i < len(s)-1 && unicode.IsUpper(r) {
			if unicode.IsLower(rune(s[i-1])) ||
				j != i-1 && unicode.IsLower(rune(s[i+1])) && unicode.IsLetter(rune(s[i-1])) {
				j = i
				b.WriteString("_")
			}
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

//call schema setter of any entity
//do not use of other model
func SchemaUpdateAny(setter ent.Mutation, entity any, ignores ...string) {
	entityMap := CallStruct(entity).Map()
	for k, v := range entityMap {
		filedName := Snake(k)
		if len(ignores) > 0 && ArrIndexOf(ignores, filedName) > -1 {
			continue
		}
		if (filedName == "parent_id" || filedName == "pid") && v.(int64) == 0 {
			continue
		}
		switch v.(type) {
		case time.Time:
			if !v.(time.Time).IsZero() {
				setter.SetField(filedName, v)
			}
			break
		default:
			if v != nil {
				setter.SetField(filedName, v)
			}
			break
		}
	}
}

func ArrayIntToInt64(source []int) []int64 {
	array := make([]int64, 0)
	for _, v := range source {
		array = append(array, int64(v))
	}
	return array
}

func ArrayStrToInt64(source []string) []int64 {
	array := make([]int64, 0)
	for _, v := range source {
		if len(v) > 0 {
			i, err := strconv.ParseInt(v, 10, 64)
			if err == nil {
				array = append(array, i)
			}
		}
	}
	return array
}

// 字符串数组追加或移除
//opt 1:add,2:delete
func ArrCodesAddOrDelete(srcArr []string, optArr []string, opt int64) []string {
	linq.From(srcArr).Distinct().ToSlice(&srcArr)
	linq.From(optArr).Distinct().ToSlice(&optArr)
	switch opt {
	case 1:
		srcArr = append(srcArr, optArr...)
		break
	case 2:
		linq.From(srcArr).WhereT(func(s string) bool {
			for _, v := range optArr {
				if v == s {
					return false
				}
			}
			return true
		}).ToSlice(&srcArr)
		break
	}
	linq.From(srcArr).Distinct().ToSlice(&srcArr)
	return srcArr
}

//JsonUnmarshalFromIOReader
func UnmarshaJsonIOData(r io.Reader, v any) error {
	buffer, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buffer, v)
	return err
}

//GenAppKeySecret
func GenAppKeyOrSecret(length int) string {
	random, _ := uuid.NewRandom()
	uuidStr := strings.ReplaceAll(random.String(), "-", "")
	return uuidStr + GetRandString(int32(length-len(uuidStr)))
}
