package domain


type AlibabaEletkActivitylinkWechatGetResultDto struct {
    /*
        执行结果     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        微信Url Scheme链接     */
    Result  *string `json:"result,omitempty" `

    /*
        错误提示     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *AlibabaEletkActivitylinkWechatGetResultDto) SetSuccess(v bool) *AlibabaEletkActivitylinkWechatGetResultDto {
    s.Success = &v
    return s
}
func (s *AlibabaEletkActivitylinkWechatGetResultDto) SetErrorCode(v string) *AlibabaEletkActivitylinkWechatGetResultDto {
    s.ErrorCode = &v
    return s
}
func (s *AlibabaEletkActivitylinkWechatGetResultDto) SetResult(v string) *AlibabaEletkActivitylinkWechatGetResultDto {
    s.Result = &v
    return s
}
func (s *AlibabaEletkActivitylinkWechatGetResultDto) SetErrorMsg(v string) *AlibabaEletkActivitylinkWechatGetResultDto {
    s.ErrorMsg = &v
    return s
}
