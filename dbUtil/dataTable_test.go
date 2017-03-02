package dbUtil

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var onceLock *sync.Once = new(sync.Once)

// 所有常规数据类型测试
func TestNormalDataType(context *testing.T) {
	db := getTestConnection()
	rows, errMsg := db.Query("select * from hello")
	if errMsg != nil {
		context.Error(errMsg)
		return
	}

	var dt *DataTable
	if dt, errMsg = NewDataTable(rows); errMsg != nil {
		context.Error(errMsg)
		return
	}

	printDataRow(dt, context)
}

// 获取数据库连接
func getTestConnection() *DbConnection {
	onceLock.Do(func() {
		if errMsg := AddConnection("GameDb", "mysql",
			"root:1234@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true"); errMsg != nil {
			panic(errors.New("无法连接到数据库" + errMsg.Error()))
		}
	})
	// user:password@tcp(localhost:5555)/dbname?charset=utf8mb4,utf8&tls=skip-verify

	return GetConnection("GameDb")
}

// 打印所有数据
func printDataRow(data *DataTable, context *testing.T) {

	txt := ""
	colNames := data.Columns()
	for i := 0; i < len(colNames); i++ {
		txt = txt + "	" + colNames[i]
	}

	for i := 0; i < data.RowCount(); i++ {
		txt = txt + "\r\n"
		for colIndex := 0; colIndex < len(colNames); colIndex++ {
			val, _ := data.CellValueByIndex(i, colIndex)
			txt = fmt.Sprintf("%s	%v", txt, val)
		}
	}

	context.Log(txt)
}
