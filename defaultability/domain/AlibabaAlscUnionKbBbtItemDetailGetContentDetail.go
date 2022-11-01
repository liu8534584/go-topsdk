package domain

type AlibabaAlscUnionKbBbtItemDetailGetContentDetail struct {
	/*
	   名称     */
	Name *string `json:"name,omitempty" `

	/*
	   单价（元）     */
	Price *string `json:"price,omitempty" `

	/*
	   数量     */
	Quantity *int64 `json:"quantity,omitempty" `

	/*
	   小计金额=数量*单价     */
	SumPrice *string `json:"sum_price,omitempty" `

	/*
	   单位     */
	Unit *string `json:"unit,omitempty" `

	/*
	   规格     */
	Spec *string `json:"spec,omitempty" `

	/*
	   文本内容     */
	TextContents *[]AlibabaAlscUnionKbBbtItemDetailGetTextContent `json:"text_contents,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetContentDetail) SetName(v string) *AlibabaAlscUnionKbBbtItemDetailGetContentDetail {
	s.Name = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetContentDetail) SetPrice(v string) *AlibabaAlscUnionKbBbtItemDetailGetContentDetail {
	s.Price = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetContentDetail) SetQuantity(v int64) *AlibabaAlscUnionKbBbtItemDetailGetContentDetail {
	s.Quantity = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetContentDetail) SetSumPrice(v string) *AlibabaAlscUnionKbBbtItemDetailGetContentDetail {
	s.SumPrice = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetContentDetail) SetUnit(v string) *AlibabaAlscUnionKbBbtItemDetailGetContentDetail {
	s.Unit = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetContentDetail) SetSpec(v string) *AlibabaAlscUnionKbBbtItemDetailGetContentDetail {
	s.Spec = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetContentDetail) SetTextContents(v []AlibabaAlscUnionKbBbtItemDetailGetTextContent) *AlibabaAlscUnionKbBbtItemDetailGetContentDetail {
	s.TextContents = &v
	return s
}
