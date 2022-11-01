package domain

type AlibabaAlscUnionKbItemDetailGetItemContentGroup struct {
	/*
	   组标题     */
	Title *string `json:"title,omitempty" `

	/*
	   组内列表     */
	ContentDetails *[]AlibabaAlscUnionKbItemDetailGetContentDetail `json:"content_details,omitempty" `
}

func (s *AlibabaAlscUnionKbItemDetailGetItemContentGroup) SetTitle(v string) *AlibabaAlscUnionKbItemDetailGetItemContentGroup {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetItemContentGroup) SetContentDetails(v []AlibabaAlscUnionKbItemDetailGetContentDetail) *AlibabaAlscUnionKbItemDetailGetItemContentGroup {
	s.ContentDetails = &v
	return s
}
