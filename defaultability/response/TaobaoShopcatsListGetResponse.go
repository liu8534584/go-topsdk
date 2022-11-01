package response

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
)

type TaobaoShopcatsListGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   店铺类目列表信息
	*/
	ShopCats []domain.TaobaoShopcatsListGetShopCat `json:"shop_cats,omitempty" `
}
