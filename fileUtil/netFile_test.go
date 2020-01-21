package fileUtil

import "testing"

func TestDownLoadNetFile(t *testing.T) {
	err := DownLoadNetFile("https://www.baidu.com/img/bd_logo1.png", "./", "baidu.png", false)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("成功了")
}
