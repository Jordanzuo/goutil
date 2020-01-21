package qcloud

// 电话结构
type telField struct {
	Nationcode string `json:"nationcode"`
	Mobile     string `json:"mobile"`
}

func newTelField(nation, number string) *telField {
	return &telField{
		Nationcode: nation,
		Mobile:     number,
	}
}
