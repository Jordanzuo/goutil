package configUtil

import (
	"fmt"

	"github.com/Jordanzuo/goutil/typeUtil"
	"github.com/Jordanzuo/goutil/xmlUtil"
)

type XmlConfig struct {
	root *xmlUtil.Node
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// bool:结果
// error:错误信息
func (this *XmlConfig) Bool(xpath string, attrName string) (bool, error) {
	val, errMsg := this.getVal(xpath, attrName)
	if errMsg != nil {
		return false, errMsg
	}

	return typeUtil.Bool(val)
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
// defaultval:默认值
//　返回值：
// bool:结果
func (this *XmlConfig) DefaultBool(xpath string, attrName string, defaultval bool) bool {
	v, err := this.Bool(xpath, attrName)
	if err != nil {
		return defaultval
	}
	return v
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// int:结果
// error:错误信息
func (this *XmlConfig) Int(xpath string, attrName string) (int, error) {
	val, errMsg := this.getVal(xpath, attrName)
	if errMsg != nil {
		return 0, errMsg
	}

	return typeUtil.Int(val)
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
// defaultval:默认值
//　返回值：
// int:结果
func (this *XmlConfig) DefaultInt(xpath string, attrName string, defaultval int) int {
	v, err := this.Int(xpath, attrName)
	if err != nil {
		return defaultval
	}

	return v
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// int64:结果
// error:错误信息
func (this *XmlConfig) Int64(xpath string, attrName string) (int64, error) {
	val, errMsg := this.getVal(xpath, attrName)
	if errMsg != nil {
		return 0, errMsg
	}

	return typeUtil.Int64(val)
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
// defaultval:默认值
//　返回值：
// int64:结果
func (this *XmlConfig) DefaultInt64(xpath string, attrName string, defaultval int64) int64 {
	v, err := this.Int64(xpath, attrName)
	if err != nil {
		return defaultval
	}
	return v

}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// float64:结果
// error:错误信息
func (this *XmlConfig) Float(xpath string, attrName string) (float64, error) {
	val, errMsg := this.getVal(xpath, attrName)
	if errMsg != nil {
		return 0, errMsg
	}

	return typeUtil.Float64(val)
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
// defaultval:默认值
//　返回值：
// float64:结果
func (this *XmlConfig) DefaultFloat(xpath string, attrName string, defaultval float64) float64 {
	v, err := this.Float(xpath, attrName)
	if err != nil {
		return defaultval
	}
	return v
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// string:结果
// error:错误信息
func (this *XmlConfig) String(xpath string, attrName string) (string, error) {
	return this.getVal(xpath, attrName)
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
// defaultval:默认值
//　返回值：
// string:结果
func (this *XmlConfig) DefaultString(xpath string, attrName string, defaultval string) string {
	v, _ := this.String(xpath, attrName)
	if v == "" {
		return defaultval
	}
	return v
}

// 获取指定位置的节点
// xpath:xpath路径
// 返回值:
// []*xmlUtil.Node：结果
func (this *XmlConfig) Nodes(xpath string) []*xmlUtil.Node {
	return this.root.SelectElements(xpath)
}

// 获取指定位置的节点
// xpath:xpath路径
// 返回值:
// *xmlUtil.Node：结果
func (this *XmlConfig) Node(xpath string) *xmlUtil.Node {
	return this.root.SelectElement(xpath)
}

// 获取指定路径的之
// xpath:xpath路径
// attrName:要获取的属性值，如果为空，则返回内部文本
func (this *XmlConfig) getVal(xpath string, attrName string) (string, error) {
	targetRoot := this.root.SelectElement(xpath)
	if targetRoot == nil {
		return "", fmt.Errorf("no find target node:%v", xpath)
	}

	val := ""
	if attrName == "" {
		val = targetRoot.InnerText()
	} else {
		val = targetRoot.SelectAttr(attrName)
	}

	return val, nil
}

// 创建新的xml配置对象
func NewXmlConfig(_root *xmlUtil.Node) *XmlConfig {
	return &XmlConfig{
		root: _root,
	}
}
