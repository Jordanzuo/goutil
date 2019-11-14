package dfaUtil

// 匹配的索引对象
type matchIndex struct {
	start int
	end int
}

func newMatchIndex(start, end int) *matchIndex {
	return &matchIndex{
		start: start,
		end: end,
	}
}

func buildMatchIndex(obj *matchIndex) *matchIndex {
	return &matchIndex {
		start: obj.start,
		end: obj.end,
	}
}