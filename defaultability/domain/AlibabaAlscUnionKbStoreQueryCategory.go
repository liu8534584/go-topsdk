package domain

type AlibabaAlscUnionKbStoreQueryCategory struct {
	/*
	   类目ID     */
	CategoryId *string `json:"category_id,omitempty" `

	/*
	   类名名称     */
	Name *string `json:"name,omitempty" `
}

func (s *AlibabaAlscUnionKbStoreQueryCategory) SetCategoryId(v string) *AlibabaAlscUnionKbStoreQueryCategory {
	s.CategoryId = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryCategory) SetName(v string) *AlibabaAlscUnionKbStoreQueryCategory {
	s.Name = &v
	return s
}
