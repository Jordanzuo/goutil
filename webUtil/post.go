package webUtil

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	ContentType_Json = map[string]string{"Content-Type": "application/json;charset=UTF-8"}
	ContentType_Form = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
)

// POST数据
// weburl：远程服务器地址
// data：post的数据集合
// header：包头集合
// 返回值：
// 返回的字节
// 错误对象
func PostWebData(weburl string, data map[string]string, header map[string]string) (result []byte, err error) {
	// 组装POST数据
	postValues := url.Values{}
	for key, value := range data {
		postValues.Set(key, value)
	}
	postDataStr := postValues.Encode()
	byteData := []byte(postDataStr)

	// 调用发送Byte数组的方法
	result, err = PostByteData(weburl, byteData, header)

	return
}

// POST Byte数组
// weburl：远程服务器地址
// data：post的Byte数组
// header：包头集合
// 返回值：
// 返回的字节
// 错误对象
func PostByteData(weburl string, data []byte, header map[string]string) (result []byte, err error) {
	// 组装POST数据
	reader := bytes.NewReader(data)

	// 构造请求对象
	var request *http.Request
	request, err = http.NewRequest("POST", weburl, reader)
	if err != nil {
		return
	}

	// 处理头部（包括默认头部，以及传入的头部集合）
	if header == nil {
		for k, v := range ContentType_Form {
			request.Header.Add(k, v)
		}
	} else {
		for k, v := range header {
			request.Header.Add(k, v)
		}
	}

	// 构造httpClient对象
	client := &http.Client{}

	// 发送数据
	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	// 读取数据
	result, err = ioutil.ReadAll(response.Body)

	return
}

// POST Byte数组
// weburl：远程服务器地址
// data：post的Byte数组
// header：包头集合
// transport: transport对象
// 返回值：
// 返回的字节
// 错误对象
func PostByteDataWithTransport(weburl string, data []byte, header map[string]string, transport *http.Transport) (result *[]byte, err error) {
	// 组装POST数据
	reader := bytes.NewReader(data)

	// 构造请求对象
	var request *http.Request
	request, err = http.NewRequest("POST", weburl, reader)
	if err != nil {
		return
	}

	// 处理头部
	if header != nil {
		for k, v := range header {
			request.Header.Add(k, v)
		}
	}

	// 构造httpClient对象
	var client *http.Client
	if transport == nil {
		client = &http.Client{}
	} else {
		client = &http.Client{Transport: transport}
	}

	// 发送数据
	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		err = err1
		return
	}

	result = &body

	return
}

// POST map类型的数据
// weburl:远程服务器地址
// data:数据
// header:包头内容
// transport:transport对象
// 返回值
// http StatusCode
// 字节数组
// 错误对象
func PostMapData(weburl string, data map[string]string, header map[string]string, transport *http.Transport) (statusCode int, result []byte, err error) {
	// 组装POST数据
	postValues := url.Values{}
	for key, value := range data {
		postValues.Set(key, value)
	}
	postDataStr := postValues.Encode()
	byteData := []byte(postDataStr)

	statusCode, result, err = PostByteData2(weburl, byteData, header, transport)
	return
}

// POST byte类型的数据（新方法）
// weburl:远程服务器地址
// data:数据
// header:包头内容
// transport:transport对象
// 返回值
// http StatusCode
// 字节数组
// 错误对象
func PostByteData2(weburl string, data []byte, header map[string]string, transport *http.Transport) (statusCode int, result []byte, err error) {
	// 组装POST数据
	reader := bytes.NewReader(data)

	// 构造请求对象
	var request *http.Request
	request, err = http.NewRequest("POST", weburl, reader)
	if err != nil {
		return
	}

	// 处理头部
	if header != nil {
		for k, v := range header {
			request.Header.Add(k, v)
		}
	}

	// 构造httpClient对象
	var client *http.Client
	if transport == nil {
		client = &http.Client{}
	} else {
		client = &http.Client{Transport: transport}
	}

	// 发送数据
	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	// 获取返回值
	statusCode = response.StatusCode
	result, err = ioutil.ReadAll(response.Body)

	return
}
