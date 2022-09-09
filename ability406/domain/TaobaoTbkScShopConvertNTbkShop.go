package domain


type TaobaoTbkScShopConvertNTbkShop struct {
    /*
        卖家id     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        店铺推广链接     */
    ClickUrl  *string `json:"click_url,omitempty" `

}

func (s *TaobaoTbkScShopConvertNTbkShop) SetUserId(v int64) *TaobaoTbkScShopConvertNTbkShop {
    s.UserId = &v
    return s
}
func (s *TaobaoTbkScShopConvertNTbkShop) SetClickUrl(v string) *TaobaoTbkScShopConvertNTbkShop {
    s.ClickUrl = &v
    return s
}
