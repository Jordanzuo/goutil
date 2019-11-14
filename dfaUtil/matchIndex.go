package dfaUtil

// Match index object
type matchIndex struct {
	start int // start index
	end   int // end index
}

// Construct from scratch
func newMatchIndex(start, end int) *matchIndex {
	return &matchIndex{
		start: start,
		end:   end,
	}
}

// Construct from existing match index object
func buildMatchIndex(obj *matchIndex) *matchIndex {
	return &matchIndex{
		start: obj.start,
		end:   obj.end,
	}
}
