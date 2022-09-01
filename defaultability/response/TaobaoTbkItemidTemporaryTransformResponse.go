package response

import (
	"topsdk/defaultability/domain"
)

type TaobaoTbkItemidTemporaryTransformResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   结果列表
	*/
	Results []domain.TaobaoTbkItemidTemporaryTransformItemIdTransformDTO `json:"results,omitempty" `
}
