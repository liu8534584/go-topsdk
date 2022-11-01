package domain

type AlibabaAlscUnionKbItemDetailGetContentDetail struct {
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
}

func (s *AlibabaAlscUnionKbItemDetailGetContentDetail) SetName(v string) *AlibabaAlscUnionKbItemDetailGetContentDetail {
	s.Name = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetContentDetail) SetPrice(v string) *AlibabaAlscUnionKbItemDetailGetContentDetail {
	s.Price = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetContentDetail) SetQuantity(v int64) *AlibabaAlscUnionKbItemDetailGetContentDetail {
	s.Quantity = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetContentDetail) SetSumPrice(v string) *AlibabaAlscUnionKbItemDetailGetContentDetail {
	s.SumPrice = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetContentDetail) SetUnit(v string) *AlibabaAlscUnionKbItemDetailGetContentDetail {
	s.Unit = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetContentDetail) SetSpec(v string) *AlibabaAlscUnionKbItemDetailGetContentDetail {
	s.Spec = &v
	return s
}
