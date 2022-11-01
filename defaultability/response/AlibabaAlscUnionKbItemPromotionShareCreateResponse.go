package response

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
)

type AlibabaAlscUnionKbItemPromotionShareCreateResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   推广链接等信息
	*/
	Result domain.AlibabaAlscUnionKbItemPromotionShareCreateExtendDTO `json:"result,omitempty" `
	/*
	   接口返回是否正确
	*/
	ResultSuccess bool `json:"result_success,omitempty" `
	/*
	   接口错误码
	*/
	BizErrorCode string `json:"biz_error_code,omitempty" `
	/*
	   接口错误描述
	*/
	BizErrorDesc string `json:"biz_error_desc,omitempty" `
}
