package domain

type AlibabaAlscUnionKbBbtItemDetailGetItemTicket struct {
	/*
	   有效期     */
	TicketPeriod *AlibabaAlscUnionKbBbtItemDetailGetTicketPeriod `json:"ticket_period,omitempty" `

	/*
	   时间规则     */
	TicketTimeRules *[]AlibabaAlscUnionKbBbtItemDetailGetTicketTimeRule `json:"ticket_time_rules,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetItemTicket) SetTicketPeriod(v AlibabaAlscUnionKbBbtItemDetailGetTicketPeriod) *AlibabaAlscUnionKbBbtItemDetailGetItemTicket {
	s.TicketPeriod = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetItemTicket) SetTicketTimeRules(v []AlibabaAlscUnionKbBbtItemDetailGetTicketTimeRule) *AlibabaAlscUnionKbBbtItemDetailGetItemTicket {
	s.TicketTimeRules = &v
	return s
}
