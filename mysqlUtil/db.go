package mysqlUtil

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Jordanzuo/goutil/logUtil"
	_ "github.com/go-sql-driver/mysql"
)

// 打开数据库连接
// connectionString：数据库连接字符串
// 返回值：
// 数据库对象
func OpenMysqlConnection(connectionString string) (*sql.DB, error) {
	connectionSlice := strings.Split(connectionString, "||")
	if len(connectionSlice) != 3 {
		return nil, fmt.Errorf("数据库连接配置不完整，当前的为：%s", connectionString)
	}

	// 建立数据库连接
	db, err := sql.Open("mysql", connectionSlice[0])
	if err != nil {
		return nil, fmt.Errorf("打开游戏数据库失败,连接字符串为：%s", connectionString)
	}

	// 设置连接池相关
	maxOpenConns_string := strings.Replace(connectionSlice[1], "MaxOpenConns=", "", 1)
	maxOpenCons, err := strconv.Atoi(maxOpenConns_string)
	if err != nil {
		return nil, fmt.Errorf("MaxOpenConns必须为int型,连接字符串为：%s", connectionString)
	}

	maxIdleConns_string := strings.Replace(connectionSlice[2], "MaxIdleConns=", "", 1)
	maxIdleConns, err := strconv.Atoi(maxIdleConns_string)
	if err != nil {
		return nil, fmt.Errorf("MaxIdleConns必须为int型,连接字符串为：%s", connectionString)
	}

	if maxOpenCons > 0 && maxIdleConns > 0 {
		db.SetMaxOpenConns(maxOpenCons)
		db.SetMaxIdleConns(maxIdleConns)

		go func() {
			// 处理内部未处理的异常，以免导致主线程退出，从而导致系统崩溃
			defer func() {
				if r := recover(); r != nil {
					logUtil.LogUnknownError(r)
				}
			}()

			for {
				// 每分钟ping一次数据库
				time.Sleep(time.Minute)

				if err := db.Ping(); err != nil {
					logUtil.Log(fmt.Sprintf("Ping数据库失败,连接字符串为：%s,错误信息为：%s", connectionString, err), logUtil.Error, true)
				}
			}
		}()
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Ping数据库失败,连接字符串为：%s,错误信息为：%s", connectionString, err)
	}

	return db
}

// 开始事务
// db:数据库对象
// 返回值：
// 事务对象
// 错误对象
func BeginTransaction(db *sql.DB) (*sql.Tx, error) {
	tx, err := db.Begin()
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
