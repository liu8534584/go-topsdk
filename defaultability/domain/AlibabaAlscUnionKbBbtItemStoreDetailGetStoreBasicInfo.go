package domain

type AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo struct {
	/*
	   门店ID     */
	StoreId *string `json:"store_id,omitempty" `

	/*
	   店名     */
	Name *string `json:"name,omitempty" `

	/*
	   cover     */
	Cover *string `json:"cover,omitempty" `

	/*
	   门店电话     */
	Mobiles *[]string `json:"mobiles,omitempty" `

	/*
	   营业信息     */
	Business *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusiness `json:"business,omitempty" `

	/*
	   位置信息     */
	Location *AlibabaAlscUnionKbBbtItemStoreDetailGetLocation `json:"location,omitempty" `

	/*
	   品牌     */
	Brand *AlibabaAlscUnionKbBbtItemStoreDetailGetBrand `json:"brand,omitempty" `

	/*
	   门店所属行业分类     */
	Categories *[]AlibabaAlscUnionKbBbtItemStoreDetailGetCategory `json:"categories,omitempty" `

	/*
	   门店资质     */
	Qualifications *[]AlibabaAlscUnionKbBbtItemStoreDetailGetImageContent `json:"qualifications,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo) SetStoreId(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo {
	s.StoreId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo) SetName(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo {
	s.Name = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo) SetCover(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo {
	s.Cover = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo) SetMobiles(v []string) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo {
	s.Mobiles = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo) SetBusiness(v AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusiness) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo {
	s.Business = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo) SetLocation(v AlibabaAlscUnionKbBbtItemStoreDetailGetLocation) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo {
	s.Location = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo) SetBrand(v AlibabaAlscUnionKbBbtItemStoreDetailGetBrand) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo {
	s.Brand = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo) SetCategories(v []AlibabaAlscUnionKbBbtItemStoreDetailGetCategory) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo {
	s.Categories = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo) SetQualifications(v []AlibabaAlscUnionKbBbtItemStoreDetailGetImageContent) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBasicInfo {
	s.Qualifications = &v
	return s
}
