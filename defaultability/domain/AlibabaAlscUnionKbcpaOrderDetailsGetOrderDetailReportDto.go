package domain

type AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto struct {
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
	   订单付款时间     */
	PayTime *string `json:"pay_time,omitempty" `

	/*
	   收货时间     */
	ReceiveTime *string `json:"receive_time,omitempty" `

	/*
	   订单结算时间     */
	SettleTime *string `json:"settle_time,omitempty" `

	/*
	   订单状态，0-已失效 1-已下单 2-已付款 4-已收货      */
	OrderState *int64 `json:"order_state,omitempty" `

	/*
	   结算状态，1.已结算 2.未结算     */
	SettleState *int64 `json:"settle_state,omitempty" `

	/*
	   结算预估收入     */
	Settle *string `json:"settle,omitempty" `

	/*
	   技术服务费     */
	Service *string `json:"service,omitempty" `

	/*
	   淘宝子单号     */
	BizOrderId *string `json:"biz_order_id,omitempty" `

	/*
	   淘宝主单号     */
	ParentOrderId *string `json:"parent_order_id,omitempty" `

	/*
	   点击时间     */
	TraceTime *string `json:"trace_time,omitempty" `

	/*
	   商品id     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   商品数量     */
	ProductNum *int64 `json:"product_num,omitempty" `

	/*
	   商品单价     */
	UnitPrice *string `json:"unit_price,omitempty" `

	/*
	   联盟补贴金额     */
	SubsidyFee *string `json:"subsidy_fee,omitempty" `

	/*
	   激励金额（第二阶段分佣金额）     */
	AwardFee *string `json:"award_fee,omitempty" `

	/*
	   技术服务费率     */
	PlatformCommissionRate *string `json:"platform_commission_rate,omitempty" `

	/*
	   技术服务费     */
	PlatformCommissionFee *string `json:"platform_commission_fee,omitempty" `

	/*
	   媒体id     */
	MediaId *int64 `json:"media_id,omitempty" `

	/*
	   媒体名称     */
	MediaName *string `json:"media_name,omitempty" `

	/*
	   推广位id     */
	AdZoneId *int64 `json:"ad_zone_id,omitempty" `

	/*
	   推广位名称     */
	AdZoneName *string `json:"ad_zone_name,omitempty" `

	/*
	   类目名称     */
	CategoryName *string `json:"category_name,omitempty" `

	/*
	   创建时间     */
	TkCreateTime *string `json:"tk_create_time,omitempty" `

	/*
	   预估收入     */
	Income *string `json:"income,omitempty" `

	/*
	   更新时间     */
	GmtModified *string `json:"gmt_modified,omitempty" `
}

func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetTitle(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetPicUrl(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.PicUrl = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetShopName(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.ShopName = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetPayAmount(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.PayAmount = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetSettleAmount(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.SettleAmount = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetPayTime(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.PayTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetReceiveTime(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.ReceiveTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetSettleTime(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.SettleTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetOrderState(v int64) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.OrderState = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetSettleState(v int64) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.SettleState = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetSettle(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.Settle = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetService(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.Service = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetBizOrderId(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.BizOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetParentOrderId(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.ParentOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetTraceTime(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.TraceTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetItemId(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetProductNum(v int64) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.ProductNum = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetUnitPrice(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.UnitPrice = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetSubsidyFee(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.SubsidyFee = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetAwardFee(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.AwardFee = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetPlatformCommissionRate(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.PlatformCommissionRate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetPlatformCommissionFee(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.PlatformCommissionFee = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetMediaId(v int64) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.MediaId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetMediaName(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.MediaName = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetAdZoneId(v int64) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.AdZoneId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetAdZoneName(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.AdZoneName = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetCategoryName(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.CategoryName = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetTkCreateTime(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.TkCreateTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetIncome(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.Income = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto) SetGmtModified(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto {
	s.GmtModified = &v
	return s
}
