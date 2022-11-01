package response

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
)

type AlibabaAlscUnionKbOrderCreateResponse struct {

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
	Data domain.AlibabaAlscUnionKbOrderCreateOrderVoucherDto `json:"data,omitempty" `
	/*
	   1
	*/
	ResultCode int64 `json:"result_code,omitempty" `
	/*
	   success
	*/
	Message string `json:"message,omitempty" `
}
