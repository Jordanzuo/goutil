package dfaUtil

type nodeFlag int

const (
	// 普通节点
	con_Normal nodeFlag = iota

	// 短终节点
	con_ShortTerminal

	// 长终节点
	con_LongTerminal
)
