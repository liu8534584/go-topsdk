package domain

type AlibabaAlscUnionKbItemDetailGetUseNote struct {
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

func (s *AlibabaAlscUnionKbItemDetailGetUseNote) SetNeedReserve(v bool) *AlibabaAlscUnionKbItemDetailGetUseNote {
	s.NeedReserve = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetUseNote) SetReserveDesc(v string) *AlibabaAlscUnionKbItemDetailGetUseNote {
	s.ReserveDesc = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetUseNote) SetLimitUserNum(v bool) *AlibabaAlscUnionKbItemDetailGetUseNote {
	s.LimitUserNum = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetUseNote) SetUserNumLimited(v int64) *AlibabaAlscUnionKbItemDetailGetUseNote {
	s.UserNumLimited = &v
	return s
}
