package request

type AlibabaAlscUnionKbcpxPositiveOrderGetRequest struct {
	/*
	   时间维度，1-付款时间 2-创建时间 3-结算时间 4-更新时间 defalutValue��1    */
	DateType *int64 `json:"date_type" required:"true" `
	/*
	   结算状态，1-已结算 2-未结算 不传-全部状态     */
	SettleState *int64 `json:"settle_state,omitempty" required:"false" `
	/*
	   查询截止时间，精确到时分秒     */
	EndDate *string `json:"end_date,omitempty" required:"false" `
	/*
	   1-CPA 2-CPS defalutValue��1    */
	BizUnit *int64 `json:"biz_unit" required:"true" `
	/*
	   每页返回数据大小，默认10，最大返回50 defalutValue��10    */
	PageSize *int64 `json:"page_size" required:"true" `
	/*
	   页码，默认第一页，取值范围1~50 defalutValue��1    */
	PageNumber *int64 `json:"page_number" required:"true" `
	/*
	   查询起始时间，精确到时分秒     */
	StartDate *string `json:"start_date" required:"true" `
	/*
	   订单状态，0-已失效 1-已下单 2-已付款 4-已收货 不传-全部状态     */
	OrderState *int64 `json:"order_state,omitempty" required:"false" `
	/*
	   场景值，7卡券订单，8卡券核销订单     */
	FlowType *int64 `json:"flow_type,omitempty" required:"false" `
	/*
	   推广位pid     */
	Pid *string `json:"pid,omitempty" required:"false" `
	/*
	   淘宝子订单号或饿了么订单号     */
	OrderId *string `json:"order_id,omitempty" required:"false" `
	/*
	   是否包含核销门店     */
	IncludeUsedStoreId *bool `json:"include_used_store_id,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionKbcpxPositiveOrderGetRequest) SetDateType(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetRequest {
	s.DateType = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetRequest) SetSettleState(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetRequest {
	s.SettleState = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetRequest) SetEndDate(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetRequest {
	s.EndDate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetRequest) SetBizUnit(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetRequest {
	s.BizUnit = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetRequest) SetPageSize(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetRequest {
	s.PageSize = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetRequest) SetPageNumber(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetRequest {
	s.PageNumber = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetRequest) SetStartDate(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetRequest {
	s.StartDate = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetRequest) SetOrderState(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetRequest {
	s.OrderState = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetRequest) SetFlowType(v int64) *AlibabaAlscUnionKbcpxPositiveOrderGetRequest {
	s.FlowType = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetRequest) SetPid(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetRequest {
	s.Pid = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetRequest) SetOrderId(v string) *AlibabaAlscUnionKbcpxPositiveOrderGetRequest {
	s.OrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbcpxPositiveOrderGetRequest) SetIncludeUsedStoreId(v bool) *AlibabaAlscUnionKbcpxPositiveOrderGetRequest {
	s.IncludeUsedStoreId = &v
	return s
}

func (req *AlibabaAlscUnionKbcpxPositiveOrderGetRequest) ToMap() map[string]interface{} {
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
	if req.BizUnit != nil {
		paramMap["biz_unit"] = *req.BizUnit
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
	if req.FlowType != nil {
		paramMap["flow_type"] = *req.FlowType
	}
	if req.Pid != nil {
		paramMap["pid"] = *req.Pid
	}
	if req.OrderId != nil {
		paramMap["order_id"] = *req.OrderId
	}
	if req.IncludeUsedStoreId != nil {
		paramMap["include_used_store_id"] = *req.IncludeUsedStoreId
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbcpxPositiveOrderGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
