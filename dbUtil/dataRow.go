package dbUtil

import (
	"errors"
	"time"
)

// 数据行结果
type DataRow struct {
	// 所属数据表
	table *DataTable

	// 行的所有值
	cells []interface{}
}

// 行的所有原始值
func (this *DataRow) CellOriginValues() []interface{} {
	return this.cells
}

// 值的个数
func (this *DataRow) Len() int {
	return len(this.cells)
}

// 单元格的字符串值（可能为nil）,如果有设置连接字符串：parseTime=true，则会有time.Time
// celIndex:单元格序号
// 返回值:
// interface{}:单元格的字符串值
// error:错误信息
func (this *DataRow) CellValue(celIndex int) (interface{}, error) {
	if len(this.cells) <= celIndex {
		return nil, errors.New("cell out of range")
	}

	// 检查是否为nil
	if this.cells[celIndex] == nil {
		return nil, nil
	}

	// 转换为字符串
	switch this.cells[celIndex].(type) {
	case []byte:
		return string(this.cells[celIndex].([]byte)), nil
	case string:
		return this.cells[celIndex].(string), nil
	case time.Time:
		return this.cells[celIndex].(time.Time), nil
	}

	return nil, errors.New("unknown value type")
}

// 单元格的字符串值（可能为nil）,如果有设置连接字符串：parseTime=true，则会有time.Time
// cellName:单元格名称
// 返回值:
// interface{}:单元格的字符串值
// error:错误信息
func (this *DataRow) CellValueByName(cellName string) (interface{}, error) {
	celIndex := this.table.cellIndex(cellName)
	if celIndex < 0 {
		return nil, errors.New("cell name no exist")
	}

	return this.CellValue(celIndex)
}

// 单元格的原始值
// celIndex:单元格序号
// 返回值:
// interface{}:单元格的字符串值
// error:错误信息
func (this *DataRow) OriginCellValue(celIndex int) (interface{}, error) {
	if len(this.cells) <= celIndex {
		return nil, errors.New("cell out of range")
	}

	return this.cells[celIndex], nil
}

// 单元格的原始值
// cellName:单元格名称
// 返回值:
// interface{}:单元格的字符串值
// error:错误信息
func (this *DataRow) OriginCellValueByName(cellName string) (interface{}, error) {
	celIndex := this.table.cellIndex(cellName)
	if celIndex < 0 {
		return nil, errors.New("cell name no exist")
	}

	return this.OriginCellValue(celIndex)
}

// 创建单元格对象
// _table:所属表对象
// _cells:单元格的值集合
func newDataRow(_table *DataTable, _cells []interface{}) *DataRow {
	return &DataRow{
		table: _table,
		cells: _cells,
	}
}
