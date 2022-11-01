package domain

type AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO struct {
	/*
	   微信推广图片下载链接     */
	ImgUrl *string `json:"img_url,omitempty" `

	/*
	   微信小程序的appid     */
	WxAppid *string `json:"wx_appid,omitempty" `

	/*
	   微信小程序的路径     */
	WxPath *string `json:"wx_path,omitempty" `

	/*
	   微信小程序码     */
	MiniQrCode *string `json:"mini_qr_code,omitempty" `

	/*
	   支付宝推广图片地址     */
	AlipayImgUrl *string `json:"alipay_img_url,omitempty" `

	/*
	   支付宝吱口令     */
	AlipayWatchword *string `json:"alipay_watchword,omitempty" `

	/*
	   支付宝吱口令的引导文案     */
	AlipayWatchwordSuggest *string `json:"alipay_watchword_suggest,omitempty" `

	/*
	   支付宝小程序码     */
	AlipayMiniQrCode *string `json:"alipay_mini_qr_code,omitempty" `

	/*
	   支付宝小程序path     */
	AlipaySchemeUrl *string `json:"alipay_scheme_url,omitempty" `

	/*
	   支付宝的h5链接     */
	H5Url *string `json:"h5_url,omitempty" `
}

func (s *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO) SetImgUrl(v string) *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO {
	s.ImgUrl = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO) SetWxAppid(v string) *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO {
	s.WxAppid = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO) SetWxPath(v string) *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO {
	s.WxPath = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO) SetMiniQrCode(v string) *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO {
	s.MiniQrCode = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO) SetAlipayImgUrl(v string) *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO {
	s.AlipayImgUrl = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO) SetAlipayWatchword(v string) *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO {
	s.AlipayWatchword = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO) SetAlipayWatchwordSuggest(v string) *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO {
	s.AlipayWatchwordSuggest = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO) SetAlipayMiniQrCode(v string) *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO {
	s.AlipayMiniQrCode = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO) SetAlipaySchemeUrl(v string) *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO {
	s.AlipaySchemeUrl = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO) SetH5Url(v string) *AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO {
	s.H5Url = &v
	return s
}
