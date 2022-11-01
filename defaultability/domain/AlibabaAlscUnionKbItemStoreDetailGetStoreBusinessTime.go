package domain

type AlibabaAlscUnionKbItemStoreDetailGetStoreBusinessTime struct {
	/*
	   营业时间文本化信息     */
	TimeTexts *[]string `json:"time_texts,omitempty" `
}

func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreBusinessTime) SetTimeTexts(v []string) *AlibabaAlscUnionKbItemStoreDetailGetStoreBusinessTime {
	s.TimeTexts = &v
	return s
}
