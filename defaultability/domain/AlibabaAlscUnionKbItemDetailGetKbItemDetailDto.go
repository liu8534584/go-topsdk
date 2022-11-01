package domain

type AlibabaAlscUnionKbItemDetailGetKbItemDetailDto struct {
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
	Images *[]AlibabaAlscUnionKbItemDetailGetImageDto `json:"images,omitempty" `

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
	ItemDetail *AlibabaAlscUnionKbItemDetailGetItemDetail `json:"item_detail,omitempty" `

	/*
	   商品子类型     */
	ItemSubType *string `json:"item_sub_type,omitempty" `

	/*
	   可核销次数     */
	UseTimes *int64 `json:"use_times,omitempty" `

	/*
	   商品可用城市     */
	ApplyCityIds *[]int32 `json:"apply_city_ids,omitempty" `

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
	ShopItemImages *[]AlibabaAlscUnionKbItemDetailGetInteger `json:"shop_item_images,omitempty" `

	/*
	   门店环境相册     */
	ShopEnvironmentImages *[]AlibabaAlscUnionKbItemDetailGetImageDto `json:"shop_environment_images,omitempty" `

	/*
	   商品可售卖的端类型。1支付宝端商品，2微信端商品，3全部     */
	ItemType *int64 `json:"item_type,omitempty" `

	/*
	   品牌     */
	Brand *AlibabaAlscUnionKbItemDetailGetBrand `json:"brand,omitempty" `
}

func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetItemId(v string) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetTitle(v string) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetSubTitle(v string) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.SubTitle = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetMainPicture(v string) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.MainPicture = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetImages(v []AlibabaAlscUnionKbItemDetailGetImageDto) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.Images = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetSaleStartTime(v int64) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.SaleStartTime = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetSaleEndTime(v int64) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.SaleEndTime = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetOriginalPriceCent(v int64) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.OriginalPriceCent = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetActivityPriceCent(v int64) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.ActivityPriceCent = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetTotalSales(v int64) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.TotalSales = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetStock(v int64) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.Stock = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetApplyShopCount(v int64) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.ApplyShopCount = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetDiscount(v string) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.Discount = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetItemDetail(v AlibabaAlscUnionKbItemDetailGetItemDetail) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.ItemDetail = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetItemSubType(v string) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.ItemSubType = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetUseTimes(v int64) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.UseTimes = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetApplyCityIds(v []int32) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.ApplyCityIds = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetNeedPhone(v bool) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.NeedPhone = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetTbCategory2Id(v string) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.TbCategory2Id = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetTbCategory2Name(v string) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.TbCategory2Name = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetTbCategory3Id(v string) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.TbCategory3Id = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetTbCategory3Name(v string) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.TbCategory3Name = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetBuyLimit(v int64) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.BuyLimit = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetShopItemImages(v []AlibabaAlscUnionKbItemDetailGetInteger) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.ShopItemImages = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetShopEnvironmentImages(v []AlibabaAlscUnionKbItemDetailGetImageDto) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.ShopEnvironmentImages = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetItemType(v int64) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.ItemType = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto) SetBrand(v AlibabaAlscUnionKbItemDetailGetBrand) *AlibabaAlscUnionKbItemDetailGetKbItemDetailDto {
	s.Brand = &v
	return s
}
