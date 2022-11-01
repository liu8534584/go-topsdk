package domain

type AlibabaAlscUnionKbOrderQueryOrderQueryDto struct {
	/*
	   淘宝子单号     */
	BizOrderId *string `json:"biz_order_id,omitempty" `
}

func (s *AlibabaAlscUnionKbOrderQueryOrderQueryDto) SetBizOrderId(v string) *AlibabaAlscUnionKbOrderQueryOrderQueryDto {
	s.BizOrderId = &v
	return s
}
