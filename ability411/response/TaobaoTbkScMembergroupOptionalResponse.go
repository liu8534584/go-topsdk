package response

import (
	"topsdk/ability411/domain"
)

type TaobaoTbkScMembergroupOptionalResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果
	*/
	Data domain.TaobaoTbkScMembergroupOptionalMapData `json:"data,omitempty" `
}
