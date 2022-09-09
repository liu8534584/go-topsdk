package domain


type TaobaoTbkScLiveInfoGetResultList struct {
    /*
        直播间状态 0=预告,1=直播中,2=已结束（回放）     */
    LiveRoomStatus  *int64 `json:"live_room_status,omitempty" `

    /*
        点赞数     */
    PraiseCount  *int64 `json:"praise_count,omitempty" `

    /*
        栏目名称，目前暂时为空     */
    Column  *string `json:"column,omitempty" `

    /*
        直播间标题     */
    Title  *string `json:"title,omitempty" `

    /*
        直播间ID     */
    LiveId  *int64 `json:"live_id,omitempty" `

    /*
        合作模式 isMediaSelf=自播，isMediaRelay=转播     */
    OutSiteCode  *string `json:"out_site_code,omitempty" `

    /*
        观看数     */
    PvCount  *int64 `json:"pv_count,omitempty" `

    /*
        主播信息     */
    Anchor  *TaobaoTbkScLiveInfoGetAnchor `json:"anchor,omitempty" `

    /*
        直播间推广链接     */
    LiveUrl  *string `json:"live_url,omitempty" `

    /*
        直播间商品列表     */
    ItemList  *[]TaobaoTbkScLiveInfoGetItemList `json:"item_list,omitempty" `

    /*
        主播位置     */
    Location  *string `json:"location,omitempty" `

    /*
        16:9封面链接     */
    Cover16x9Url  *string `json:"cover16x9_url,omitempty" `

    /*
        1:1封面链接     */
    Cover1x1Url  *string `json:"cover1x1_url,omitempty" `

    /*
        频道名称，目前暂时为空     */
    Channel  *string `json:"channel,omitempty" `

    /*
        主播身份 daren=达人，shop=商家，目前暂时为空     */
    AnchorType  *string `json:"anchor_type,omitempty" `

}

func (s *TaobaoTbkScLiveInfoGetResultList) SetLiveRoomStatus(v int64) *TaobaoTbkScLiveInfoGetResultList {
    s.LiveRoomStatus = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetResultList) SetPraiseCount(v int64) *TaobaoTbkScLiveInfoGetResultList {
    s.PraiseCount = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetResultList) SetColumn(v string) *TaobaoTbkScLiveInfoGetResultList {
    s.Column = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetResultList) SetTitle(v string) *TaobaoTbkScLiveInfoGetResultList {
    s.Title = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetResultList) SetLiveId(v int64) *TaobaoTbkScLiveInfoGetResultList {
    s.LiveId = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetResultList) SetOutSiteCode(v string) *TaobaoTbkScLiveInfoGetResultList {
    s.OutSiteCode = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetResultList) SetPvCount(v int64) *TaobaoTbkScLiveInfoGetResultList {
    s.PvCount = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetResultList) SetAnchor(v TaobaoTbkScLiveInfoGetAnchor) *TaobaoTbkScLiveInfoGetResultList {
    s.Anchor = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetResultList) SetLiveUrl(v string) *TaobaoTbkScLiveInfoGetResultList {
    s.LiveUrl = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetResultList) SetItemList(v []TaobaoTbkScLiveInfoGetItemList) *TaobaoTbkScLiveInfoGetResultList {
    s.ItemList = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetResultList) SetLocation(v string) *TaobaoTbkScLiveInfoGetResultList {
    s.Location = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetResultList) SetCover16x9Url(v string) *TaobaoTbkScLiveInfoGetResultList {
    s.Cover16x9Url = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetResultList) SetCover1x1Url(v string) *TaobaoTbkScLiveInfoGetResultList {
    s.Cover1x1Url = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetResultList) SetChannel(v string) *TaobaoTbkScLiveInfoGetResultList {
    s.Channel = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetResultList) SetAnchorType(v string) *TaobaoTbkScLiveInfoGetResultList {
    s.AnchorType = &v
    return s
}
