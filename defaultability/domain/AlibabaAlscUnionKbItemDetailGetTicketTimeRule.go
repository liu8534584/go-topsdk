package domain

type AlibabaAlscUnionKbItemDetailGetTicketTimeRule struct {
	/*
	   时间规则生效模式（"IN":"包含","EX":"排除）     */
	RuleApplyMode *string `json:"rule_apply_mode,omitempty" `

	/*
	   时分维度的规则（10:00~12:00）     */
	HourMinRules *[]string `json:"hour_min_rules,omitempty" `

	/*
	   星期维度的规则（周一到周日分别是：1~7）     */
	WeekRules *[]string `json:"week_rules,omitempty" `

	/*
	   日维度的规则:某天到某天     */
	DateRules *[]string `json:"date_rules,omitempty" `
}

func (s *AlibabaAlscUnionKbItemDetailGetTicketTimeRule) SetRuleApplyMode(v string) *AlibabaAlscUnionKbItemDetailGetTicketTimeRule {
	s.RuleApplyMode = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetTicketTimeRule) SetHourMinRules(v []string) *AlibabaAlscUnionKbItemDetailGetTicketTimeRule {
	s.HourMinRules = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetTicketTimeRule) SetWeekRules(v []string) *AlibabaAlscUnionKbItemDetailGetTicketTimeRule {
	s.WeekRules = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetTicketTimeRule) SetDateRules(v []string) *AlibabaAlscUnionKbItemDetailGetTicketTimeRule {
	s.DateRules = &v
	return s
}
