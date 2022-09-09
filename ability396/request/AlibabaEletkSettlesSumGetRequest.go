package request


type AlibabaEletkSettlesSumGetRequest struct {
    /*
        月份     */
    Month  *int64 `json:"month" required:"true" `
}

func (s *AlibabaEletkSettlesSumGetRequest) SetMonth(v int64) *AlibabaEletkSettlesSumGetRequest {
    s.Month = &v
    return s
}

func (req *AlibabaEletkSettlesSumGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Month != nil) {
        paramMap["month"] = *req.Month
    }
    return paramMap
}

func (req *AlibabaEletkSettlesSumGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}