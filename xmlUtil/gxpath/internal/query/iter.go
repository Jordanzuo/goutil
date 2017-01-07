package query

import "github.com/Jordanzuo/goutil/xmlUtil/gxpath/xpath"

type Iterator interface {
	Current() xpath.NodeNavigator
}
