package response

import (
	"github.com/liu8534584/topsdk/ability2973/domain"
)

type TaobaoTbkScVegasSendStatusResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果描述信息
	*/
	ResultMsg string `json:"result_msg,omitempty" `
	/*
	   返回结果封装对象
	*/
	Data domain.TaobaoTbkScVegasSendStatusData `json:"data,omitempty" `
}
