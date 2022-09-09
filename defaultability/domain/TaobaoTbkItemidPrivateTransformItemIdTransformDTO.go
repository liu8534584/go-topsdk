package domain


type TaobaoTbkItemidPrivateTransformItemIdTransformDTO struct {
    /*
        原商品id     */
    OriginalItemId  *string `json:"original_item_id,omitempty" `

    /*
        升级后的商品id-b段     */
    ItemId  *string `json:"item_id,omitempty" `

}

func (s *TaobaoTbkItemidPrivateTransformItemIdTransformDTO) SetOriginalItemId(v string) *TaobaoTbkItemidPrivateTransformItemIdTransformDTO {
    s.OriginalItemId = &v
    return s
}
func (s *TaobaoTbkItemidPrivateTransformItemIdTransformDTO) SetItemId(v string) *TaobaoTbkItemidPrivateTransformItemIdTransformDTO {
    s.ItemId = &v
    return s
}
