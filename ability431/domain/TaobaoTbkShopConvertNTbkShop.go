package domain


type TaobaoTbkShopConvertNTbkShop struct {
    /*
        卖家ID     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        淘客地址     */
    ClickUrl  *string `json:"click_url,omitempty" `

}

func (s *TaobaoTbkShopConvertNTbkShop) SetUserId(v int64) *TaobaoTbkShopConvertNTbkShop {
    s.UserId = &v
    return s
}
func (s *TaobaoTbkShopConvertNTbkShop) SetClickUrl(v string) *TaobaoTbkShopConvertNTbkShop {
    s.ClickUrl = &v
    return s
}
