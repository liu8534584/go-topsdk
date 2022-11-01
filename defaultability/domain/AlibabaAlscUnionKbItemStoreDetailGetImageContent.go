package domain

type AlibabaAlscUnionKbItemStoreDetailGetImageContent struct {
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

func (s *AlibabaAlscUnionKbItemStoreDetailGetImageContent) SetTitle(v string) *AlibabaAlscUnionKbItemStoreDetailGetImageContent {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetImageContent) SetDesc(v string) *AlibabaAlscUnionKbItemStoreDetailGetImageContent {
	s.Desc = &v
	return s
}
func (s *AlibabaAlscUnionKbItemStoreDetailGetImageContent) SetUrls(v []string) *AlibabaAlscUnionKbItemStoreDetailGetImageContent {
	s.Urls = &v
	return s
}
