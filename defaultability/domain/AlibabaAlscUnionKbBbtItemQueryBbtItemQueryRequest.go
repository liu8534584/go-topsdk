package domain

type AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest struct {
	/*
	   页码 defalutValue:1    */
	PageNumber *int64 `json:"page_number,omitempty" `

	/*
	   排序类型（0-时间倒序，1-佣金比例倒序） defalutValue:0    */
	SortType *int64 `json:"sort_type,omitempty" `

	/*
	   会话ID     */
	SessionId *string `json:"session_id,omitempty" `

	/*
	   二级类目ID     */
	Category2Id *string `json:"category2_id,omitempty" `

	/*
	   每页数目 defalutValue:20    */
	PageSize *int64 `json:"page_size,omitempty" `

	/*
	   城市ID     */
	CityId *string `json:"city_id,omitempty" `

	/*
	   是否返回需要手机号的商品，false仅返回不需要手机号的品；true全部返回     */
	IncludePhone *bool `json:"include_phone,omitempty" `

	/*
	   三方供给标识，","隔开，不为空时include_phone必须为true     */
	TripartiteAppkeys *string `json:"tripartite_appkeys,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest) SetPageNumber(v int64) *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest {
	s.PageNumber = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest) SetSortType(v int64) *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest {
	s.SortType = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest) SetSessionId(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest {
	s.SessionId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest) SetCategory2Id(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest {
	s.Category2Id = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest) SetPageSize(v int64) *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest {
	s.PageSize = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest) SetCityId(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest {
	s.CityId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest) SetIncludePhone(v bool) *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest {
	s.IncludePhone = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest) SetTripartiteAppkeys(v string) *AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest {
	s.TripartiteAppkeys = &v
	return s
}
