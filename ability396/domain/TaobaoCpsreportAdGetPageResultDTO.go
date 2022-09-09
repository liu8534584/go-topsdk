package domain


type TaobaoCpsreportAdGetPageResultDTO struct {
    /*
        系统自动生成     */
    ResultList  *[]TaobaoCpsreportAdGetBuyShopOrderTagDTO `json:"result_list,omitempty" `

    /*
        系统自动生成     */
    Success  *bool `json:"success,omitempty" `

    /*
        系统自动生成     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        总条数     */
    TotalCount  *int64 `json:"total_count,omitempty" `

    /*
        系统自动生成     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

}

func (s *TaobaoCpsreportAdGetPageResultDTO) SetResultList(v []TaobaoCpsreportAdGetBuyShopOrderTagDTO) *TaobaoCpsreportAdGetPageResultDTO {
    s.ResultList = &v
    return s
}
func (s *TaobaoCpsreportAdGetPageResultDTO) SetSuccess(v bool) *TaobaoCpsreportAdGetPageResultDTO {
    s.Success = &v
    return s
}
func (s *TaobaoCpsreportAdGetPageResultDTO) SetErrorCode(v string) *TaobaoCpsreportAdGetPageResultDTO {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoCpsreportAdGetPageResultDTO) SetTotalCount(v int64) *TaobaoCpsreportAdGetPageResultDTO {
    s.TotalCount = &v
    return s
}
func (s *TaobaoCpsreportAdGetPageResultDTO) SetErrorMsg(v string) *TaobaoCpsreportAdGetPageResultDTO {
    s.ErrorMsg = &v
    return s
}
