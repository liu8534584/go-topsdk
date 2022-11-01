package domain

type AlibabaAlscUnionKbOrderRefundOrderRefundDto struct {
	/*
	   用户退款原因，必填     */
	Reason *string `json:"reason,omitempty" `

	/*
	   本地生活订单号，，必填     */
	BizOrderId *string `json:"biz_order_id,omitempty" `

	/*
	   退款明细     */
	VoucherList *[]AlibabaAlscUnionKbOrderRefundOrderVoucherDetailDto `json:"voucher_list,omitempty" `

	/*
	   扩展参数，json格式     */
	ExtInfo *string `json:"ext_info,omitempty" `
}

func (s *AlibabaAlscUnionKbOrderRefundOrderRefundDto) SetReason(v string) *AlibabaAlscUnionKbOrderRefundOrderRefundDto {
	s.Reason = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderRefundOrderRefundDto) SetBizOrderId(v string) *AlibabaAlscUnionKbOrderRefundOrderRefundDto {
	s.BizOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderRefundOrderRefundDto) SetVoucherList(v []AlibabaAlscUnionKbOrderRefundOrderVoucherDetailDto) *AlibabaAlscUnionKbOrderRefundOrderRefundDto {
	s.VoucherList = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderRefundOrderRefundDto) SetExtInfo(v string) *AlibabaAlscUnionKbOrderRefundOrderRefundDto {
	s.ExtInfo = &v
	return s
}
