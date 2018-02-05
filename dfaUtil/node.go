package dfaUtil

// 节点类型
type node struct {
	// 字符
	ch rune

	// 节点状态枚举
	flag nodeFlag

	// 子节点集合
	children map[rune]*node
}

// 获取子节点
func (this *node) getChild(ch rune) (child *node, exists bool) {
	child, exists = this.children[ch]
	return
}

// 添加子节点
func (this *node) addChild(child *node) {
	this.children[child.ch] = child
}

// 创建新节点
func newNode(_ch rune, _flag nodeFlag) *node {
	return &node{
		ch:       _ch,
		flag:     _flag,
		children: make(map[rune]*node),
	}
}
