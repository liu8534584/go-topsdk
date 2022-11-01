package domain

type AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest struct {
	/*
	   会话ID（分页场景首次请求结果返回，后续请求必须携带，服务根据session_id相同请求次数自动翻页返回）     */
	SessionId *string `json:"session_id,omitempty" `

	/*
	   渠道PID     */
	Pid *string `json:"pid,omitempty" `

	/*
	   经度     */
	Longitude *string `json:"longitude,omitempty" `

	/*
	   纬度     */
	Latitude *string `json:"latitude,omitempty" `

	/*
	   城市编码     */
	CityId *string `json:"city_id,omitempty" `

	/*
	   排序类型，默认normal，排序规则包括:{"normal":"商品创建时间倒序","distance":"距离由近到远","commission":"佣金倒序","monthlySale":"月销量","couponAmount":"叠加券金额倒序","activityReward":"奖励金金额倒序"}     */
	SortType *string `json:"sort_type,omitempty" `

	/*
	   是否参与奖励金活动（默认false不做过滤）     */
	InActivity *bool `json:"in_activity,omitempty" `

	/*
	   否当前有c端奖励金活动库存（默认false不做过滤）     */
	HasBonusStock *bool `json:"has_bonus_stock,omitempty" `

	/*
	   店铺佣金比例下限，代表筛选店铺全店佣金大于等于0.01的店铺     */
	MinCommissionRate *string `json:"min_commission_rate,omitempty" `

	/*
	   每页数量（1~20，默认20）     */
	PageSize *int64 `json:"page_size,omitempty" `

	/*
	   三方扩展id     */
	Sid *string `json:"sid,omitempty" `

	/*
	   指定召回供给枚举     */
	BizType *string `json:"biz_type,omitempty" `

	/*
	   in_activity=false的条件下，召回的非奖励金活动cps商家是否需要带出叠加券     */
	RecallOverlayCoupon *bool `json:"recall_overlay_coupon,omitempty" `

	/*
	   filter_has_overlay_coupon=true的条件下，限定只召回带叠加券的cps商户     */
	FilterHasOverlayCoupon *bool `json:"filter_has_overlay_coupon,omitempty" `

	/*
	   filter_has_overlay_coupon=true的情况下，限定的最小叠加券券金额，单位元     */
	MinOverlayCouponAmount *string `json:"min_overlay_coupon_amount,omitempty" `

	/*
	   以一级类目进行类目限定，以,或者|进行类目分隔     */
	FilterFirstCategories *string `json:"filter_first_categories,omitempty" `
}

func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetSessionId(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.SessionId = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetPid(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.Pid = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetLongitude(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.Longitude = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetLatitude(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.Latitude = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetCityId(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.CityId = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetSortType(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.SortType = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetInActivity(v bool) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.InActivity = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetHasBonusStock(v bool) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.HasBonusStock = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetMinCommissionRate(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.MinCommissionRate = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetPageSize(v int64) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.PageSize = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetSid(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.Sid = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetBizType(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.BizType = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetRecallOverlayCoupon(v bool) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.RecallOverlayCoupon = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetFilterHasOverlayCoupon(v bool) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.FilterHasOverlayCoupon = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetMinOverlayCouponAmount(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.MinOverlayCouponAmount = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) SetFilterFirstCategories(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest {
	s.FilterFirstCategories = &v
	return s
}
