package response

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
)

type AlibabaAlscUnionKbOrderRefundResponse struct {

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
	Data domain.AlibabaAlscUnionKbOrderRefundOrderVoucherDto `json:"data,omitempty" `
	/*
	   1
	*/
	ResultCode int64 `json:"result_code,omitempty" `
	/*
	   发生错误时对应错误原因
	*/
	Message string `json:"message,omitempty" `
}
