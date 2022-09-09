package domain


type TaobaoTbkScVegasSendStatusMapData struct {
    /*
        若该用户当前无待核销的红包，则返回1，若当前有待核销的红包，则返回0     */
    IsNewUser  *string `json:"is_new_user,omitempty" `

}

func (s *TaobaoTbkScVegasSendStatusMapData) SetIsNewUser(v string) *TaobaoTbkScVegasSendStatusMapData {
    s.IsNewUser = &v
    return s
}
