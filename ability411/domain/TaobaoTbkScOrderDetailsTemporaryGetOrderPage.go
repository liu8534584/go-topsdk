package domain


type TaobaoTbkScOrderDetailsTemporaryGetOrderPage struct {
    /*
        PublisherOrderDto     */
    Results  *[]TaobaoTbkScOrderDetailsTemporaryGetPublisherOrderDto `json:"results,omitempty" `

    /*
        是否还有上一页     */
    HasPre  *bool `json:"has_pre,omitempty" `

    /*
        位点字段，由调用方原样传递     */
    PositionIndex  *string `json:"position_index,omitempty" `

    /*
        是否还有下一页     */
    HasNext  *bool `json:"has_next,omitempty" `

    /*
        页码     */
    PageNo  *int64 `json:"page_no,omitempty" `

    /*
        页大小     */
    PageSize  *int64 `json:"page_size,omitempty" `

}

func (s *TaobaoTbkScOrderDetailsTemporaryGetOrderPage) SetResults(v []TaobaoTbkScOrderDetailsTemporaryGetPublisherOrderDto) *TaobaoTbkScOrderDetailsTemporaryGetOrderPage {
    s.Results = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetOrderPage) SetHasPre(v bool) *TaobaoTbkScOrderDetailsTemporaryGetOrderPage {
    s.HasPre = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetOrderPage) SetPositionIndex(v string) *TaobaoTbkScOrderDetailsTemporaryGetOrderPage {
    s.PositionIndex = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetOrderPage) SetHasNext(v bool) *TaobaoTbkScOrderDetailsTemporaryGetOrderPage {
    s.HasNext = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetOrderPage) SetPageNo(v int64) *TaobaoTbkScOrderDetailsTemporaryGetOrderPage {
    s.PageNo = &v
    return s
}
func (s *TaobaoTbkScOrderDetailsTemporaryGetOrderPage) SetPageSize(v int64) *TaobaoTbkScOrderDetailsTemporaryGetOrderPage {
    s.PageSize = &v
    return s
}
