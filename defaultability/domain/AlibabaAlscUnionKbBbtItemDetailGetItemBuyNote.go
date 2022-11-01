package domain

type AlibabaAlscUnionKbBbtItemDetailGetItemBuyNote struct {
	/*
	   商家须知     */
	ShopInfo *AlibabaAlscUnionKbBbtItemDetailGetShopInfo `json:"shop_info,omitempty" `

	/*
	   使用须知     */
	UseNote *AlibabaAlscUnionKbBbtItemDetailGetUseNote `json:"use_note,omitempty" `

	/*
	   更多须知内容     */
	ExtraNotes *[]AlibabaAlscUnionKbBbtItemDetailGetTextContent `json:"extra_notes,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetItemBuyNote) SetShopInfo(v AlibabaAlscUnionKbBbtItemDetailGetShopInfo) *AlibabaAlscUnionKbBbtItemDetailGetItemBuyNote {
	s.ShopInfo = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetItemBuyNote) SetUseNote(v AlibabaAlscUnionKbBbtItemDetailGetUseNote) *AlibabaAlscUnionKbBbtItemDetailGetItemBuyNote {
	s.UseNote = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetItemBuyNote) SetExtraNotes(v []AlibabaAlscUnionKbBbtItemDetailGetTextContent) *AlibabaAlscUnionKbBbtItemDetailGetItemBuyNote {
	s.ExtraNotes = &v
	return s
}
