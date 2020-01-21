package typeUtil

import (
	"fmt"
	"reflect"
)

// map转换为字符串(如果类型不匹配)
// data:map数据
// separator1:间隔符1
// separator1:间隔符2
// 返回值:
// result:转换后的字符串
// err:错误信息
func MapToString(data interface{}, separator1, separator2 string) (result string, err error) {
	if data == nil {
		return
	}

	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Map {
		err = fmt.Errorf("只能转换Map类型的，当前类型是:%v", val.Kind().String())
		return
	}

	if val.Len() <= 0 {
		return
	}

	for _, keyItem := range val.MapKeys() {
		valItem := val.MapIndex(keyItem)
		result = result + fmt.Sprintf("%v%s%v%s", keyItem.Interface(), separator1, valItem.Interface(), separator2)
	}

	result = result[:len(result)-1]

	return
}

// map转换为字符串(如果类型不匹配)
// data:map数据
// separator1:间隔符1
// separator1:间隔符2
// valGetFunc:结果值获取函数
// 返回值:
// result:转换后的字符串
// err:错误信息
func MapToString2(data interface{}, separator1, separator2 string, valGetFunc func(val interface{}) interface{}) (result string, err error) {
	if data == nil {
		return
	}

	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Map {
		err = fmt.Errorf("只能转换Map类型的，当前类型是:%v", val.Kind().String())
		return
	}

	if val.Len() <= 0 {
		return
	}

	for _, keyItem := range val.MapKeys() {
		valItem := val.MapIndex(keyItem)
		result = result + fmt.Sprintf("%v%s%v%s", keyItem.Interface(), separator1, valGetFunc(valItem.Interface()), separator2)
	}

	result = result[:len(result)-1]

	return
}
