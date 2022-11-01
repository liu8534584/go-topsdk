package domain

type AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO struct {
	/*
	   商品标题     */
	Title *string `json:"title,omitempty" `

	/*
	   商品图片下载url     */
	PicUrl *string `json:"pic_url,omitempty" `

	/*
	   店铺名称     */
	ShopName *string `json:"shop_name,omitempty" `

	/*
	   付款金额     */
	PayAmount *string `json:"pay_amount,omitempty" `

	/*
	   结算金额，针对CPS订单     */
	SettleAmount *string `json:"settle_amount,omitempty" `

	/*
	   点击时间     */
	TraceTime *string `json:"trace_time,omitempty" `

	/*
	   创建时间     */
	TkCreateTime *string `json:"tk_create_time,omitempty" `

	/*
	   付款时间     */
	PayTime *string `json:"pay_time,omitempty" `

	/*
	   收货时间     */
	ReceiveTime *string `json:"receive_time,omitempty" `

	/*
	   结算时间     */
	SettleTime *string `json:"settle_time,omitempty" `

	/*
	   预估收入     */
	Income *string `json:"income,omitempty" `

	/*
	   结算预估收入     */
	Settle *string `json:"settle,omitempty" `

	/*
	   商品ID     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   商品数量     */
	ProductNum *int64 `json:"product_num,omitempty" `

	/*
	   商品单价     */
	UnitPrice *string `json:"unit_price,omitempty" `

	/*
	   商品类目     */
	CategoryName *string `json:"category_name,omitempty" `

	/*
	   淘宝子单号     */
	BizOrderId *int64 `json:"biz_order_id,omitempty" `

	/*
	   淘宝主单号     */
	ParentOrderId *int64 `json:"parent_order_id,omitempty" `

	/*
	   主商品ID，针对CPS订单     */
	MainItemId *string `json:"main_item_id,omitempty" `

	/*
	   主商品名称，针对CPS订单     */
	MainItemTitle *string `json:"main_item_title,omitempty" `

	/*
	   订单状态，0-已失效 1-已下单 2-已付款 4-已收货     */
	OrderState *int64 `json:"order_state,omitempty" `

	/*
	   订单补充状态，针对CPS订单，考虑到存在已付款的cps订单发生售中退款，不参与结算的情况需要渠道知晓     */
	OrderItemStatusName *string `json:"order_item_status_name,omitempty" `

	/*
	   结算状态，1-已结算 2-未结算     */
	SettleState *int64 `json:"settle_state,omitempty" `

	/*
	   结算基数，针对CPS订单，等于付款金额+平台补贴     */
	FullSettleAmount *string `json:"full_settle_amount,omitempty" `

	/*
	   佣金比率，针对CPS订单     */
	CommissionRate *string `json:"commission_rate,omitempty" `

	/*
	   佣金金额，针对CPS订单     */
	CommissionFee *string `json:"commission_fee,omitempty" `

	/*
	   补贴比率，针对CPS订单     */
	SubsidyRate *string `json:"subsidy_rate,omitempty" `

	/*
	   补贴金额     */
	SubsidyFee *string `json:"subsidy_fee,omitempty" `

	/*
	   收入比率，针对CPS订单，等于佣金比率+补贴比率     */
	IncomeRate *string `json:"income_rate,omitempty" `

	/*
	   分成比率，针对CPS订单     */
	StratifyRate *string `json:"stratify_rate,omitempty" `

	/*
	   提成，针对CPS订单，等于收入比率*分成比率     */
	DeductRate *string `json:"deduct_rate,omitempty" `

	/*
	   技术服务费率     */
	PlatformCommissionRate *string `json:"platform_commission_rate,omitempty" `

	/*
	   技术服务费     */
	PlatformCommissionFee *string `json:"platform_commission_fee,omitempty" `

	/*
	   淘宝直播费率，针对CPS订单     */
	ChannelRate *string `json:"channel_rate,omitempty" `

	/*
	   淘宝直播费，针对CPS订单     */
	ChannelFee *string `json:"channel_fee,omitempty" `

	/*
	   媒体ID     */
	MediaId *string `json:"media_id,omitempty" `

	/*
	   媒体名称     */
	MediaName *string `json:"media_name,omitempty" `

	/*
	   推广位ID     */
	AdZoneId *string `json:"ad_zone_id,omitempty" `

	/*
	   推广位名称     */
	AdZoneName *string `json:"ad_zone_name,omitempty" `

	/*
	   招商服务费     */
	ActivityFee *string `json:"activity_fee,omitempty" `

	/*
	   招商服务费中的技术服务费     */
	ActivityServiceFee *string `json:"activity_service_fee,omitempty" `

	/*
	   招商服务费中的技术服务率     */
	ActivityServiceRate *string `json:"activity_service_rate,omitempty" `

	/*
	   更新时间     */
	GmtModified *string `json:"gmt_modified,omitempty" `

	/*
	   售中退 或 售后退     */
	Tag *string `json:"tag,omitempty" `

	/*
	   会员标识     */
	Sid *string `json:"sid,omitempty" `

	/*
	   1口碑，2饿了么     */
	PlatformType *int64 `json:"platform_type,omitempty" `

	/*
	   活动ID     */
	ActivityId *int64 `json:"activity_id,omitempty" `

	/*
	   核销门店（已加密）     */
	UsedStoreId *string `json:"used_store_id,omitempty" `

	/*
	   pid     */
	Pid *string `json:"pid,omitempty" `

	/*
	   卡券订单号     */
	RelationOrderId *int64 `json:"relation_order_id,omitempty" `

	/*
	   场景值，7卡券订单，8卡券核销订单     */
	FlowType *int64 `json:"flow_type,omitempty" `

	/*
	   0已失效，1已下单，2已付款，3售中退，4已收货，5售后退     */
	OrderItemStatus *int64 `json:"order_item_status,omitempty" `
}

func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetTitle(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetPicUrl(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.PicUrl = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetShopName(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.ShopName = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetPayAmount(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.PayAmount = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetSettleAmount(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.SettleAmount = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetTraceTime(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.TraceTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetTkCreateTime(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.TkCreateTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetPayTime(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.PayTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetReceiveTime(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.ReceiveTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetSettleTime(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.SettleTime = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetIncome(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.Income = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetSettle(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.Settle = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetItemId(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetProductNum(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.ProductNum = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetUnitPrice(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.UnitPrice = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetCategoryName(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.CategoryName = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetBizOrderId(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.BizOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetParentOrderId(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.ParentOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetMainItemId(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.MainItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetMainItemTitle(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.MainItemTitle = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetOrderState(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.OrderState = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetOrderItemStatusName(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.OrderItemStatusName = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetSettleState(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.SettleState = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetFullSettleAmount(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.FullSettleAmount = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetCommissionRate(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.CommissionRate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetCommissionFee(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.CommissionFee = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetSubsidyRate(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.SubsidyRate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetSubsidyFee(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.SubsidyFee = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetIncomeRate(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.IncomeRate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetStratifyRate(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.StratifyRate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetDeductRate(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.DeductRate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetPlatformCommissionRate(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.PlatformCommissionRate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetPlatformCommissionFee(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.PlatformCommissionFee = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetChannelRate(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.ChannelRate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetChannelFee(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.ChannelFee = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetMediaId(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.MediaId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetMediaName(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.MediaName = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetAdZoneId(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.AdZoneId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetAdZoneName(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.AdZoneName = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetActivityFee(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.ActivityFee = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetActivityServiceFee(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.ActivityServiceFee = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetActivityServiceRate(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.ActivityServiceRate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetGmtModified(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.GmtModified = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetTag(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.Tag = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetSid(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.Sid = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetPlatformType(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.PlatformType = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetActivityId(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.ActivityId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetUsedStoreId(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.UsedStoreId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetPid(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.Pid = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetRelationOrderId(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.RelationOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetFlowType(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.FlowType = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO) SetOrderItemStatus(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO {
	s.OrderItemStatus = &v
	return s
}
