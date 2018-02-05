package webUtil

import (
	"testing"
)

func TestGetWebData(t *testing.T) {
	weburl := "http://www.baidu.com"
	resp, err := GetWebData(weburl, nil)
	if err != nil {
		t.Errorf("测试错误，返回的结果为：%s", err)
	}

	if len(resp) == 0 {
		t.Errorf("返回的数据为空，期望不为空")
	}
}

func TestGetWebData2(t *testing.T) {
	weburl := "http://www.baidu.com"
	data := make(map[string]string)
	status, resp, err := GetWebData2(weburl, data, nil, nil)
	if status != 200 || err != nil {
		t.Errorf("Test failed. status:%d, err:%s", status, err)
	}

	if len(resp) == 0 {
		t.Errorf("The result is empty, but we expect not empty")
	}

	data["Name"] = "Jordan"
	data["Age"] = "32"
	status, resp, err = GetWebData2(weburl, data, nil, nil)
	if status != 200 || err != nil {
		t.Errorf("Test failed. status:%d, err:%s", status, err)
	}

	if len(resp) == 0 {
		t.Errorf("The result is empty, but we expect not empty")
	}
}
