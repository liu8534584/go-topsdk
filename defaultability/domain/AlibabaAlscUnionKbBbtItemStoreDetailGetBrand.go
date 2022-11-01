package domain

type AlibabaAlscUnionKbBbtItemStoreDetailGetBrand struct {
	/*
	   品牌Id     */
	BrandId *string `json:"brand_id,omitempty" `

	/*
	   品牌名     */
	BrandName *string `json:"brand_name,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetBrand) SetBrandId(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetBrand {
	s.BrandId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetBrand) SetBrandName(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetBrand {
	s.BrandName = &v
	return s
}
