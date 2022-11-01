package domain

type AlibabaAlscUnionKbBbtItemDetailGetItemContent struct {
	/*
	   商品内容详情组     */
	ContentGroups *[]AlibabaAlscUnionKbBbtItemDetailGetItemContentGroup `json:"content_groups,omitempty" `

	/*
	   图文详情     */
	ImageContents *[]AlibabaAlscUnionKbBbtItemDetailGetImageContent `json:"image_contents,omitempty" `

	/*
	   商品说明     */
	TextContents *[]AlibabaAlscUnionKbBbtItemDetailGetTextContent `json:"text_contents,omitempty" `

	/*
	   补充说明     */
	Remarks *[]string `json:"remarks,omitempty" `

	/*
	   商家公告     */
	Announcement *string `json:"announcement,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetItemContent) SetContentGroups(v []AlibabaAlscUnionKbBbtItemDetailGetItemContentGroup) *AlibabaAlscUnionKbBbtItemDetailGetItemContent {
	s.ContentGroups = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetItemContent) SetImageContents(v []AlibabaAlscUnionKbBbtItemDetailGetImageContent) *AlibabaAlscUnionKbBbtItemDetailGetItemContent {
	s.ImageContents = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetItemContent) SetTextContents(v []AlibabaAlscUnionKbBbtItemDetailGetTextContent) *AlibabaAlscUnionKbBbtItemDetailGetItemContent {
	s.TextContents = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetItemContent) SetRemarks(v []string) *AlibabaAlscUnionKbBbtItemDetailGetItemContent {
	s.Remarks = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetItemContent) SetAnnouncement(v string) *AlibabaAlscUnionKbBbtItemDetailGetItemContent {
	s.Announcement = &v
	return s
}
