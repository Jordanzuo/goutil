package mathUtil

import (
	"fmt"
)

// int类型区间对象，表示连续的int类型区间
type IntRegion struct {
	Lower int
	Upper int
}

func (this *IntRegion) String() string {
	return fmt.Sprintf("%d-%d", this.Lower, this.Upper)
}

// 是否包含指定的值
func (this *IntRegion) Contains(value int) bool {
	return this.Lower <= value && value <= this.Upper
}

// 是否是有序的
func (this *IntRegion) IsSorted() bool {
	return this.Lower < this.Upper
}

// 创建int类型区间对象
func NewIntRegion(lower, upper int) *IntRegion {
	return &IntRegion{
		Lower: lower,
		Upper: upper,
	}
}
