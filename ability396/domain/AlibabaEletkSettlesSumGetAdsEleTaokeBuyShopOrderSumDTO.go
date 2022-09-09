package domain


type AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO struct {
    /*
        总结算金额（单位分）     */
    SettleTotalAmount  *string `json:"settle_total_amount,omitempty" `

    /*
        新客复购数     */
    NewuserRebuyNums  *int64 `json:"newuser_rebuy_nums,omitempty" `

    /*
        新客首单数     */
    AccFstordNums  *int64 `json:"acc_fstord_nums,omitempty" `

    /*
        CPS结算订单gmv（单位分）     */
    CpsOrderGmv  *string `json:"cps_order_gmv,omitempty" `

    /*
        CPS结算订单数     */
    CpsOrderNums  *int64 `json:"cps_order_nums,omitempty" `

    /*
        渠道号     */
    TrackId  *int64 `json:"track_id,omitempty" `

}

func (s *AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO) SetSettleTotalAmount(v string) *AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO {
    s.SettleTotalAmount = &v
    return s
}
func (s *AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO) SetNewuserRebuyNums(v int64) *AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO {
    s.NewuserRebuyNums = &v
    return s
}
func (s *AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO) SetAccFstordNums(v int64) *AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO {
    s.AccFstordNums = &v
    return s
}
func (s *AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO) SetCpsOrderGmv(v string) *AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO {
    s.CpsOrderGmv = &v
    return s
}
func (s *AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO) SetCpsOrderNums(v int64) *AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO {
    s.CpsOrderNums = &v
    return s
}
func (s *AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO) SetTrackId(v int64) *AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO {
    s.TrackId = &v
    return s
}
