package dbUtil

import (
	"database/sql"
	"errors"
)

// 数据表结构
type DataTable struct {
	// 行对象集合
	rowData []*DataRow

	// 列名称集合
	columnNames map[string]int
}

// 数据表初始化
// rows:原始的数据行信息
// 返回值：
// error:初始化的错误信息
func (this *DataTable) init(rows *sql.Rows) error {
	defer func() {
		rows.Close()
	}()

	// 读取列信息和保存列名称
	tmpColumns, errMsg := rows.Columns()
	if errMsg != nil {
		return errMsg
	}
	this.columnNames = make(map[string]int)
	for index, val := range tmpColumns {
		this.columnNames[val] = index
	}

	// 读取行数据
	this.rowData = make([]*DataRow, 0)
	columnCount := len(this.columnNames)

	args := make([]interface{}, columnCount)
	for rows.Next() {
		values := make([]interface{}, columnCount)
		for i := 0; i < columnCount; i++ {
			args[i] = &values[i]
		}
		rows.Scan(args...)

		this.rowData = append(this.rowData, newDataRow(this, values))
	}

	return nil
}

// 获取原始单元格值（一般为:string或[]byte）
// rowIndex:行序号
// cellIndex:单元格序号
// 返回值:
// interface{}:原始单元格值（一般为:string或[]byte）
// error:获取错误信息
func (this *DataTable) OriginCellValueByIndex(rowIndex int, cellIndex int) (interface{}, error) {
	if len(this.rowData) <= rowIndex {
		return nil, errors.New("row out of range")
	}

	rowItem := this.rowData[rowIndex]
	if len(rowItem.cells) <= cellIndex {
		return nil, errors.New("column out of range")
	}

	return rowItem.OriginCellValue(cellIndex)
}

// 获取原始单元格值（一般为:string或[]byte）
// rowIndex:行序号
// cellIndex:单元格序号
// 返回值:
// interface{}:原始单元格值（一般为:string或[]byte）
// error:获取错误信息
func (this *DataTable) OriginCellValueByCellName(rowIndex int, cellName string) (interface{}, error) {
	if len(this.rowData) <= rowIndex {
		return nil, errors.New("row out of range")
	}

	rowItem := this.rowData[rowIndex]

	return rowItem.OriginCellValueByName(cellName)
}

// 获取字符串的单元格值（有可能为nil）
// rowIndex:行序号
// cellIndex:单元格序号
// 返回值:
// interface{}:字符串的单元格值（有可能为nil）
// error:获取错误信息
func (this *DataTable) CellValueByIndex(rowIndex int, cellIndex int) (interface{}, error) {
	if len(this.rowData) <= rowIndex {
		return nil, errors.New("row out of range")
	}

	rowItem := this.rowData[rowIndex]
	if len(rowItem.cells) <= cellIndex {
		return nil, errors.New("column out of range")
	}

	return rowItem.CellValue(cellIndex)
}

// 获取字符串的单元格值（有可能为nil）
// rowIndex:行序号
// cellIndex:单元格序号
// 返回值:
// interface{}:字符串的单元格值（有可能为nil）
// error:获取错误信息
func (this *DataTable) CellValueByName(rowIndex int, cellName string) (interface{}, error) {
	if len(this.rowData) <= rowIndex {
		return nil, errors.New("row out of range")
	}

	rowItem := this.rowData[rowIndex]

	return rowItem.CellValueByName(cellName)
}

// 获取行对象
// rowIndex:行序号
// 返回值:
// *DataRow:行对象
// error:错误信息
func (this *DataTable) Row(rowIndex int) (*DataRow, error) {
	if len(this.rowData) <= rowIndex {
		return nil, errors.New("row out of range")
	}

	return this.rowData[rowIndex], nil
}

// 根据列名获取列序号
// cellName:列名
// 返回值:
// int:列序号
func (this *DataTable) cellIndex(cellName string) int {
	cellIndex, isExist := this.columnNames[cellName]
	if isExist == false {
		return -1
	}

	return cellIndex
}

// 获取所有列的名字
// 返回值:
// []string:列字段名集合
func (this *DataTable) Columns() []string {
	result := make([]string, len(this.columnNames))
	for key, val := range this.columnNames {
		result[val] = key
	}

	return result
}

// 获取列数量
// 返回值:
// int:列数量
func (this *DataTable) ColumnCount() int {
	return len(this.columnNames)
}

// 获取数据行数
// 返回值:
// int:行数
func (this *DataTable) RowCount() int {
	return len(this.rowData)
}

// 新建数据表对象
// rows:数据行对象
// 返回值：
// *DataTable:数据表对象
// error:错误信息
func NewDataTable(rows *sql.Rows) (*DataTable, error) {
	table := &DataTable{}
	errMsg := table.init(rows)
	if errMsg != nil {
		return nil, errMsg
	}

	return table, nil
}
