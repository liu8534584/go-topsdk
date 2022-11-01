package response

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
)

type AlibabaAlscUnionKbOrderPayResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回对象
	*/
	Data domain.AlibabaAlscUnionKbOrderPayOrderVoucherDto `json:"data,omitempty" `
	/*
	   1
	*/
	ResultCode string `json:"result_code,omitempty" `
	/*
	   success
	*/
	Message string `json:"message,omitempty" `
}
