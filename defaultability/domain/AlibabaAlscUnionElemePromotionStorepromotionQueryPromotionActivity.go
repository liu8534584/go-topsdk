package domain

type AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity struct {
	/*
	   活动Id     */
	ActivityId *string `json:"activity_id,omitempty" `

	/*
	   营销计划服务费（分）     */
	ServiceFeeCent *int64 `json:"service_fee_cent,omitempty" `

	/*
	   奖励金红包面额（分）     */
	BonusCent *int64 `json:"bonus_cent,omitempty" `

	/*
	   活动的日库存     */
	DailyQuantity *int64 `json:"daily_quantity,omitempty" `

	/*
	   活动日剩余库存     */
	DailySellableQuantity *int64 `json:"daily_sellable_quantity,omitempty" `

	/*
	   起始时间（秒）     */
	StartTime *int64 `json:"start_time,omitempty" `

	/*
	   结束时间（秒）     */
	EndTime *int64 `json:"end_time,omitempty" `

	/*
	   奖励金门槛 (分)     */
	BountyMinLimitCent *int64 `json:"bounty_min_limit_cent,omitempty" `
}

func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity) SetActivityId(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity {
	s.ActivityId = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity) SetServiceFeeCent(v int64) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity {
	s.ServiceFeeCent = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity) SetBonusCent(v int64) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity {
	s.BonusCent = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity) SetDailyQuantity(v int64) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity {
	s.DailyQuantity = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity) SetDailySellableQuantity(v int64) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity {
	s.DailySellableQuantity = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity) SetStartTime(v int64) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity {
	s.StartTime = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity) SetEndTime(v int64) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity {
	s.EndTime = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity) SetBountyMinLimitCent(v int64) *AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionActivity {
	s.BountyMinLimitCent = &v
	return s
}
