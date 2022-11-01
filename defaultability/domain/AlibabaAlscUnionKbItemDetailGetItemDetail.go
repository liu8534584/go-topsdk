package domain

type AlibabaAlscUnionKbItemDetailGetItemDetail struct {
	/*
	   内容详情     */
	ItemContent *AlibabaAlscUnionKbItemDetailGetItemContent `json:"item_content,omitempty" `

	/*
	   购买须知     */
	ItemBuyNote *AlibabaAlscUnionKbItemDetailGetItemBuyNote `json:"item_buy_note,omitempty" `

	/*
	   凭证     */
	ItemTicket *AlibabaAlscUnionKbItemDetailGetItemTicket `json:"item_ticket,omitempty" `
}

func (s *AlibabaAlscUnionKbItemDetailGetItemDetail) SetItemContent(v AlibabaAlscUnionKbItemDetailGetItemContent) *AlibabaAlscUnionKbItemDetailGetItemDetail {
	s.ItemContent = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetItemDetail) SetItemBuyNote(v AlibabaAlscUnionKbItemDetailGetItemBuyNote) *AlibabaAlscUnionKbItemDetailGetItemDetail {
	s.ItemBuyNote = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetItemDetail) SetItemTicket(v AlibabaAlscUnionKbItemDetailGetItemTicket) *AlibabaAlscUnionKbItemDetailGetItemDetail {
	s.ItemTicket = &v
	return s
}
