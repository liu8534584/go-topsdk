package response

import (
    "github.com/liu8534584/topsdk/defaultability/domain"
)

type TaobaoSellercatsListGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        卖家自定义类目
    */
    SellerCats  []domain.TaobaoSellercatsListGetSellerCat `json:"seller_cats,omitempty" `
}
