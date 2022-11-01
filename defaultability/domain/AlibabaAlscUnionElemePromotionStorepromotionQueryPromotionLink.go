package domain

type AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionLink struct {
	/*
	   小程序appId     */
	WxAppid *string `json:"wx_appid,omitempty" `

	/*
	   微信小程序path链接     */
	WxPath *string `json:"wx_path,omitempty" `

	/*
	   小程序appId-立减活动     */
	ReduceWxAppid *string `json:"reduce_wx_appid,omitempty" `

	/*
	   微信小程序path链接-立减活动     */
	ReduceWxPath *string `json:"reduce_wx_path,omitempty" `
}

func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionLink) SetWxAppid(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionLink {
	s.WxAppid = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionLink) SetWxPath(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionLink {
	s.WxPath = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionLink) SetReduceWxAppid(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionLink {
	s.ReduceWxAppid = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionLink) SetReduceWxPath(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionLink {
	s.ReduceWxPath = &v
	return s
}
