package domain


type TaobaoTbkScLiveMaterialGetItemMapData struct {
    /*
        商品图片     */
    Pic  *string `json:"pic,omitempty" `

    /*
        商品价格     */
    PromotionPrice  *string `json:"promotion_price,omitempty" `

    /*
        商品标题     */
    ItemName  *string `json:"item_name,omitempty" `

    /*
        商品ID     */
    ItemId  *int64 `json:"item_id,omitempty" `

}

func (s *TaobaoTbkScLiveMaterialGetItemMapData) SetPic(v string) *TaobaoTbkScLiveMaterialGetItemMapData {
    s.Pic = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetItemMapData) SetPromotionPrice(v string) *TaobaoTbkScLiveMaterialGetItemMapData {
    s.PromotionPrice = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetItemMapData) SetItemName(v string) *TaobaoTbkScLiveMaterialGetItemMapData {
    s.ItemName = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetItemMapData) SetItemId(v int64) *TaobaoTbkScLiveMaterialGetItemMapData {
    s.ItemId = &v
    return s
}
