package domain

type AlibabaAlscUnionKbBbtItemDetailGetItemDetail struct {
	/*
	   内容详情     */
	ItemContent *AlibabaAlscUnionKbBbtItemDetailGetItemContent `json:"item_content,omitempty" `

	/*
	   购买须知     */
	ItemBuyNote *AlibabaAlscUnionKbBbtItemDetailGetItemBuyNote `json:"item_buy_note,omitempty" `

	/*
	   凭证     */
	ItemTicket *AlibabaAlscUnionKbBbtItemDetailGetItemTicket `json:"item_ticket,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetItemDetail) SetItemContent(v AlibabaAlscUnionKbBbtItemDetailGetItemContent) *AlibabaAlscUnionKbBbtItemDetailGetItemDetail {
	s.ItemContent = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetItemDetail) SetItemBuyNote(v AlibabaAlscUnionKbBbtItemDetailGetItemBuyNote) *AlibabaAlscUnionKbBbtItemDetailGetItemDetail {
	s.ItemBuyNote = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetItemDetail) SetItemTicket(v AlibabaAlscUnionKbBbtItemDetailGetItemTicket) *AlibabaAlscUnionKbBbtItemDetailGetItemDetail {
	s.ItemTicket = &v
	return s
}
