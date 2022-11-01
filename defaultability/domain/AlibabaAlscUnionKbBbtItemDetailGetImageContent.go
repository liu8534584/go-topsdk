package domain

type AlibabaAlscUnionKbBbtItemDetailGetImageContent struct {
	/*
	   标题     */
	Title *string `json:"title,omitempty" `

	/*
	   描述     */
	Desc *string `json:"desc,omitempty" `

	/*
	   图片列表     */
	Urls *[]string `json:"urls,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetImageContent) SetTitle(v string) *AlibabaAlscUnionKbBbtItemDetailGetImageContent {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetImageContent) SetDesc(v string) *AlibabaAlscUnionKbBbtItemDetailGetImageContent {
	s.Desc = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetImageContent) SetUrls(v []string) *AlibabaAlscUnionKbBbtItemDetailGetImageContent {
	s.Urls = &v
	return s
}
