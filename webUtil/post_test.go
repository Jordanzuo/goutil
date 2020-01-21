package webUtil

import (
	"/zlibUtil"
	"testing"
)

func TestPostWebData(t *testing.T) {
	weburl := "http://managecentertest.qyc.moqikaka.com/API/ServerActivate.ashx"
	postDict := make(map[string]string, 0)
	postDict["ServerGroupID"] = "20002"
	resp, err := PostWebData(weburl, postDict, nil)
	if err != nil {
		t.Errorf("测试错误，返回的结果为：%s", err)
	}

	if len(resp) == 0 {
		t.Errorf("返回的数据为空，期望不为空")
	}

	// 将收到的数据进行zlib解压缩
	_, err = zlibUtil.Decompress(resp)
	if err != nil {
		t.Errorf("Error:%s", err)
	}
}
