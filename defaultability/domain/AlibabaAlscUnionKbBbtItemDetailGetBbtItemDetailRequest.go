package domain

type AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailRequest struct {
	/*
	   品ID     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   城市ID     */
	CityId *string `json:"city_id,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailRequest) SetItemId(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailRequest {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailRequest) SetCityId(v string) *AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailRequest {
	s.CityId = &v
	return s
}
