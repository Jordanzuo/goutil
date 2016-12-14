package configUtil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 读取JSON格式的配置文件
// config_file_path：配置文件路径
// 返回值：
// 配置内容的map格式数组
// 错误对象
func ReadJsonConfig_Array(config_file_path string) ([]map[string]interface{}, error) {
	// 读取配置文件（一次性读取整个文件，则使用ioutil）
	bytes, err := ioutil.ReadFile(config_file_path)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件的内容出错:%s", err)
	}

	// 使用json反序列化
	config := make([]map[string]interface{}, 0, 4)
	if err = json.Unmarshal(bytes, &config); err != nil {
		return nil, fmt.Errorf("反序列化配置文件的内容出错:%s", err)
	}

	return config, nil
}

func getConfigValue(config []map[string]interface{}, configName string) (configValue interface{}, exists bool, err error) {
	for _, configItem := range config {
		if configValue, exists = configItem[configName]; exists {
			break
		}
	}

	if !exists {
		err = fmt.Errorf("不存在名为%s的配置或配置为空", configName)
	}

	return
}

// 从config配置中获取int类型的配置值
// config：从config文件中反序列化出来的map对象
// configName：配置名称
// 返回值：
// 配置值
// 错误对象
func ReadIntJsonValue_Array(config []map[string]interface{}, configName string) (int, error) {
	configValue, exists, err := getConfigValue(config, configName)
	if !exists {
		return 0, err
	}

	configValue_float, ok := configValue.(float64)
	if !ok {
		return 0, fmt.Errorf("%s必须为int型", configName)
	}

	return int(configValue_float), nil
}

// 从config配置中获取string类型的配置值
// config：从config文件中反序列化出来的map对象
// configName：配置名称
// 返回值：
// 配置值
// 错误对象
func ReadStringJsonValue_Array(config []map[string]interface{}, configName string) (string, error) {
	configValue, exists, err := getConfigValue(config, configName)
	if !exists {
		return "", err
	}

	configValue_string, ok := configValue.(string)
	if !ok {
		return "", fmt.Errorf("%s必须为string型", configName)
	}

	return configValue_string, nil
}

// 从config配置中获取string类型的配置值
// config：从config文件中反序列化出来的map对象
// configName：配置名称
// 返回值：
// 配置值
// 错误对象
func ReadBoolJsonValue_Array(config []map[string]interface{}, configName string) (bool, error) {
	configValue, exists, err := getConfigValue(config, configName)
	if !exists {
		return false, err
	}

	configValue_bool, ok := configValue.(bool)
	if !ok {
		return false, fmt.Errorf("%s必须为bool型", configName)
	}

	return configValue_bool, nil
}
