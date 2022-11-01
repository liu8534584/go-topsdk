package domain

type AlibabaAlscUnionKbBbtItemDetailGetShopInfo struct {
	/*
	   免费wifi     */
	FreeWifi *bool `json:"free_wifi,omitempty" `

	/*
	   免费停车     */
	FreePark *bool `json:"free_park,omitempty" `

	/*
	   免费停车小时数     */
	FreeParkHours *int64 `json:"free_park_hours,omitempty" `

	/*
	   停车收费金额     */
	ParkFeePerHour *string `json:"park_fee_per_hour,omitempty" `

	/*
	   每段时间的封顶金额  例如 24小时封顶xx元     */
	ParkFeeUpperBoundPerDay *string `json:"park_fee_upper_bound_per_day,omitempty" `

	/*
	   提供发票     */
	SupplyInvoice *bool `json:"supply_invoice,omitempty" `

	/*
	   发票类型:电子发票或纸质发票     */
	InvoiceTypes *[]string `json:"invoice_types,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetShopInfo) SetFreeWifi(v bool) *AlibabaAlscUnionKbBbtItemDetailGetShopInfo {
	s.FreeWifi = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetShopInfo) SetFreePark(v bool) *AlibabaAlscUnionKbBbtItemDetailGetShopInfo {
	s.FreePark = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetShopInfo) SetFreeParkHours(v int64) *AlibabaAlscUnionKbBbtItemDetailGetShopInfo {
	s.FreeParkHours = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetShopInfo) SetParkFeePerHour(v string) *AlibabaAlscUnionKbBbtItemDetailGetShopInfo {
	s.ParkFeePerHour = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetShopInfo) SetParkFeeUpperBoundPerDay(v string) *AlibabaAlscUnionKbBbtItemDetailGetShopInfo {
	s.ParkFeeUpperBoundPerDay = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetShopInfo) SetSupplyInvoice(v bool) *AlibabaAlscUnionKbBbtItemDetailGetShopInfo {
	s.SupplyInvoice = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetShopInfo) SetInvoiceTypes(v []string) *AlibabaAlscUnionKbBbtItemDetailGetShopInfo {
	s.InvoiceTypes = &v
	return s
}
