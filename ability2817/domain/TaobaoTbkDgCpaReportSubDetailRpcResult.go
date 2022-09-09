package domain


type TaobaoTbkDgCpaReportSubDetailRpcResult struct {
    /*
        结果数据     */
    Data  *TaobaoTbkDgCpaReportSubDetailPageResult `json:"data,omitempty" `

    /*
        执行结果     */
    Success  *bool `json:"success,omitempty" `

    /*
        结果code     */
    ResultCode  *int64 `json:"result_code,omitempty" `

    /*
        业务错误描述     */
    BizErrorDesc  *string `json:"biz_error_desc,omitempty" `

    /*
        异常错误的详细原因     */
    ResultMsg  *string `json:"result_msg,omitempty" `

    /*
        业务错误代码     */
    BizErrorCode  *int64 `json:"biz_error_code,omitempty" `

}

func (s *TaobaoTbkDgCpaReportSubDetailRpcResult) SetData(v TaobaoTbkDgCpaReportSubDetailPageResult) *TaobaoTbkDgCpaReportSubDetailRpcResult {
    s.Data = &v
    return s
}
func (s *TaobaoTbkDgCpaReportSubDetailRpcResult) SetSuccess(v bool) *TaobaoTbkDgCpaReportSubDetailRpcResult {
    s.Success = &v
    return s
}
func (s *TaobaoTbkDgCpaReportSubDetailRpcResult) SetResultCode(v int64) *TaobaoTbkDgCpaReportSubDetailRpcResult {
    s.ResultCode = &v
    return s
}
func (s *TaobaoTbkDgCpaReportSubDetailRpcResult) SetBizErrorDesc(v string) *TaobaoTbkDgCpaReportSubDetailRpcResult {
    s.BizErrorDesc = &v
    return s
}
func (s *TaobaoTbkDgCpaReportSubDetailRpcResult) SetResultMsg(v string) *TaobaoTbkDgCpaReportSubDetailRpcResult {
    s.ResultMsg = &v
    return s
}
func (s *TaobaoTbkDgCpaReportSubDetailRpcResult) SetBizErrorCode(v int64) *TaobaoTbkDgCpaReportSubDetailRpcResult {
    s.BizErrorCode = &v
    return s
}
