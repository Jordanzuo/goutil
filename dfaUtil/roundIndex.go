package dfaUtil

// 每轮的索引对象
type roundIndex struct {
	// 索引值
	index int

	// 节点状态枚举
	flag nodeFlag
}

func newRoundIndex(_index int, _flag nodeFlag) *roundIndex {
	return &roundIndex{
		index: _index,
		flag:  _flag,
	}
}
