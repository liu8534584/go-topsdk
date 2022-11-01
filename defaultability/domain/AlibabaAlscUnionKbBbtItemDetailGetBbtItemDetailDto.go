package domain

type AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto struct {
	/*
	   商品ID     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   标题     */
	Title *string `json:"title,omitempty" `

	/*
	   副标题     */
	SubTitle *string `json:"sub_title,omitempty" `

	/*
	   主图     */
	MainPicture *string `json:"main_picture,omitempty" `

	/*
	   相册     */
	Images *[]AlibabaAlscUnionKbBbtItemDetailGetImageDto `json:"images,omitempty" `

	/*
	   售卖起始时间（秒）     */
	SaleStartTime *int64 `json:"sale_start_time,omitempty" `

	/*
	   售卖结束时间（秒）     */
	SaleEndTime *int64 `json:"sale_end_time,omitempty" `

	/*
	   原价（分）     */
	OriginalPriceCent *int64 `json:"original_price_cent,omitempty" `

	/*
	   活动价（分）     */
	ActivityPriceCent *int64 `json:"activity_price_cent,omitempty" `

	/*
	   佣金比例     */
	CommissionRate *string `json:"commission_rate,omitempty" `

	/*
	   总销量     */
	TotalSales *int64 `json:"total_sales,omitempty" `

	/*
	   库存     */
	Stock *int64 `json:"stock,omitempty" `

	/*
	   适用门店数量（city_id不为空则返回当前城市可用门店数，否则返回全部可用门店数）     */
	ApplyShopCount *int64 `json:"apply_shop_count,omitempty" `

	/*
	   折扣     */
	Discount *string `json:"discount,omitempty" `

	/*
	   商品详情     */
	ItemDetail *AlibabaAlscUnionKbBbtItemDetailGetItemDetail `json:"item_detail,omitempty" `

	/*
	   商品子类型     */
	ItemSubType *string `json:"item_sub_type,omitempty" `

	/*
	   可核销次数     */
	UseTimes *int64 `json:"use_times,omitempty" `

	/*
	   商品可用城市     */
	ApplyCityIds *[]string `json:"apply_city_ids,omitempty" `

	/*
	   当前商品购买是否需要手机号     */
	NeedPhone *bool `json:"need_phone,omitempty" `

	/*
	   淘宝二级类目ID     */
	TbCategory2Id *string `json:"tb_category_2_id,omitempty" `

	/*
	   淘宝二级类目名称     */
	TbCategory2Name *string `json:"tb_category_2_name,omitempty" `

	/*
	   淘宝三级类目ID     */
	TbCategory3Id *string `json:"tb_category_3_id,omitempty" `

	/*
	   淘宝三级类目名称     */
	TbCategory3Name *string `json:"tb_category_3_name,omitempty" `

	/*
	   限购份数（-1表示不限购）     */
	BuyLimit *int64 `json:"buy_limit,omitempty" `

	/*
	   门店商品相册     */
	ShopItemImages *[]AlibabaAlscUnionKbBbtItemDetailGetImageDto `json:"shop_item_images,omitempty" `

	/*
	   门店环境相册     */
	ShopEnvironmentImages *[]AlibabaAlscUnionKbBbtItemDetailGetImageDto `json:"shop_environment_images,omitempty" `

	/*
	   品牌     */
	Brand *AlibabaAlscUnionKbBbtItemDetailGetBrand `json:"brand,omitempty" `

	/*
	   三方服务商名称     */
	TripartiteName *string `json:"tripartite_name,omitempty" `

	/*
	   三方服务商appKey     */
	TripartiteAppkey *string `json:"tripartite_appkey,omitempty" `

	/*
	   三方站点名称     */
	TripartiteSiteName *string `json:"tripartite_site_name,omitempty" `

	/*
	   免责声明     */
	Disclaimer *string `json:"disclaimer,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetItemId(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetTitle(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetSubTitle(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.SubTitle = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetMainPicture(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.MainPicture = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetImages(v []AlibabaAlscUnionKbBbtItemDetailGetImageDto) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.Images = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetSaleStartTime(v int64) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.SaleStartTime = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetSaleEndTime(v int64) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.SaleEndTime = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetOriginalPriceCent(v int64) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.OriginalPriceCent = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetActivityPriceCent(v int64) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.ActivityPriceCent = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetCommissionRate(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.CommissionRate = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetTotalSales(v int64) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.TotalSales = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetStock(v int64) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.Stock = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetApplyShopCount(v int64) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.ApplyShopCount = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetDiscount(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.Discount = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetItemDetail(v AlibabaAlscUnionKbBbtItemDetailGetItemDetail) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.ItemDetail = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetItemSubType(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.ItemSubType = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetUseTimes(v int64) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.UseTimes = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetApplyCityIds(v []string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.ApplyCityIds = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetNeedPhone(v bool) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.NeedPhone = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetTbCategory2Id(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.TbCategory2Id = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetTbCategory2Name(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.TbCategory2Name = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetTbCategory3Id(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.TbCategory3Id = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetTbCategory3Name(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.TbCategory3Name = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetBuyLimit(v int64) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.BuyLimit = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetShopItemImages(v []AlibabaAlscUnionKbBbtItemDetailGetImageDto) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.ShopItemImages = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetShopEnvironmentImages(v []AlibabaAlscUnionKbBbtItemDetailGetImageDto) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.ShopEnvironmentImages = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetBrand(v AlibabaAlscUnionKbBbtItemDetailGetBrand) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.Brand = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetTripartiteName(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.TripartiteName = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetTripartiteAppkey(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.TripartiteAppkey = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetTripartiteSiteName(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.TripartiteSiteName = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto) SetDisclaimer(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailDto {
	s.Disclaimer = &v
	return s
}
