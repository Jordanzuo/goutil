/*
xml操作工具类：
此操作工具类来源于:https://github.com/antchfx/xquery
根据实际情况。我去掉了对golang.org/x/net/html/charset的依赖，并添加了各种xml加载函数

由于go默认编码的原因，如果文本编码不是utf-8 将会出现乱码

使用方式
	root := xmlUtil.LoadFromString(xml)
	nodes:= root.SelectElements(xpath)
*/
package xmlUtil
