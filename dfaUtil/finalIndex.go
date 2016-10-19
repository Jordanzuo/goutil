package dfaUtil

// 最终的索引对象
type finalIndex struct {
	// 索引下限
	lowerIndex int

	// 索引上限
	upperIndex int
}

func newFinalIndex(_lowerIndex, _upperIndex int) *finalIndex {
	return &finalIndex{
		lowerIndex: _lowerIndex,
		upperIndex: _upperIndex,
	}
}
