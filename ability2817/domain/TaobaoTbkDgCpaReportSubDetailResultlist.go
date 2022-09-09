package domain


type TaobaoTbkDgCpaReportSubDetailResultlist struct {
    /*
        结算明细mapjson串     */
    Settle  *string `json:"settle,omitempty" `

    /*
        预估明细mapjson串     */
    Pre  *string `json:"pre,omitempty" `

    /*
        成员id     */
    Id  *string `json:"id,omitempty" `

}

func (s *TaobaoTbkDgCpaReportSubDetailResultlist) SetSettle(v string) *TaobaoTbkDgCpaReportSubDetailResultlist {
    s.Settle = &v
    return s
}
func (s *TaobaoTbkDgCpaReportSubDetailResultlist) SetPre(v string) *TaobaoTbkDgCpaReportSubDetailResultlist {
    s.Pre = &v
    return s
}
func (s *TaobaoTbkDgCpaReportSubDetailResultlist) SetId(v string) *TaobaoTbkDgCpaReportSubDetailResultlist {
    s.Id = &v
    return s
}
