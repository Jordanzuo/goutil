package configUtil

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Jordanzuo/goutil/typeUtil"
	"github.com/Jordanzuo/goutil/xmlUtil"
)

type XmlConfig struct {
	root *xmlUtil.Node
}

// 从文件加载
// xmlFilePath:xml文件路径
// 返回值:
// error:错误信息
func (this *XmlConfig) LoadFromFile(xmlFilePath string) error {
	if this.root != nil {
		return fmt.Errorf("There has been an xml file loaded.")
	}

	root, err := xmlUtil.LoadFromFile(xmlFilePath)
	if err != nil {
		return err
	}

	this.root = root

	return nil
}

// 从node节点加载（会取其根节点）
// xmlRoot:xml节点
// 返回值:
// error:错误信息
func (this *XmlConfig) LoadFromXmlNode(xmlRoot *xmlUtil.Node) error {
	if this.root != nil {
		return fmt.Errorf("There has been an xml file loaded.")
	}

	if xmlRoot == nil {
		return fmt.Errorf("xmlRoot is nil")
	}

	this.root = xmlRoot

	return nil
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// bool:结果
// error:错误信息
func (this *XmlConfig) Bool(xpath string, attrName string) (bool, error) {
	value, err := this.getVal(xpath, attrName)
	if err != nil {
		return false, err
	}

	return typeUtil.Bool(value)
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
// defaultVal:默认值
//　返回值：
// bool:结果
func (this *XmlConfig) DefaultBool(xpath string, attrName string, defaultVal bool) bool {
	value, err := this.Bool(xpath, attrName)
	if err != nil {
		return defaultVal
	}

	return value
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// int:结果
// error:错误信息
func (this *XmlConfig) Int(xpath string, attrName string) (int, error) {
	value, err := this.getVal(xpath, attrName)
	if err != nil {
		return 0, err
	}

	return typeUtil.Int(value)
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
// defaultVal:默认值
//　返回值：
// int:结果
func (this *XmlConfig) DefaultInt(xpath string, attrName string, defaultVal int) int {
	value, err := this.Int(xpath, attrName)
	if err != nil {
		return defaultVal
	}

	return value
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// int64:结果
// error:错误信息
func (this *XmlConfig) Int64(xpath string, attrName string) (int64, error) {
	value, err := this.getVal(xpath, attrName)
	if err != nil {
		return 0, err
	}

	return typeUtil.Int64(value)
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
// defaultVal:默认值
//　返回值：
// int64:结果
func (this *XmlConfig) DefaultInt64(xpath string, attrName string, defaultVal int64) int64 {
	value, err := this.Int64(xpath, attrName)
	if err != nil {
		return defaultVal
	}

	return value

}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
//　返回值：
// float64:结果
// error:错误信息
func (this *XmlConfig) Float(xpath string, attrName string) (float64, error) {
	value, err := this.getVal(xpath, attrName)
	if err != nil {
		return 0, err
	}

	return typeUtil.Float64(value)
}

// 获取指定xpath路径下的值
// xpath:xpath路径
// attrName:属性名，如果为空，则返回节点的内部文本
// defaultVal:默认值
//　返回值：
// float64:结果
func (this *XmlConfig) DefaultFloat(xpath string, attrName string, defaultVal float64) float64 {
	value, err := this.Float(xpath, attrName)
	if err != nil {
		return defaultVal
	}

	return value
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
// defaultVal:默认值
//　返回值：
// string:结果
func (this *XmlConfig) DefaultString(xpath string, attrName string, defaultVal string) string {
	value, err := this.String(xpath, attrName)
	if err != nil {
		return defaultVal
	}

	return value
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

// 反序列化指定的整个节点
// xpath:xml的path
// data:反序列化得到的数据
// 返回值:
// error:错误信息
func (this *XmlConfig) Unmarshal(xpath string, data interface{}) error {
	nodeItem := this.Node(xpath)
	if nodeItem == nil {
		return fmt.Errorf("节点不存在,XPATH:%s", xpath)
	}

	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	dataType := value.Type()

	// 依次设置字段值
	var err error
	fieldCount := value.NumField()
	for i := 0; i < fieldCount; i++ {
		fieldItem := value.Field(i)
		fieldName := dataType.Field(i).Name

		// 读取数据
		var valueString string
		tmpXpath := fmt.Sprintf("%s/%s", xpath, fieldName)
		if valueString, err = this.getVal(tmpXpath, ""); err != nil {
			valueString, err = this.getVal(xpath, fieldName)
			if err != nil {
				// 压根儿无此字段的配置数据，则略过
				continue
			}
		}

		// 字符串转换成目标值
		fieldValue, err := typeUtil.Convert(valueString, fieldItem.Kind())
		if err != nil {
			return fmt.Errorf("读取字段失败, DataType:%s FieldName:%s Value:%v 错误信息:%v ", dataType.Name(), fieldName, valueString, err)
		}

		// 设置到字段上面
		valType := reflect.ValueOf(fieldValue)
		if valType.Type() == fieldItem.Type() {
			fieldItem.Set(valType)
		} else {
			fieldItem.Set(valType.Convert(fieldItem.Type()))
		}
	}

	return nil
}

// 获取指定路径的之
// xpath:xpath路径
// attrName:要获取的属性值，如果为空，则返回内部文本
func (this *XmlConfig) getVal(xpath string, attrName string) (val string, err error) {
	targetRoot := this.root.SelectElement(xpath)
	if targetRoot == nil {
		err = fmt.Errorf("no find target node:%v", xpath)
		return
	}

	if attrName == "" {
		val = strings.TrimSpace(targetRoot.InnerText())
		return
	}

	exist := false
	val, exist = targetRoot.SelectAttr(attrName)
	if exist == false {
		err = fmt.Errorf("no find target attr, node:%v attr:%v", xpath, attrName)
		return
	}

	return
}

// 创建新的xml配置对象
func NewXmlConfig() *XmlConfig {
	return &XmlConfig{}
}
