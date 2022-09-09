package domain


type TaobaoTbkDgCpaReportSubDetailDisplaylist struct {
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

func (s *TaobaoTbkDgCpaReportSubDetailDisplaylist) SetFieldName(v string) *TaobaoTbkDgCpaReportSubDetailDisplaylist {
    s.FieldName = &v
    return s
}
func (s *TaobaoTbkDgCpaReportSubDetailDisplaylist) SetDisplayName(v string) *TaobaoTbkDgCpaReportSubDetailDisplaylist {
    s.DisplayName = &v
    return s
}
func (s *TaobaoTbkDgCpaReportSubDetailDisplaylist) SetDesc(v string) *TaobaoTbkDgCpaReportSubDetailDisplaylist {
    s.Desc = &v
    return s
}
