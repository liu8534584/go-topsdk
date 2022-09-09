package domain


type TaobaoTbkScLiveInfoGetAnchor struct {
    /*
        主播昵称     */
    UserNick  *string `json:"user_nick,omitempty" `

    /*
        主播头像     */
    UserLogo  *string `json:"user_logo,omitempty" `

}

func (s *TaobaoTbkScLiveInfoGetAnchor) SetUserNick(v string) *TaobaoTbkScLiveInfoGetAnchor {
    s.UserNick = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetAnchor) SetUserLogo(v string) *TaobaoTbkScLiveInfoGetAnchor {
    s.UserLogo = &v
    return s
}
