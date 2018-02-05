package xmlUtil

import (
	"bytes"
	"container/list"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"
)

// A NodeType is the type of a Node.
type NodeType uint

const (
	// 文档对象节点（根节点）
	DocumentNode NodeType = iota

	// 头不声明节点
	DeclarationNode

	// 元素节点
	ElementNode

	// 节点文本
	TextNode

	// 注释
	CommentNode
)

// xml节点对象
type Node struct {
	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

	Type      NodeType
	NodeName  string
	Namespace string
	Attr      []xml.Attr

	level int // node level in the tree
}

// InnerText returns the text between the start and end tags of the object.
func (n *Node) InnerText() string {
	if n.Type == TextNode || n.Type == CommentNode {
		return n.NodeName
	}

	var buf bytes.Buffer
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		// filt commentnode
		if child.Type == CommentNode {
			continue
		}

		buf.WriteString(child.InnerText())
	}

	return buf.String()
}

func outputXML(buf *bytes.Buffer, n *Node) {
	if n.Type == TextNode || n.Type == CommentNode {
		buf.WriteString(strings.TrimSpace(n.NodeName))
		return
	}
	buf.WriteString("<" + n.NodeName)
	for _, attr := range n.Attr {
		if attr.Name.Space != "" {
			buf.WriteString(fmt.Sprintf(` %s:%s="%s"`, attr.Name.Space, attr.Name.Local, attr.Value))
		} else {
			buf.WriteString(fmt.Sprintf(` %s="%s"`, attr.Name.Local, attr.Value))
		}
	}
	buf.WriteString(">")
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		outputXML(buf, child)
	}
	buf.WriteString(fmt.Sprintf("</%s>", n.NodeName))
}

// OutputXML returns the text that including tags name.
func (n *Node) OutputXML() string {
	var buf bytes.Buffer
	outputXML(&buf, n)
	return buf.String()
}

// get all children
func (n *Node) Children() []*Node {
	childrenList := make([]*Node, 0)
	nowChild := n.FirstChild
	for {
		if nowChild == nil {
			break
		}

		childrenList = append(childrenList, nowChild)
		if nowChild == n.LastChild {
			break
		}

		nowChild = nowChild.NextSibling
	}

	return childrenList
}

// get children len
func (n *Node) ChildrenLen() int {
	var count int = 0
	nowChild := n.FirstChild
	for {
		if nowChild == nil {
			break
		}

		count += 1
		if nowChild == n.LastChild {
			break
		}

		nowChild = nowChild.NextSibling
	}

	return count
}

// get all attribute
func (n *Node) ALLAttribute() []xml.Attr {
	if n.Attr == nil {
		return nil
	}

	return n.Attr[:]
}

// 获取属性个数
func (n *Node) AttributeLen() int {
	if n.Attr == nil {
		return 0
	}

	return len(n.Attr)
}

// 输出所有(主要用于测试)
func (this *Node) OutALL() {
	stack := list.New()
	tmpItem := this
	for {
		if tmpItem != nil {
			stack.PushBack(tmpItem)
			tmpItem = tmpItem.NextSibling
		}

		break
	}

	for {
		if stack.Len() <= 0 {
			break
		}

		nowNode := stack.Front().Value.(*Node)
		stack.Remove(stack.Front())
		for _, item := range nowNode.Children() {
			stack.PushFront(item)
		}

		fmt.Println("name:", nowNode.NodeName, " level: ", nowNode.level, " attr:", nowNode.Attr)
	}
}

// SelectElements finds child elements with the specified name.
func (n *Node) SelectElements(name string) []*Node {
	return Find(n, name)
}

// SelectElements finds child elements with the specified name.
func (n *Node) SelectElement(name string) *Node {
	return FindOne(n, name)
}

// SelectAttr returns the attribute value with the specified name.
func (n *Node) SelectAttr(name string) (string, bool) {
	var local, space string
	local = name
	if i := strings.Index(name, ":"); i > 0 {
		space = name[:i]
		local = name[i+1:]
	}
	for _, attr := range n.Attr {
		if attr.Name.Local == local && attr.Name.Space == space {
			return attr.Value, true
		}
	}
	return "", false
}

// 给节点添加属性值
func addAttr(n *Node, key, val string) {
	var attr xml.Attr
	if i := strings.Index(key, ":"); i > 0 {
		attr = xml.Attr{
			Name:  xml.Name{Space: key[:i], Local: key[i+1:]},
			Value: val,
		}
	} else {
		attr = xml.Attr{
			Name:  xml.Name{Local: key},
			Value: val,
		}
	}

	n.Attr = append(n.Attr, attr)
}

// 给节点添加子节点
func addChild(parent, n *Node) {
	n.Parent = parent
	if parent.FirstChild == nil {
		parent.FirstChild = n
	} else {
		parent.LastChild.NextSibling = n
		n.PrevSibling = parent.LastChild
	}

	parent.LastChild = n
}

// 给节点添加兄弟节点
func addSibling(sibling, n *Node) {
	n.Parent = sibling.Parent
	sibling.NextSibling = n
	n.PrevSibling = sibling
	if sibling.Parent != nil {
		sibling.Parent.LastChild = n
	}
}

// 从reader里面加载xml文档
func LoadFromReader(r io.Reader) (*Node, error) {
	var (
		decoder  = xml.NewDecoder(r) //// xml解码对象
		doc      = &Node{Type: DocumentNode}
		level    = 0
		declared = false
	)
	var prev *Node = doc
	for {
		tok, err := decoder.Token()
		switch {
		case err == io.EOF:
			goto quit
		case err != nil:
			return nil, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			//if !declared { // if have no xml node，we also need add children
			//	return nil, errors.New("xml: document is invalid")
			//}
			// if there is no xml node.then create it
			if declared == false {
				level++

				tmpNode := &Node{Type: DeclarationNode, level: level}
				addAttr(tmpNode, "version", "1.0")
				addAttr(tmpNode, "encoding", "UTF-8")
				declared = true
				if level == prev.level {
					addSibling(prev, tmpNode)
				} else if level > prev.level {
					addChild(prev, tmpNode)
				}

				prev = tmpNode
			}
			node := &Node{
				Type:      ElementNode,
				NodeName:  tok.Name.Local,
				Namespace: tok.Name.Space,
				Attr:      tok.Attr,
				level:     level,
			}
			//fmt.Println(fmt.Sprintf("start > %s : %d", node.Data, level))
			if level == prev.level {
				addSibling(prev, node)
			} else if level > prev.level {
				addChild(prev, node)
			} else if level < prev.level {
				for i := prev.level - level; i > 1; i-- {
					prev = prev.Parent
				}
				addSibling(prev.Parent, node)
			}
			prev = node
			level++
		case xml.EndElement:
			level--
		case xml.CharData:
			node := &Node{Type: TextNode, NodeName: string(tok), level: level}
			if level == prev.level {
				addSibling(prev, node)
			} else if level > prev.level {
				addChild(prev, node)
			}
		case xml.Comment:
			node := &Node{Type: CommentNode, NodeName: string(tok), level: level}
			if level == prev.level {
				addSibling(prev, node)
			} else if level > prev.level {
				addChild(prev, node)
			}
		case xml.ProcInst: // Processing Instruction
			if declared || (!declared && tok.Target != "xml") {
				return nil, errors.New("xml: document is invalid")
			}
			level++
			node := &Node{Type: DeclarationNode, level: level}
			pairs := strings.Split(string(tok.Inst), " ")
			for _, pair := range pairs {
				pair = strings.TrimSpace(pair)
				if i := strings.Index(pair, "="); i > 0 {
					addAttr(node, pair[:i], strings.Trim(pair[i+1:], `"`))
				}
			}
			declared = true
			if level == prev.level {
				addSibling(prev, node)
			} else if level > prev.level {
				addChild(prev, node)
			}
			prev = node
		case xml.Directive:
		}

	}
quit:
	return doc, nil
}
