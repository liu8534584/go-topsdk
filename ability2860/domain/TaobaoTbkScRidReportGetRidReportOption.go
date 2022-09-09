package domain


type TaobaoTbkScRidReportGetRidReportOption struct {
    /*
        请求日期（天维度）     */
    BizDate  *string `json:"biz_date,omitempty" `

    /*
        关系id     */
    RelationId  *string `json:"relation_id,omitempty" `

    /*
        请求日期（月维度，暂不支持）     */
    BizMonth  *string `json:"biz_month,omitempty" `

    /*
        mm_xxx_22_xxx三段式的第二段数字     */
    SiteId  *string `json:"site_id,omitempty" `

}

func (s *TaobaoTbkScRidReportGetRidReportOption) SetBizDate(v string) *TaobaoTbkScRidReportGetRidReportOption {
    s.BizDate = &v
    return s
}
func (s *TaobaoTbkScRidReportGetRidReportOption) SetRelationId(v string) *TaobaoTbkScRidReportGetRidReportOption {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkScRidReportGetRidReportOption) SetBizMonth(v string) *TaobaoTbkScRidReportGetRidReportOption {
    s.BizMonth = &v
    return s
}
func (s *TaobaoTbkScRidReportGetRidReportOption) SetSiteId(v string) *TaobaoTbkScRidReportGetRidReportOption {
    s.SiteId = &v
    return s
}
