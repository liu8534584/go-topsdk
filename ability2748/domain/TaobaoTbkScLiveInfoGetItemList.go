package domain


type TaobaoTbkScLiveInfoGetItemList struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        商品标题     */
    ItemName  *string `json:"item_name,omitempty" `

    /*
        商品价格     */
    PromotionPrice  *string `json:"promotion_price,omitempty" `

    /*
        商品图像     */
    Pic  *string `json:"pic,omitempty" `

}

func (s *TaobaoTbkScLiveInfoGetItemList) SetItemId(v int64) *TaobaoTbkScLiveInfoGetItemList {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetItemList) SetItemName(v string) *TaobaoTbkScLiveInfoGetItemList {
    s.ItemName = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetItemList) SetPromotionPrice(v string) *TaobaoTbkScLiveInfoGetItemList {
    s.PromotionPrice = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetItemList) SetPic(v string) *TaobaoTbkScLiveInfoGetItemList {
    s.Pic = &v
    return s
}
