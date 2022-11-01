package domain

type AlibabaAlscUnionKbItemDetailGetInteger struct {
	/*
	   图片名     */
	Name *string `json:"name,omitempty" `

	/*
	   图片地址     */
	Url *string `json:"url,omitempty" `
}

func (s *AlibabaAlscUnionKbItemDetailGetInteger) SetName(v string) *AlibabaAlscUnionKbItemDetailGetInteger {
	s.Name = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetInteger) SetUrl(v string) *AlibabaAlscUnionKbItemDetailGetInteger {
	s.Url = &v
	return s
}
