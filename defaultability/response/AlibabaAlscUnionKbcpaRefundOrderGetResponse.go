package response

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
)

type AlibabaAlscUnionKbcpaRefundOrderGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回维权订单list
	*/
	Result []domain.AlibabaAlscUnionKbcpaRefundOrderGetRefundOrderDetailReportDTO `json:"result,omitempty" `
	/*
	   错误描述
	*/
	BizErrorDesc string `json:"biz_error_desc,omitempty" `
	/*
	   当前条件下订单列表总数
	*/
	TotalCount int64 `json:"total_count,omitempty" `
	/*
	   错误码
	*/
	BizErrorCode string `json:"biz_error_code,omitempty" `
}
