package toolElse

import (
	"log"
	"os"
	"reflect"

	"github.com/axgle/mahonia"
)

//转码
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

//把GBK转成UTF8
func toChinese(in string) string {
	enc := mahonia.NewDecoder("gbk")
	targetStr := enc.ConvertString(in)
	return targetStr
}

/**
判断是否有错误
如果有错误则报错且完全退出
*/
func CheckError(err error, tips ...string) {
	if err != nil {
		for _, tip := range tips {
			log.Print(tip)
		}
		log.Println(err.Error())
		os.Exit(1)
	}
}

func CheckErrorResult(err error, tips ...string) bool {
	if err != nil {
		for _, tip := range tips {
			log.Print(tip)
		}
		log.Println(err.Error())
		return true
	}
	return false
}

//是否是同一个类型 作用和java instanceOf相同
func InstanceOf(v interface{}, t interface{}) bool {
	if v == nil || t == nil {
		return false
	}
	tt := reflect.TypeOf(t)
	tv := reflect.TypeOf(v)

	if tt.Kind() == reflect.Ptr {
		tt = tt.Elem()
	}
	if tv.Kind() == reflect.Ptr {
		tv = tv.Elem()
	}

	if tt.Kind() == reflect.Interface {
		if tv.Kind() != reflect.Interface {
			return reflect.PtrTo(tv).Implements(tt)
		} else {
			return tv.Implements(tt)
		}
	} else {
		return tt == tv || hasEmbedded(tv, tt)
	}
}

func hasEmbedded(tv reflect.Type, tt reflect.Type) bool {
	if tt.Kind() == reflect.Ptr {
		tt = tt.Elem()
	}
	if tv.Kind() == reflect.Ptr {
		tv = tv.Elem()
	}

	for i := 0; i < tv.NumField(); i++ {
		tsf := tv.Field(i)
		if tsf.Anonymous && (tt == tsf.Type || hasEmbedded(tsf.Type, tt)) {
			return true
		}
	}
	return false
}
