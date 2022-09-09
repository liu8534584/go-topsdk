package response

import (
    "github.com/liu8534584/topsdk/ability427/domain"
)

type TaobaoTbkScAdzoneCreateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        MapData
    */
    Data  domain.TaobaoTbkScAdzoneCreateMapData `json:"data,omitempty" `
}
