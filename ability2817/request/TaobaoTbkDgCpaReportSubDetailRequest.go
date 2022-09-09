package request


type TaobaoTbkDgCpaReportSubDetailRequest struct {
    /*
        活动id     */
    EventId  *int64 `json:"event_id" required:"true" `
    /*
        策略名称     */
    CampaignCode  *string `json:"campaign_code" required:"true" `
    /*
        翻页参数 defalutValue��10    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        翻页参数 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
}

func (s *TaobaoTbkDgCpaReportSubDetailRequest) SetEventId(v int64) *TaobaoTbkDgCpaReportSubDetailRequest {
    s.EventId = &v
    return s
}
func (s *TaobaoTbkDgCpaReportSubDetailRequest) SetCampaignCode(v string) *TaobaoTbkDgCpaReportSubDetailRequest {
    s.CampaignCode = &v
    return s
}
func (s *TaobaoTbkDgCpaReportSubDetailRequest) SetPageSize(v int64) *TaobaoTbkDgCpaReportSubDetailRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoTbkDgCpaReportSubDetailRequest) SetPageNo(v int64) *TaobaoTbkDgCpaReportSubDetailRequest {
    s.PageNo = &v
    return s
}

func (req *TaobaoTbkDgCpaReportSubDetailRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.EventId != nil) {
        paramMap["event_id"] = *req.EventId
    }
    if(req.CampaignCode != nil) {
        paramMap["campaign_code"] = *req.CampaignCode
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    return paramMap
}

func (req *TaobaoTbkDgCpaReportSubDetailRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}