package domain

type AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO struct {
	/*
	   商品标题     */
	Title *string `json:"title,omitempty" `

	/*
	   商品图片url     */
	PicUrl *string `json:"pic_url,omitempty" `

	/*
	   店铺名称     */
	ShopName *string `json:"shop_name,omitempty" `

	/*
	   付款金额     */
	PayAmount *string `json:"pay_amount,omitempty" `

	/*
	   结算金额     */
	SettleAmount *string `json:"settle_amount,omitempty" `

	/*
	   订单结算时间     */
	SettleTime *string `json:"settle_time,omitempty" `

	/*
	   维权创建时间     */
	ExplainStartTime *string `json:"explain_start_time,omitempty" `

	/*
	   维权结束时间     */
	ExplainEndTime *string `json:"explain_end_time,omitempty" `

	/*
	   维权状态，0-维权成功 1-维权创建 2-维权关闭 3-维权失败     */
	ExplainState *int64 `json:"explain_state,omitempty" `

	/*
	   渠道应结算金额     */
	Settle *string `json:"settle,omitempty" `

	/*
	   申诉之后的佣金返回状态 1-已结算 2-未结算     */
	SettleState *int64 `json:"settle_state,omitempty" `

	/*
	   淘宝子单号     */
	BizOrderId *string `json:"biz_order_id,omitempty" `

	/*
	   父订单号     */
	ParentOrderId *string `json:"parent_order_id,omitempty" `

	/*
	   更新时间     */
	GmtModified *string `json:"gmt_modified,omitempty" `
}

func (s *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO) SetTitle(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO) SetPicUrl(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO {
	s.PicUrl = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO) SetShopName(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO {
	s.ShopName = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO) SetPayAmount(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO {
	s.PayAmount = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO) SetSettleAmount(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO {
	s.SettleAmount = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO) SetSettleTime(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO {
	s.SettleTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO) SetExplainStartTime(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO {
	s.ExplainStartTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO) SetExplainEndTime(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO {
	s.ExplainEndTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO) SetExplainState(v int64) *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO {
	s.ExplainState = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO) SetSettle(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO {
	s.Settle = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO) SetSettleState(v int64) *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO {
	s.SettleState = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO) SetBizOrderId(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO {
	s.BizOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO) SetParentOrderId(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO {
	s.ParentOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO) SetGmtModified(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO {
	s.GmtModified = &v
	return s
}
