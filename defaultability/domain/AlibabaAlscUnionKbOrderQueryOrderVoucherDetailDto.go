package domain

import (
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto struct {
	/*
	   商品ID     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   商品名称     */
	Title *string `json:"title,omitempty" `

	/*
	   凭证ID     */
	VoucherId *string `json:"voucher_id,omitempty" `

	/*
	   核销码。使用时需要生成二维码，并提示“请商家用阿里本地通（原口碑掌柜）核销”，否则用户核销时会被商家拒绝     */
	TicketCode *string `json:"ticket_code,omitempty" `

	/*
	   凭证状态。可用EFFECTIVE、已用USED、失效CANCELED     */
	VoucherStatus *string `json:"voucher_status,omitempty" `

	/*
	   总份数     */
	TotalAmount *int64 `json:"total_amount,omitempty" `

	/*
	   已使用份数     */
	UsedAmount *int64 `json:"used_amount,omitempty" `

	/*
	   已退款份数（售中、售后）     */
	RefundAmount *int64 `json:"refund_amount,omitempty" `

	/*
	   凭证生效时间     */
	EffectTime *util.LocalTime `json:"effect_time,omitempty" `

	/*
	   凭证过期时间     */
	ExpireTime *util.LocalTime `json:"expire_time,omitempty" `

	/*
	   扩展字段     */
	ExtInfo *string `json:"ext_info,omitempty" `

	/*
	   凭证退款类型。售中退BEFORE_USE_REFUND、售后退AFTER_USE_REFUND、过期退EXPIRATION_REFUND、冲正REVERSE     */
	RefundType *string `json:"refund_type,omitempty" `

	/*
	   http://xxx.com     */
	TicketUrl *string `json:"ticket_url,omitempty" `
}

func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto) SetItemId(v string) *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto) SetTitle(v string) *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto) SetVoucherId(v string) *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto {
	s.VoucherId = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto) SetTicketCode(v string) *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto {
	s.TicketCode = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto) SetVoucherStatus(v string) *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto {
	s.VoucherStatus = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto) SetTotalAmount(v int64) *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto {
	s.TotalAmount = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto) SetUsedAmount(v int64) *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto {
	s.UsedAmount = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto) SetRefundAmount(v int64) *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto {
	s.RefundAmount = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto) SetEffectTime(v util.LocalTime) *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto {
	s.EffectTime = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto) SetExpireTime(v util.LocalTime) *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto {
	s.ExpireTime = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto) SetExtInfo(v string) *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto {
	s.ExtInfo = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto) SetRefundType(v string) *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto {
	s.RefundType = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto) SetTicketUrl(v string) *AlibabaAlscUnionKbOrderQueryOrderVoucherDetailDto {
	s.TicketUrl = &v
	return s
}
