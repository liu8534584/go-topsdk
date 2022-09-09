package request


type TaobaoTbkDgCpaReportDetailRequest struct {
    /*
        活动id     */
    EventId  *int64 `json:"event_id" required:"true" `
}

func (s *TaobaoTbkDgCpaReportDetailRequest) SetEventId(v int64) *TaobaoTbkDgCpaReportDetailRequest {
    s.EventId = &v
    return s
}

func (req *TaobaoTbkDgCpaReportDetailRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.EventId != nil) {
        paramMap["event_id"] = *req.EventId
    }
    return paramMap
}

func (req *TaobaoTbkDgCpaReportDetailRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}