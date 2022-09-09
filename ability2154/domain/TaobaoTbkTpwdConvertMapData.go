package domain


type TaobaoTbkTpwdConvertMapData struct {
    /*
        商品Id     */
    NumIid  *string `json:"num_iid,omitempty" `

    /*
        商品淘客转链     */
    ClickUrl  *string `json:"click_url,omitempty" `

    /*
        店铺卖家ID     */
    SellerId  *string `json:"seller_id,omitempty" `

    /*
        入参淘口令对应原始链接     */
    OriginUrl  *string `json:"origin_url,omitempty" `

    /*
        入参淘口令推广链接中的pid，如果不属于当前调用的推广者则展示“0”     */
    OriginPid  *string `json:"origin_pid,omitempty" `

    /*
        1-动态ID转链场景，2-消费者比价场景，3-商品库导购场景     */
    BizSceneId  *string `json:"biz_scene_id,omitempty" `

}

func (s *TaobaoTbkTpwdConvertMapData) SetNumIid(v string) *TaobaoTbkTpwdConvertMapData {
    s.NumIid = &v
    return s
}
func (s *TaobaoTbkTpwdConvertMapData) SetClickUrl(v string) *TaobaoTbkTpwdConvertMapData {
    s.ClickUrl = &v
    return s
}
func (s *TaobaoTbkTpwdConvertMapData) SetSellerId(v string) *TaobaoTbkTpwdConvertMapData {
    s.SellerId = &v
    return s
}
func (s *TaobaoTbkTpwdConvertMapData) SetOriginUrl(v string) *TaobaoTbkTpwdConvertMapData {
    s.OriginUrl = &v
    return s
}
func (s *TaobaoTbkTpwdConvertMapData) SetOriginPid(v string) *TaobaoTbkTpwdConvertMapData {
    s.OriginPid = &v
    return s
}
func (s *TaobaoTbkTpwdConvertMapData) SetBizSceneId(v string) *TaobaoTbkTpwdConvertMapData {
    s.BizSceneId = &v
    return s
}
