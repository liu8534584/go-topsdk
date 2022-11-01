package domain

type AlibabaAlscUnionKbCommonEncryptBlowfishModel struct {
	/*
	   待加密字符串     */
	Text *string `json:"text,omitempty" `
}

func (s *AlibabaAlscUnionKbCommonEncryptBlowfishModel) SetText(v string) *AlibabaAlscUnionKbCommonEncryptBlowfishModel {
	s.Text = &v
	return s
}
