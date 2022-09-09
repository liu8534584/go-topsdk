package domain


type TaobaoTbkDgCpaReportSubDetailResult struct {
    /*
        策略名称     */
    CampaignCode  *string `json:"campaign_code,omitempty" `

    /*
        明细结果     */
    ResultList  *[]TaobaoTbkDgCpaReportSubDetailResultlist `json:"result_list,omitempty" `

    /*
        改策略字段说明     */
    DisplayList  *[]TaobaoTbkDgCpaReportSubDetailDisplaylist `json:"display_list,omitempty" `

}

func (s *TaobaoTbkDgCpaReportSubDetailResult) SetCampaignCode(v string) *TaobaoTbkDgCpaReportSubDetailResult {
    s.CampaignCode = &v
    return s
}
func (s *TaobaoTbkDgCpaReportSubDetailResult) SetResultList(v []TaobaoTbkDgCpaReportSubDetailResultlist) *TaobaoTbkDgCpaReportSubDetailResult {
    s.ResultList = &v
    return s
}
func (s *TaobaoTbkDgCpaReportSubDetailResult) SetDisplayList(v []TaobaoTbkDgCpaReportSubDetailDisplaylist) *TaobaoTbkDgCpaReportSubDetailResult {
    s.DisplayList = &v
    return s
}
