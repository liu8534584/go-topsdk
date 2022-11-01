package domain

type AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto struct {
	/*
	   活动ID     */
	Id *string `json:"id,omitempty" `

	/*
	   标题     */
	Title *string `json:"title,omitempty" `

	/*
	   描述     */
	Description *string `json:"description,omitempty" `

	/*
	   活动创意图片     */
	Picture *string `json:"picture,omitempty" `

	/*
	   起始时间（秒）     */
	StartTime *int64 `json:"start_time,omitempty" `

	/*
	   结束时间（秒）     */
	EndTime *int64 `json:"end_time,omitempty" `

	/*
	   推广链接     */
	Link *AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink `json:"link,omitempty" `
}

func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto) SetId(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto {
	s.Id = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto) SetTitle(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto) SetDescription(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto {
	s.Description = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto) SetPicture(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto {
	s.Picture = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto) SetStartTime(v int64) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto {
	s.StartTime = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto) SetEndTime(v int64) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto {
	s.EndTime = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto) SetLink(v AlibabaAlscUnionElemePromotionOfficialactivityGetPromotionLink) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityPromotionDto {
	s.Link = &v
	return s
}
