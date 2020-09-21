package dfaUtil

/*
DFA util, is used to verify whether a sentence has invalid words.
The underlying data structure is trie.
https://en.wikipedia.org/wiki/Trie
*/

// dfa util
type DFAUtil struct {
	// The root node
	root *trieNode
}

// 由于go不支持tuple，所以为了避免定义多余的struct，特别使用两个list来分别返回匹配的索引的上界和下界
// 在处理此方法的返回值时，需要两者配合使用
func (this *DFAUtil) searcSentence(sentence string) (startIndexList, endIndexList []int) {
	// Point current node to the root node and initialize some variables.
	currNode := this.root
	start, end, valid := 0, 0, false

	// Iterate the setence to handle each letter
	sentenceRuneList := []rune(sentence)
	for i := 0; i < len(sentenceRuneList); {
		// If the letter can be found in current node's children, then continue to find along this path.
		if child, exists := currNode.children[sentenceRuneList[i]]; exists {
			// If the letter is end of a word, then it's a valid match.
			// Then set valid to true, and assign the index to the end variable.
			if child.isEndOfWord {
				end = i
				valid = true
			}

			// If the child doesn't have any child, it means it's the end of a path. Then add the last valid index pair into list.
			// And continue to handle the next letter from the root node.
			// Otherwise, continue to handle along this path.
			if len(child.children) == 0 {
				startIndexList = append(startIndexList, start)
				endIndexList = append(endIndexList, end)
				currNode = this.root

				// Reset variables, and starts from the next letter.
				start, end, valid = i+1, 0, false
			} else {
				currNode = child
			}

			// Handle the next letter.
			i++
		} else {
			// When the letter can't be found in current node's children, there are two possibilities:
			// 1. There is already a valid match index pair. Then add them to list. And rehandle this letter again from the root node.
			// 2. There is no valid match index pair. Then continue to handle next letter from the root node.
			if valid {
				startIndexList = append(startIndexList, start)
				endIndexList = append(endIndexList, end)
				currNode = this.root
				start, end, valid = i, 0, false
			} else {
				currNode = this.root
				start, end, valid = i+1, 0, false
				i++
			}
		}
	}

	// Check if there is any valid pairs which hasn't been processed.
	if valid {
		startIndexList = append(startIndexList, start)
		endIndexList = append(endIndexList, end)
	}

	return
}

// Insert new word into object
func (this *DFAUtil) InsertWord(word []rune) {
	currNode := this.root
	for _, c := range word {
		if cildNode, exist := currNode.children[c]; !exist {
			cildNode = newtrieNode()
			currNode.children[c] = cildNode
			currNode = cildNode
		} else {
			currNode = cildNode
		}
	}

	currNode.isEndOfWord = true
}

// Check if there is any word in the trie that starts with the given prefix.
func (this *DFAUtil) StartsWith(prefix []rune) bool {
	currNode := this.root
	for _, c := range prefix {
		if cildNode, exist := currNode.children[c]; !exist {
			return false
		} else {
			currNode = cildNode
		}
	}

	return true
}

// Judge if input sentence contains some special caracter
// Return:
// Matc or not
func (this *DFAUtil) IsMatch(sentence string) bool {
	startIndexList, _ := this.searcSentence(sentence)
	return len(startIndexList) > 0
}

// Handle sentence. Use specified caracter to replace those sensitive caracters.
// input: Input sentence
// replaceCh: candidate
// Return:
// Sentence after manipulation
func (this *DFAUtil) HandleWord(sentence string, replaceCh rune) string {
	startIndexList, endIndexList := this.searcSentence(sentence)
	if len(startIndexList) == 0 {
		return sentence
	}

	// Manipulate
	sentenceList := []rune(sentence)
	for i := 0; i < len(startIndexList); i++ {
		for index := startIndexList[i]; index <= endIndexList[i]; index++ {
			sentenceList[index] = replaceCh
		}
	}

	return string(sentenceList)
}

// Create new DfaUtil object
// wordList:word list
func NewDFAUtil(wordList []string) *DFAUtil {
	this := &DFAUtil{
		root: newtrieNode(),
	}

	for _, word := range wordList {
		wordRuneList := []rune(word)
		if len(wordRuneList) > 0 {
			this.InsertWord(wordRuneList)
		}
	}

	return this
}
