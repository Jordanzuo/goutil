package appChargeUtil

import (
	"encoding/json"
	"fmt"
	"github.com/Jordanzuo/goutil/typeUtil"
	"strconv"
)

// APP Store充值收据对象
type Receipt struct {
	// Bvrs
	Bvrs string

	// BundleIdentifier
	BundleIdentifier string

	// 产品Id
	ProductId string

	// 交易Id
	TransactionId string

	// 数量
	Quantity int

	// 状态
	Status int
}

// BundleIdentifier是否有效
// bundleIdentifierList：配置的BundleIdentifier列表
// 返回值：
// 是否有效
func (this *Receipt) IsBundleIdentifierValid(bundleIdentifierList []string) bool {
	for _, item := range bundleIdentifierList {
		if this.BundleIdentifier == item {
			return true
		}
	}

	return false
}

// ProductId是否有效
// productId：输入的ProductId
// 返回值：
// 是否有效
func (this *Receipt) IsProductIdValid(productId string) bool {
	return this.ProductId == productId
}

// 转换为字符串
// 返回值：
// 字符串
func (this *Receipt) String() string {
	return fmt.Sprintf("{Bvrs=%s,BundleIdentifier=%s,ProductId=%s,TransactionId=%s,Quantity=%d,Status=%d}", this.Bvrs, this.BundleIdentifier, this.ProductId, this.TransactionId, this.Quantity, this.Status)
}

// 创建新的收据对象
// receiptInfo：收据信息
// 返回值：
// 收据对象
// 错误对象
func newReceipt(receiptInfo string) (receiptObj *Receipt, err error) {
	// 将接收的数据转化为map类型的对象
	receiptDataMap := make(map[string]interface{})
	if err = json.Unmarshal([]byte(receiptInfo), &receiptDataMap); err != nil {
		return
	}
	mapData := typeUtil.NewMapData(receiptDataMap)

	// 创建空对象
	receiptObj = &Receipt{}

	// 定义、并判断返回状态
	var status int
	if status, err = mapData.Int("status"); err != nil {
		return
	}
	if status != 0 {
		err = fmt.Errorf("状态:%s不正确", status)
		return
	} else {
		receiptObj.Status = status
	}

	// Receipt is actually a child
	var ok bool
	if receiptDataMap, ok = mapData["receipt"].(map[string]interface{}); !ok {
		err = fmt.Errorf("receipt错误")
		return
	}
	mapData = typeUtil.NewMapData(receiptDataMap)

	// 用返回值对本对象的属性进行赋值
	if receiptObj.Bvrs, err = mapData.String("bvrs"); err != nil {
		return
	}
	if receiptObj.BundleIdentifier, err = mapData.String("bid"); err != nil {
		return
	}
	if receiptObj.ProductId, err = mapData.String("product_id"); err != nil {
		return
	}
	if receiptObj.TransactionId, err = mapData.String("transaction_id"); err != nil {
		return
	}
	if quality, err1 := mapData.String("quantity"); err1 != nil {
		err = err1
		return
	} else {
		if receiptObj.Quantity, err = strconv.Atoi(quality); err != nil {
			return
		}
	}

	return
}
