package domain

type AlibabaAlscUnionKbStoreItemQueryKbShopItemDto struct {
	/*
	   商品ID     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   商品标题     */
	Title *string `json:"title,omitempty" `

	/*
	   商品图     */
	MainPicture *string `json:"main_picture,omitempty" `

	/*
	   活动价（分）     */
	ActivityPriceCent *int64 `json:"activity_price_cent,omitempty" `

	/*
	   券后价（分）     */
	PriceWithCouponCent *int64 `json:"price_with_coupon_cent,omitempty" `

	/*
	   推广链接     */
	Link *AlibabaAlscUnionKbStoreItemQueryPromotionLink `json:"link,omitempty" `
}

func (s *AlibabaAlscUnionKbStoreItemQueryKbShopItemDto) SetItemId(v string) *AlibabaAlscUnionKbStoreItemQueryKbShopItemDto {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreItemQueryKbShopItemDto) SetTitle(v string) *AlibabaAlscUnionKbStoreItemQueryKbShopItemDto {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreItemQueryKbShopItemDto) SetMainPicture(v string) *AlibabaAlscUnionKbStoreItemQueryKbShopItemDto {
	s.MainPicture = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreItemQueryKbShopItemDto) SetActivityPriceCent(v int64) *AlibabaAlscUnionKbStoreItemQueryKbShopItemDto {
	s.ActivityPriceCent = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreItemQueryKbShopItemDto) SetPriceWithCouponCent(v int64) *AlibabaAlscUnionKbStoreItemQueryKbShopItemDto {
	s.PriceWithCouponCent = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreItemQueryKbShopItemDto) SetLink(v AlibabaAlscUnionKbStoreItemQueryPromotionLink) *AlibabaAlscUnionKbStoreItemQueryKbShopItemDto {
	s.Link = &v
	return s
}
