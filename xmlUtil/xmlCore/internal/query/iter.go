package query

import "github.com/Jordanzuo/goutil/xmlUtil/xmlCore/xpath"

type Iterator interface {
	Current() xpath.NodeNavigator
}
