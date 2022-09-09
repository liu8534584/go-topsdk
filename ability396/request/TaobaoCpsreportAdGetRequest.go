package request


type TaobaoCpsreportAdGetRequest struct {
    /*
        开始启始下标（注意不是页码，从0开始）     */
    Start  *int64 `json:"start" required:"true" `
    /*
        每页条数     */
    Limit  *int64 `json:"limit" required:"true" `
    /*
        开始时间     */
    StartDate  *string `json:"start_date" required:"true" `
    /*
        结束时间     */
    EndDate  *string `json:"end_date" required:"true" `
}

func (s *TaobaoCpsreportAdGetRequest) SetStart(v int64) *TaobaoCpsreportAdGetRequest {
    s.Start = &v
    return s
}
func (s *TaobaoCpsreportAdGetRequest) SetLimit(v int64) *TaobaoCpsreportAdGetRequest {
    s.Limit = &v
    return s
}
func (s *TaobaoCpsreportAdGetRequest) SetStartDate(v string) *TaobaoCpsreportAdGetRequest {
    s.StartDate = &v
    return s
}
func (s *TaobaoCpsreportAdGetRequest) SetEndDate(v string) *TaobaoCpsreportAdGetRequest {
    s.EndDate = &v
    return s
}

func (req *TaobaoCpsreportAdGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Start != nil) {
        paramMap["start"] = *req.Start
    }
    if(req.Limit != nil) {
        paramMap["limit"] = *req.Limit
    }
    if(req.StartDate != nil) {
        paramMap["start_date"] = *req.StartDate
    }
    if(req.EndDate != nil) {
        paramMap["end_date"] = *req.EndDate
    }
    return paramMap
}

func (req *TaobaoCpsreportAdGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}