package domain

type AlibabaAlscUnionKbBbtItemDetailGetUseNote struct {
	/*
	   需要预约     */
	NeedReserve *bool `json:"need_reserve,omitempty" `

	/*
	   预约说明     */
	ReserveDesc *string `json:"reserve_desc,omitempty" `

	/*
	   是否限制使用用户数     */
	LimitUserNum *bool `json:"limit_user_num,omitempty" `

	/*
	   限制多少人使用     */
	UserNumLimited *int64 `json:"user_num_limited,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetUseNote) SetNeedReserve(v bool) *AlibabaAlscUnionKbBbtItemDetailGetUseNote {
	s.NeedReserve = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetUseNote) SetReserveDesc(v string) *AlibabaAlscUnionKbBbtItemDetailGetUseNote {
	s.ReserveDesc = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetUseNote) SetLimitUserNum(v bool) *AlibabaAlscUnionKbBbtItemDetailGetUseNote {
	s.LimitUserNum = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetUseNote) SetUserNumLimited(v int64) *AlibabaAlscUnionKbBbtItemDetailGetUseNote {
	s.UserNumLimited = &v
	return s
}
