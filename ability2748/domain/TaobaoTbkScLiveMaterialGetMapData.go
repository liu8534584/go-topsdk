package domain


type TaobaoTbkScLiveMaterialGetMapData struct {
    /*
        观看数     */
    PvCount  *int64 `json:"pv_count,omitempty" `

    /*
        点赞数     */
    PraiseCount  *int64 `json:"praise_count,omitempty" `

    /*
        主播位置     */
    Location  *string `json:"location,omitempty" `

    /*
        直播间推广链接     */
    LiveUrl  *string `json:"live_url,omitempty" `

    /*
        主播身份 daren=达人，shop=商家，目前暂时为空     */
    AnchorType  *string `json:"anchor_type,omitempty" `

    /*
        主播信息     */
    Anchor  *TaobaoTbkScLiveMaterialGetAnchor `json:"anchor,omitempty" `

    /*
        直播间商品列表     */
    ItemList  *[]TaobaoTbkScLiveMaterialGetItemMapData `json:"item_list,omitempty" `

    /*
        16:9封面链接     */
    Cover16x9Url  *string `json:"cover16x9_url,omitempty" `

    /*
        1:1封面链接     */
    Cover1x1Url  *string `json:"cover1x1_url,omitempty" `

    /*
        直播间标题     */
    Title  *string `json:"title,omitempty" `

    /*
        直播间ID     */
    LiveId  *int64 `json:"live_id,omitempty" `

    /*
        直播间状态 0=预告,1=直播中,2=已结束（回放）     */
    LiveRoomStatus  *int64 `json:"live_room_status,omitempty" `

}

func (s *TaobaoTbkScLiveMaterialGetMapData) SetPvCount(v int64) *TaobaoTbkScLiveMaterialGetMapData {
    s.PvCount = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetMapData) SetPraiseCount(v int64) *TaobaoTbkScLiveMaterialGetMapData {
    s.PraiseCount = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetMapData) SetLocation(v string) *TaobaoTbkScLiveMaterialGetMapData {
    s.Location = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetMapData) SetLiveUrl(v string) *TaobaoTbkScLiveMaterialGetMapData {
    s.LiveUrl = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetMapData) SetAnchorType(v string) *TaobaoTbkScLiveMaterialGetMapData {
    s.AnchorType = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetMapData) SetAnchor(v TaobaoTbkScLiveMaterialGetAnchor) *TaobaoTbkScLiveMaterialGetMapData {
    s.Anchor = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetMapData) SetItemList(v []TaobaoTbkScLiveMaterialGetItemMapData) *TaobaoTbkScLiveMaterialGetMapData {
    s.ItemList = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetMapData) SetCover16x9Url(v string) *TaobaoTbkScLiveMaterialGetMapData {
    s.Cover16x9Url = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetMapData) SetCover1x1Url(v string) *TaobaoTbkScLiveMaterialGetMapData {
    s.Cover1x1Url = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetMapData) SetTitle(v string) *TaobaoTbkScLiveMaterialGetMapData {
    s.Title = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetMapData) SetLiveId(v int64) *TaobaoTbkScLiveMaterialGetMapData {
    s.LiveId = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetMapData) SetLiveRoomStatus(v int64) *TaobaoTbkScLiveMaterialGetMapData {
    s.LiveRoomStatus = &v
    return s
}
