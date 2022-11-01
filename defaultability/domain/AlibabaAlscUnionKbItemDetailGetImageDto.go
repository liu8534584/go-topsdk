package domain

type AlibabaAlscUnionKbItemDetailGetImageDto struct {
	/*
	   图片名     */
	Name *string `json:"name,omitempty" `

	/*
	   图片地址     */
	Url *string `json:"url,omitempty" `
}

func (s *AlibabaAlscUnionKbItemDetailGetImageDto) SetName(v string) *AlibabaAlscUnionKbItemDetailGetImageDto {
	s.Name = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetImageDto) SetUrl(v string) *AlibabaAlscUnionKbItemDetailGetImageDto {
	s.Url = &v
	return s
}
