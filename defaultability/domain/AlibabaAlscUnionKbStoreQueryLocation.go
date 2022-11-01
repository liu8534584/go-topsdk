package domain

type AlibabaAlscUnionKbStoreQueryLocation struct {
	/*
	   经度     */
	Longitude *string `json:"longitude,omitempty" `

	/*
	   纬度     */
	Latitude *string `json:"latitude,omitempty" `

	/*
	   地址     */
	Address *string `json:"address,omitempty" `
}

func (s *AlibabaAlscUnionKbStoreQueryLocation) SetLongitude(v string) *AlibabaAlscUnionKbStoreQueryLocation {
	s.Longitude = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryLocation) SetLatitude(v string) *AlibabaAlscUnionKbStoreQueryLocation {
	s.Latitude = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryLocation) SetAddress(v string) *AlibabaAlscUnionKbStoreQueryLocation {
	s.Address = &v
	return s
}
