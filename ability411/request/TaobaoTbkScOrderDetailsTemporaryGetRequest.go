package request

type TaobaoTbkScOrderDetailsTemporaryGetRequest struct {
	/*
	   查询时间类型，1：按照订单淘客创建时间查询，2:按照订单淘客付款时间查询，3:按照订单淘客结算时间查询，4:按照订单更新时间； defalutValue��1    */
	QueryType *int32 `json:"query_type,omitempty" required:"false" `
	/*
	   位点，除第一页之外，都需要传递；前端原样返回。     */
	PositionIndex *string `json:"position_index,omitempty" required:"false" `
	/*
	   页大小，默认20，1~100 defalutValue��20    */
	PageSize *int32 `json:"page_size,omitempty" required:"false" `
	/*
	   推广者角色类型,2:二方，3:三方，不传，表示所有角色     */
	MemberType *int32 `json:"member_type,omitempty" required:"false" `
	/*
	   淘客订单状态，12-付款，13-关闭，14-确认收货，3-结算成功;不传，表示所有状态     */
	TkStatus *int32 `json:"tk_status,omitempty" required:"false" `
	/*
	   订单查询结束时间，订单开始时间至订单结束时间，中间时间段日常要求不超过3个小时，但如618、双11、年货节等大促期间预估时间段不可超过20分钟，超过会提示错误，调用时请务必注意时间段的选择，以保证亲能正常调用！     */
	EndTime *string `json:"end_time" required:"true" `
	/*
	   订单查询开始时间     */
	StartTime *string `json:"start_time" required:"true" `
	/*
	   跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页 defalutValue��1    */
	JumpType *int32 `json:"jump_type,omitempty" required:"false" `
	/*
	   第几页，默认1，1~100 defalutValue��1    */
	PageNo *int32 `json:"page_no,omitempty" required:"false" `
	/*
	   场景订单场景类型，1:常规订单，2:渠道订单，3:会员运营订单，默认为1 defalutValue��1    */
	OrderScene *int32 `json:"order_scene,omitempty" required:"false" `
	/*
	   member组ID     */
	MemberGroupId *string `json:"member_group_id,omitempty" required:"false" `
}

func (s *TaobaoTbkScOrderDetailsTemporaryGetRequest) SetQueryType(v int32) *TaobaoTbkScOrderDetailsTemporaryGetRequest {
	s.QueryType = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetRequest) SetPositionIndex(v string) *TaobaoTbkScOrderDetailsTemporaryGetRequest {
	s.PositionIndex = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetRequest) SetPageSize(v int32) *TaobaoTbkScOrderDetailsTemporaryGetRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetRequest) SetMemberType(v int32) *TaobaoTbkScOrderDetailsTemporaryGetRequest {
	s.MemberType = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetRequest) SetTkStatus(v int32) *TaobaoTbkScOrderDetailsTemporaryGetRequest {
	s.TkStatus = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetRequest) SetEndTime(v string) *TaobaoTbkScOrderDetailsTemporaryGetRequest {
	s.EndTime = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetRequest) SetStartTime(v string) *TaobaoTbkScOrderDetailsTemporaryGetRequest {
	s.StartTime = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetRequest) SetJumpType(v int32) *TaobaoTbkScOrderDetailsTemporaryGetRequest {
	s.JumpType = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetRequest) SetPageNo(v int32) *TaobaoTbkScOrderDetailsTemporaryGetRequest {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetRequest) SetOrderScene(v int32) *TaobaoTbkScOrderDetailsTemporaryGetRequest {
	s.OrderScene = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetRequest) SetMemberGroupId(v string) *TaobaoTbkScOrderDetailsTemporaryGetRequest {
	s.MemberGroupId = &v
	return s
}

func (req *TaobaoTbkScOrderDetailsTemporaryGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryType != nil {
		paramMap["query_type"] = *req.QueryType
	}
	if req.PositionIndex != nil {
		paramMap["position_index"] = *req.PositionIndex
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.MemberType != nil {
		paramMap["member_type"] = *req.MemberType
	}
	if req.TkStatus != nil {
		paramMap["tk_status"] = *req.TkStatus
	}
	if req.EndTime != nil {
		paramMap["end_time"] = *req.EndTime
	}
	if req.StartTime != nil {
		paramMap["start_time"] = *req.StartTime
	}
	if req.JumpType != nil {
		paramMap["jump_type"] = *req.JumpType
	}
	if req.PageNo != nil {
		paramMap["page_no"] = *req.PageNo
	}
	if req.OrderScene != nil {
		paramMap["order_scene"] = *req.OrderScene
	}
	if req.MemberGroupId != nil {
		paramMap["member_group_id"] = *req.MemberGroupId
	}
	return paramMap
}

func (req *TaobaoTbkScOrderDetailsTemporaryGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
