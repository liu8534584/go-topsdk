package domain

type AlibabaAlscUnionKbItemStoreDetailGetRegion struct {
	/*
	   省份编码     */
	ProvinceId *string `json:"province_id,omitempty" `

	/*
	   省份名称     */
	ProvinceName *string `json:"province_name,omitempty" `

	/*
	   城市编码     */
	CityId *string `json:"city_id,omitempty" `

	/*
	   城市名称     */
	CityName *string `json:"city_name,omitempty" `

	/*
	   行政区编码     */
	DistrictId *string `json:"district_id,omitempty" `

	/*
	   行政区名称     */
	DistrictName *string `json:"district_name,omitempty" `
}

func (s *AlibabaAlscUnionKbItemStoreDetailGetRegion) SetProvinceId(v string) *AlibabaAlscUnionKbItemStoreDetailGetRegion {
	s.ProvinceId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetRegion) SetProvinceName(v string) *AlibabaAlscUnionKbItemStoreDetailGetRegion {
	s.ProvinceName = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetRegion) SetCityId(v string) *AlibabaAlscUnionKbItemStoreDetailGetRegion {
	s.CityId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetRegion) SetCityName(v string) *AlibabaAlscUnionKbItemStoreDetailGetRegion {
	s.CityName = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetRegion) SetDistrictId(v string) *AlibabaAlscUnionKbItemStoreDetailGetRegion {
	s.DistrictId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetRegion) SetDistrictName(v string) *AlibabaAlscUnionKbItemStoreDetailGetRegion {
	s.DistrictName = &v
	return s
}
