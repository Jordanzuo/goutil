package appChargeUtil

import (
	"encoding/json"
	"fmt"

	"github.com/Jordanzuo/goutil/typeUtil"
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
	// 创建空对象
	receiptObj = &Receipt{}

	// 将接收的数据转化为map类型的对象
	receiptDataMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(receiptInfo), &receiptDataMap)
	if err != nil {
		return
	}
	mapData := typeUtil.NewMapData(receiptDataMap)

	// 定义、并判断返回状态
	receiptObj.Status, err = mapData.Int("status")
	if err != nil {
		return
	}
	if receiptObj.Status != 0 {
		err = fmt.Errorf("状态:%d不正确", receiptObj.Status)
		return
	}

	// Receipt is actually a child
	receiptDataMap, ok := mapData["receipt"].(map[string]interface{})
	if !ok {
		err = fmt.Errorf("receipt错误")
		return
	}
	mapData = typeUtil.NewMapData(receiptDataMap)

	// 用返回值对本对象的属性进行赋值
	receiptObj.Bvrs, err = mapData.String("bvrs")
	if err != nil {
		return
	}
	receiptObj.BundleIdentifier, err = mapData.String("bid")
	if err != nil {
		return
	}
	receiptObj.ProductId, err = mapData.String("product_id")
	if err != nil {
		return
	}
	receiptObj.TransactionId, err = mapData.String("transaction_id")
	if err != nil {
		return
	}
	receiptObj.Quantity, err = mapData.Int("quantity")
	if err != nil {
		return
	}

	return
}
