package mysqlUtil

import (
	"fmt"
	"strconv"
	"strings"
)

// 转换数据库连接字符串（从C#->Go)
// connString：C#的数据库连接字符串
// 返回值：
// Go的数据库连接字符串
func ConvertConnectionStringFromCSharpToGo(connString string) string {
	goConnString := "{userid}:{password}@tcp({datasource}:{port})/{database}?charset={charset}&parseTime=true&loc=Local&timeout={timeout}s||MaxOpenConns={MaxOpenConns}||MaxIdleConns={MaxIdleConns}"

	// 将字符串按;进行切割
	connStringList := strings.Split(connString, ";")
	var datasource string      // DataSource=10.162.2.205;
	var port string            // port=3306;
	var userid string          // UserId=admin;
	var password string        // Password=MOQIkaka$#@!1234;
	var database string        // Database=s201_dzz_log;
	var charset string         // charset=utf8;
	var pooling string         // pooling=true;
	var minimumpoolsize string // MinimumPoolSize=20; or min pool size=20;
	var maximumpoolsize string // maximumpoolsize=200; or max pool size=200;
	var commandtimeout string  // command timeout=60;

	// 遍历处理
	for _, item := range connStringList {
		// 将字符串按=进行切割
		subItemList := strings.Split(item, "=")

		// 对每一项进行判断
		if len(subItemList) != 2 {
			continue
		}

		// 先转换为小写字母
		switch strings.ToLower(subItemList[0]) {
		case "datasource":
			datasource = subItemList[1]
		case "port":
			port = subItemList[1]
		case "userid":
			userid = subItemList[1]
		case "password":
			password = subItemList[1]
		case "database":
			database = subItemList[1]
		case "charset":
			charset = subItemList[1]
		case "pooling":
			pooling = subItemList[1]
		case "minimumpoolsize", "min pool size":
			minimumpoolsize = subItemList[1]
		case "maximumpoolsize", "max pool size":
			maximumpoolsize = subItemList[1]
		case "command timeout":
			commandtimeout = subItemList[1]
		}
	}

	// 替换占位符
	goConnString = strings.Replace(goConnString, "{userid}", userid, 1)
	goConnString = strings.Replace(goConnString, "{password}", password, 1)
	goConnString = strings.Replace(goConnString, "{datasource}", datasource, 1)
	goConnString = strings.Replace(goConnString, "{port}", port, 1)
	goConnString = strings.Replace(goConnString, "{database}", database, 1)
	goConnString = strings.Replace(goConnString, "{charset}", charset, 1)
	goConnString = strings.Replace(goConnString, "{timeout}", commandtimeout, 1)
	if pooling == "true" {
		goConnString = strings.Replace(goConnString, "{MaxOpenConns}", maximumpoolsize, 1)
		goConnString = strings.Replace(goConnString, "{MaxIdleConns}", minimumpoolsize, 1)
	} else {
		goConnString = strings.Replace(goConnString, "{MaxOpenConns}", "0", 1)
		goConnString = strings.Replace(goConnString, "{MaxIdleConns}", "0", 1)
	}

	return goConnString
}

// 解析连接字符串(obsolete，建议使用NewDBConfig2)
// connString：数据库连接字符串
// 返回值
// 有效的连接字符串
// 最大开启连接数量
// 最大空闲连接数量
// 错误对象
func ParseConnectionString(connString string) (conn string, maxOpenConns int, maxIdleConns int, err error) {
	connSlice := strings.Split(connString, "||")
	length := len(connSlice)
	if length != 1 && length != 3 {
		err = fmt.Errorf("connString:%s格式不正确，length:%d", connString, length)
		return
	}

	// 获取连接字符串
	conn = connSlice[0]
	if conn == "" {
		err = fmt.Errorf("connString:%s格式不正确，length:%d", connString, length)
		return
	}

	// 如果只配置了连接字符串，则MaxOpenConns、MaxIdleConns取默认值
	if length == 1 {
		return
	}

	// 获取连接池相关
	maxOpenConns_string := strings.Replace(connSlice[1], "MaxOpenConns=", "", 1)
	maxOpenConns, err = strconv.Atoi(maxOpenConns_string)
	if err != nil {
		err = fmt.Errorf("MaxOpenConns必须为int型,连接字符串为：%s", connString)
		return
	}

	maxIdleConns_string := strings.Replace(connSlice[2], "MaxIdleConns=", "", 1)
	maxIdleConns, err = strconv.Atoi(maxIdleConns_string)
	if err != nil {
		err = fmt.Errorf("MaxIdleConns必须为int型,连接字符串为：%s", connString)
		return
	}

	return
}

// 连接字符串是否为C#格式
func IsCSharpStyle(connString string) bool {
	lowerString := strings.ToLower(connString)
	if strings.Contains(lowerString, "datasource") && strings.Contains(lowerString, "port") {
		return true
	}

	return false
}

// 连接字符串是否为Go格式
func IsGoStyle(connString string) bool {
	lowerString := strings.ToLower(connString)
	if strings.Contains(lowerString, "@tcp") {
		return true
	}

	return false
}
