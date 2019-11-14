package dfaUtil

/*
Use dfa algorithm to verify whether a sentence has invalid words.
The underlying data structure is trie.
https://en.wikipedia.org/wiki/Trie
*/

// dfa助手对象
type DFAUtil struct {
	// 根节点
	root *TrieNode
}

func (this *DFAUtil) insert(word []rune) {
	currNode := this.root
	for _, ch := range word {
		if childNode, exist := currNode.children[ch]; !exist {
			childNode = newTrieNode()
			currNode.children[ch] = childNode
			currNode = childNode
		} else {
			currNode = childNode
		}
	}

	currNode.isEndOfWord = true
}

func (this *DFAUtil) search(word []rune) bool {
	currNode := this.root
	for _, ch := range word {
		if childNode, exist := currNode.children[ch]; !exist {
			return false
		} else {
			currNode = childNode
		}
	}

	return currNode.isEndOfWord
}

// 搜索特殊字符
// input:输入数据
// 返回值:
// 特殊字符所处的索引区间列表
func (this *DFAUtil) SearchWord(sentence string) (matchIndexList []*matchIndex) {
	start, end := 0, 1
	sentenceRuneList := []rune(sentence)
	
	// A word or phrase can be matched for more than one path, so we need to get the longest one
	// We iterate the sentence from the beginning to the end.
	// If a word or phrase is matched, we still need to check whether a longer one will match or not.
	var currMatchIndex *matchIndex
	for end <= len(sentenceRuneList) {
		if this.search(sentenceRuneList[start:end+1]) {
			// When matching, we create a new MatchIndex object or extend the current end index
			if currMatchIndex == nil {
				currMatchIndex = newMatchIndex(start, end)
			} else {
				currMatchIndex.end = end
			}
		} else {
			// When not matching, we add the previous matched object into the match list
			if currMatchIndex != nil {
				matchIndexList = append(matchIndexList, buildMatchIndex(currMatchIndex))
				currMatchIndex = nil
			}

			// Move the start to the beginning of next word or phrase
			start = end
		}
		end += 1
	}

	return
}

// 判断输入数据是否匹配特殊字符
// input:输入数据
// 返回值:
// 是否匹配
func (this *DFAUtil) IsMatch(sentence string) bool {
	return len(this.SearchWord(sentence)) > 0
}

// 处理单词，用指定字符替换特殊字符
// input:输入数据
// replaceCh:替换的字符
// 返回值:
// 处理后的字符串
func (this *DFAUtil) HandleWord(sentence string, replaceCh rune) string {
	// 最终的索引范围列表<索引下限、索引上限>
	matchIndexList := this.SearchWord(sentence)

	// 判断是否需要替换
	if len(matchIndexList) == 0 {
		return sentence
	}

	// 获取输入内容的字符串数组
	sentenceList := []rune(sentence)
	for _, matchIndexObj := range matchIndexList {
		for index := matchIndexObj.start; index <= matchIndexObj.end; index ++ {
			sentenceList[index] = replaceCh
		}
	}

	return string(sentenceList)
}

// 创建新的dfaUtil对象
// wordList:词语列表
func NewDFAUtil(wordList []string) *DFAUtil {
	this := &DFAUtil{
		root: newTrieNode(),
	}

	for _, word := range wordList {
		wordRuneList := []rune(word)
		if len(wordRuneList) > 0 {
			this.insert(wordRuneList)
		}
	}

	return this
}
