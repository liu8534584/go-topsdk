package domain


type TaobaoTbkScVegasSendStatusData struct {
    /*
        返回结果封装对象     */
    ResultList  *[]TaobaoTbkScVegasSendStatusMapData `json:"result_list,omitempty" `

}

func (s *TaobaoTbkScVegasSendStatusData) SetResultList(v []TaobaoTbkScVegasSendStatusMapData) *TaobaoTbkScVegasSendStatusData {
    s.ResultList = &v
    return s
}
