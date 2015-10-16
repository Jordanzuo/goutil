package webUtil

import (
	"io/ioutil"
	"net/http"
)

// GET数据
// weburl：远程服务器地址
// headers：包头
// 返回值：
// 返回的字节
func GetWebData(weburl string, headers map[string]string) ([]byte, error) {
	// 构造请求对象
	httpRequest, _ := http.NewRequest("GET", weburl, nil)

	// 处理包头
	if headers != nil {
		for key, value := range headers {
			httpRequest.Header.Add(key, value)
		}
	}

	// 构造httpClient对象
	httpClient := &http.Client{}

	// 发送数据
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	// 读取数据
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
