package domain

type AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusiness struct {
	/*
	   营业状态（0-营业中，1-暂停营业，2-休息中，-1-关店）     */
	BusinessStatus *int64 `json:"business_status,omitempty" `

	/*
	   营业状态描述     */
	BusinessDesc *string `json:"business_desc,omitempty" `

	/*
	   营业时间     */
	BusinessTime *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusinessTime `json:"business_time,omitempty" `

	/*
	   店铺公告     */
	Promotion *string `json:"promotion,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusiness) SetBusinessStatus(v int64) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusiness {
	s.BusinessStatus = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusiness) SetBusinessDesc(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusiness {
	s.BusinessDesc = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusiness) SetBusinessTime(v AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusinessTime) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusiness {
	s.BusinessTime = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusiness) SetPromotion(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetStoreBusiness {
	s.Promotion = &v
	return s
}
