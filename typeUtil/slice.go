package typeUtil

import (
	"fmt"
	"reflect"
)

// 把一个集合转换成字符串
// data:slice类型的集合
// separator:分隔符
// 返回值:
// result:转换后的字符串
// err:错误信息对象
func SliceToString(data interface{}, separator string) (result string, err error) {
	if data == nil {
		return
	}

	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		err = fmt.Errorf("目标类型不正确，只能是slice或array 当前类型是:%v", value.Kind().String())
		return
	}

	if value.Len() <= 0 {
		return
	}

	for i := 0; i < value.Len(); i++ {
		valItem := value.Index(i)
		result = result + fmt.Sprintf("%s%v", separator, valItem.Interface())
	}
	result = result[1:]

	return
}
