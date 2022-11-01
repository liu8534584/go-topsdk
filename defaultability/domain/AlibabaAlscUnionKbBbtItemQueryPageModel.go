package domain

type AlibabaAlscUnionKbBbtItemQueryPageModel struct {
	/*
	   会话ID     */
	SessionId *string `json:"session_id,omitempty" `

	/*
	   分页记录     */
	Records *[]AlibabaAlscUnionKbBbtItemQueryBbtItemDto `json:"records,omitempty" `

	/*
	   页码     */
	PageNumber *int64 `json:"page_number,omitempty" `

	/*
	   每页数目     */
	PageSize *int64 `json:"page_size,omitempty" `

	/*
	   总数     */
	Total *int64 `json:"total,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemQueryPageModel) SetSessionId(v string) *AlibabaAlscUnionKbBbtItemQueryPageModel {
	s.SessionId = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryPageModel) SetRecords(v []AlibabaAlscUnionKbBbtItemQueryBbtItemDto) *AlibabaAlscUnionKbBbtItemQueryPageModel {
	s.Records = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryPageModel) SetPageNumber(v int64) *AlibabaAlscUnionKbBbtItemQueryPageModel {
	s.PageNumber = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryPageModel) SetPageSize(v int64) *AlibabaAlscUnionKbBbtItemQueryPageModel {
	s.PageSize = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemQueryPageModel) SetTotal(v int64) *AlibabaAlscUnionKbBbtItemQueryPageModel {
	s.Total = &v
	return s
}
