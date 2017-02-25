package appChargeUtil

import (
	"fmt"
	"github.com/Jordanzuo/goutil/webUtil"
)

const (
	con_SandBoxUrl    = "https://sandbox.itunes.apple.com/verifyReceipt"
	con_ProductionUrl = "https://buy.itunes.apple.com/verifyReceipt"
)

// 验证充值是否有效
// bundleIdentifierList：配置的BundleIdentifier列表
// productId：输入的ProductId
// receiptData：订单数据
// isSandBox：是否为沙盒模式
// 返回值：
// 充值收据对象
// 是否有效
func ValidateCharge(bundleIdentifierList []string, productId, receiptData string, isSandBox bool) (receiptObj *Receipt, isValid bool) {
	// 判断参数是否为空
	if len(bundleIdentifierList) == 0 || productId == "" || receiptData == "" {
		return
	}

	// 获取Receipt对象
	var err error
	if receiptObj, err = getReceipt(receiptData, isSandBox); err != nil {
		fmt.Printf("err:%s\n", err)
		return
	}

	if receiptObj.IsBundleIdentifierValid(bundleIdentifierList) == false {
		return
	}

	if receiptObj.IsProductIdValid(productId) == false {
		return
	}

	isValid = true

	return
}

func getReceipt(receiptData string, isSandBox bool) (receiptObj *Receipt, err error) {
	weburl := con_ProductionUrl
	if isSandBox {
		weburl = con_SandBoxUrl
	}
	data := []byte(convertReceiptToPost(receiptData))
	var returnBytes []byte
	if returnBytes, err = webUtil.PostByteData(weburl, data, nil); err != nil {
		return
	}

	if len(returnBytes) == 0 {
		err = fmt.Errorf("返回的数据为空")
		return
	}

	receiptObj, err = newReceipt(string(returnBytes))

	return
}

func convertReceiptToPost(receiptData string) string {
	return fmt.Sprintf("{\"receipt-data\":\"%s\"}", receiptData)
}
