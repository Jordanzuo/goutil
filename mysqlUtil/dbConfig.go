package mysqlUtil

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// 数据库配置对象
type DBConfig struct {
	// 连接字符串
	ConnectionString string

	// 最大开启连接数量
	MaxOpenConns int

	// 最大空闲连接数量
	MaxIdleConns int
}

func (this *DBConfig) String() string {
	bytes, _ := json.Marshal(this)
	return string(bytes)
}

// 创建数据库配置对象
// connectionString：连接字符串
// maxOpenConns：最大开启连接数
// maxIdleConns：最大空闲连接数
// 返回值：
// 数据库配置对象
func NewDBConfig(connectionString string, maxOpenConns, maxIdleConns int) *DBConfig {
	return &DBConfig{
		ConnectionString: connectionString,
		MaxOpenConns:     maxOpenConns,
		MaxIdleConns:     maxIdleConns,
	}
}

// 创建数据库配置对象
// dbConfigStr：数据库配置字符串，格式：root:moqikaka3306@tcp(10.1.0.10:3306)/gameserver_data?charset=utf8&parseTime=true&loc=Local&timeout=60s||MaxOpenConns=10||MaxIdleConns=5
// 返回值：
// 数据库配置对象
// 错误对象
func NewDBConfig2(dbConfigStr string) (*DBConfig, error) {
	connectionString := ""
	maxOpenConns := 0
	maxIdleConns := 0

	// 按照||来进行分割
	connSlice := strings.Split(dbConfigStr, "||")
	length := len(connSlice)
	if length != 1 && length != 3 {
		return nil, fmt.Errorf("dbConfigStr:%s格式不正确，length:%d", dbConfigStr, length)
	}

	// 获取连接字符串
	connectionString = connSlice[0]
	if connectionString == "" {
		return nil, fmt.Errorf("dbConfigStr:%s格式不正确，length:%d", dbConfigStr, length)
	}

	var err error

	// 如果配置了MaxOpenConns、MaxIdleConns，则进行解析
	if length == 3 {
		// 获取连接池相关
		maxOpenConns_string := strings.Replace(connSlice[1], "MaxOpenConns=", "", 1)
		maxOpenConns, err = strconv.Atoi(maxOpenConns_string)
		if err != nil {
			err = fmt.Errorf("MaxOpenConns必须为int型,连接字符串为：%s", connectionString)
			return nil, err
		}

		maxIdleConns_string := strings.Replace(connSlice[2], "MaxIdleConns=", "", 1)
		maxIdleConns, err = strconv.Atoi(maxIdleConns_string)
		if err != nil {
			err = fmt.Errorf("MaxIdleConns必须为int型,连接字符串为：%s", connectionString)
			return nil, err
		}
	}

	return NewDBConfig(connectionString, maxOpenConns, maxIdleConns), nil
}
