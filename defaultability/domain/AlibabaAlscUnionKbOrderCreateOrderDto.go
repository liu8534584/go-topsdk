package domain

type AlibabaAlscUnionKbOrderCreateOrderDto struct {
	/*
	   渠道订单号，需保证全局唯一     */
	OuterOrderId *string `json:"outer_order_id,omitempty" `

	/*
	   商品ID     */
	ItemId *string `json:"item_id,omitempty" `

	/*
	   购买数量     */
	Quantity *int64 `json:"quantity,omitempty" `

	/*
	   等同sell_price     */
	PayOrderFee *int64 `json:"pay_order_fee,omitempty" `

	/*
	   商品的原价*份数，单位分     */
	OrderFee *int64 `json:"order_fee,omitempty" `

	/*
	   商品的活动价*份数，单位分     */
	SellPrice *int64 `json:"sell_price,omitempty" `

	/*
	   商品名称     */
	Title *string `json:"title,omitempty" `

	/*
	   扩展参数，json格式     */
	ExtInfo *string `json:"ext_info,omitempty" `

	/*
	   true预下单不支付，false下单并支付     */
	SkipPay *bool `json:"skip_pay,omitempty" `

	/*
	   加密后的手机号     */
	Phone *string `json:"phone,omitempty" `
}

func (s *AlibabaAlscUnionKbOrderCreateOrderDto) SetOuterOrderId(v string) *AlibabaAlscUnionKbOrderCreateOrderDto {
	s.OuterOrderId = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderCreateOrderDto) SetItemId(v string) *AlibabaAlscUnionKbOrderCreateOrderDto {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderCreateOrderDto) SetQuantity(v int64) *AlibabaAlscUnionKbOrderCreateOrderDto {
	s.Quantity = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderCreateOrderDto) SetPayOrderFee(v int64) *AlibabaAlscUnionKbOrderCreateOrderDto {
	s.PayOrderFee = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderCreateOrderDto) SetOrderFee(v int64) *AlibabaAlscUnionKbOrderCreateOrderDto {
	s.OrderFee = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderCreateOrderDto) SetSellPrice(v int64) *AlibabaAlscUnionKbOrderCreateOrderDto {
	s.SellPrice = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderCreateOrderDto) SetTitle(v string) *AlibabaAlscUnionKbOrderCreateOrderDto {
	s.Title = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderCreateOrderDto) SetExtInfo(v string) *AlibabaAlscUnionKbOrderCreateOrderDto {
	s.ExtInfo = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderCreateOrderDto) SetSkipPay(v bool) *AlibabaAlscUnionKbOrderCreateOrderDto {
	s.SkipPay = &v
	return s
}
func (s *AlibabaAlscUnionKbOrderCreateOrderDto) SetPhone(v string) *AlibabaAlscUnionKbOrderCreateOrderDto {
	s.Phone = &v
	return s
}
