package request

type AlibabaAlscUnionKbcpaOrderDetailsGetRequest struct {
	/*
	   时间维度，1-付款时间 2-创建时间 3-结算时间 4-更新时间 defalutValue��1    */
	DateType *int64 `json:"date_type" required:"true" `
	/*
	   结算状态，1-已结算 2-未结算 不传-所有状态     */
	SettleState *int64 `json:"settle_state,omitempty" required:"false" `
	/*
	   查询结束时间     */
	EndDate *string `json:"end_date,omitempty" required:"false" `
	/*
	   每页返回数据大小，默认10，最大返回50 defalutValue��10    */
	PageSize *int64 `json:"page_size" required:"true" `
	/*
	   页码，默认第一页，取值范围1~50 defalutValue��1    */
	PageNumber *int64 `json:"page_number" required:"true" `
	/*
	   查询开始时间     */
	StartDate *string `json:"start_date" required:"true" `
	/*
	   订单状态，0-已失效 1-已下单 2-已付款 4-已收货 不传-全部状态     */
	OrderState *int64 `json:"order_state,omitempty" required:"false" `
	/*
	   推广位pid     */
	Pid *string `json:"pid,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionKbcpaOrderDetailsGetRequest) SetDateType(v int64) *AlibabaAlscUnionKbcpaOrderDetailsGetRequest {
	s.DateType = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetRequest) SetSettleState(v int64) *AlibabaAlscUnionKbcpaOrderDetailsGetRequest {
	s.SettleState = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetRequest) SetEndDate(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetRequest {
	s.EndDate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetRequest) SetPageSize(v int64) *AlibabaAlscUnionKbcpaOrderDetailsGetRequest {
	s.PageSize = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetRequest) SetPageNumber(v int64) *AlibabaAlscUnionKbcpaOrderDetailsGetRequest {
	s.PageNumber = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetRequest) SetStartDate(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetRequest {
	s.StartDate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetRequest) SetOrderState(v int64) *AlibabaAlscUnionKbcpaOrderDetailsGetRequest {
	s.OrderState = &v
	return s
}
func (s *AlibabaAlscUnionKbcpaOrderDetailsGetRequest) SetPid(v string) *AlibabaAlscUnionKbcpaOrderDetailsGetRequest {
	s.Pid = &v
	return s
}

func (req *AlibabaAlscUnionKbcpaOrderDetailsGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.DateType != nil {
		paramMap["date_type"] = *req.DateType
	}
	if req.SettleState != nil {
		paramMap["settle_state"] = *req.SettleState
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
	if req.OrderState != nil {
		paramMap["order_state"] = *req.OrderState
	}
	if req.Pid != nil {
		paramMap["pid"] = *req.Pid
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbcpaOrderDetailsGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
