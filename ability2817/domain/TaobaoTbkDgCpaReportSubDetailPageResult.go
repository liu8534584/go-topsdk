package domain


type TaobaoTbkDgCpaReportSubDetailPageResult struct {
    /*
        结果集合     */
    Results  *[]TaobaoTbkDgCpaReportSubDetailResult `json:"results,omitempty" `

}

func (s *TaobaoTbkDgCpaReportSubDetailPageResult) SetResults(v []TaobaoTbkDgCpaReportSubDetailResult) *TaobaoTbkDgCpaReportSubDetailPageResult {
    s.Results = &v
    return s
}
