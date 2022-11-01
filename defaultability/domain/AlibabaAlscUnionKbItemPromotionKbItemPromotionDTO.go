package domain

type AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO struct {
	/*
	   商品可售卖截止时间,时间戳(秒)     */
	ItemSaleEndTime *int64 `json:"item_sale_end_time,omitempty" `

	/*
	   原始价格，单位元     */
	OriginalPrice *string `json:"original_price,omitempty" `

	/*
	   口碑微信小程序appId     */
	WxAppId *string `json:"wx_app_id,omitempty" `

	/*
	   折扣     */
	Discount *string `json:"discount,omitempty" `

	/*
	   商品标题     */
	Title *string `json:"title,omitempty" `

	/*
	   月销量     */
	ThirtySoldQuantity *int64 `json:"thirty_sold_quantity,omitempty" `

	/*
	   商品id     */
	ItemId *int64 `json:"item_id,omitempty" `

	/*
	   售卖价格,折扣后价格     */
	Price *string `json:"price,omitempty" `

	/*
	   商品图片     */
	ImageUrl *string `json:"image_url,omitempty" `

	/*
	   点击商品后，微信小程序的承接页  	     */
	WxPath *string `json:"wx_path,omitempty" `

	/*
	   预估佣金，单位元     */
	Commission *string `json:"commission,omitempty" `

	/*
	   商品可适用门店数量     */
	ApplyShopNum *int64 `json:"apply_shop_num,omitempty" `

	/*
	   库存     */
	Stock *int64 `json:"stock,omitempty" `

	/*
	   商品可售卖开始时间,单位元     */
	ItemSaleStartTime *int64 `json:"item_sale_start_time,omitempty" `

	/*
	   可使用城市列表     */
	ValidCities *[]string `json:"valid_cities,omitempty" `

	/*
	   核销后奖励佣金,单位元;cpa业务类型返回     */
	BonusCommission *string `json:"bonus_commission,omitempty" `

	/*
	   总销量     */
	TotalSales *int64 `json:"total_sales,omitempty" `

	/*
	   商品可售卖的端类型。1支付宝端商品，2微信端商品，3全部     */
	ItemType *int64 `json:"item_type,omitempty" `
}

func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetItemSaleEndTime(v int64) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.ItemSaleEndTime = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetOriginalPrice(v string) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.OriginalPrice = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetWxAppId(v string) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.WxAppId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetDiscount(v string) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.Discount = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetTitle(v string) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetThirtySoldQuantity(v int64) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.ThirtySoldQuantity = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetItemId(v int64) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetPrice(v string) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.Price = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetImageUrl(v string) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.ImageUrl = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetWxPath(v string) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.WxPath = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetCommission(v string) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.Commission = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetApplyShopNum(v int64) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.ApplyShopNum = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetStock(v int64) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.Stock = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetItemSaleStartTime(v int64) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.ItemSaleStartTime = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetValidCities(v []string) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.ValidCities = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetBonusCommission(v string) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.BonusCommission = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetTotalSales(v int64) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.TotalSales = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO) SetItemType(v int64) *AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO {
	s.ItemType = &v
	return s
}
