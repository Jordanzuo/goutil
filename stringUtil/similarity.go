/*
	使用编辑距离Edit Distance的方式来计算两个字符串的相似类。
	参考资料：https://en.wikipedia.org/wiki/Edit_distance
*/
package stringUtil

// 计算两个字符串的相似度
// word1: 第一个字符串
// word2: 第二个字符串
// 返回值:
// 两个字符串的距离，表示两个字符串经过多少次变换，可以变成同一个字符串
// 两个字符串的相似度，范围为[0, 1]
func Similarity(word1, word2 string) (distance int, similarity float64) {
	// 内部方法，用于计算最小值、最大值
	min := func(nums ...int) int {
		_min := nums[0]
		for _, v := range nums {
			if v < _min {
				_min = v
			}
		}
		return _min
	}
	max := func(nums ...int) int {
		_max := nums[0]
		for _, v := range nums {
			if v > _max {
				_max = v
			}
		}
		return _max
	}

	// 如果有单词为空，或者单词相同，则直接计算出结果，无需进一步计算
	m, n := len(word1), len(word2)
	if m == 0 {
		distance = n
		return
	}
	if n == 0 {
		distance = m
		return
	}
	if m == n && word1 == word2 {
		distance = 0
		similarity = 1
		return
	}

	// 使用动态规划计算编辑距离(Edit Distance)
	// Step 1: define the data structure
	dp := make([][]int, m+1, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1, n+1)
	}

	// Step 2: init the data
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 {
				dp[i][j] = j
			}
			if j == 0 {
				dp[i][j] = i
			}
		}
	}

	// Step 3: state transition equation
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	// 得到距离并计算相似度
	distance = dp[m][n]
	maxLength := max(m, n)
	similarity = float64(maxLength-distance) / float64(maxLength)

	return
}
