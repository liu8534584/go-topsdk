package response

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
)

type AlibabaAlscUnionKbcpaOrderDetailsGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   订单明细列表
	*/
	Result []domain.AlibabaAlscUnionKbcpaOrderDetailsGetOrderDetailReportDto `json:"result,omitempty" `
	/*
	   接口正常标识
	*/
	ResultSuccess bool `json:"result_success,omitempty" `
	/*
	   错误描述
	*/
	BizErrorDesc string `json:"biz_error_desc,omitempty" `
	/*
	   当前条件下订单总量
	*/
	TotalCount int64 `json:"total_count,omitempty" `
	/*
	   错误码
	*/
	BizErrorCode string `json:"biz_error_code,omitempty" `
}
