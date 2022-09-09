package response

import (
    "github.com/liu8534584/topsdk/ability1826/domain"
)

type TaobaoTbkDgVegasTljCreateResponse struct {

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
    Result  domain.TaobaoTbkDgVegasTljCreateResult `json:"result,omitempty" `
}
