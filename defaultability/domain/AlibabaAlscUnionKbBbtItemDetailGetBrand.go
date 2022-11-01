package domain

type AlibabaAlscUnionKbBbtItemDetailGetBrand struct {
	/*
	   品牌Id     */
	BrandId *string `json:"brand_id,omitempty" `

	/*
	   品牌名     */
	BrandName *string `json:"brand_name,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetBrand) SetBrandId(v string) *AlibabaAlscUnionKbBbtItemDetailGetBrand {
	s.BrandId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBrand) SetBrandName(v string) *AlibabaAlscUnionKbBbtItemDetailGetBrand {
	s.BrandName = &v
	return s
}
