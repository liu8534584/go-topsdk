package domain

type AlibabaAlscUnionKbItemQueryKbItemPromotionDTO struct {
	/*
	   商品ID     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   标题     */
	Title *string `json:"title,omitempty" `

	/*
	   商品图     */
	MainPicture *string `json:"main_picture,omitempty" `

	/*
	   原价（分）     */
	OriginalPriceCent *int64 `json:"original_price_cent,omitempty" `

	/*
	   活动价（分）     */
	ActivityPriceCent *int64 `json:"activity_price_cent,omitempty" `

	/*
	   券后价（分）     */
	PriceWithCouponCent *int64 `json:"price_with_coupon_cent,omitempty" `

	/*
	   券价格（分）     */
	CouponPriceCent *int64 `json:"coupon_price_cent,omitempty" `

	/*
	   九十天销量     */
	NinetySales *int64 `json:"ninety_sales,omitempty" `

	/*
	   总销量     */
	TotalSales *int64 `json:"total_sales,omitempty" `

	/*
	   推广链接     */
	Link *AlibabaAlscUnionKbItemQueryPromotionLink `json:"link,omitempty" `
}

func (s *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO) SetItemId(v string) *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO) SetTitle(v string) *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO) SetMainPicture(v string) *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO {
	s.MainPicture = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO) SetOriginalPriceCent(v int64) *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO {
	s.OriginalPriceCent = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO) SetActivityPriceCent(v int64) *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO {
	s.ActivityPriceCent = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO) SetPriceWithCouponCent(v int64) *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO {
	s.PriceWithCouponCent = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO) SetCouponPriceCent(v int64) *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO {
	s.CouponPriceCent = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO) SetNinetySales(v int64) *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO {
	s.NinetySales = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO) SetTotalSales(v int64) *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO {
	s.TotalSales = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO) SetLink(v AlibabaAlscUnionKbItemQueryPromotionLink) *AlibabaAlscUnionKbItemQueryKbItemPromotionDTO {
	s.Link = &v
	return s
}
