package domain

type AlibabaAlscUnionKbOrderPayOrderVoucherDto struct {
	/*
	   渠道订单号     */
	OuterOrderId *string `json:"outer_order_id,omitempty" `

	/*
	   淘宝子单号     */
	BizOrderId *string `json:"biz_order_id,omitempty" `

	/*
	   订单状态，创建成功后的状态     */
	OrderStatus *string `json:"order_status,omitempty" `
}

func (s *AlibabaAlscUnionKbOrderPayOrderVoucherDto) SetOuterOrderId(v string) *AlibabaAlscUnionKbOrderPayOrderVoucherDto {
	s.OuterOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderPayOrderVoucherDto) SetBizOrderId(v string) *AlibabaAlscUnionKbOrderPayOrderVoucherDto {
	s.BizOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderPayOrderVoucherDto) SetOrderStatus(v string) *AlibabaAlscUnionKbOrderPayOrderVoucherDto {
	s.OrderStatus = &v
	return s
}
