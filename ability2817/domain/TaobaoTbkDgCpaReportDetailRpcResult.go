package domain


type TaobaoTbkDgCpaReportDetailRpcResult struct {
    /*
        结果数据     */
    Data  *TaobaoTbkDgCpaReportDetailCpaReportDto `json:"data,omitempty" `

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

func (s *TaobaoTbkDgCpaReportDetailRpcResult) SetData(v TaobaoTbkDgCpaReportDetailCpaReportDto) *TaobaoTbkDgCpaReportDetailRpcResult {
    s.Data = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailRpcResult) SetSuccess(v bool) *TaobaoTbkDgCpaReportDetailRpcResult {
    s.Success = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailRpcResult) SetResultCode(v int64) *TaobaoTbkDgCpaReportDetailRpcResult {
    s.ResultCode = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailRpcResult) SetBizErrorDesc(v string) *TaobaoTbkDgCpaReportDetailRpcResult {
    s.BizErrorDesc = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailRpcResult) SetResultMsg(v string) *TaobaoTbkDgCpaReportDetailRpcResult {
    s.ResultMsg = &v
    return s
}
func (s *TaobaoTbkDgCpaReportDetailRpcResult) SetBizErrorCode(v int64) *TaobaoTbkDgCpaReportDetailRpcResult {
    s.BizErrorCode = &v
    return s
}
