package mysqlUtil

import (
	"database/sql"
	"fmt"

	"github.com/Jordanzuo/goutil/logUtil"
	_ "github.com/go-sql-driver/mysql"
)

// 打开数据库连接
// connectionString：数据库连接字符串，格式：root:moqikaka3306@tcp(10.1.0.10:3306)/gameserver_data?charset=utf8&parseTime=true&loc=Local&timeout=60s||MaxOpenConns=10||MaxIdleConns=5
// 返回值：
// 数据库对象
// 错误对象
func OpenMysqlConnection(connectionString string) (dbObj *sql.DB, err error) {
	dbConfigObj, err1 := NewDBConfig2(connectionString)
	if err1 != nil {
		err = err1
		return
	}

	dbObj, err = OpenMysqlConnection3(dbConfigObj)
	return
}

// 打开数据库连接
// connectionString：数据库连接字符串
// maxOpenConns：最大打开的连接数
// maxIdleConns：最大处于闲置状态的连接数
// 返回值：
// 数据库对象
// 错误对象
func OpenMysqlConnection2(connectionString string, maxOpenConns, maxIdleConns int) (dbObj *sql.DB, err error) {
	dbConfigObj := NewDBConfig(connectionString, maxOpenConns, maxIdleConns)
	dbObj, err = OpenMysqlConnection3(dbConfigObj)
	return
}

// 建立Mysql数据库连接
// dbConfigObj：数据库配置对象
// 返回值：
// 数据库对象
// 错误对象
func OpenMysqlConnection3(dbConfigObj *DBConfig) (dbObj *sql.DB, err error) {
	// 建立数据库连接
	logUtil.DebugLog("开始连接Mysql数据库")
	dbObj, err = sql.Open("mysql", dbConfigObj.ConnectionString)
	if err != nil {
		err = fmt.Errorf("打开游戏数据库失败,连接字符串为：%s", dbConfigObj.ConnectionString)
		return
	}
	logUtil.DebugLog("连接Mysql数据库成功")

	if dbConfigObj.MaxOpenConns > 0 && dbConfigObj.MaxIdleConns > 0 {
		dbObj.SetMaxOpenConns(dbConfigObj.MaxOpenConns)
		dbObj.SetMaxIdleConns(dbConfigObj.MaxIdleConns)
	}

	if err = dbObj.Ping(); err != nil {
		err = fmt.Errorf("Ping数据库失败,连接字符串为：%s,错误信息为：%s", dbConfigObj.ConnectionString, err)
		return
	}

	return
}

// 测试数据库连接
// dbObj:数据库连对象
// 返回值：
// 错误对象
func TestConnection(dbObj *sql.DB) error {
	command := "SHOW DATABASES;"
	rows, err := dbObj.Query(command)
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

// 开始事务
// db:数据库对象
// 返回值：
// 事务对象
// 错误对象
func BeginTransaction(dbObj *sql.DB) (*sql.Tx, error) {
	tx, err := dbObj.Begin()
	if err != nil {
		logUtil.Log(fmt.Sprintf("开启事务失败，错误信息：%s", err), logUtil.Error, true)
	}

	return tx, err
}

// 提交事务
// tx:事务对象
// 返回值：
// 错误对象
func CommitTransaction(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		logUtil.Log(fmt.Sprintf("提交事务失败，错误信息：%s", err), logUtil.Error, true)
	}

	return err
}

// 记录Prepare错误
// command：执行的SQL语句
// err：错误对象
func WritePrepareError(command string, err error) {
	logUtil.Log(fmt.Sprintf("Prepare失败，错误信息：%s，command:%s", err, command), logUtil.Error, true)
}

// 记录Query错误
// command：执行的SQL语句
// err：错误对象
func WriteQueryError(command string, err error) {
	logUtil.Log(fmt.Sprintf("Query失败，错误信息：%s，command:%s", err, command), logUtil.Error, true)
}

// 记录Exec错误
// command：执行的SQL语句
// err：错误对象
func WriteExecError(command string, err error) {
	logUtil.Log(fmt.Sprintf("Exec失败，错误信息：%s，command:%s", err, command), logUtil.Error, true)
}

// 记录Scan错误
// command：执行的SQL语句
// err：错误对象
func WriteScanError(command string, err error) {
	logUtil.Log(fmt.Sprintf("Scan失败，错误信息：%s，command:%s", err, command), logUtil.Error, true)
}
