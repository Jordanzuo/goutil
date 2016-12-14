package dfaUtil

// 节点类型
type node struct {
	// 字符
	ch rune

	// 节点状态枚举
	flag nodeFlag

	// 节点列表
	nodeList []*node
}

// 创建新节点
func newNode(_ch rune, _flag nodeFlag) *node {
	return &node{
		ch:   _ch,
		flag: _flag,
	}
}
