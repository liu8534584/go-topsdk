package domain

type AlibabaAlscUnionKbOrderQueryOrderVoucherDto struct {
	/*
	   淘宝子单号     */
	BizOrderId *string `json:"biz_order_id,omitempty" `

	/*
	   凭证列表     */
	VoucherList *[]AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto `json:"voucher_list,omitempty" `

	/*
	   订单状态。当前只有PAID一个状态     */
	OrderStatus *string `json:"order_status,omitempty" `
}

func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDto) SetBizOrderId(v string) *AlibabaAlscUnionKbOrderQueryOrderVoucherDto {
	s.BizOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDto) SetVoucherList(v []AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto) *AlibabaAlscUnionKbOrderQueryOrderVoucherDto {
	s.VoucherList = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDto) SetOrderStatus(v string) *AlibabaAlscUnionKbOrderQueryOrderVoucherDto {
	s.OrderStatus = &v
	return s
}
