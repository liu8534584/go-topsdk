package domain


type TaobaoTbkDgRidReportGetData struct {
    /*
        非自购订单购买用户数(选定时间内去重)     */
    PromoUv  *string `json:"promo_uv,omitempty" `

    /*
        非自购订单金额     */
    PromoAlipayAmt  *string `json:"promo_alipay_amt,omitempty" `

    /*
        自购订单购买用户数(选定时间内去重)     */
    IndividualUv  *string `json:"individual_uv,omitempty" `

    /*
        非自购订单笔数     */
    PromoAlipayNum  *string `json:"promo_alipay_num,omitempty" `

    /*
        购买用户数(选定时间内去重)     */
    PurchaseUv  *string `json:"purchase_uv,omitempty" `

    /*
        购买订单笔数     */
    PurchaseAlipayNum  *string `json:"purchase_alipay_num,omitempty" `

    /*
        自购订单笔数     */
    IndividualAlipayNum  *string `json:"individual_alipay_num,omitempty" `

    /*
        购买订单金额     */
    PurchaseAlipayAmt  *string `json:"purchase_alipay_amt,omitempty" `

    /*
        自购订单金额     */
    IndividualAlipayAmt  *string `json:"individual_alipay_amt,omitempty" `

}

func (s *TaobaoTbkDgRidReportGetData) SetPromoUv(v string) *TaobaoTbkDgRidReportGetData {
    s.PromoUv = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetData) SetPromoAlipayAmt(v string) *TaobaoTbkDgRidReportGetData {
    s.PromoAlipayAmt = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetData) SetIndividualUv(v string) *TaobaoTbkDgRidReportGetData {
    s.IndividualUv = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetData) SetPromoAlipayNum(v string) *TaobaoTbkDgRidReportGetData {
    s.PromoAlipayNum = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetData) SetPurchaseUv(v string) *TaobaoTbkDgRidReportGetData {
    s.PurchaseUv = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetData) SetPurchaseAlipayNum(v string) *TaobaoTbkDgRidReportGetData {
    s.PurchaseAlipayNum = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetData) SetIndividualAlipayNum(v string) *TaobaoTbkDgRidReportGetData {
    s.IndividualAlipayNum = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetData) SetPurchaseAlipayAmt(v string) *TaobaoTbkDgRidReportGetData {
    s.PurchaseAlipayAmt = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetData) SetIndividualAlipayAmt(v string) *TaobaoTbkDgRidReportGetData {
    s.IndividualAlipayAmt = &v
    return s
}
