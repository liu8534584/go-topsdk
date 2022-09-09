package domain


type TaobaoTbkItemidTemporaryTransformItemIdTransformDTO struct {
    /*
        原商品id     */
    OriginalItemId  *string `json:"original_item_id,omitempty" `

    /*
        升级后的商品id     */
    ItemId  *string `json:"item_id,omitempty" `

}

func (s *TaobaoTbkItemidTemporaryTransformItemIdTransformDTO) SetOriginalItemId(v string) *TaobaoTbkItemidTemporaryTransformItemIdTransformDTO {
    s.OriginalItemId = &v
    return s
}
func (s *TaobaoTbkItemidTemporaryTransformItemIdTransformDTO) SetItemId(v string) *TaobaoTbkItemidTemporaryTransformItemIdTransformDTO {
    s.ItemId = &v
    return s
}
