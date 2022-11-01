package domain

type AlibabaAlscUnionKbOrderCreateOrderVoucherDto struct {
	/*
	   外部订单号     */
	OuterOrderId *string `json:"outer_order_id,omitempty" `

	/*
	   淘宝子单号     */
	BizOrderId *int64 `json:"biz_order_id,omitempty" `

	/*
	   订单状态，创建成功后的状态，固定值     */
	OrderStatus *string `json:"order_status,omitempty" `
}

func (s *AlibabaAlscUnionKbOrderCreateOrderVoucherDto) SetOuterOrderId(v string) *AlibabaAlscUnionKbOrderCreateOrderVoucherDto {
	s.OuterOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderCreateOrderVoucherDto) SetBizOrderId(v int64) *AlibabaAlscUnionKbOrderCreateOrderVoucherDto {
	s.BizOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderCreateOrderVoucherDto) SetOrderStatus(v string) *AlibabaAlscUnionKbOrderCreateOrderVoucherDto {
	s.OrderStatus = &v
	return s
}
