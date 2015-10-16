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
