package webUtil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GET数据
// weburl：远程服务器地址
// header：包头集合
// 返回值：
// 返回的字节
// 错误对象
func GetWebData(weburl string, header map[string]string) (result []byte, err error) {
	// 构造请求对象
	var request *http.Request
	request, err = http.NewRequest("GET", weburl, nil)
	if err != nil {
		return
	}

	// 处理包头
	if header != nil {
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
	if response.StatusCode == http.StatusOK {
		result, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return
		}
	}

	return
}

// GET数据（新方法）
// weburl:远程服务器地址
// data：数据集合
// header:包头内容
// transport:transport对象
// 返回值
// http StatusCode
// 字节数组
// 错误对象
func GetWebData2(weburl string, data map[string]string, header map[string]string, transport *http.Transport) (statusCode int, result []byte, err error) {
	// 处理data，将data以key=value的形式拼接到weburl后，形成一个完整的url
	dataStr := ""
	count := 0
	for k, v := range data {
		if count == len(data)-1 {
			dataStr += fmt.Sprintf("%s=%s", k, v)
		} else {
			dataStr += fmt.Sprintf("%s=%s&", k, v)
		}
		count += 1
	}

	if dataStr != "" {
		if strings.Contains(weburl, "?") {
			weburl = fmt.Sprintf("%s&%s", weburl, dataStr)
		} else {
			weburl = fmt.Sprintf("%s?%s", weburl, dataStr)
		}
	}

	// 构造请求对象
	var request *http.Request
	request, err = http.NewRequest("GET", weburl, nil)
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
