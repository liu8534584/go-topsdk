package domain

type AlibabaAlscUnionKbStoreItemQueryPageModel struct {
	/*
	   总数     */
	Total *int64 `json:"total,omitempty" `

	/*
	   分页记录     */
	Records *[]AlibabaAlscUnionKbStoreItemQueryKbShopItemDto `json:"records,omitempty" `
}

func (s *AlibabaAlscUnionKbStoreItemQueryPageModel) SetTotal(v int64) *AlibabaAlscUnionKbStoreItemQueryPageModel {
	s.Total = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreItemQueryPageModel) SetRecords(v []AlibabaAlscUnionKbStoreItemQueryKbShopItemDto) *AlibabaAlscUnionKbStoreItemQueryPageModel {
	s.Records = &v
	return s
}
