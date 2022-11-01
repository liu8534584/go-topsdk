package domain

type AlibabaAlscUnionKbBbtItemStoreDetailGetCategory struct {
	/*
	   分类ID     */
	CategoryId *string `json:"category_id,omitempty" `

	/*
	   分类名称     */
	Name *string `json:"name,omitempty" `

	/*
	   父分类ID     */
	ParentCategoryId *string `json:"parent_category_id,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetCategory) SetCategoryId(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetCategory {
	s.CategoryId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetCategory) SetName(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetCategory {
	s.Name = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetCategory) SetParentCategoryId(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetCategory {
	s.ParentCategoryId = &v
	return s
}
