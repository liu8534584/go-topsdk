package response

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
)

type AlibabaAlscUnionKbcpxPositiveOrderGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   所有订单list
	*/
	Result []domain.AlibabaAlscUnionKbcpxPositiveOrderGetOrderDetailReportDTO `json:"result,omitempty" `
	/*
	   接口返回是否正确
	*/
	ResultSuccess bool `json:"result_success,omitempty" `
	/*
	   接口错误描述
	*/
	BizErrorDesc string `json:"biz_error_desc,omitempty" `
	/*
	   当前查询条件下总的记录
	*/
	TotalCount int64 `json:"total_count,omitempty" `
	/*
	   接口错误码
	*/
	BizErrorCode int64 `json:"biz_error_code,omitempty" `
}
