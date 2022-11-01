package request

type AlibabaAlscUnionKbcpaPunishOrderGetRequest struct {
	/*
	   时间维度，1.订单结算时间 2.维权创建时间 3.维权完成时间     */
	DateType *int64 `json:"date_type,omitempty" required:"false" `
	/*
	   截止查询时间     */
	EndDate *string `json:"end_date,omitempty" required:"false" `
	/*
	   每页返回数据大小，默认20，最大返回50     */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
	/*
	   页码，默认第一页，取值范围1~50     */
	PageNumber *int64 `json:"page_number,omitempty" required:"false" `
	/*
	   开始查询时间     */
	StartDate *string `json:"start_date,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionKbcpaPunishOrderGetRequest) SetDateType(v int64) *AlibabaAlscUnionKbcpaPunishOrderGetRequest {
	s.DateType = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetRequest) SetEndDate(v string) *AlibabaAlscUnionKbcpaPunishOrderGetRequest {
	s.EndDate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetRequest) SetPageSize(v int64) *AlibabaAlscUnionKbcpaPunishOrderGetRequest {
	s.PageSize = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetRequest) SetPageNumber(v int64) *AlibabaAlscUnionKbcpaPunishOrderGetRequest {
	s.PageNumber = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaPunishOrderGetRequest) SetStartDate(v string) *AlibabaAlscUnionKbcpaPunishOrderGetRequest {
	s.StartDate = &v
	return s
}

func (req *AlibabaAlscUnionKbcpaPunishOrderGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.DateType != nil {
		paramMap["date_type"] = *req.DateType
	}
	if req.EndDate != nil {
		paramMap["end_date"] = *req.EndDate
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.PageNumber != nil {
		paramMap["page_number"] = *req.PageNumber
	}
	if req.StartDate != nil {
		paramMap["start_date"] = *req.StartDate
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbcpaPunishOrderGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
