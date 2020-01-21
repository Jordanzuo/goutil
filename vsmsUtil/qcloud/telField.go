package qcloud

type telField struct {
	Nationcode string `json:"nationcode"`
	Mobile     string `json:"mobile"`
}

func newTelField(nation, mobile string) *telField {
	return &telField{
		Nationcode: nation,
		Mobile:     mobile,
	}
}
