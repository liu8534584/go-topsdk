package domain

type AlibabaAlscUnionElemePromotionStorepromotionGetPromotionItem struct {
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

func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionItem) SetTitle(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionItem {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionItem) SetOriginPrice(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionItem {
	s.OriginPrice = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionItem) SetPrice(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionItem {
	s.Price = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionItem) SetPicture(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetPromotionItem {
	s.Picture = &v
	return s
}
