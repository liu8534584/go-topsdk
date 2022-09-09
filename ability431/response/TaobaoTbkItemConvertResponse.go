package response

import (
    "github.com/liu8534584/topsdk/ability431/domain"
)

type TaobaoTbkItemConvertResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        淘宝客商品
    */
    Results  []domain.TaobaoTbkItemConvertNTbkItem `json:"results,omitempty" `
}
