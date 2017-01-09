package xmlUtil

import (
	"fmt"

	"github.com/Jordanzuo/goutil/xmlUtil/gxpath"
	"github.com/Jordanzuo/goutil/xmlUtil/gxpath/xpath"
)

// xml节点查找结构(用于遍历xml节点)
type xmlNodeNavigator struct {
	root, curr *Node
	attr       int
}

// 节点类型
func (x *xmlNodeNavigator) NodeType() xpath.NodeType {
	switch x.curr.Type {
	case CommentNode:
		return xpath.CommentNode
	case TextNode:
		return xpath.TextNode
	case DeclarationNode, DocumentNode:
		return xpath.RootNode
	case ElementNode:
		if x.attr != -1 {
			return xpath.AttributeNode
		}
		return xpath.ElementNode
	}
	panic(fmt.Sprintf("unknown XML node type: %v", x.curr.Type))
}

// 当前查找的节点名或属性名
func (x *xmlNodeNavigator) LocalName() string {
	if x.attr != -1 {
		return x.curr.Attr[x.attr].Name.Local
	}
	return x.curr.NodeName

}

// 名节点前缀
func (x *xmlNodeNavigator) Prefix() string {
	return x.curr.Namespace
}

// 节点值或属性值
func (x *xmlNodeNavigator) Value() string {
	switch x.curr.Type {
	case CommentNode:
		return x.curr.NodeName
	case ElementNode:
		if x.attr != -1 {
			return x.curr.Attr[x.attr].Value
		}
		return x.curr.InnerText()
	case TextNode:
		return x.curr.NodeName
	}
	return ""
}

// 创建一个拷贝对象
func (x *xmlNodeNavigator) Copy() xpath.NodeNavigator {
	n := *x
	return &n
}

// 移动到根节点
func (x *xmlNodeNavigator) MoveToRoot() {
	x.curr = x.root
}

// 移动到父节点
func (x *xmlNodeNavigator) MoveToParent() bool {
	if node := x.curr.Parent; node != nil {
		x.curr = node
		return true
	}
	return false
}

// 移动到下一个属性
func (x *xmlNodeNavigator) MoveToNextAttribute() bool {
	if x.attr >= len(x.curr.Attr)-1 {
		return false
	}
	x.attr++
	return true
}

// 移动到子节点
func (x *xmlNodeNavigator) MoveToChild() bool {
	if node := x.curr.FirstChild; node != nil {
		x.curr = node
		return true
	}
	return false
}

// 移动到第一个节点
func (x *xmlNodeNavigator) MoveToFirst() bool {
	if x.curr.PrevSibling == nil {
		return false
	}
	for {
		node := x.curr.PrevSibling
		if node == nil {
			break
		}
		x.curr = node
	}
	return true
}

// 节点的值
func (x *xmlNodeNavigator) String() string {
	return x.Value()
}

// 移动到下一个兄弟节点
func (x *xmlNodeNavigator) MoveToNext() bool {
	if node := x.curr.NextSibling; node != nil {
		x.curr = node
		return true
	}
	return false
}

// 移动到上一个兄弟节点
func (x *xmlNodeNavigator) MoveToPrevious() bool {
	if node := x.curr.PrevSibling; node != nil {
		x.curr = node
		return true
	}
	return false
}

// 移动到指定节点
func (x *xmlNodeNavigator) MoveTo(other xpath.NodeNavigator) bool {
	node, ok := other.(*xmlNodeNavigator)
	if !ok || node.root != x.root {
		return false
	}

	x.curr = node.curr
	x.attr = node.attr
	return true
}

// CreateXPathNavigator creates a new xpath.NodeNavigator for the specified html.Node.
func CreateXPathNavigator(top *Node) xpath.NodeNavigator {
	return &xmlNodeNavigator{curr: top, root: top, attr: -1}
}

// 按照xpath查找所有匹配的节点
// top:根节点
// expr:xpath表达式
// 返回值:
// []*Node:结果
func Find(top *Node, expr string) []*Node {
	t := gxpath.Select(CreateXPathNavigator(top), expr)
	var elems []*Node
	for t.MoveNext() {
		elems = append(elems, (t.Current().(*xmlNodeNavigator)).curr)
	}
	return elems
}

// 按照xpath查找第一个匹配的节点
// top:根节点
// expr:xpath表达式
// 返回值:
// *Node:查找到的第一个节点
func FindOne(top *Node, expr string) *Node {
	t := gxpath.Select(CreateXPathNavigator(top), expr)
	var elem *Node
	if t.MoveNext() {
		elem = (t.Current().(*xmlNodeNavigator)).curr
	}
	return elem
}

// FindEach searches the html.Node and calls functions cb.
func FindEach(top *Node, expr string, cb func(int, *Node)) {
	t := gxpath.Select(CreateXPathNavigator(top), expr)
	var i int
	for t.MoveNext() {
		cb(i, (t.Current().(*xmlNodeNavigator)).curr)
		i++
	}
}
