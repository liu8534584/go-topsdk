package domain

type TaobaoTbkScOrderDetailsTemporaryGetServiceFeeDto struct {
	/*
	   专项服务费率     */
	ShareRelativeRate *string `json:"share_relative_rate,omitempty" `

	/*
	   结算专项服务费     */
	ShareFee *string `json:"share_fee,omitempty" `

	/*
	   预估专项服务费     */
	SharePreFee *string `json:"share_pre_fee,omitempty" `

	/*
	   专项服务费来源，122-渠道     */
	TkShareRoleType *int32 `json:"tk_share_role_type,omitempty" `
}

func (s *TaobaoTbkScOrderDetailsTemporaryGetServiceFeeDto) SetShareRelativeRate(v string) *TaobaoTbkScOrderDetailsTemporaryGetServiceFeeDto {
	s.ShareRelativeRate = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetServiceFeeDto) SetShareFee(v string) *TaobaoTbkScOrderDetailsTemporaryGetServiceFeeDto {
	s.ShareFee = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetServiceFeeDto) SetSharePreFee(v string) *TaobaoTbkScOrderDetailsTemporaryGetServiceFeeDto {
	s.SharePreFee = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetServiceFeeDto) SetTkShareRoleType(v int32) *TaobaoTbkScOrderDetailsTemporaryGetServiceFeeDto {
	s.TkShareRoleType = &v
	return s
}
