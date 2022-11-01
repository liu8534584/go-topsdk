package domain

import (
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionElemePromotionStorepromotionQueryOverlayCouponDTO struct {
	/*
	   券id     */
	TemplateId *int64 `json:"template_id,omitempty" `

	/*
	   券金额，单位元     */
	Amount *string `json:"amount,omitempty" `

	/*
	   活动有效期天（发出去的有效期）     */
	ValidPeriod *int64 `json:"valid_period,omitempty" `

	/*
	   活动开始时间     */
	StartTime *util.LocalTime `json:"start_time,omitempty" `

	/*
	   活动结束时间     */
	EndTime *util.LocalTime `json:"end_time,omitempty" `
}

func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryOverlayCouponDTO) SetTemplateId(v int64) *AlibabaAlscUnionElemePromotionStorepromotionQueryOverlayCouponDTO {
	s.TemplateId = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryOverlayCouponDTO) SetAmount(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryOverlayCouponDTO {
	s.Amount = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryOverlayCouponDTO) SetValidPeriod(v int64) *AlibabaAlscUnionElemePromotionStorepromotionQueryOverlayCouponDTO {
	s.ValidPeriod = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryOverlayCouponDTO) SetStartTime(v util.LocalTime) *AlibabaAlscUnionElemePromotionStorepromotionQueryOverlayCouponDTO {
	s.StartTime = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryOverlayCouponDTO) SetEndTime(v util.LocalTime) *AlibabaAlscUnionElemePromotionStorepromotionQueryOverlayCouponDTO {
	s.EndTime = &v
	return s
}
