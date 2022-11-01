package response

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
)

type AlibabaAlscUnionKbOrderQueryResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   订单对象
	*/
	Data domain.AlibabaAlscUnionKbOrderQueryOrderVoucherDto `json:"data,omitempty" `
	/*
	   错误码
	*/
	ResultCode int64 `json:"result_code,omitempty" `
	/*
	   如果有错误，对应错误原因
	*/
	Message string `json:"message,omitempty" `
}
