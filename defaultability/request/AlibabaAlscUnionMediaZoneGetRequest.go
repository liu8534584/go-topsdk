package request

type AlibabaAlscUnionMediaZoneGetRequest struct {
	/*
	   页码，从1开始 defalutValue��1    */
	Page *int64 `json:"page" required:"true" `
	/*
	   每页展示条数 defalutValue��10    */
	Limit *int64 `json:"limit" required:"true" `
}

func (s *AlibabaAlscUnionMediaZoneGetRequest) SetPage(v int64) *AlibabaAlscUnionMediaZoneGetRequest {
	s.Page = &v
	return s
}
func (s *AlibabaAlscUnionMediaZoneGetRequest) SetLimit(v int64) *AlibabaAlscUnionMediaZoneGetRequest {
	s.Limit = &v
	return s
}

func (req *AlibabaAlscUnionMediaZoneGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Page != nil {
		paramMap["page"] = *req.Page
	}
	if req.Limit != nil {
		paramMap["limit"] = *req.Limit
	}
	return paramMap
}

func (req *AlibabaAlscUnionMediaZoneGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
