package domain

type AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusinessTime struct {
	/*
	   营业时间文本化信息     */
	TimeTexts *[]string `json:"time_texts,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusinessTime) SetTimeTexts(v []string) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusinessTime {
	s.TimeTexts = &v
	return s
}
