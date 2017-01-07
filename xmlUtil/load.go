package xmlUtil

import (
	"bytes"
	"io/ioutil"
	"regexp"
)

// 从文件加载
// filePath:文件路径
// 返回值:
// *Node:根节点对象
// error:错误信息
func LoadFromFile(filePath string) (*Node, error) {
	data, errMsg := ioutil.ReadFile(filePath)
	if errMsg != nil {
		return nil, errMsg
	}

	return LoadFromByte(data)
}

// 从字节数组加载
// data:文档数据
// 返回值:
// *Node:根节点对象
// error:错误信息
func LoadFromByte(data []byte) (*Node, error) {
	return LoadFromString(string(data))
}

// 文档字符串
// doc:文档字符串
// 返回值:
// *Node:根节点对象
// error:错误信息
func LoadFromString(doc string) (*Node, error) {
	// xml.Decoder doesn't properly handle whitespace in some doc
	// see songTextString.xml test case ...
	reg, _ := regexp.Compile("[ \t\n\r]*<")
	doc = reg.ReplaceAllString(doc, "<")

	b := bytes.NewBufferString(doc)

	return LoadFromReader(b)
}
