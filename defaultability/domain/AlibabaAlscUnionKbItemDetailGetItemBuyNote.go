package domain

type AlibabaAlscUnionKbItemDetailGetItemBuyNote struct {
	/*
	   商家须知     */
	ShopInfo *AlibabaAlscUnionKbItemDetailGetShopInfo `json:"shop_info,omitempty" `

	/*
	   使用须知     */
	UseNote *AlibabaAlscUnionKbItemDetailGetUseNote `json:"use_note,omitempty" `

	/*
	   更多须知内容     */
	ExtraNotes *[]AlibabaAlscUnionKbItemDetailGetTextContent `json:"extra_notes,omitempty" `
}

func (s *AlibabaAlscUnionKbItemDetailGetItemBuyNote) SetShopInfo(v AlibabaAlscUnionKbItemDetailGetShopInfo) *AlibabaAlscUnionKbItemDetailGetItemBuyNote {
	s.ShopInfo = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetItemBuyNote) SetUseNote(v AlibabaAlscUnionKbItemDetailGetUseNote) *AlibabaAlscUnionKbItemDetailGetItemBuyNote {
	s.UseNote = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetItemBuyNote) SetExtraNotes(v []AlibabaAlscUnionKbItemDetailGetTextContent) *AlibabaAlscUnionKbItemDetailGetItemBuyNote {
	s.ExtraNotes = &v
	return s
}
