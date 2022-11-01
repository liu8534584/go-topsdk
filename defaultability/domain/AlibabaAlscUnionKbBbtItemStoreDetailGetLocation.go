package domain

type AlibabaAlscUnionKbBbtItemStoreDetailGetLocation struct {
	/*
	   地址     */
	Address *string `json:"address,omitempty" `

	/*
	   地址备注(如交通信息等)     */
	AddressMemo *string `json:"address_memo,omitempty" `

	/*
	   所属区域     */
	Region *AlibabaAlscUnionKbBbtItemStoreDetailGetRegion `json:"region,omitempty" `

	/*
	   经度     */
	Longitude *string `json:"longitude,omitempty" `

	/*
	   纬度     */
	Latitude *string `json:"latitude,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetLocation) SetAddress(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetLocation {
	s.Address = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetLocation) SetAddressMemo(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetLocation {
	s.AddressMemo = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetLocation) SetRegion(v AlibabaAlscUnionKbBbtItemStoreDetailGetRegion) *AlibabaAlscUnionKbBbtItemStoreDetailGetLocation {
	s.Region = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetLocation) SetLongitude(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetLocation {
	s.Longitude = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetLocation) SetLatitude(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetLocation {
	s.Latitude = &v
	return s
}
