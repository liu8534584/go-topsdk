package request


type AlibabaEletkSettlesTagGetRequest struct {
    /*
        月份     */
    Month  *int64 `json:"month" required:"true" `
    /*
        起始下标（从0开始），注意不是页码     */
    Start  *int64 `json:"start" required:"true" `
    /*
        条数     */
    Limit  *int64 `json:"limit" required:"true" `
}

func (s *AlibabaEletkSettlesTagGetRequest) SetMonth(v int64) *AlibabaEletkSettlesTagGetRequest {
    s.Month = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetRequest) SetStart(v int64) *AlibabaEletkSettlesTagGetRequest {
    s.Start = &v
    return s
}
func (s *AlibabaEletkSettlesTagGetRequest) SetLimit(v int64) *AlibabaEletkSettlesTagGetRequest {
    s.Limit = &v
    return s
}

func (req *AlibabaEletkSettlesTagGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Month != nil) {
        paramMap["month"] = *req.Month
    }
    if(req.Start != nil) {
        paramMap["start"] = *req.Start
    }
    if(req.Limit != nil) {
        paramMap["limit"] = *req.Limit
    }
    return paramMap
}

func (req *AlibabaEletkSettlesTagGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}