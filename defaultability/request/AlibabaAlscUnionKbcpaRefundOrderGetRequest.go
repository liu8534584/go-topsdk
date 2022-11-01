package request

type AlibabaAlscUnionKbcpaRefundOrderGetRequest struct {
	/*
	   时间维度，1.订单结算时间 2.维权创建时间 3.维权完成时间 4更新时间     */
	DateType *int64 `json:"date_type" required:"true" `
	/*
	   查询结束时间     */
	EndDate *string `json:"end_date,omitempty" required:"false" `
	/*
	   每页返回数据大小，默认20，最大返回50     */
	PageSize *int64 `json:"page_size" required:"true" `
	/*
	   页码，默认第一页，取值范围1~50     */
	PageNumber *int64 `json:"page_number" required:"true" `
	/*
	   查询开始时间     */
	StartDate *string `json:"start_date" required:"true" `
	/*
	   推广位pid     */
	Pid *string `json:"pid,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionKbcpaRefundOrderGetRequest) SetDateType(v int64) *AlibabaAlscUnionKbcpaRefundOrderGetRequest {
	s.DateType = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRequest) SetEndDate(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRequest {
	s.EndDate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRequest) SetPageSize(v int64) *AlibabaAlscUnionKbcpaRefundOrderGetRequest {
	s.PageSize = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRequest) SetPageNumber(v int64) *AlibabaAlscUnionKbcpaRefundOrderGetRequest {
	s.PageNumber = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRequest) SetStartDate(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRequest {
	s.StartDate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaRefundOrderGetRequest) SetPid(v string) *AlibabaAlscUnionKbcpaRefundOrderGetRequest {
	s.Pid = &v
	return s
}

func (req *AlibabaAlscUnionKbcpaRefundOrderGetRequest) ToMap() map[string]interface{} {
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
	if req.Pid != nil {
		paramMap["pid"] = *req.Pid
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbcpaRefundOrderGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
