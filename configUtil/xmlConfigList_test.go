package configUtil

import (
	"fmt"
	"testing"

	"github.com/Jordanzuo/goutil/xmlUtil"
)

// bool值读取测试
func TestBoolList(t *testing.T) {
	xmlConfigData, errMsg := getxmlConfigListData()
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()

		return
	}

	booList, errMsg := xmlConfigData.BoolList("html/body/ul/li/a", "id")
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()
		return
	}
	fmt.Println("TestBoolList读取到的值:", booList)
}

// int值读取测试
func TestIntList(t *testing.T) {
	xmlConfigData, errMsg := getxmlConfigListData()
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()

		return
	}

	valList, errMsg := xmlConfigData.IntList("html/body/ul/li/a", "id")
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()
		return
	}
	fmt.Println("TestInt读取到的值:", valList)
}

// int64值读取测试
func TestInt64List(t *testing.T) {
	xmlConfigData, errMsg := getxmlConfigListData()
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()

		return
	}

	valList, errMsg := xmlConfigData.Int64List("html/body/ul/li/a", "id")
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()
		return
	}
	fmt.Println("TestInt64读取到的值:", valList)
}

// Float值读取测试
func TestFloatList(t *testing.T) {
	xmlConfigData, errMsg := getxmlConfigListData()
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()

		return
	}

	valList, errMsg := xmlConfigData.FloatList("html/body/ul/li/a", "id")
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()
		return
	}
	fmt.Println("TestFloat读取到的值:", valList)
}

// 字符串读取测试
func TestStringList(t *testing.T) {
	xmlConfigData, errMsg := getxmlConfigListData()
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()

		return
	}

	valList, errMsg := xmlConfigData.StringList("html/body/ul/li/a", "id")
	if errMsg != nil {
		t.Error(errMsg)
		t.Fail()
		return
	}
	fmt.Println("TestString读取到的值:", valList)
}

func getxmlConfigListData() (xmlConfigData *XmlConfig, errMsg error) {
	content := `
	<html lang="en">
		   <head>
			   <title>Hello</title>
			   <meta name="language" content="en"/>
		   </head>
		   <body IsPost='true'>
				<h1> This is a H1 </h1>
				<ul>
					<li><a id="1" dd='1.1' href="/">Home</a></li>
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
	var root *xmlUtil.Node
	root, errMsg = xmlUtil.LoadFromString(content)
	if errMsg == nil {
		xmlConfigData = NewXmlConfig()
		xmlConfigData.LoadFromXmlNode(root)
	}

	return
}
