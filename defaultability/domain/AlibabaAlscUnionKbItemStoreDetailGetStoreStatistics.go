package domain

type AlibabaAlscUnionKbItemStoreDetailGetStoreStatistics struct {
	/*
	   服务评级     */
	ServiceRating *string `json:"service_rating,omitempty" `

	/*
	   人均价格     */
	AvgPrice *string `json:"avg_price,omitempty" `
}

func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreStatistics) SetServiceRating(v string) *AlibabaAlscUnionKbItemStoreDetailGetStoreStatistics {
	s.ServiceRating = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreStatistics) SetAvgPrice(v string) *AlibabaAlscUnionKbItemStoreDetailGetStoreStatistics {
	s.AvgPrice = &v
	return s
}
