package response

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
)

type AlibabaAlscUnionMediaZoneAddResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   推广位信息
	*/
	Result domain.AlibabaAlscUnionMediaZoneAddZoneInfoDTO `json:"result,omitempty" `
	/*
	   返回成功或者失败
	*/
	ResultSuccess bool `json:"result_success,omitempty" `
	/*
	   业务异常
	*/
	BizErrorCode int64 `json:"biz_error_code,omitempty" `
	/*
	   推广位名称空
	*/
	BizErrorDesc string `json:"biz_error_desc,omitempty" `
}
