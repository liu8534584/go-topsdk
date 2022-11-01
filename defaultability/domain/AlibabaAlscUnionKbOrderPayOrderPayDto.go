package domain

type AlibabaAlscUnionKbOrderPayOrderPayDto struct {
	/*
	   渠道订单号     */
	OuterOrderId *string `json:"outer_order_id,omitempty" `

	/*
	   淘宝子单号     */
	BizOrderId *string `json:"biz_order_id,omitempty" `
}

func (s *AlibabaAlscUnionKbOrderPayOrderPayDto) SetOuterOrderId(v string) *AlibabaAlscUnionKbOrderPayOrderPayDto {
	s.OuterOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderPayOrderPayDto) SetBizOrderId(v string) *AlibabaAlscUnionKbOrderPayOrderPayDto {
	s.BizOrderId = &v
	return s
}
