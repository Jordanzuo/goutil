package mysqlUtil

import (
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
	var minimumpoolsize string // MinimumPoolSize=20;
	var maximumpoolsize string // maximumpoolsize=200;
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
		case "minimumpoolsize":
			minimumpoolsize = subItemList[1]
		case "maximumpoolsize":
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
