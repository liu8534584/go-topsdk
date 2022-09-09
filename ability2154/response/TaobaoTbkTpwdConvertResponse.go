package response

import (
    "github.com/liu8534584/topsdk/ability2154/domain"
)

type TaobaoTbkTpwdConvertResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        data
    */
    Data  domain.TaobaoTbkTpwdConvertMapData `json:"data,omitempty" `
}
