package domain

type AlibabaAlscUnionKbBbtItemDetailGetTextContent struct {
	/*
	   标题     */
	Title *string `json:"title,omitempty" `

	/*
	   描述     */
	Desc *string `json:"desc,omitempty" `

	/*
	   内容     */
	Contents *[]string `json:"contents,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetTextContent) SetTitle(v string) *AlibabaAlscUnionKbBbtItemDetailGetTextContent {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetTextContent) SetDesc(v string) *AlibabaAlscUnionKbBbtItemDetailGetTextContent {
	s.Desc = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetTextContent) SetContents(v []string) *AlibabaAlscUnionKbBbtItemDetailGetTextContent {
	s.Contents = &v
	return s
}
