package domain

type AlibabaAlscUnionKbItemStoreDetailGetStoreBusiness struct {
	/*
	   营业状态（0-营业中，1-暂停营业，2-休息中，-1-关店）     */
	BusinessStatus *int64 `json:"business_status,omitempty" `

	/*
	   营业状态描述     */
	BusinessDesc *string `json:"business_desc,omitempty" `

	/*
	   营业时间     */
	BusinessTime *AlibabaAlscUnionKbItemStoreDetailGetStoreBusinessTime `json:"business_time,omitempty" `

	/*
	   店铺公告     */
	Promotion *string `json:"promotion,omitempty" `
}

func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreBusiness) SetBusinessStatus(v int64) *AlibabaAlscUnionKbItemStoreDetailGetStoreBusiness {
	s.BusinessStatus = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreBusiness) SetBusinessDesc(v string) *AlibabaAlscUnionKbItemStoreDetailGetStoreBusiness {
	s.BusinessDesc = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreBusiness) SetBusinessTime(v AlibabaAlscUnionKbItemStoreDetailGetStoreBusinessTime) *AlibabaAlscUnionKbItemStoreDetailGetStoreBusiness {
	s.BusinessTime = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetStoreBusiness) SetPromotion(v string) *AlibabaAlscUnionKbItemStoreDetailGetStoreBusiness {
	s.Promotion = &v
	return s
}
