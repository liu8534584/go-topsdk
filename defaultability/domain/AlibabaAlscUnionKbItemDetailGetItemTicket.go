package domain

type AlibabaAlscUnionKbItemDetailGetItemTicket struct {
	/*
	   有效期     */
	TicketPeriod *AlibabaAlscUnionKbItemDetailGetTicketPeriod `json:"ticket_period,omitempty" `

	/*
	   时间规则     */
	TicketTimeRules *[]AlibabaAlscUnionKbItemDetailGetTicketTimeRule `json:"ticket_time_rules,omitempty" `
}

func (s *AlibabaAlscUnionKbItemDetailGetItemTicket) SetTicketPeriod(v AlibabaAlscUnionKbItemDetailGetTicketPeriod) *AlibabaAlscUnionKbItemDetailGetItemTicket {
	s.TicketPeriod = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetItemTicket) SetTicketTimeRules(v []AlibabaAlscUnionKbItemDetailGetTicketTimeRule) *AlibabaAlscUnionKbItemDetailGetItemTicket {
	s.TicketTimeRules = &v
	return s
}
