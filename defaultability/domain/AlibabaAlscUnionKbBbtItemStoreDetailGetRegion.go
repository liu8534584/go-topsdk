package domain

type AlibabaAlscUnionKbBbtItemStoreDetailGetRegion struct {
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

func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetRegion) SetProvinceId(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetRegion {
	s.ProvinceId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetRegion) SetProvinceName(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetRegion {
	s.ProvinceName = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetRegion) SetCityId(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetRegion {
	s.CityId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetRegion) SetCityName(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetRegion {
	s.CityName = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetRegion) SetDistrictId(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetRegion {
	s.DistrictId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetRegion) SetDistrictName(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetRegion {
	s.DistrictName = &v
	return s
}
