package stringUtil

// 表情符号集合
var emojiData map[rune]rune

func init() {
	emojiData = make(map[rune]rune, 0)

	addEmojiChar()
}

// 添加一个范围的unicode
// start:unicode起始位置
// endlist:unicode结束位置
func addUnicodeRange(start rune, endlist ...rune) {
	if len(endlist) <= 0 {
		// 添加单个的
		emojiData[start] = start
		return
	}

	end := endlist[0]
	if start > end {
		return
	}

	// 添加范围的
	for i := start; i <= end; i++ {
		emojiData[i] = i
	}
}

// 增加emoji表情符号
// 表情字符大全参考：
// https://zh.wikipedia.org/wiki/%E7%B9%AA%E6%96%87%E5%AD%97
// 对应unicode版本号：Unicode 10.0版本
func addEmojiChar() {
	addUnicodeRange(0x00A9)
	addUnicodeRange(0x00AE)

	addUnicodeRange(0x203C)
	addUnicodeRange(0x2049)
	addUnicodeRange(0x2122)
	addUnicodeRange(0x2139)
	addUnicodeRange(0x2194, 0x2199)
	addUnicodeRange(0x21A9, 0x21AA)
	addUnicodeRange(0x231A, 0x231B)
	addUnicodeRange(0x2328)
	addUnicodeRange(0x23CF)
	addUnicodeRange(0x23E9, 0x23F3)
	addUnicodeRange(0x23F8, 0x23FA)
	addUnicodeRange(0x24C2)
	addUnicodeRange(0x25AA, 0x25AB)
	addUnicodeRange(0x25B6)
	addUnicodeRange(0x25C0)
	addUnicodeRange(0x25FB, 0x25FE)

	addUnicodeRange(0x2600, 0x2604)
	addUnicodeRange(0x260E)
	addUnicodeRange(0x2611)
	addUnicodeRange(0x2614, 0x2615)
	addUnicodeRange(0x2618)
	addUnicodeRange(0x261D)
	addUnicodeRange(0x2620)
	addUnicodeRange(0x2622, 0x2623)
	addUnicodeRange(0x2626)
	addUnicodeRange(0x262A)
	addUnicodeRange(0x262E, 0x262F)
	addUnicodeRange(0x2638, 0x263A)
	addUnicodeRange(0x2640)
	addUnicodeRange(0x2642)
	addUnicodeRange(0x2648, 0x2653)
	addUnicodeRange(0x2660)
	addUnicodeRange(0x2663)
	addUnicodeRange(0x2665, 0x2666)
	addUnicodeRange(0x2668)
	addUnicodeRange(0x267B)
	addUnicodeRange(0x267F)
	addUnicodeRange(0x2692, 0x2697)
	addUnicodeRange(0x2699)
	addUnicodeRange(0x269B, 0x269C)

	addUnicodeRange(0x26A0, 0x26A1)
	addUnicodeRange(0x26AA, 0x26AB)
	addUnicodeRange(0x26B0, 0x26B1)
	addUnicodeRange(0x26BD, 0x26BE)
	addUnicodeRange(0x26C4, 0x26C5)
	addUnicodeRange(0x26C8)
	addUnicodeRange(0x26CE, 0x26CF)
	addUnicodeRange(0x26D1)
	addUnicodeRange(0x26D3, 0x26D4)
	addUnicodeRange(0x26E9, 0x26EA)
	addUnicodeRange(0x26F0, 0x26F5)
	addUnicodeRange(0x26F7, 0x26FA)
	addUnicodeRange(0x26FD)

	addUnicodeRange(0x2702)
	addUnicodeRange(0x2705)
	addUnicodeRange(0x2708, 0x270D)
	addUnicodeRange(0x270F)
	addUnicodeRange(0x2712)
	addUnicodeRange(0x2714)
	addUnicodeRange(0x2716)
	addUnicodeRange(0x271D)
	addUnicodeRange(0x2721)
	addUnicodeRange(0x2728)
	addUnicodeRange(0x2733, 0x2734)
	addUnicodeRange(0x2744)
	addUnicodeRange(0x2747)
	addUnicodeRange(0x274C)
	addUnicodeRange(0x274E)
	addUnicodeRange(0x2753, 0x2755)
	addUnicodeRange(0x2757)
	addUnicodeRange(0x2763, 0x2764)
	addUnicodeRange(0x2795, 0x2797)
	addUnicodeRange(0x27A1)
	addUnicodeRange(0x27B0)
	addUnicodeRange(0x27BF)
	addUnicodeRange(0x2934, 0x2935)
	addUnicodeRange(0x2B05, 0x2B07)
	addUnicodeRange(0x2B1B, 0x2B1C)
	addUnicodeRange(0x2B50)
	addUnicodeRange(0x2B55)
	addUnicodeRange(0x3030)
	addUnicodeRange(0x303D)
	addUnicodeRange(0x3297, 0x3299)
	addUnicodeRange(0x3299)
	addUnicodeRange(0x1F004)
	addUnicodeRange(0x1F0CF)
	addUnicodeRange(0x1F170, 0x1F171)
	addUnicodeRange(0x1F17E, 0x1F17F)
	addUnicodeRange(0x1F18E)

	addUnicodeRange(0x1F191, 0x1F19A)
	addUnicodeRange(0x1F201, 0x1F202)
	addUnicodeRange(0x1F21A)
	addUnicodeRange(0x1F22F)
	addUnicodeRange(0x1F232, 0x1F23A)
	addUnicodeRange(0x1F250, 0x1F251)
	addUnicodeRange(0x1F300, 0x1F321)
	addUnicodeRange(0x1F324, 0x1F393)
	addUnicodeRange(0x1F396, 0x1F397)
	addUnicodeRange(0x1F399, 0x1F39B)
	addUnicodeRange(0x1F39E, 0x1F3F0)
	addUnicodeRange(0x1F3F3, 0x1F3F5)
	addUnicodeRange(0x1F3F7, 0x1F53D)

	addUnicodeRange(0x1F549, 0x1F54E)
	addUnicodeRange(0x1F550, 0x1F567)
	addUnicodeRange(0x1F56F, 0x1F570)
	addUnicodeRange(0x1F573, 0x1F57A)
	addUnicodeRange(0x1F587)
	addUnicodeRange(0x1F58A, 0x1F58D)
	addUnicodeRange(0x1F590)
	addUnicodeRange(0x1F595, 0x1F596)
	addUnicodeRange(0x1F5A4, 0x1F5A5)
	addUnicodeRange(0x1F5A8)
	addUnicodeRange(0x1F5B1, 0x1F5B2)
	addUnicodeRange(0x1F5BC)
	addUnicodeRange(0x1F5C2, 0x1F5C4)
	addUnicodeRange(0x1F5D1, 0x1F5D3)
	addUnicodeRange(0x1F5DC, 0x1F5DE)
	addUnicodeRange(0x1F5E1)
	addUnicodeRange(0x1F5E3)
	addUnicodeRange(0x1F5E8)
	addUnicodeRange(0x1F5EF)
	addUnicodeRange(0x1F5F3)
	addUnicodeRange(0x1F5FA, 0x1F6C5)
	addUnicodeRange(0x1F6CB, 0x1F6D2)
	addUnicodeRange(0x1F6E0, 0x1F6E5)
	addUnicodeRange(0x1F6E8)
	addUnicodeRange(0x1F6EB, 0x1F6EC)
	addUnicodeRange(0x1F6F0)
	addUnicodeRange(0x1F6F3, 0x1F6F8)
	addUnicodeRange(0x1F910, 0x1F93A)
	addUnicodeRange(0x1F93B, 0x1F93E)
	addUnicodeRange(0x1F940, 0x1F945)
	addUnicodeRange(0x1F947, 0x1F94C)
	addUnicodeRange(0x1F950, 0x1F96B)
	addUnicodeRange(0x1F980, 0x1F997)
	addUnicodeRange(0x1F9C0)
	addUnicodeRange(0x1F9D0, 0x1F9E6)
}

// 检查是否含有表情字符
// val:待查看的字符串
// 返回值:
// 是否包含有表情字符
func IfHaveEmoji(val string) bool {
	// 由于golang在内存中本来就是使用的Unicode，所以可以直接进行匹配操作
	for _, charItem := range val {
		if _, eixst := emojiData[charItem]; eixst {
			return true
		}
	}

	return false
}
