package domain


type TaobaoTbkDgRidReportGetResult struct {
    /*
        真正的业务数据结构     */
    Data  *TaobaoTbkDgRidReportGetData `json:"data,omitempty" `

    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        接口返回值信息     */
    ResultCode  *int64 `json:"result_code,omitempty" `

    /*
        业务错误描述     */
    BizErrorDesc  *string `json:"biz_error_desc,omitempty" `

    /*
        返回信息     */
    ResultMsg  *string `json:"result_msg,omitempty" `

    /*
        业务错误码 101, 102,103     */
    BizErrorCode  *int64 `json:"biz_error_code,omitempty" `

}

func (s *TaobaoTbkDgRidReportGetResult) SetData(v TaobaoTbkDgRidReportGetData) *TaobaoTbkDgRidReportGetResult {
    s.Data = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetResult) SetSuccess(v bool) *TaobaoTbkDgRidReportGetResult {
    s.Success = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetResult) SetResultCode(v int64) *TaobaoTbkDgRidReportGetResult {
    s.ResultCode = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetResult) SetBizErrorDesc(v string) *TaobaoTbkDgRidReportGetResult {
    s.BizErrorDesc = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetResult) SetResultMsg(v string) *TaobaoTbkDgRidReportGetResult {
    s.ResultMsg = &v
    return s
}
func (s *TaobaoTbkDgRidReportGetResult) SetBizErrorCode(v int64) *TaobaoTbkDgRidReportGetResult {
    s.BizErrorCode = &v
    return s
}
