package domain

type AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink struct {
	/*
	   小程序appId     */
	WxAppid *string `json:"wx_appid,omitempty" `

	/*
	   微信小程序path链接     */
	WxPath *string `json:"wx_path,omitempty" `

	/*
	   推广图片地址     */
	Picture *string `json:"picture,omitempty" `

	/*
	   支付宝小程序推广链接     */
	AlipayMiniUrl *string `json:"alipay_mini_url,omitempty" `

	/*
	   h5推广地址     */
	H5Url *string `json:"h5_url,omitempty" `

	/*
	   淘宝二维码图片地址     */
	TbQrCode *string `json:"tb_qr_code,omitempty" `

	/*
	   微信独立二维码     */
	MiniQrcode *string `json:"mini_qrcode,omitempty" `

	/*
	   淘宝独立二维码     */
	TbMiniQrcode *string `json:"tb_mini_qrcode,omitempty" `

	/*
	   饿了么唤端链接     */
	EleSchemeUrl *string `json:"ele_scheme_url,omitempty" `

	/*
	   h5推广地址短链     */
	H5ShortLink *string `json:"h5_short_link,omitempty" `
}

func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetWxAppid(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
	s.WxAppid = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetWxPath(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
	s.WxPath = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetPicture(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
	s.Picture = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetAlipayMiniUrl(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
	s.AlipayMiniUrl = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetH5Url(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
	s.H5Url = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetTbQrCode(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
	s.TbQrCode = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetMiniQrcode(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
	s.MiniQrcode = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetTbMiniQrcode(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
	s.TbMiniQrcode = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetEleSchemeUrl(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
	s.EleSchemeUrl = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) SetH5ShortLink(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink {
	s.H5ShortLink = &v
	return s
}
