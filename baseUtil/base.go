/*
一个处理不同进制的工具包；用于将十进制和其它进制进行互相转换
*/
package baseUtil

import (
	"fmt"
	"math"
)

// 进制对象定义
type Base struct {
	elementList []string
	base        uint64
}

// 将10进制的uint64类型数据转换为字符串形式
// source:10进制的uint64类型数据
// 返回值:
// 对应进制的字符串形式
func (this *Base) Transform(source uint64) (result string) {
	quotient, remainder := uint64(0), source

	for {
		quotient, remainder = remainder/this.base, remainder%this.base
		result = this.elementList[remainder] + result
		if quotient == 0 {
			break
		}
		remainder = quotient
	}

	return
}

// 将字符串解析为10进制的uint64类型
// source:对应进制的字符串形式
// 返回值:10进制的uint64类型数据
func (this *Base) Parse(source string) (result uint64) {
	if source == "" {
		return
	}

	sourceList := make([]string, 0, len(source))
	for _, v := range source {
		sourceList = append(sourceList, string(v))
	}

	for idx, exp := len(sourceList)-1, 0; idx >= 0; idx, exp = idx-1, exp+1 {
		sourceItem := sourceList[idx]

		// Find the source item in the elementList
		for i, v := range this.elementList {
			if sourceItem == v {
				result += uint64(float64(i) * math.Pow(float64(this.base), float64(exp)))
			}
		}
	}

	return
}

// 以指定的任意非重复的数组，来指定基于的进制数
func New(elements string) (baseObj *Base, err error) {
	if len(elements) == 0 {
		err = fmt.Errorf("输入的字符数串为空")
		return
	}

	elementList := make([]string, 0, len(elements))
	elementMap := make(map[rune]struct{}, len(elements))
	for _, v := range elements {
		if _, exist := elementMap[v]; exist {
			err = fmt.Errorf("输入的字符串中含有重复的字符:%s", string(v))
			return
		} else {
			elementMap[v] = struct{}{}
			elementList = append(elementList, string(v))
		}
	}

	baseObj = &Base{
		elementList: elementList,
		base:        uint64(len(elementList)),
	}

	return
}

// 包含01
func NewBase2() (baseObj *Base, err error) {
	return New("01")
}

// 包含0-7
func NewBase8() (baseObj *Base, err error) {
	return New("01234567")
}

// 包含0-9,a-x
func NewBase16() (baseObj *Base, err error) {
	return New("0123456789abcdef")
}

// 包含a-z
func NewBase26() (baseObj *Base, err error) {
	return New("abcdefghijklmnopqrstuvwxyz")
}

// 包含0-9,a-z
func NewBase36() (baseObj *Base, err error) {
	return New("0123456789abcdefghijklmnopqrstuvwxyz")
}

// 包含0-9,a-z,A-Z
func NewBase62() (baseObj *Base, err error) {
	return New("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
