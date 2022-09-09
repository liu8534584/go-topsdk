package domain


type TaobaoTbkScRidReportGetResult struct {
    /*
        真正的业务数据结构     */
    Data  *TaobaoTbkScRidReportGetData `json:"data,omitempty" `

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

func (s *TaobaoTbkScRidReportGetResult) SetData(v TaobaoTbkScRidReportGetData) *TaobaoTbkScRidReportGetResult {
    s.Data = &v
    return s
}
func (s *TaobaoTbkScRidReportGetResult) SetSuccess(v bool) *TaobaoTbkScRidReportGetResult {
    s.Success = &v
    return s
}
func (s *TaobaoTbkScRidReportGetResult) SetResultCode(v int64) *TaobaoTbkScRidReportGetResult {
    s.ResultCode = &v
    return s
}
func (s *TaobaoTbkScRidReportGetResult) SetBizErrorDesc(v string) *TaobaoTbkScRidReportGetResult {
    s.BizErrorDesc = &v
    return s
}
func (s *TaobaoTbkScRidReportGetResult) SetResultMsg(v string) *TaobaoTbkScRidReportGetResult {
    s.ResultMsg = &v
    return s
}
func (s *TaobaoTbkScRidReportGetResult) SetBizErrorCode(v int64) *TaobaoTbkScRidReportGetResult {
    s.BizErrorCode = &v
    return s
}
