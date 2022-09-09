package domain


type TaobaoTbkDgCpaReportDetailResultlist struct {
    /*
        预估明细     */
    Pre  *string `json:"pre,omitempty" `

    /*
        结算明细     */
    Settle  *string `json:"settle,omitempty" `

}

func (s *TaobaoTbkDgCpaReportDetailResultlist) SetPre(v string) *TaobaoTbkDgCpaReportDetailResultlist {
    s.Pre = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailResultlist) SetSettle(v string) *TaobaoTbkDgCpaReportDetailResultlist {
    s.Settle = &v
    return s
}
