package domain

type AlibabaAlscUnionKbItemStoreRelationQueryKbItemShopRelationRequest struct {
	/*
	   商品ID     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   城市ID     */
	CityId *string `json:"city_id,omitempty" `

	/*
	   业务类型（cps/cpa）     */
	BizType *string `json:"biz_type,omitempty" `
}

func (s *AlibabaAlscUnionKbItemStoreRelationQueryKbItemShopRelationRequest) SetItemId(v string) *AlibabaAlscUnionKbItemStoreRelationQueryKbItemShopRelationRequest {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreRelationQueryKbItemShopRelationRequest) SetCityId(v string) *AlibabaAlscUnionKbItemStoreRelationQueryKbItemShopRelationRequest {
	s.CityId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreRelationQueryKbItemShopRelationRequest) SetBizType(v string) *AlibabaAlscUnionKbItemStoreRelationQueryKbItemShopRelationRequest {
	s.BizType = &v
	return s
}
