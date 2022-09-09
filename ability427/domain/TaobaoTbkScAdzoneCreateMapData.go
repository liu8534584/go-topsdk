package domain


type TaobaoTbkScAdzoneCreateMapData struct {
    /*
        完整的pid     */
    Model  *string `json:"model,omitempty" `

}

func (s *TaobaoTbkScAdzoneCreateMapData) SetModel(v string) *TaobaoTbkScAdzoneCreateMapData {
    s.Model = &v
    return s
}
