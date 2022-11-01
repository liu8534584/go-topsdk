package domain

type AlibabaAlscUnionKbItemDetailGetImageContent struct {
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

func (s *AlibabaAlscUnionKbItemDetailGetImageContent) SetTitle(v string) *AlibabaAlscUnionKbItemDetailGetImageContent {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetImageContent) SetDesc(v string) *AlibabaAlscUnionKbItemDetailGetImageContent {
	s.Desc = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetImageContent) SetUrls(v []string) *AlibabaAlscUnionKbItemDetailGetImageContent {
	s.Urls = &v
	return s
}
