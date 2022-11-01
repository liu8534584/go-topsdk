package response

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
)

type AlibabaAlscUnionKbItemPromotionFilterListResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   筛选项列表
	*/
	Result []domain.AlibabaAlscUnionKbItemPromotionFilterListFilterTableNameDTO `json:"result,omitempty" `
	/*
	   错误描述码
	*/
	BizErrorCode int64 `json:"biz_error_code,omitempty" `
	/*
	   错误信息描述
	*/
	BizErrorMsg string `json:"biz_error_msg,omitempty" `
}
