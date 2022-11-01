package domain

type AlibabaAlscUnionKbItemStoreDetailGetKbShopDetailDto struct {
	/*
	   门店基本信息     */
	BasicInfo *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo `json:"basic_info,omitempty" `

	/*
	   门店统计信息     */
	Statistics *AlibabaAlscUnionKbItemStoreDetailGetStoreStatistics `json:"statistics,omitempty" `
}

func (s *AlibabaAlscUnionKbItemStoreDetailGetKbShopDetailDto) SetBasicInfo(v AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo) *AlibabaAlscUnionKbItemStoreDetailGetKbShopDetailDto {
	s.BasicInfo = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetKbShopDetailDto) SetStatistics(v AlibabaAlscUnionKbItemStoreDetailGetStoreStatistics) *AlibabaAlscUnionKbItemStoreDetailGetKbShopDetailDto {
	s.Statistics = &v
	return s
}
