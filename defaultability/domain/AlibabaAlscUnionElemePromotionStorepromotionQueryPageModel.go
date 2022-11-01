package domain

type AlibabaAlscUnionElemePromotionStorepromotionQueryPageModel struct {
	/*
	   会话ID（下次请求作为请求参数，用于标记分页会话自动翻页）     */
	SessionId *string `json:"session_id,omitempty" `

	/*
	   分页记录     */
	Records *[]AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto `json:"records,omitempty" `

	/*
	   每页数目     */
	PageSize *int64 `json:"page_size,omitempty" `
}

func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPageModel) SetSessionId(v string) *AlibabaAlscUnionElemePromotionStorepromotionQueryPageModel {
	s.SessionId = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPageModel) SetRecords(v []AlibabaAlscUnionElemePromotionStorepromotionQueryStorePromotionDto) *AlibabaAlscUnionElemePromotionStorepromotionQueryPageModel {
	s.Records = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryPageModel) SetPageSize(v int64) *AlibabaAlscUnionElemePromotionStorepromotionQueryPageModel {
	s.PageSize = &v
	return s
}
