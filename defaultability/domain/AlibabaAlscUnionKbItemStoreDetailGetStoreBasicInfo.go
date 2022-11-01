package domain

type AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo struct {
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
	Business *AlibabaAlscUnionKbItemStoreDetailGetStoreBusiness `json:"business,omitempty" `

	/*
	   位置信息     */
	Location *AlibabaAlscUnionKbItemStoreDetailGetLocation `json:"location,omitempty" `

	/*
	   品牌     */
	Brand *AlibabaAlscUnionKbItemStoreDetailGetBrand `json:"brand,omitempty" `

	/*
	   门店所属行业分类     */
	Categories *[]AlibabaAlscUnionKbItemStoreDetailGetCategory `json:"categories,omitempty" `

	/*
	   门店资质     */
	Qualifications *[]AlibabaAlscUnionKbItemStoreDetailGetImageContent `json:"qualifications,omitempty" `
}

func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo) SetStoreId(v string) *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo {
	s.StoreId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo) SetName(v string) *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo {
	s.Name = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo) SetCover(v string) *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo {
	s.Cover = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo) SetMobiles(v []string) *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo {
	s.Mobiles = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo) SetBusiness(v AlibabaAlscUnionKbItemStoreDetailGetStoreBusiness) *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo {
	s.Business = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo) SetLocation(v AlibabaAlscUnionKbItemStoreDetailGetLocation) *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo {
	s.Location = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo) SetBrand(v AlibabaAlscUnionKbItemStoreDetailGetBrand) *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo {
	s.Brand = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo) SetCategories(v []AlibabaAlscUnionKbItemStoreDetailGetCategory) *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo {
	s.Categories = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo) SetQualifications(v []AlibabaAlscUnionKbItemStoreDetailGetImageContent) *AlibabaAlscUnionKbItemStoreDetailGetStoreBasicInfo {
	s.Qualifications = &v
	return s
}
