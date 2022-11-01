package domain

type AlibabaAlscUnionKbBbtItemQueryBrand struct {
	/*
	   品牌Id     */
	BrandId *string `json:"brand_id,omitempty" `

	/*
	   品牌名     */
	BrandName *string `json:"brand_name,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemQueryBrand) SetBrandId(v string) *AlibabaAlscUnionKbBbtItemQueryBrand {
	s.BrandId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBrand) SetBrandName(v string) *AlibabaAlscUnionKbBbtItemQueryBrand {
	s.BrandName = &v
	return s
}
