package domain


type AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO struct {
    /*
        是否使用淘客红包     */
    IsUseTaokeCoupon14d  *int64 `json:"is_use_taoke_coupon14d,omitempty" `

    /*
        是否新客复购单     */
    IsNewuserRebuy14d  *int64 `json:"is_newuser_rebuy14d,omitempty" `

    /*
        是否新客首单     */
    IsAccFstord7d  *int64 `json:"is_acc_fstord7d,omitempty" `

    /*
        是否有在会场内该门店的点击记录（7天内订单，超过7天的则不计入）     */
    IsShopClk7d  *int64 `json:"is_shop_clk7d,omitempty" `

    /*
        订单id     */
    OrderId  *string `json:"order_id,omitempty" `

    /*
        渠道号     */
    TrackId  *int64 `json:"track_id,omitempty" `

    /*
        是否CPS订单     */
    CpsOrderFlag  *bool `json:"cps_order_flag,omitempty" `

    /*
        是否领券后7天内首单     */
    IsCouponDrawAccFstord7d  *int64 `json:"is_coupon_draw_acc_fstord7d,omitempty" `

    /*
        rid     */
    Rid  *string `json:"rid,omitempty" `

    /*
        支付金额     */
    PayAmt  *string `json:"pay_amt,omitempty" `

    /*
        创建时间     */
    CreateTime  *string `json:"create_time,omitempty" `

    /*
        城市名称     */
    OrderCityName  *string `json:"order_city_name,omitempty" `

    /*
        track_pid     */
    TrackPid  *string `json:"track_pid,omitempty" `

    /*
        是否不结算订单     */
    AbnorFlag  *bool `json:"abnor_flag,omitempty" `

    /*
        是否拉新订单     */
    NewOrderFlag  *bool `json:"new_order_flag,omitempty" `

    /*
        扩展字段（order_type：new_retail（新零售）、take_out（餐饮）、 coupon（卡券）； use_order_id：卡券对应的核销外卖订单的id； use_order_time：卡券对应的核销外卖订单的订单日期）     */
    ExtraInfo  *string `json:"extra_info,omitempty" `

}

func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetIsUseTaokeCoupon14d(v int64) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.IsUseTaokeCoupon14d = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetIsNewuserRebuy14d(v int64) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.IsNewuserRebuy14d = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetIsAccFstord7d(v int64) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.IsAccFstord7d = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetIsShopClk7d(v int64) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.IsShopClk7d = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetOrderId(v string) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.OrderId = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetTrackId(v int64) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.TrackId = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetCpsOrderFlag(v bool) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.CpsOrderFlag = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetIsCouponDrawAccFstord7d(v int64) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.IsCouponDrawAccFstord7d = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetRid(v string) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.Rid = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetPayAmt(v string) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.PayAmt = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetCreateTime(v string) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.CreateTime = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetOrderCityName(v string) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.OrderCityName = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetTrackPid(v string) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.TrackPid = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetAbnorFlag(v bool) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.AbnorFlag = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetNewOrderFlag(v bool) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.NewOrderFlag = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) SetExtraInfo(v string) *AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO {
    s.ExtraInfo = &v
    return s
}
