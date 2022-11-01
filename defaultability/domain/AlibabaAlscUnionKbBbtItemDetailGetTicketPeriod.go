package domain

type AlibabaAlscUnionKbBbtItemDetailGetTicketPeriod struct {
	/*
	   有效周期类型     */
	PeriodType *string `json:"period_type,omitempty" `

	/*
	   相对有效期，单位：天     */
	Period *int64 `json:"period,omitempty" `

	/*
	   是否自然日     */
	NatureDay *bool `json:"nature_day,omitempty" `

	/*
	   起始时间（秒）     */
	StartTime *int64 `json:"start_time,omitempty" `

	/*
	   终止时间（秒）     */
	EndTime *int64 `json:"end_time,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetTicketPeriod) SetPeriodType(v string) *AlibabaAlscUnionKbBbtItemDetailGetTicketPeriod {
	s.PeriodType = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetTicketPeriod) SetPeriod(v int64) *AlibabaAlscUnionKbBbtItemDetailGetTicketPeriod {
	s.Period = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetTicketPeriod) SetNatureDay(v bool) *AlibabaAlscUnionKbBbtItemDetailGetTicketPeriod {
	s.NatureDay = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetTicketPeriod) SetStartTime(v int64) *AlibabaAlscUnionKbBbtItemDetailGetTicketPeriod {
	s.StartTime = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetTicketPeriod) SetEndTime(v int64) *AlibabaAlscUnionKbBbtItemDetailGetTicketPeriod {
	s.EndTime = &v
	return s
}
