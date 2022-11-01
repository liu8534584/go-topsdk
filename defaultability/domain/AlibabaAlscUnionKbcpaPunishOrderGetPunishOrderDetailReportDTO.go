package domain

type AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO struct {
	/*
	   商品标题     */
	Title *string `json:"title,omitempty" `

	/*
	   商品图片     */
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
	   维权状态，0.待申诉 1.待审核 2.审核中 3.申诉成功 4.申诉失败 5.申诉过期     */
	ExplainState *int64 `json:"explain_state,omitempty" `

	/*
	   渠道应结算金额     */
	Settle *string `json:"settle,omitempty" `

	/*
	   返佣完成状态 0-待返回 1-已返回     */
	ReturnCommissionState *int64 `json:"return_commission_state,omitempty" `

	/*
	   订单流水号     */
	SerialNo *string `json:"serial_no,omitempty" `
}

func (s *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO) SetTitle(v string) *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO) SetPicUrl(v string) *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO {
	s.PicUrl = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO) SetShopName(v string) *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO {
	s.ShopName = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO) SetPayAmount(v string) *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO {
	s.PayAmount = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO) SetSettleAmount(v string) *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO {
	s.SettleAmount = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO) SetSettleTime(v string) *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO {
	s.SettleTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO) SetExplainStartTime(v string) *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO {
	s.ExplainStartTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO) SetExplainEndTime(v string) *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO {
	s.ExplainEndTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO) SetExplainState(v int64) *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO {
	s.ExplainState = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO) SetSettle(v string) *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO {
	s.Settle = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO) SetReturnCommissionState(v int64) *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO {
	s.ReturnCommissionState = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO) SetSerialNo(v string) *AlibabaAlscUnionKbcpaPunishOrderGetPunishOrderDetailReportDTO {
	s.SerialNo = &v
	return s
}
