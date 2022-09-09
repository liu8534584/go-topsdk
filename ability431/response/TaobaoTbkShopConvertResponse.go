package response

import (
    "github.com/liu8534584/topsdk/ability431/domain"
)

type TaobaoTbkShopConvertResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        淘宝客店铺
    */
    Results  []domain.TaobaoTbkShopConvertNTbkShop `json:"results,omitempty" `
}
