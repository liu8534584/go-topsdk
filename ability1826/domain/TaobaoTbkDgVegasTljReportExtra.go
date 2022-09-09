package domain


type TaobaoTbkDgVegasTljReportExtra struct {
    /*
        领取率，领取淘礼金个数/创建淘礼金个数     */
    GetRate  *string `json:"get_rate,omitempty" `

    /*
        使用率，使用淘礼金个数/领取淘礼金个数     */
    UseRate  *string `json:"use_rate,omitempty" `

    /*
        引导付款笔数，同一个红包，若因消费者付款使用后取消订单或退货退款，产生二次红包使用行为，引导付款笔数也会记录两单     */
    AlipayNum  *int64 `json:"alipay_num,omitempty" `

    /*
        引导付款金额，同一个红包，若因消费者付款使用后取消订单或退货退款，产生二次红包使用行为，引导付款笔数也会记录两单     */
    AlipayAmt  *string `json:"alipay_amt,omitempty" `

    /*
        付款佣金，下单付款，产生二次红包使用行为，会记录2次     */
    PrePubShareFeeForDisp  *string `json:"pre_pub_share_fee_for_disp,omitempty" `

    /*
        结算佣金，确认收货，产生二次红包使用行为，会记录2次     */
    CmSettleAmt  *string `json:"cm_settle_amt,omitempty" `

    /*
        领取淘礼金个数     */
    WinPv  *int64 `json:"win_pv,omitempty" `

    /*
        领取淘礼金金额     */
    WinSumAmt  *string `json:"win_sum_amt,omitempty" `

    /*
        未领取淘礼金个数，过了领取有效期或者暂停后没有被领取的红包个数     */
    RemainingNum  *int64 `json:"remaining_num,omitempty" `

    /*
        未领取金额，过了领取有效期或者暂停后没有被领取的红包金额     */
    RemainingAmt  *string `json:"remaining_amt,omitempty" `

    /*
        使用淘礼金个数，同一个红包，若因消费者付款使用后取消订单或退货退款，产生二次红包使用行为，使用淘礼金个数一天内会去重，所以相当于不会重记     */
    UseNum  *int64 `json:"use_num,omitempty" `

    /*
        使用淘礼金金额，若红包被重复使用（1)淘礼金红包被拆分，并且产生部分退款，会保留部分退款的订单淘礼金金额；若全部退款，会保留订单全部淘礼金金额），因此，已使用金额可能大于消费者实际使用金额（使用红包后，若产生红包退回后再次使用，已使用金额会被二次计算，已使用数量不会）     */
    UseSumAmt  *string `json:"use_sum_amt,omitempty" `

    /*
        退款淘礼金个数，红包使用后，由于订单取消，退货退款等行为带来的淘礼金红包退回数量，退款红包数量单日内不重复计算，跨天重复计算     */
    RefundNum  *int64 `json:"refund_num,omitempty" `

    /*
        退款淘礼金金额，红包使用后，由于订单取消，退货退款等行为行为带来的淘礼金红包退回数量 （退款红包若产生多次使用，退款红包金额会被多次计算，退款红包数量单日内不重复计算，跨天重复计算）     */
    RefundSumAmt  *string `json:"refund_sum_amt,omitempty" `

}

func (s *TaobaoTbkDgVegasTljReportExtra) SetGetRate(v string) *TaobaoTbkDgVegasTljReportExtra {
    s.GetRate = &v
    return s
}
func (s *TaobaoTbkDgVegasTljReportExtra) SetUseRate(v string) *TaobaoTbkDgVegasTljReportExtra {
    s.UseRate = &v
    return s
}
func (s *TaobaoTbkDgVegasTljReportExtra) SetAlipayNum(v int64) *TaobaoTbkDgVegasTljReportExtra {
    s.AlipayNum = &v
    return s
}
func (s *TaobaoTbkDgVegasTljReportExtra) SetAlipayAmt(v string) *TaobaoTbkDgVegasTljReportExtra {
    s.AlipayAmt = &v
    return s
}
func (s *TaobaoTbkDgVegasTljReportExtra) SetPrePubShareFeeForDisp(v string) *TaobaoTbkDgVegasTljReportExtra {
    s.PrePubShareFeeForDisp = &v
    return s
}
func (s *TaobaoTbkDgVegasTljReportExtra) SetCmSettleAmt(v string) *TaobaoTbkDgVegasTljReportExtra {
    s.CmSettleAmt = &v
    return s
}
func (s *TaobaoTbkDgVegasTljReportExtra) SetWinPv(v int64) *TaobaoTbkDgVegasTljReportExtra {
    s.WinPv = &v
    return s
}
func (s *TaobaoTbkDgVegasTljReportExtra) SetWinSumAmt(v string) *TaobaoTbkDgVegasTljReportExtra {
    s.WinSumAmt = &v
    return s
}
func (s *TaobaoTbkDgVegasTljReportExtra) SetRemainingNum(v int64) *TaobaoTbkDgVegasTljReportExtra {
    s.RemainingNum = &v
    return s
}
func (s *TaobaoTbkDgVegasTljReportExtra) SetRemainingAmt(v string) *TaobaoTbkDgVegasTljReportExtra {
    s.RemainingAmt = &v
    return s
}
func (s *TaobaoTbkDgVegasTljReportExtra) SetUseNum(v int64) *TaobaoTbkDgVegasTljReportExtra {
    s.UseNum = &v
    return s
}
func (s *TaobaoTbkDgVegasTljReportExtra) SetUseSumAmt(v string) *TaobaoTbkDgVegasTljReportExtra {
    s.UseSumAmt = &v
    return s
}
func (s *TaobaoTbkDgVegasTljReportExtra) SetRefundNum(v int64) *TaobaoTbkDgVegasTljReportExtra {
    s.RefundNum = &v
    return s
}
func (s *TaobaoTbkDgVegasTljReportExtra) SetRefundSumAmt(v string) *TaobaoTbkDgVegasTljReportExtra {
    s.RefundSumAmt = &v
    return s
}
