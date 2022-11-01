package domain

type AlibabaAlscUnionKbItemDetailGetTicketPeriod struct {
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

func (s *AlibabaAlscUnionKbItemDetailGetTicketPeriod) SetPeriodType(v string) *AlibabaAlscUnionKbItemDetailGetTicketPeriod {
	s.PeriodType = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetTicketPeriod) SetPeriod(v int64) *AlibabaAlscUnionKbItemDetailGetTicketPeriod {
	s.Period = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetTicketPeriod) SetNatureDay(v bool) *AlibabaAlscUnionKbItemDetailGetTicketPeriod {
	s.NatureDay = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetTicketPeriod) SetStartTime(v int64) *AlibabaAlscUnionKbItemDetailGetTicketPeriod {
	s.StartTime = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetTicketPeriod) SetEndTime(v int64) *AlibabaAlscUnionKbItemDetailGetTicketPeriod {
	s.EndTime = &v
	return s
}
