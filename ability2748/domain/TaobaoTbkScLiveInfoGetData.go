package domain


type TaobaoTbkScLiveInfoGetData struct {
    /*
        列表中的直播间数量     */
    TotalCount  *int64 `json:"total_count,omitempty" `

    /*
        直播间列表     */
    ResultList  *[]TaobaoTbkScLiveInfoGetResultList `json:"result_list,omitempty" `

}

func (s *TaobaoTbkScLiveInfoGetData) SetTotalCount(v int64) *TaobaoTbkScLiveInfoGetData {
    s.TotalCount = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetData) SetResultList(v []TaobaoTbkScLiveInfoGetResultList) *TaobaoTbkScLiveInfoGetData {
    s.ResultList = &v
    return s
}
