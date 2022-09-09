package domain


type TaobaoTbkDgCpaReportDetailDisplaylist struct {
    /*
        字段名称     */
    FieldName  *string `json:"field_name,omitempty" `

    /*
        展现名称     */
    DisplayName  *string `json:"display_name,omitempty" `

    /*
        字段描述     */
    Desc  *string `json:"desc,omitempty" `

}

func (s *TaobaoTbkDgCpaReportDetailDisplaylist) SetFieldName(v string) *TaobaoTbkDgCpaReportDetailDisplaylist {
    s.FieldName = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailDisplaylist) SetDisplayName(v string) *TaobaoTbkDgCpaReportDetailDisplaylist {
    s.DisplayName = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailDisplaylist) SetDesc(v string) *TaobaoTbkDgCpaReportDetailDisplaylist {
    s.Desc = &v
    return s
}
