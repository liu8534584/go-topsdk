package domain

type AlibabaAlscUnionKbItemStoreDetailGetLocation struct {
	/*
	   地址     */
	Address *string `json:"address,omitempty" `

	/*
	   地址备注(如交通信息等)     */
	AddressMemo *string `json:"address_memo,omitempty" `

	/*
	   所属区域     */
	Region *AlibabaAlscUnionKbItemStoreDetailGetRegion `json:"region,omitempty" `

	/*
	   经度     */
	Longitude *string `json:"longitude,omitempty" `

	/*
	   纬度     */
	Latitude *string `json:"latitude,omitempty" `
}

func (s *AlibabaAlscUnionKbItemStoreDetailGetLocation) SetAddress(v string) *AlibabaAlscUnionKbItemStoreDetailGetLocation {
	s.Address = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetLocation) SetAddressMemo(v string) *AlibabaAlscUnionKbItemStoreDetailGetLocation {
	s.AddressMemo = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetLocation) SetRegion(v AlibabaAlscUnionKbItemStoreDetailGetRegion) *AlibabaAlscUnionKbItemStoreDetailGetLocation {
	s.Region = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetLocation) SetLongitude(v string) *AlibabaAlscUnionKbItemStoreDetailGetLocation {
	s.Longitude = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetLocation) SetLatitude(v string) *AlibabaAlscUnionKbItemStoreDetailGetLocation {
	s.Latitude = &v
	return s
}
