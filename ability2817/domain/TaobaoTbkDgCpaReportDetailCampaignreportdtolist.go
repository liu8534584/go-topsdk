package domain


type TaobaoTbkDgCpaReportDetailCampaignreportdtolist struct {
    /*
        策略名称     */
    CampaignCode  *string `json:"campaign_code,omitempty" `

    /*
        该策略下结算金额     */
    SettleFee  *string `json:"settle_fee,omitempty" `

    /*
        该策略下预估金额     */
    PreFee  *string `json:"pre_fee,omitempty" `

    /*
        结果数据     */
    ResultList  *[]TaobaoTbkDgCpaReportDetailResultlist `json:"result_list,omitempty" `

    /*
        字段说明     */
    DisplayList  *[]TaobaoTbkDgCpaReportDetailDisplaylist `json:"display_list,omitempty" `

}

func (s *TaobaoTbkDgCpaReportDetailCampaignreportdtolist) SetCampaignCode(v string) *TaobaoTbkDgCpaReportDetailCampaignreportdtolist {
    s.CampaignCode = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailCampaignreportdtolist) SetSettleFee(v string) *TaobaoTbkDgCpaReportDetailCampaignreportdtolist {
    s.SettleFee = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailCampaignreportdtolist) SetPreFee(v string) *TaobaoTbkDgCpaReportDetailCampaignreportdtolist {
    s.PreFee = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailCampaignreportdtolist) SetResultList(v []TaobaoTbkDgCpaReportDetailResultlist) *TaobaoTbkDgCpaReportDetailCampaignreportdtolist {
    s.ResultList = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailCampaignreportdtolist) SetDisplayList(v []TaobaoTbkDgCpaReportDetailDisplaylist) *TaobaoTbkDgCpaReportDetailCampaignreportdtolist {
    s.DisplayList = &v
    return s
}
