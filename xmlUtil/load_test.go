package xmlUtil

import (
	"fmt"
	"strings"
	"testing"
)

// 测试从文件加载
func TestLoadFromFile(t *testing.T) {
	root, errMsg := LoadFromFile("sample.xml")
	if errMsg != nil {
		t.Error("文件加载失败：", errMsg)
		t.Fail()
		return
	}

	node := root.SelectElement("html/head/title")
	if node == nil {
		t.Error("读取节点失败：", "html/head/title")
	}

	fmt.Println("节点值:", strings.TrimSpace(node.InnerText()))
}

// 测试从字符串加载
func TestLoadFromString(t *testing.T) {
	var xml string = `
		<html lang="en">
		   <head>
			   <title>Hello</title>
			   <meta name="language" content="en"/>
		   </head>
		   <body>
				<h1> This is a H1 </h1>
				<ul>
					<li><a id="1" href="/">Home</a></li>
					<li><a id="2" href="/about">about</a></li>
					<li><a id="3" href="/account">login</a></li>
					<li></li>
				</ul>
				<p>
					Hello,This is an example for gxpath.
				</p>
				<footer>footer script</footer>
		   </body>
		</html>
`
	root, errMsg := LoadFromString(xml)
	if errMsg != nil {
		t.Error("文件加载失败：", errMsg)
		t.Fail()
		return
	}

	node := root.SelectElement("html/head/title")
	if node == nil {
		t.Error("读取节点失败：", "html/head/title")
	}

	fmt.Println("节点值:", strings.TrimSpace(node.InnerText()))
}
