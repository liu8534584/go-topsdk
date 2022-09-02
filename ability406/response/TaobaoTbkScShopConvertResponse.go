package response

import (
	"github.com/liu8534584/topsdk/ability406/domain"
)

type TaobaoTbkScShopConvertResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   加入淘宝客的店铺
	*/
	Results []domain.TaobaoTbkScShopConvertNTbkShop `json:"results,omitempty" `
}
