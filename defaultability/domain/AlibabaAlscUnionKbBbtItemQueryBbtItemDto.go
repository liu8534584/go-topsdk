package domain

type AlibabaAlscUnionKbBbtItemQueryBbtItemDto struct {
	/*
	   商品ID     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   标题     */
	Title *string `json:"title,omitempty" `

	/*
	   主图     */
	MainPicture *string `json:"main_picture,omitempty" `

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
	   适用门店数量（city_id不为空则返回当前城市可用门店数，否则返回全部可用门店数）     */
	ApplyShopCount *int64 `json:"apply_shop_count,omitempty" `

	/*
	   销量     */
	TotalSales *int64 `json:"total_sales,omitempty" `

	/*
	   库存     */
	Stock *int64 `json:"stock,omitempty" `

	/*
	   商品子类型     */
	ItemSubType *string `json:"item_sub_type,omitempty" `

	/*
	   使用次数     */
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
	   品牌     */
	Brand *AlibabaAlscUnionKbBbtItemQueryBrand `json:"brand,omitempty" `

	/*
	   三方服务商名称     */
	TripartiteName *string `json:"tripartite_name,omitempty" `

	/*
	   三方服务商appKey     */
	TripartiteAppkey *string `json:"tripartite_appkey,omitempty" `

	/*
	   三方站点名称     */
	TripartiteSiteName *string `json:"tripartite_site_name,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetItemId(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetTitle(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetMainPicture(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.MainPicture = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetSaleStartTime(v int64) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.SaleStartTime = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetSaleEndTime(v int64) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.SaleEndTime = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetOriginalPriceCent(v int64) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.OriginalPriceCent = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetActivityPriceCent(v int64) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.ActivityPriceCent = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetCommissionRate(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.CommissionRate = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetApplyShopCount(v int64) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.ApplyShopCount = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetTotalSales(v int64) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.TotalSales = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetStock(v int64) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.Stock = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetItemSubType(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.ItemSubType = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetUseTimes(v int64) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.UseTimes = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetApplyCityIds(v []string) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.ApplyCityIds = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetNeedPhone(v bool) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.NeedPhone = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetTbCategory2Id(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.TbCategory2Id = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetTbCategory2Name(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.TbCategory2Name = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetTbCategory3Id(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.TbCategory3Id = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetTbCategory3Name(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.TbCategory3Name = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetBuyLimit(v int64) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.BuyLimit = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetBrand(v AlibabaAlscUnionKbBbtItemQueryBrand) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.Brand = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetTripartiteName(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.TripartiteName = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetTripartiteAppkey(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.TripartiteAppkey = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemDto) SetTripartiteSiteName(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemDto {
	s.TripartiteSiteName = &v
	return s
}
