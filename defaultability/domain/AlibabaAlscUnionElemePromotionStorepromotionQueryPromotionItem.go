package domain

type AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionItem struct {
	/*
	   标题     */
	Title *string `json:"title,omitempty" `

	/*
	   原价     */
	OriginPrice *string `json:"origin_price,omitempty" `

	/*
	   现价     */
	Price *string `json:"price,omitempty" `

	/*
	   图片     */
	Picture *string `json:"picture,omitempty" `
}

func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionItem) SetTitle(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionItem {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionItem) SetOriginPrice(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionItem {
	s.OriginPrice = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionItem) SetPrice(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionItem {
	s.Price = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionItem) SetPicture(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionItem {
	s.Picture = &v
	return s
}
