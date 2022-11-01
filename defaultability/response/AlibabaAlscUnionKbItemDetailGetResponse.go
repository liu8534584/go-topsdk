package response

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
)

type AlibabaAlscUnionKbItemDetailGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   商品详情
	*/
	Data domain.AlibabaAlscUnionKbItemDetailGetKbItemDetailDto `json:"data,omitempty" `
	/*
	   返回码
	*/
	ResultCode int64 `json:"result_code,omitempty" `
	/*
	   返回消息
	*/
	Message string `json:"message,omitempty" `
}
