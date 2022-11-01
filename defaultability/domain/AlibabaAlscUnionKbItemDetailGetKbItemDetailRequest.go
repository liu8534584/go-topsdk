package domain

type AlibabaAlscUnionKbItemDetailGetKbItemDetailRequest struct {
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

func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailRequest) SetItemId(v string) *AlibabaAlscUnionKbItemDetailGetKbItemDetailRequest {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailRequest) SetCityId(v string) *AlibabaAlscUnionKbItemDetailGetKbItemDetailRequest {
	s.CityId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetKbItemDetailRequest) SetBizType(v string) *AlibabaAlscUnionKbItemDetailGetKbItemDetailRequest {
	s.BizType = &v
	return s
}
