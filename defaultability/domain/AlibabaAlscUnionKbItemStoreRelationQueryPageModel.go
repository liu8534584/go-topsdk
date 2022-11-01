package domain

type AlibabaAlscUnionKbItemStoreRelationQueryPageModel struct {
	/*
	   分页记录     */
	Records *[]AlibabaAlscUnionKbItemStoreRelationQueryKbItemShopRelationDto `json:"records,omitempty" `

	/*
	   总数     */
	Total *int64 `json:"total,omitempty" `
}

func (s *AlibabaAlscUnionKbItemStoreRelationQueryPageModel) SetRecords(v []AlibabaAlscUnionKbItemStoreRelationQueryKbItemShopRelationDto) *AlibabaAlscUnionKbItemStoreRelationQueryPageModel {
	s.Records = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreRelationQueryPageModel) SetTotal(v int64) *AlibabaAlscUnionKbItemStoreRelationQueryPageModel {
	s.Total = &v
	return s
}
