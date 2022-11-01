package domain

type AlibabaAlscUnionKbOrderRefundOrderVoucherDto struct {
	/*
	   淘宝子单号     */
	BizOrderId *int64 `json:"biz_order_id,omitempty" `

	/*
	   订单状态     */
	OrderStatus *string `json:"order_status,omitempty" `
}

func (s *AlibabaAlscUnionKbOrderRefundOrderVoucherDto) SetBizOrderId(v int64) *AlibabaAlscUnionKbOrderRefundOrderVoucherDto {
	s.BizOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderRefundOrderVoucherDto) SetOrderStatus(v string) *AlibabaAlscUnionKbOrderRefundOrderVoucherDto {
	s.OrderStatus = &v
	return s
}
