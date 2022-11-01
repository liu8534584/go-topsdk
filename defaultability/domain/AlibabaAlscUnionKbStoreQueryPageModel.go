package domain

type AlibabaAlscUnionKbStoreQueryPageModel struct {
	/*
	   会话ID（下次请求作为请求参数，用于标记分页会话自动翻页）     */
	SessionId *string `json:"session_id,omitempty" `

	/*
	   页码     */
	PageNumber *int64 `json:"page_number,omitempty" `

	/*
	   每页数目     */
	PageSize *int64 `json:"page_size,omitempty" `

	/*
	   总数     */
	Total *int64 `json:"total,omitempty" `

	/*
	   分页记录     */
	Records *[]AlibabaAlscUnionKbStoreQueryInteger `json:"records,omitempty" `
}

func (s *AlibabaAlscUnionKbStoreQueryPageModel) SetSessionId(v string) *AlibabaAlscUnionKbStoreQueryPageModel {
	s.SessionId = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryPageModel) SetPageNumber(v int64) *AlibabaAlscUnionKbStoreQueryPageModel {
	s.PageNumber = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryPageModel) SetPageSize(v int64) *AlibabaAlscUnionKbStoreQueryPageModel {
	s.PageSize = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryPageModel) SetTotal(v int64) *AlibabaAlscUnionKbStoreQueryPageModel {
	s.Total = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryPageModel) SetRecords(v []AlibabaAlscUnionKbStoreQueryInteger) *AlibabaAlscUnionKbStoreQueryPageModel {
	s.Records = &v
	return s
}
