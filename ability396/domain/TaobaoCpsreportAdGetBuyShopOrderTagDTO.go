package domain


type TaobaoCpsreportAdGetBuyShopOrderTagDTO struct {
    /*
        渠道号     */
    ChannelNo  *string `json:"channel_no,omitempty" `

    /*
        track_pid     */
    TrackPid  *string `json:"track_pid,omitempty" `

    /*
        rid     */
    Rid  *string `json:"rid,omitempty" `

    /*
        订单日期     */
    CreateTime  *string `json:"create_time,omitempty" `

    /*
        是否结佣     */
    Commissed  *bool `json:"commissed,omitempty" `

    /*
        是否领券后7天内首单     */
    IsCouponDrawAccFstord7d  *int64 `json:"is_coupon_draw_acc_fstord7d,omitempty" `

    /*
        是否有在会场内该门店的点击记录（7天内订单，超过7天的则不计入）     */
    IsShopClk7d  *int64 `json:"is_shop_clk7d,omitempty" `

    /*
        是否使用饿了么渠道优惠券     */
    IsUseTaokeCoupon14d  *int64 `json:"is_use_taoke_coupon14d,omitempty" `

    /*
        是否新客复购单     */
    IsNewuserRebuy14d  *int64 `json:"is_newuser_rebuy14d,omitempty" `

    /*
        是否新客首单     */
    IsAccFstord7d  *int64 `json:"is_acc_fstord7d,omitempty" `

    /*
        券核销金额     */
    CouponAmt  *string `json:"coupon_amt,omitempty" `

    /*
        下单的城市名称     */
    OrderCityName  *string `json:"order_city_name,omitempty" `

    /*
        支付金额     */
    PayAmt  *string `json:"pay_amt,omitempty" `

    /*
        订单id     */
    OrderId  *string `json:"order_id,omitempty" `

    /*
        扩展字段（order_type：new_retail（新零售）、take_out（餐饮）、 coupon（卡券）； use_order_id：卡券对应的核销外卖订单的id； use_order_time：卡券对应的核销外卖订单的订单日期）     */
    ExtraInfo  *string `json:"extra_info,omitempty" `

}

func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetChannelNo(v string) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.ChannelNo = &v
    return s
}
func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetTrackPid(v string) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.TrackPid = &v
    return s
}
func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetRid(v string) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.Rid = &v
    return s
}
func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetCreateTime(v string) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.CreateTime = &v
    return s
}
func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetCommissed(v bool) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.Commissed = &v
    return s
}
func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetIsCouponDrawAccFstord7d(v int64) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.IsCouponDrawAccFstord7d = &v
    return s
}
func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetIsShopClk7d(v int64) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.IsShopClk7d = &v
    return s
}
func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetIsUseTaokeCoupon14d(v int64) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.IsUseTaokeCoupon14d = &v
    return s
}
func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetIsNewuserRebuy14d(v int64) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.IsNewuserRebuy14d = &v
    return s
}
func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetIsAccFstord7d(v int64) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.IsAccFstord7d = &v
    return s
}
func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetCouponAmt(v string) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.CouponAmt = &v
    return s
}
func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetOrderCityName(v string) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.OrderCityName = &v
    return s
}
func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetPayAmt(v string) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.PayAmt = &v
    return s
}
func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetOrderId(v string) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.OrderId = &v
    return s
}
func (s *TaobaoCpsreportAdGetBuyShopOrderTagDTO) SetExtraInfo(v string) *TaobaoCpsreportAdGetBuyShopOrderTagDTO {
    s.ExtraInfo = &v
    return s
}
