package domain

type AlibabaAlscUnionKbItemPromotionFilterListFilterTableNameDTO struct {
	/*
	   筛选项key值     */
	FilterKey *string `json:"filter_key,omitempty" `

	/*
	   筛选项展示名称     */
	FilterName *string `json:"filter_name,omitempty" `
}

func (s *AlibabaAlscUnionKbItemPromotionFilterListFilterTableNameDTO) SetFilterKey(v string) *AlibabaAlscUnionKbItemPromotionFilterListFilterTableNameDTO {
	s.FilterKey = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionFilterListFilterTableNameDTO) SetFilterName(v string) *AlibabaAlscUnionKbItemPromotionFilterListFilterTableNameDTO {
	s.FilterName = &v
	return s
}
