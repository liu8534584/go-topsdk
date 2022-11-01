package domain

type AlibabaAlscUnionKbItemQueryPromotionLink struct {
	/*
	   支付宝schema     */
	AlipaySchemeUrl *string `json:"alipay_scheme_url,omitempty" `

	/*
	   支付宝小程序h5 url     */
	AlipayH5Url *string `json:"alipay_h5_url,omitempty" `
}

func (s *AlibabaAlscUnionKbItemQueryPromotionLink) SetAlipaySchemeUrl(v string) *AlibabaAlscUnionKbItemQueryPromotionLink {
	s.AlipaySchemeUrl = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryPromotionLink) SetAlipayH5Url(v string) *AlibabaAlscUnionKbItemQueryPromotionLink {
	s.AlipayH5Url = &v
	return s
}
