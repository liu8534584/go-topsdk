package domain

type AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto struct {
	/*
	   门店名称     */
	Title *string `json:"title,omitempty" `

	/*
	   门店logo     */
	ShopLogo *string `json:"shop_logo,omitempty" `

	/*
	   模糊销量     */
	IndistinctMonthlySales *string `json:"indistinct_monthly_sales,omitempty" `

	/*
	   佣金比例     */
	CommissionRate *string `json:"commission_rate,omitempty" `

	/*
	   店铺类型（"activityCps":活动cps,"ordinaryCps":基础cps）     */
	BizType *string `json:"biz_type,omitempty" `

	/*
	   活动数据     */
	Activity *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionActivity `json:"activity,omitempty" `

	/*
	   推广链接     */
	Link *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink `json:"link,omitempty" `

	/*
	   一级类目ID     */
	Category1Id *string `json:"category_1_id,omitempty" `

	/*
	   起送价（元）     */
	DeliveryPrice *string `json:"delivery_price,omitempty" `

	/*
	   推荐理由     */
	RecommendReasons *[]string `json:"recommend_reasons,omitempty" `

	/*
	   服务评级     */
	ServiceRating *string `json:"service_rating,omitempty" `

	/*
	   推荐商品     */
	Items *[]AlibabaAlscUnionElemePromotionStorepromotionGetPromotionItem `json:"items,omitempty" `

	/*
	   一级类目名称     */
	Category1Name *string `json:"category_1_name,omitempty" `

	/*
	   叠加券活动数据     */
	OverlayCoupon *AlibabaAlscUnionElemePromotionStorepromotionGetOverlayCouponDTO `json:"overlay_coupon,omitempty" `
}

func (s *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto) SetTitle(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto) SetShopLogo(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto {
	s.ShopLogo = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto) SetIndistinctMonthlySales(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto {
	s.IndistinctMonthlySales = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto) SetCommissionRate(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto {
	s.CommissionRate = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto) SetBizType(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto {
	s.BizType = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto) SetActivity(v AlibabaAlscUnionElemePromotionStorepromotionGetPromotionActivity) *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto {
	s.Activity = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto) SetLink(v AlibabaAlscUnionElemePromotionStorepromotionGetPromotionLink) *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto {
	s.Link = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto) SetCategory1Id(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto {
	s.Category1Id = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto) SetDeliveryPrice(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto {
	s.DeliveryPrice = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto) SetRecommendReasons(v []string) *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto {
	s.RecommendReasons = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto) SetServiceRating(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto {
	s.ServiceRating = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto) SetItems(v []AlibabaAlscUnionElemePromotionStorepromotionGetPromotionItem) *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto {
	s.Items = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto) SetCategory1Name(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto {
	s.Category1Name = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto) SetOverlayCoupon(v AlibabaAlscUnionElemePromotionStorepromotionGetOverlayCouponDTO) *AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto {
	s.OverlayCoupon = &v
	return s
}
