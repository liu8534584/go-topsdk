package domain

type AlibabaAlscUnionKbStoreQueryInteger struct {
	/*
	   门店ID     */
	StoreId *string `json:"store_id,omitempty" `

	/*
	   门店名称     */
	Name *string `json:"name,omitempty" `

	/*
	   封面图     */
	Cover *string `json:"cover,omitempty" `

	/*
	   门店所属行业分类     */
	Categories *[]AlibabaAlscUnionKbStoreQueryCategory `json:"categories,omitempty" `

	/*
	   位置信息     */
	Location *AlibabaAlscUnionKbStoreQueryLocation `json:"location,omitempty" `

	/*
	   距离（米）     */
	Distance *int64 `json:"distance,omitempty" `
}

func (s *AlibabaAlscUnionKbStoreQueryInteger) SetStoreId(v string) *AlibabaAlscUnionKbStoreQueryInteger {
	s.StoreId = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryInteger) SetName(v string) *AlibabaAlscUnionKbStoreQueryInteger {
	s.Name = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryInteger) SetCover(v string) *AlibabaAlscUnionKbStoreQueryInteger {
	s.Cover = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryInteger) SetCategories(v []AlibabaAlscUnionKbStoreQueryCategory) *AlibabaAlscUnionKbStoreQueryInteger {
	s.Categories = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryInteger) SetLocation(v AlibabaAlscUnionKbStoreQueryLocation) *AlibabaAlscUnionKbStoreQueryInteger {
	s.Location = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryInteger) SetDistance(v int64) *AlibabaAlscUnionKbStoreQueryInteger {
	s.Distance = &v
	return s
}
