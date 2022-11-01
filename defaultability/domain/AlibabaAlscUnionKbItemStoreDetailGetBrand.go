package domain

type AlibabaAlscUnionKbItemStoreDetailGetBrand struct {
	/*
	   品牌Id     */
	BrandId *string `json:"brand_id,omitempty" `

	/*
	   品牌名     */
	BrandName *string `json:"brand_name,omitempty" `
}

func (s *AlibabaAlscUnionKbItemStoreDetailGetBrand) SetBrandId(v string) *AlibabaAlscUnionKbItemStoreDetailGetBrand {
	s.BrandId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetBrand) SetBrandName(v string) *AlibabaAlscUnionKbItemStoreDetailGetBrand {
	s.BrandName = &v
	return s
}
