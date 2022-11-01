package domain

type AlibabaAlscUnionKbItemDetailGetTextContent struct {
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

func (s *AlibabaAlscUnionKbItemDetailGetTextContent) SetTitle(v string) *AlibabaAlscUnionKbItemDetailGetTextContent {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetTextContent) SetDesc(v string) *AlibabaAlscUnionKbItemDetailGetTextContent {
	s.Desc = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetTextContent) SetContents(v []string) *AlibabaAlscUnionKbItemDetailGetTextContent {
	s.Contents = &v
	return s
}
