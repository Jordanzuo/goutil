package dfaUtil

/*
DFA util, is used to verify whether a sentence has invalid words.
The underlying data structure is trie.
https://en.wikipedia.org/wiki/Trie
*/

// dfa util
type DFAUtil struct {
	// The root node
	root *TrieNode
}

func (this *DFAUtil) insertWord(word []rune) {
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

// Search and make sure if a word is existed in the underlying trie.
func (this *DFAUtil) searchWord(word []rune) bool {
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

// Search a whole sentence and get all the matching words and their indices
// Return:
// A list of all the match index object
func (this *DFAUtil) searchSentence(sentence string) (matchIndexList []*matchIndex) {
	start, end := 0, 1
	sentenceRuneList := []rune(sentence)
	
	// A word or phrase can be matched for more than one path, so we need to get the longest one
	// We iterate the sentence from the beginning to the end.
	// If a word or phrase is matched, we still need to check whether a longer one will match or not.
	var currMatchIndex *matchIndex
	for end <= len(sentenceRuneList) {
		if this.searchWord(sentenceRuneList[start:end+1]) {
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

// Judge if input sentence contains some special character
// Return:
// Match or not
func (this *DFAUtil) IsMatch(sentence string) bool {
	return len(this.searchSentence(sentence)) > 0
}

// Handle sentence. Use specified character to replace those sensitive characters.
// input: Input sentence
// replaceCh: candidate
// Return:
// Sentence after manipulation
func (this *DFAUtil) HandleWord(sentence string, replaceCh rune) string {
	matchIndexList := this.searchSentence(sentence)
	if len(matchIndexList) == 0 {
		return sentence
	}

	// Manipulate
	sentenceList := []rune(sentence)
	for _, matchIndexObj := range matchIndexList {
		for index := matchIndexObj.start; index <= matchIndexObj.end; index ++ {
			sentenceList[index] = replaceCh
		}
	}

	return string(sentenceList)
}

// Create new DfaUtil object
// wordList:word list
func NewDFAUtil(wordList []string) *DFAUtil {
	this := &DFAUtil{
		root: newTrieNode(),
	}

	for _, word := range wordList {
		wordRuneList := []rune(word)
		if len(wordRuneList) > 0 {
			this.insertWord(wordRuneList)
		}
	}

	return this
}
