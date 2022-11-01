package response

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
)

type AlibabaAlscUnionMediaZoneGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   接口正常标识
	*/
	ResultSuccess bool `json:"result_success,omitempty" `
	/*
	   错误描述
	*/
	BizErrorDesc string `json:"biz_error_desc,omitempty" `
	/*
	   总数量
	*/
	TotalCount int64 `json:"total_count,omitempty" `
	/*
	   错误码
	*/
	BizErrorCode int64 `json:"biz_error_code,omitempty" `
	/*
	   资源位明细列表
	*/
	Result []domain.AlibabaAlscUnionMediaZoneGetZoneInfoDTO `json:"result,omitempty" `
}
