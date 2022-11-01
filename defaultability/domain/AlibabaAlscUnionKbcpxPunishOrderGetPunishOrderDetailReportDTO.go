package domain

type AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO struct {
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
	   订单结算时间     */
	SettleTime *string `json:"settle_time,omitempty" `

	/*
	   维权创建时间     */
	ExplainStartTime *string `json:"explain_start_time,omitempty" `

	/*
	   维权结束时间     */
	ExplainEndTime *string `json:"explain_end_time,omitempty" `

	/*
	   订单流水号     */
	SerialNo *string `json:"serial_no,omitempty" `

	/*
	   主商品ID，针对CPS订单     */
	MainItemId *string `json:"main_item_id,omitempty" `

	/*
	   主商品信息，针对CPS订单     */
	MainItemTitle *string `json:"main_item_title,omitempty" `

	/*
	   维权状态，0.待申诉 1.待审核 2.审核中 3.申诉成功 4.申诉失败 5.申诉过期     */
	ExplainState *string `json:"explain_state,omitempty" `

	/*
	   维权成功后的返佣状态，0-待返回 1-已返回     */
	ReturnCommissionState *int64 `json:"return_commission_state,omitempty" `

	/*
	   渠道应结算金额     */
	Settle *string `json:"settle,omitempty" `

	/*
	   更新时间     */
	GmtModified *string `json:"gmt_modified,omitempty" `

	/*
	   会员标识     */
	Sid *string `json:"sid,omitempty" `

	/*
	   1口碑，2饿了么     */
	PlatformType *int64 `json:"platform_type,omitempty" `

	/*
	   活动ID     */
	ActivityId *int64 `json:"activity_id,omitempty" `
}

func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetTitle(v string) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetPicUrl(v string) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.PicUrl = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetShopName(v string) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.ShopName = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetPayAmount(v string) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.PayAmount = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetSettleTime(v string) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.SettleTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetExplainStartTime(v string) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.ExplainStartTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetExplainEndTime(v string) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.ExplainEndTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetSerialNo(v string) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.SerialNo = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetMainItemId(v string) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.MainItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetMainItemTitle(v string) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.MainItemTitle = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetExplainState(v string) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.ExplainState = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetReturnCommissionState(v int64) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.ReturnCommissionState = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetSettle(v string) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.Settle = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetGmtModified(v string) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.GmtModified = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetSid(v string) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.Sid = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetPlatformType(v int64) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.PlatformType = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO) SetActivityId(v int64) *AlibabaAlscUnionKbcpxPunishOrderGetPunishOrderDetailReportDTO {
	s.ActivityId = &v
	return s
}
