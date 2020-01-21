package configUtil

import (
	"testing"
)

var (
	config_Array []map[string]interface{}
	err_Array    error
)

func TestReadJsonConfig_Array(t *testing.T) {
	config_Array, err_Array = ReadJsonConfig_Array("testdata/jsonConfigArray.ini")
	if err_Array != nil {
		t.Errorf("读取JSON配置失败，错误信息为：%s", err_Array)
	}
}

func TestReadIntJsonValue_Array(t *testing.T) {
	actualValue, err_Array := ReadIntJsonValue_Array(config_Array, "ServerGroupId")
	if err_Array != nil {
		t.Errorf("读取JSON配置失败，错误信息为：%s", err_Array)
	}

	expectedValue := 1
	if actualValue != expectedValue {
		t.Errorf("期望的值为%d，实际的值为%d", expectedValue, actualValue)
	}
}

func TestReadStringJsonValue_Array(t *testing.T) {
	actualValue, err_Array := ReadStringJsonValue_Array(config_Array, "ChatDBConnection")
	if err_Array != nil {
		t.Errorf("读取JSON配置失败，错误信息为：%s", err_Array)
	}

	expectedValue := "root:moqikaka@tcp(192.168.1.226:3306)/chatserver?charset=utf8&parseTime=true&loc=Local&timeout=30s"
	if actualValue != expectedValue {
		t.Errorf("期望的值为%s，实际的值为%s", expectedValue, actualValue)
	}
}

func TestReadBoolJsonValue_Array(t *testing.T) {
	actualValue, err_Array := ReadBoolJsonValue_Array(config_Array, "IfRecordMessage")
	if err_Array != nil {
		t.Errorf("读取JSON配置失败，错误信息为：%s", err_Array)
	}

	expectedValue := true
	if actualValue != expectedValue {
		t.Errorf("期望的值为%v，实际的值为%v", expectedValue, actualValue)
	}
}
