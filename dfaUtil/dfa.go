package dfaUtil

// dfa助手对象
type DFAUtil struct {
	// 根节点
	root *node
}

// 插入子节点
// currNode:当前节点
// chArr:字符集合
// index:字符位于字符集合的索引
func (this *DFAUtil) insertNode(currNode *node, chArr []rune, index int) {
	// 判断字符是否已经存在于当前节点的子节点中
	childNode, exists := currNode.getChild(chArr[index])

	// 如果存在，且index已经是chArr最后一个字符，则说明子节点应该是一个短节点；如果该子节点的Flag=Normal，则将其设置为ShortTerminal
	if exists && index == len(chArr)-1 && childNode.flag == con_Normal {
		childNode.flag = con_ShortTerminal
	}

	// 如果不存在此节点，则创建节点，并添加到当前节点的子节点列表中
	//  1、判断是否为最后一个字符，如果是则设置为长终节点；
	//  2、同时需要判断当前节点是否为长终节点，如果是则需要将其置为短终节点（因为具有子节点的一定不能是长终节点，而只能是短终节点）
	if !exists {
		childNode = newNode(chArr[index], con_Normal)

		// 注释1
		if index == len(chArr)-1 {
			childNode.flag = con_LongTerminal
		}

		// 注释2
		if currNode.flag == con_LongTerminal {
			currNode.flag = con_ShortTerminal
		}

		//添加子节点
		currNode.addChild(childNode)
	}

	// 自增并向后继续插入
	index++
	if index < len(chArr) {
		// 使用的是尾递归，以节约内存和效率
		this.insertNode(childNode, chArr, index)
	}
}

// 搜索特殊字符
// input:输入数据
// 返回值:
// 特殊字符所处的索引区间列表
func (this *DFAUtil) SearchWord(input string) []*finalIndex {
	// 最终的索引范围列表<索引下限、索引上限>
	finalIndexList := make([]*finalIndex, 0, 32)

	// 每一轮是屏蔽词的<索引、Flag>列表
	roundIndexList := make([]*roundIndex, 0, 8)

	// 将输入字符串转换为字符集合
	chArr := []rune(input)

	// 定义寻找短路径的方法
	searchShortPath := func(_roundIndexList []*roundIndex) *roundIndex {
		var lastItem *roundIndex = nil
		for _, item := range _roundIndexList {
			if item.flag == con_ShortTerminal {
				lastItem = item
			}
		}

		return lastItem
	}

	//从根节点开始遍历
	nodeObj := this.root
	exists := false
	for index := 0; index < len(chArr); index++ {
		nodeObj, exists = nodeObj.getChild(chArr[index])
		if !exists {
			// 如果没有找到匹配的节点，说明此轮判断结束
			// 需要判断是否有符合条件的短搜索路径
			if len(roundIndexList) > 0 {
				//查找最后一个NodeFlag==1的数据项
				//1、如果找到了，则需要当作有效路径来处理
				//2、如果没有找到，则丢弃，并向后倒退roundIndexList.Count个位置，以避免当前位置的字符被误判的情况
				//例如有两个关键字叠加“习近平的”，“近平”
				//相当于每次只判断一个字符
				lastItem := searchShortPath(roundIndexList)
				if lastItem != nil {
					//将本轮的索引列表添加到最终的索引列表中
					finalIndexList = append(finalIndexList, newFinalIndex(roundIndexList[0].index, lastItem.index))
				} else {
					index = index - len(roundIndexList)
				}
			}

			// 清空本轮的索引列表
			roundIndexList = roundIndexList[:0]

			// 重新从根节点开始计算
			nodeObj = this.root
		} else if nodeObj.flag == con_LongTerminal {
			// 如果是长终节点，说明已经找到最长路径了，于是本轮结束
			// 将对应的index添加到本轮的索引列表中
			roundIndexList = append(roundIndexList, newRoundIndex(index, nodeObj.flag))

			// 将本轮的索引列表添加到最终的索引列表中
			finalIndexList = append(finalIndexList, newFinalIndex(roundIndexList[0].index, roundIndexList[len(roundIndexList)-1].index))

			// 清空本轮的索引列表
			roundIndexList = roundIndexList[:0]

			// 重新从根节点开始计算
			nodeObj = this.root
		} else {
			// 如果是非终节点，则添加到本轮的列表中，然后继续往后搜索
			// 如果是短终节点，则添加到本轮的列表中，然后继续往后搜索，以便于查找是否还有更长的搜索路径
			// 将对应的index添加到本轮的索引列表中
			roundIndexList = append(roundIndexList, newRoundIndex(index, nodeObj.flag))
		}

		// 最后一个字处理完成后，需要再进行一次判断，以避免当前位置的字符被误判的情况
		// 例如有两个关键字叠加“习近平的”，“近平”
		// 相当于每次只判断一个字符
		if index == len(chArr)-1 {
			// 判断是否存在没有完全匹配的数据
			if len(roundIndexList) > 0 {
				lastItem := searchShortPath(roundIndexList)
				if lastItem != nil {
					// 将本轮的索引列表添加到最终的索引列表中
					finalIndexList = append(finalIndexList, newFinalIndex(roundIndexList[0].index, lastItem.index))
					break
				} else {
					index = index - (len(roundIndexList) - 1)
				}
			}

			// 清空本轮的索引列表
			roundIndexList = roundIndexList[:0]

			// 重新从根节点开始计算
			nodeObj = this.root
		}
	}

	return finalIndexList
}

// 判断输入数据是否匹配特殊字符
// input:输入数据
// 返回值:
// 是否匹配
func (this *DFAUtil) IsMatch(input string) bool {
	return len(this.SearchWord(input)) > 0
}

// 处理单词，用指定字符替换特殊字符
// input:输入数据
// replaceCh:替换的字符
// 返回值:
// 处理后的字符串
func (this *DFAUtil) HandleWord(input string, replaceCh rune) string {
	// 最终的索引范围列表<索引下限、索引上限>
	finalIndexList := this.SearchWord(input)

	// 判断是否需要替换
	if len(finalIndexList) == 0 {
		return input
	}

	// 定义最终的索引范围是否包含指定索引的方法
	isContains := func(index int) bool {
		for _, finalIndexObj := range finalIndexList {
			if finalIndexObj.lowerIndex <= index && index <= finalIndexObj.upperIndex {
				return true
			}
		}

		return false
	}

	// 获取输入内容的字符串数组
	chArr := []rune(input)

	// 进行替换
	for index := 0; index < len(chArr); index++ {
		// 判断当前的索引是否位于可被替换的区间内
		if isContains(index) {
			chArr[index] = replaceCh
		}
	}

	return string(chArr)
}

// 创建新的dfaUtil对象
// wordList:词语列表
func NewDFAUtil(wordList []string) *DFAUtil {
	this := &DFAUtil{
		root: newNode('R', con_Normal),
	}

	for _, word := range wordList {
		chArr := []rune(word)
		if len(chArr) > 0 {
			this.insertNode(this.root, chArr, 0)
		}
	}

	return this
}
