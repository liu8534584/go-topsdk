package domain


type TaobaoTbkItemConvertNTbkItem struct {
    /*
        淘客地址     */
    ClickUrl  *string `json:"click_url,omitempty" `

    /*
        商品ID     */
    NumIid  *string `json:"num_iid,omitempty" `

    /*
        原始输入的商品ID     */
    InputNumIid  *string `json:"input_num_iid,omitempty" `

}

func (s *TaobaoTbkItemConvertNTbkItem) SetClickUrl(v string) *TaobaoTbkItemConvertNTbkItem {
    s.ClickUrl = &v
    return s
}
func (s *TaobaoTbkItemConvertNTbkItem) SetNumIid(v string) *TaobaoTbkItemConvertNTbkItem {
    s.NumIid = &v
    return s
}
func (s *TaobaoTbkItemConvertNTbkItem) SetInputNumIid(v string) *TaobaoTbkItemConvertNTbkItem {
    s.InputNumIid = &v
    return s
}
