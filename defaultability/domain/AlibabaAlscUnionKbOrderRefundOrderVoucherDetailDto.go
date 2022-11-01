package domain

type AlibabaAlscUnionKbOrderRefundOrderVoucherDetailDto struct {
	/*
	   商品ID，必填     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   凭证ID，必填     */
	VoucherId *string `json:"voucher_id,omitempty" `
}

func (s *AlibabaAlscUnionKbOrderRefundOrderVoucherDetailDto) SetItemId(v string) *AlibabaAlscUnionKbOrderRefundOrderVoucherDetailDto {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderRefundOrderVoucherDetailDto) SetVoucherId(v string) *AlibabaAlscUnionKbOrderRefundOrderVoucherDetailDto {
	s.VoucherId = &v
	return s
}
