package domain

type AlibabaAlscUnionElemePromotionStorepromotionGetSingleStorePromotionRequest struct {
	/*
	   渠道PID     */
	Pid *string `json:"pid,omitempty" `

	/*
	   门店ID（加密）     */
	ShopId *string `json:"shop_id,omitempty" `

	/*
	   活动ID     */
	ActivityId *string `json:"activity_id,omitempty" `

	/*
	   三方扩展id     */
	Sid *string `json:"sid,omitempty" `

	/*
	   是否返回微信推广图片     */
	IncludeWxImg *bool `json:"include_wx_img,omitempty" `
}

func (s *AlibabaAlscUnionElemePromotionStorepromotionGetSingleStorePromotionRequest) SetPid(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetSingleStorePromotionRequest {
	s.Pid = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetSingleStorePromotionRequest) SetShopId(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetSingleStorePromotionRequest {
	s.ShopId = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetSingleStorePromotionRequest) SetActivityId(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetSingleStorePromotionRequest {
	s.ActivityId = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetSingleStorePromotionRequest) SetSid(v string) *AlibabaAlscUnionElemePromotionStorepromotionGetSingleStorePromotionRequest {
	s.Sid = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionGetSingleStorePromotionRequest) SetIncludeWxImg(v bool) *AlibabaAlscUnionElemePromotionStorepromotionGetSingleStorePromotionRequest {
	s.IncludeWxImg = &v
	return s
}
