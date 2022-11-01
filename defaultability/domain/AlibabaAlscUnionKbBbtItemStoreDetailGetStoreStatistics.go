package domain

type AlibabaAlscUnionKbBbtItemStoreDetailGetStoreStatistics struct {
	/*
	   服务评级     */
	ServiceRating *string `json:"service_rating,omitempty" `

	/*
	   人均价格     */
	AvgPrice *string `json:"avg_price,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreStatistics) SetServiceRating(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreStatistics {
	s.ServiceRating = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreStatistics) SetAvgPrice(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreStatistics {
	s.AvgPrice = &v
	return s
}
