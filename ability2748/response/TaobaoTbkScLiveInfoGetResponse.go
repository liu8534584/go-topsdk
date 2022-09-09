package response

import (
    "github.com/liu8534584/topsdk/ability2748/domain"
)

type TaobaoTbkScLiveInfoGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        出参数据包
    */
    Data  domain.TaobaoTbkScLiveInfoGetData `json:"data,omitempty" `
}
