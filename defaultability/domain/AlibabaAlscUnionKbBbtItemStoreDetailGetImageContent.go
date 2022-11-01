package domain

type AlibabaAlscUnionKbBbtItemStoreDetailGetImageContent struct {
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

func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetImageContent) SetTitle(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetImageContent {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetImageContent) SetDesc(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetImageContent {
	s.Desc = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetImageContent) SetUrls(v []string) *AlibabaAlscUnionKbBbtItemStoreDetailGetImageContent {
	s.Urls = &v
	return s
}
