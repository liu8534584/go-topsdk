package domain


type AlibabaEletkSettlesSumGetResultDTO struct {
    /*
        接口返回实体信息     */
    Result  *AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO `json:"result,omitempty" `

    /*
        执行结果     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误描述     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *AlibabaEletkSettlesSumGetResultDTO) SetResult(v AlibabaEletkSettlesSumGetAdsEleTaokeBuyShopOrderSumDTO) *AlibabaEletkSettlesSumGetResultDTO {
    s.Result = &v
    return s
}
func (s *AlibabaEletkSettlesSumGetResultDTO) SetSuccess(v bool) *AlibabaEletkSettlesSumGetResultDTO {
    s.Success = &v
    return s
}
func (s *AlibabaEletkSettlesSumGetResultDTO) SetErrorCode(v string) *AlibabaEletkSettlesSumGetResultDTO {
    s.ErrorCode = &v
    return s
}
func (s *AlibabaEletkSettlesSumGetResultDTO) SetErrorMsg(v string) *AlibabaEletkSettlesSumGetResultDTO {
    s.ErrorMsg = &v
    return s
}
