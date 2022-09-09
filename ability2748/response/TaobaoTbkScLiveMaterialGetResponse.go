package response

import (
    "github.com/liu8534584/topsdk/ability2748/domain"
)

type TaobaoTbkScLiveMaterialGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        直播间列表
    */
    ResultList  []domain.TaobaoTbkScLiveMaterialGetMapData `json:"result_list,omitempty" `
    /*
        列表中的直播间数量
    */
    TotalCount  int64 `json:"total_count,omitempty" `
}
