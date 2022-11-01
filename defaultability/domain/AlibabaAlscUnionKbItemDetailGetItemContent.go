package domain

type AlibabaAlscUnionKbItemDetailGetItemContent struct {
	/*
	   商品内容详情组     */
	ContentGroups *[]AlibabaAlscUnionKbItemDetailGetItemContentGroup `json:"content_groups,omitempty" `

	/*
	   图文详情     */
	ImageContents *[]AlibabaAlscUnionKbItemDetailGetImageContent `json:"image_contents,omitempty" `

	/*
	   商品说明     */
	TextContents *[]AlibabaAlscUnionKbItemDetailGetTextContent `json:"text_contents,omitempty" `

	/*
	   补充说明     */
	Remarks *[]string `json:"remarks,omitempty" `

	/*
	   商家公告     */
	Announcement *string `json:"announcement,omitempty" `
}

func (s *AlibabaAlscUnionKbItemDetailGetItemContent) SetContentGroups(v []AlibabaAlscUnionKbItemDetailGetItemContentGroup) *AlibabaAlscUnionKbItemDetailGetItemContent {
	s.ContentGroups = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetItemContent) SetImageContents(v []AlibabaAlscUnionKbItemDetailGetImageContent) *AlibabaAlscUnionKbItemDetailGetItemContent {
	s.ImageContents = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetItemContent) SetTextContents(v []AlibabaAlscUnionKbItemDetailGetTextContent) *AlibabaAlscUnionKbItemDetailGetItemContent {
	s.TextContents = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetItemContent) SetRemarks(v []string) *AlibabaAlscUnionKbItemDetailGetItemContent {
	s.Remarks = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetItemContent) SetAnnouncement(v string) *AlibabaAlscUnionKbItemDetailGetItemContent {
	s.Announcement = &v
	return s
}
