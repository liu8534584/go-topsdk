package domain


type TaobaoShopSellerGetShop struct {
    /*
        店铺编号     */
    Sid  *int64 `json:"sid,omitempty" `

    /*
        店铺标题     */
    Title  *string `json:"title,omitempty" `

    /*
        店标地址。返回相对路径，可以用"http://logo.taobao.com/shop-logo"来拼接成绝对路径     */
    PicPath  *string `json:"pic_path,omitempty" `

}

func (s *TaobaoShopSellerGetShop) SetSid(v int64) *TaobaoShopSellerGetShop {
    s.Sid = &v
    return s
}
func (s *TaobaoShopSellerGetShop) SetTitle(v string) *TaobaoShopSellerGetShop {
    s.Title = &v
    return s
}
func (s *TaobaoShopSellerGetShop) SetPicPath(v string) *TaobaoShopSellerGetShop {
    s.PicPath = &v
    return s
}
