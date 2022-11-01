package domain

type AlibabaAlscUnionKbItemDetailGetShopInfo struct {
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

func (s *AlibabaAlscUnionKbItemDetailGetShopInfo) SetFreeWifi(v bool) *AlibabaAlscUnionKbItemDetailGetShopInfo {
	s.FreeWifi = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetShopInfo) SetFreePark(v bool) *AlibabaAlscUnionKbItemDetailGetShopInfo {
	s.FreePark = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetShopInfo) SetFreeParkHours(v int64) *AlibabaAlscUnionKbItemDetailGetShopInfo {
	s.FreeParkHours = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetShopInfo) SetParkFeePerHour(v string) *AlibabaAlscUnionKbItemDetailGetShopInfo {
	s.ParkFeePerHour = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetShopInfo) SetParkFeeUpperBoundPerDay(v string) *AlibabaAlscUnionKbItemDetailGetShopInfo {
	s.ParkFeeUpperBoundPerDay = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetShopInfo) SetSupplyInvoice(v bool) *AlibabaAlscUnionKbItemDetailGetShopInfo {
	s.SupplyInvoice = &v
	return s
}
func (s *AlibabaAlscUnionKbItemDetailGetShopInfo) SetInvoiceTypes(v []string) *AlibabaAlscUnionKbItemDetailGetShopInfo {
	s.InvoiceTypes = &v
	return s
}
