package domain

type AlibabaAlscUnionKbBbtItemDetailGetItemContentGroup struct {
	/*
	   组标题     */
	Title *string `json:"title,omitempty" `

	/*
	   组内列表     */
	ContentDetails *[]AlibabaAlscUnionKbBbtItemDetailGetContentDetail `json:"content_details,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetItemContentGroup) SetTitle(v string) *AlibabaAlscUnionKbBbtItemDetailGetItemContentGroup {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetItemContentGroup) SetContentDetails(v []AlibabaAlscUnionKbBbtItemDetailGetContentDetail) *AlibabaAlscUnionKbBbtItemDetailGetItemContentGroup {
	s.ContentDetails = &v
	return s
}
