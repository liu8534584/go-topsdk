package domain


type TaobaoTbkDgCpaReportDetailCpaReportDto struct {
    /*
        数据更新时间     */
    TheDate  *string `json:"the_date,omitempty" `

    /*
        结算维度奖励总金额     */
    SettleTotal  *string `json:"settle_total,omitempty" `

    /*
        预估维度奖励总金额     */
    PreTotal  *string `json:"pre_total,omitempty" `

    /*
        结果数据     */
    CampaignReportdtoList  *[]TaobaoTbkDgCpaReportDetailCampaignreportdtolist `json:"campaign_reportdto_list,omitempty" `

}

func (s *TaobaoTbkDgCpaReportDetailCpaReportDto) SetTheDate(v string) *TaobaoTbkDgCpaReportDetailCpaReportDto {
    s.TheDate = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailCpaReportDto) SetSettleTotal(v string) *TaobaoTbkDgCpaReportDetailCpaReportDto {
    s.SettleTotal = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailCpaReportDto) SetPreTotal(v string) *TaobaoTbkDgCpaReportDetailCpaReportDto {
    s.PreTotal = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailCpaReportDto) SetCampaignReportdtoList(v []TaobaoTbkDgCpaReportDetailCampaignreportdtolist) *TaobaoTbkDgCpaReportDetailCpaReportDto {
    s.CampaignReportdtoList = &v
    return s
}
