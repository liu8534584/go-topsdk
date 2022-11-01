package domain

type AlibabaAlscUnionKbItemDetailGetBrand struct {
	/*
	   品牌Id     */
	BrandId *string `json:"brand_id,omitempty" `

	/*
	   品牌名     */
	BrandName *string `json:"brand_name,omitempty" `
}

func (s *AlibabaAlscUnionKbItemDetailGetBrand) SetBrandId(v string) *AlibabaAlscUnionKbItemDetailGetBrand {
	s.BrandId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetBrand) SetBrandName(v string) *AlibabaAlscUnionKbItemDetailGetBrand {
	s.BrandName = &v
	return s
}
