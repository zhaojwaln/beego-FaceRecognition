package common

/**
Auth: ghostwwl
Email: ghostwwl@gmail.com
Note: 可能经常用到的小函数了 对应已有的 ghostlib.py
**/

import (
	"encoding/hex"
	"crypto/md5"
	"unicode/utf8"
	"fmt"
	"reflect"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/axgle/mahonia"
)

const (
	GFrmDateTimeMillisecond = "2006-01-02 15:04:05.99999"
	GFrmDateTime = "2006-01-02 15:04:05"
	GFrmTime     = "15:04:05"
	GFrmDay      = "2006-01-02"
)


func Try(do_handdler func(), catch_handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch_handler(err)
		}
	}()
	do_handdler()
}


func Struct2Map(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		result[objT.Field(i).Name] = objV.Field(i).Interface()
	}
	return result
}
func GetMd5(instr string) string {
	sum := md5.Sum([]byte(instr))
	return hex.EncodeToString(sum[:])
}

func Utf8Strlen(instr string) int {
	return utf8.RuneCountInString(instr)
}

func Utf8SubStr(instr string, start_pos, sublen int) string {
	if sublen < 1 {
		return ""
	}

	rune_str := []rune(instr)
	real_len := len(rune_str)
	
	if real_len == 0 {
		return ""
	}

	if start_pos < 0 {
		start_pos = 0
	}
	if start_pos >= real_len {
		return ""
	}
	end_pos := start_pos + sublen
	if end_pos > real_len {
		return string(rune_str[start_pos:])
	} else {
		return string(rune_str[start_pos:end_pos])
	}
}


func CurrentTimeStamp() int64 {
	return time.Now().Unix()
}

func CurrentDateTime() string {
	return time.Now().Format(GFrmDateTime)
}

func CurrentDate() string {
	return time.Now().Format(GFrmDay)
}

func CurrentTime() string {
	return time.Now().Format(GFrmTime)
}

func DtimeToSecond(szday string) int64 {
	loc, err := time.LoadLocation("Local")
	the_time, err := time.ParseInLocation(GFrmDateTime, szday, loc)
	if err == nil {
		unix_time := the_time.Unix()
		return unix_time
	}
	return -1
}

func SecondToDtime(isecond int64) string {
	return time.Unix(isecond, 0).Format(GFrmDateTime)
}

func ConvertStrEncode(inStr, inCharset, outCharset string) string {
	if outCharset == "" {
		outCharset = inCharset
	}

	inCharset = strings.ToLower(inCharset)
	outCharset = strings.ToLower(outCharset)

	if inCharset == outCharset {
		return inStr
	}
	
	if inCharset == "gbk" || inCharset == "gb2312" {
		inCharset = "gb18030"
	}

	// 输入字符串解码为utf-8
	var destr string
	if inCharset != "utf8" && inCharset != "utf-8" {
		destr = mahonia.NewDecoder(inCharset).ConvertString(inStr)
	} else {
		destr = inStr
	}

	if outCharset == "utf8" || outCharset == "utf-8" {
		return destr
	}
	// 转换为 outCharset
	return mahonia.NewEncoder(outCharset).ConvertString(destr)
}

func UrlEncode(instr string) string {
	return url.QueryEscape(instr)
}

func UrlDecode(instr string) (string, error) {
	return url.QueryUnescape(instr)
}

func InitPostData(inmap map[string]interface{}) url.Values {
	result := url.Values{}
	//fmt.Println(inmap)
	for k, v := range inmap {
		result.Add(k, ToString(v))
	}

	return result
}

func ToString(v interface{}) string {
	switch result := v.(type) {
	case string:
		return result
	case []byte:
		return string(result)
	default:
		if v != nil {
			return fmt.Sprintf("%v", result)
		}
	}
	return ""
}

func ToInt(v interface{}) int {
	switch result := v.(type) {
	case int:
		return result
	case int32:
		return int(result)
	case int64:
		return int(result)
	default:
		if d := ToString(v); d != "" {
			value, _ := strconv.Atoi(d)
			return value
		}
	}
	return 0
}

func ToInt64(v interface{}) int64 {
	switch result := v.(type) {
	case int:
		return int64(result)
	case int32:
		return int64(result)
	case int64:
		return result
	default:

		if d := ToString(v); d != "" {
			value, _ := strconv.ParseInt(d, 10, 64)
			return value
		}
	}
	return -1
}

func ToFloat64(v interface{}) float64 {
	switch result := v.(type) {
	case float64:
		return result
	default:
		if d := ToString(v); d != "" {
			value, _ := strconv.ParseFloat(d, 64)
			return value
		}
	}
	return 0
}

func ToBool(v interface{}) bool {
	switch result := v.(type) {
	case bool:
		return result
	default:
		if d := ToString(v); d != "" {
			value, _ := strconv.ParseBool(d)
			return value
		}
	}
	return false
}

func Msg(message string, level int) {
	var msg_head string = ""
	switch level {
	case 1:
		msg_head = fmt.Sprintf("%-7s %s]", "INFO", time.Now())
	case 2:
		msg_head = fmt.Sprintf("%-7s %s]", "WARNING", time.Now())
	case 3:
		msg_head = fmt.Sprintf("%-7s %s]", "ERROR", time.Now())
	case 4:
		msg_head = fmt.Sprintf("%-7s %s]", "DEBUG", time.Now())
	case 5:
		msg_head = fmt.Sprintf("%-7s %s]", "FATAL", time.Now())
	}
	fmt.Printf("%s %s\n", msg_head, message)
}