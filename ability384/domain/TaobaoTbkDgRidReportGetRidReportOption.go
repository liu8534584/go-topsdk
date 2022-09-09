package domain


type TaobaoTbkDgRidReportGetRidReportOption struct {
    /*
        请求日期（天维度）     */
    BizDate  *string `json:"biz_date,omitempty" `

    /*
        渠道关系id     */
    RelationId  *string `json:"relation_id,omitempty" `

    /*
        请求日期（月维度，暂不支持）     */
    BizMonth  *string `json:"biz_month,omitempty" `

}

func (s *TaobaoTbkDgRidReportGetRidReportOption) SetBizDate(v string) *TaobaoTbkDgRidReportGetRidReportOption {
    s.BizDate = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetRidReportOption) SetRelationId(v string) *TaobaoTbkDgRidReportGetRidReportOption {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetRidReportOption) SetBizMonth(v string) *TaobaoTbkDgRidReportGetRidReportOption {
    s.BizMonth = &v
    return s
}
