package domain

type AlibabaAlscUnionKbBbtItemStoreRelationQueryBbtItemShopRelationRequest struct {
	/*
	   商品ID     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   城市ID     */
	CityId *string `json:"city_id,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemStoreRelationQueryBbtItemShopRelationRequest) SetItemId(v string) *AlibabaAlscUnionKbBbtItemStoreRelationQueryBbtItemShopRelationRequest {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreRelationQueryBbtItemShopRelationRequest) SetCityId(v string) *AlibabaAlscUnionKbBbtItemStoreRelationQueryBbtItemShopRelationRequest {
	s.CityId = &v
	return s
}
