package typeUtil

// KeyValue数据集合
type MapData map[string]interface{}

// 创建新的MapData
// 返回新的Map对象
func NewMapData() MapData {
	return MapData(make(map[string]interface{}))
}

// 类型转换为int
// 返回值:
// int:结果
// bool:转换是否成功
func (this MapData) Int(key string) (int, bool) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return 0, false
	}

	return Int(val)
}

// 类型转换为int
// 返回值:
// float64:结果
// bool:转换是否成功
func (this MapData) Float64(key string) (float64, bool) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return 0, false
	}

	return Float64(val)
}

// 类型转换为bool
// 返回值:
// bool:结果
// bool:转换是否成功
func (this MapData) Bool(key string) (bool, bool) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return false, false
	}

	return Bool(val)
}

// 类型转换为字符串
// 返回值:
// string:结果
// bool:转换是否成功
func (this MapData) String(key string) (string, bool) {
	val, isExist := this[key]
	if isExist == false || val == nil {
		return "", false
	}

	return String(val)
}
