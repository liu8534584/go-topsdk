package domain


type TaobaoTbkScLiveMaterialGetAnchor struct {
    /*
        主播头像     */
    UserLogo  *string `json:"user_logo,omitempty" `

    /*
        主播昵称     */
    UserNick  *string `json:"user_nick,omitempty" `

}

func (s *TaobaoTbkScLiveMaterialGetAnchor) SetUserLogo(v string) *TaobaoTbkScLiveMaterialGetAnchor {
    s.UserLogo = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetAnchor) SetUserNick(v string) *TaobaoTbkScLiveMaterialGetAnchor {
    s.UserNick = &v
    return s
}
