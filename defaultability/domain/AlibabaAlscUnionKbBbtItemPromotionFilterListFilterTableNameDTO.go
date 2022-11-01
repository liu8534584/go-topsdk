package domain

type AlibabaAlscUnionKbBbtItemPromotionFilterListFilterTableNameDTO struct {
	/*
	   筛选项key值     */
	FilterKey *string `json:"filter_key,omitempty" `

	/*
	   筛选项展示名称     */
	FilterName *string `json:"filter_name,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemPromotionFilterListFilterTableNameDTO) SetFilterKey(v string) *AlibabaAlscUnionKbBbtItemPromotionFilterListFilterTableNameDTO {
	s.FilterKey = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemPromotionFilterListFilterTableNameDTO) SetFilterName(v string) *AlibabaAlscUnionKbBbtItemPromotionFilterListFilterTableNameDTO {
	s.FilterName = &v
	return s
}
