package configUtil

import (
	"testing"
)

var (
	config map[string]interface{}
	err    error
)

func TestReadJsonConfig(t *testing.T) {
	config, err = ReadJsonConfig("testdata/jsonConfig.ini")
	if err != nil {
		t.Errorf("读取JSON配置失败，错误信息为：%s", err)
	}
}

func TestReadIntJsonValue(t *testing.T) {
	actualValue, err := ReadIntJsonValue(config, "ServerGroupId")
	if err != nil {
		t.Errorf("读取JSON配置失败，错误信息为：%s", err)
	}

	expectedValue := 1
	if actualValue != expectedValue {
		t.Errorf("期望的值为%d，实际的值为%d", expectedValue, actualValue)
	}
}

func TestReadStringJsonValue(t *testing.T) {
	actualValue, err := ReadStringJsonValue(config, "ChatDBConnection")
	if err != nil {
		t.Errorf("读取JSON配置失败，错误信息为：%s", err)
	}

	expectedValue := "root:moqikaka@tcp(192.168.1.226:3306)/chatserver?charset=utf8&parseTime=true&loc=Local&timeout=30s"
	if actualValue != expectedValue {
		t.Errorf("期望的值为%s，实际的值为%s", expectedValue, actualValue)
	}
}

func TestReadBoolJsonValue(t *testing.T) {
	actualValue, err := ReadBoolJsonValue(config, "IfRecordMessage")
	if err != nil {
		t.Errorf("读取JSON配置失败，错误信息为：%s", err)
	}

	expectedValue := true
	if actualValue != expectedValue {
		t.Errorf("期望的值为%v，实际的值为%v", expectedValue, actualValue)
	}
}
