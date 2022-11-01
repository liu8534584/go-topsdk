package response

import ()

type AlibabaAlscUnionKbCommonEncryptResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   加密后的字符串，有效期1小时
	*/
	Data string `json:"data,omitempty" `
	/*
	   1
	*/
	ResultCode int64 `json:"result_code,omitempty" `
	/*
	   success
	*/
	Message string `json:"message,omitempty" `
}
