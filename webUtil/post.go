package webUtil

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

// POST数据
// weburl：远程服务器地址
// postDict：post的数据字典
// headers：包头
// 返回值：
// 返回的字节
// 错误对象
func PostWebData(weburl string, postDict map[string]string, headers map[string]string) ([]byte, error) {
	// 组装POST数据
	postValues := url.Values{}
	for key, value := range postDict {
		postValues.Set(key, value)
	}

	postDataStr := postValues.Encode()
	postDataBytes := []byte(postDataStr)
	postDataBytesReader := bytes.NewReader(postDataBytes)

	// 构造请求对象
	httpRequest, _ := http.NewRequest("POST", weburl, postDataBytesReader)

	// 处理头部
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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

// POST Byte数组
// weburl：远程服务器地址
// data：post的Byte数组
// headers：包头
// 返回值：
// 返回的字节
// 错误对象
func PostByteData(weburl string, data []byte, headers map[string]string) ([]byte, error) {
	// 组装POST数据
	postDataBytesReader := bytes.NewReader(data)

	// 构造请求对象
	httpRequest, _ := http.NewRequest("POST", weburl, postDataBytesReader)

	// 处理头部
	httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
