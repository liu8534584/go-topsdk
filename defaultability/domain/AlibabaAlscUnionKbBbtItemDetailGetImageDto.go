package domain

type AlibabaAlscUnionKbBbtItemDetailGetImageDto struct {
	/*
	   图片名     */
	Name *string `json:"name,omitempty" `

	/*
	   图片地址     */
	Url *string `json:"url,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetImageDto) SetName(v string) *AlibabaAlscUnionKbBbtItemDetailGetImageDto {
	s.Name = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetImageDto) SetUrl(v string) *AlibabaAlscUnionKbBbtItemDetailGetImageDto {
	s.Url = &v
	return s
}
