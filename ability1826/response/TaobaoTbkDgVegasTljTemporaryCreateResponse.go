package response

import (
    "github.com/liu8534584/topsdk/ability1826/domain"
)

type TaobaoTbkDgVegasTljTemporaryCreateResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        result
    */
    Result  domain.TaobaoTbkDgVegasTljTemporaryCreateResult `json:"result,omitempty" `
}
