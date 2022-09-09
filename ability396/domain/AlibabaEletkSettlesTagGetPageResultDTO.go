package domain


type AlibabaEletkSettlesTagGetPageResultDTO struct {
    /*
        系统自动生成     */
    ResultList  *[]AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO `json:"result_list,omitempty" `

    /*
        成功or失败     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        总数     */
    TotalCount  *int64 `json:"total_count,omitempty" `

    /*
        错误信息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *AlibabaEletkSettlesTagGetPageResultDTO) SetResultList(v []AlibabaEletkSettlesTagGetAdsEleTaokeBuyShopOrderTagDTO) *AlibabaEletkSettlesTagGetPageResultDTO {
    s.ResultList = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetPageResultDTO) SetSuccess(v bool) *AlibabaEletkSettlesTagGetPageResultDTO {
    s.Success = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetPageResultDTO) SetErrorCode(v string) *AlibabaEletkSettlesTagGetPageResultDTO {
    s.ErrorCode = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetPageResultDTO) SetTotalCount(v int64) *AlibabaEletkSettlesTagGetPageResultDTO {
    s.TotalCount = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetPageResultDTO) SetErrorMsg(v string) *AlibabaEletkSettlesTagGetPageResultDTO {
    s.ErrorMsg = &v
    return s
}
