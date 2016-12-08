package jsonUtil

import (
	"encoding/json"
	"strings"
)

// 使用Number类型来反序列化字符串
// 当被序列化为interface{}类型时，如果int型的长度大于7，则会被使用科学计数法进行表示
// 当反序列化时，会无法转换为int类型，会导致错误
// 所以需要使用Number类型
// s:输入字符串
// 返回值:
// 反序列化后的数据
// 错误对象
func UnMarshalWithNumberType(s string) (interface{}, error) {
	// 构造decode对象
	var decode = json.NewDecoder(strings.NewReader(s))
	decode.UseNumber()

	// decode
	var result interface{}
	if err := decode.Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// 深拷贝对象
// src:源数据
// 返回值:
// 新对象
// 错误对象
func DeepClone(src interface{}) (interface{}, error) {
	var byteSlice []byte
	var err error

	// 先序列化为[]byte
	if byteSlice, err = json.Marshal(src); err != nil {
		return nil, err
	}

	// 再反序列化成对象
	var result interface{}
	if err := json.Unmarshal(byteSlice, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// 使用Number类型来深拷贝对象
// src:源数据
// 返回值:
// 新对象
// 错误对象
func DeepCloneWithNumberType(src interface{}) (interface{}, error) {
	var byteSlice []byte
	var err error

	// 先序列化为[]byte
	if byteSlice, err = json.Marshal(src); err != nil {
		return nil, err
	}

	// 构造decode对象
	var decode = json.NewDecoder(strings.NewReader(string(byteSlice)))
	decode.UseNumber()

	// decode
	var result interface{}
	if err := decode.Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
