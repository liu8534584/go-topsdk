package response

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
)

type AlibabaAlscUnionKbItemPromotionResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   翻页时带入下一页
	*/
	SessionId string `json:"session_id,omitempty" `
	/*
	   错误详情描述
	*/
	BizErrorDesc string `json:"biz_error_desc,omitempty" `
	/*
	   当前条件下可推广品总数
	*/
	TotalCount int64 `json:"total_count,omitempty" `
	/*
	   推广品List
	*/
	Items []domain.AlibabaAlscUnionKbItemPromotionKbItemPromotionDTO `json:"items,omitempty" `
	/*
	   商品错误描述码
	*/
	BizErrorCode string `json:"biz_error_code,omitempty" `
}
