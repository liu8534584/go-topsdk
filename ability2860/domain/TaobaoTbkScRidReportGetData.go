package domain


type TaobaoTbkScRidReportGetData struct {
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

func (s *TaobaoTbkScRidReportGetData) SetPromoUv(v string) *TaobaoTbkScRidReportGetData {
    s.PromoUv = &v
    return s
}
func (s *TaobaoTbkScRidReportGetData) SetPromoAlipayAmt(v string) *TaobaoTbkScRidReportGetData {
    s.PromoAlipayAmt = &v
    return s
}
func (s *TaobaoTbkScRidReportGetData) SetIndividualUv(v string) *TaobaoTbkScRidReportGetData {
    s.IndividualUv = &v
    return s
}
func (s *TaobaoTbkScRidReportGetData) SetPromoAlipayNum(v string) *TaobaoTbkScRidReportGetData {
    s.PromoAlipayNum = &v
    return s
}
func (s *TaobaoTbkScRidReportGetData) SetPurchaseUv(v string) *TaobaoTbkScRidReportGetData {
    s.PurchaseUv = &v
    return s
}
func (s *TaobaoTbkScRidReportGetData) SetPurchaseAlipayNum(v string) *TaobaoTbkScRidReportGetData {
    s.PurchaseAlipayNum = &v
    return s
}
func (s *TaobaoTbkScRidReportGetData) SetIndividualAlipayNum(v string) *TaobaoTbkScRidReportGetData {
    s.IndividualAlipayNum = &v
    return s
}
func (s *TaobaoTbkScRidReportGetData) SetPurchaseAlipayAmt(v string) *TaobaoTbkScRidReportGetData {
    s.PurchaseAlipayAmt = &v
    return s
}
func (s *TaobaoTbkScRidReportGetData) SetIndividualAlipayAmt(v string) *TaobaoTbkScRidReportGetData {
    s.IndividualAlipayAmt = &v
    return s
}
