package domain

type AlibabaAlscUnionKbBbtItemStoreRelationQueryPageModel struct {
	/*
	   分页记录     */
	Records *[]AlibabaAlscUnionKbBbtItemStoreRelationQueryBbtItemShopRelationDto `json:"records,omitempty" `

	/*
	   总数     */
	Total *int64 `json:"total,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemStoreRelationQueryPageModel) SetRecords(v []AlibabaAlscUnionKbBbtItemStoreRelationQueryBbtItemShopRelationDto) *AlibabaAlscUnionKbBbtItemStoreRelationQueryPageModel {
	s.Records = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreRelationQueryPageModel) SetTotal(v int64) *AlibabaAlscUnionKbBbtItemStoreRelationQueryPageModel {
	s.Total = &v
	return s
}
