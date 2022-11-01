package domain

type AlibabaAlscUnionKbStoreItemQueryPromotionLink struct {
	/*
	   支付宝schema     */
	AlipaySchemeUrl *string `json:"alipay_scheme_url,omitempty" `

	/*
	   支付宝小程序h5 url     */
	AlipayH5Url *string `json:"alipay_h5_url,omitempty" `
}

func (s *AlibabaAlscUnionKbStoreItemQueryPromotionLink) SetAlipaySchemeUrl(v string) *AlibabaAlscUnionKbStoreItemQueryPromotionLink {
	s.AlipaySchemeUrl = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreItemQueryPromotionLink) SetAlipayH5Url(v string) *AlibabaAlscUnionKbStoreItemQueryPromotionLink {
	s.AlipayH5Url = &v
	return s
}
