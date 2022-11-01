package domain

type AlibabaAlscUnionKbItemStoreDetailGetCategory struct {
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

func (s *AlibabaAlscUnionKbItemStoreDetailGetCategory) SetCategoryId(v string) *AlibabaAlscUnionKbItemStoreDetailGetCategory {
	s.CategoryId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetCategory) SetName(v string) *AlibabaAlscUnionKbItemStoreDetailGetCategory {
	s.Name = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetCategory) SetParentCategoryId(v string) *AlibabaAlscUnionKbItemStoreDetailGetCategory {
	s.ParentCategoryId = &v
	return s
}
