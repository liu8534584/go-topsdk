package domain

type AlibabaAlscUnionKbBbtItemStoreDetailGetBbtShopDetailDto struct {
	/*
	   门店基本信息     */
	BasicInfo *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo `json:"basic_info,omitempty" `

	/*
	   门店统计信息     */
	Statistics *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreStatistics `json:"statistics,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetBbtShopDetailDto) SetBasicInfo(v AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo) *AlibabaAlscUnionKbBbtItemStoreDetailGetBbtShopDetailDto {
	s.BasicInfo = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetBbtShopDetailDto) SetStatistics(v AlibabaAlscUnionKbBbtItemStoreDetailGetStoreStatistics) *AlibabaAlscUnionKbBbtItemStoreDetailGetBbtShopDetailDto {
	s.Statistics = &v
	return s
}
