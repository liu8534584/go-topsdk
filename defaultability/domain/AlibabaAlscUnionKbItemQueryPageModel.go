package domain

type AlibabaAlscUnionKbItemQueryPageModel struct {
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
	Records *[]AlibabaAlscUnionKbItemQueryKbItemPromotionDTO `json:"records,omitempty" `
}

func (s *AlibabaAlscUnionKbItemQueryPageModel) SetSessionId(v string) *AlibabaAlscUnionKbItemQueryPageModel {
	s.SessionId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryPageModel) SetPageNumber(v int64) *AlibabaAlscUnionKbItemQueryPageModel {
	s.PageNumber = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryPageModel) SetPageSize(v int64) *AlibabaAlscUnionKbItemQueryPageModel {
	s.PageSize = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryPageModel) SetTotal(v int64) *AlibabaAlscUnionKbItemQueryPageModel {
	s.Total = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryPageModel) SetRecords(v []AlibabaAlscUnionKbItemQueryKbItemPromotionDTO) *AlibabaAlscUnionKbItemQueryPageModel {
	s.Records = &v
	return s
}
