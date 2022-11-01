package request

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionKbOrderRefundRequest struct {
	/*
	   退款对象     */
	OrderRefundDto *domain.AlibabaAlscUnionKbOrderRefundOrderRefundDto `json:"order_refund_dto" required:"true" `
}

func (s *AlibabaAlscUnionKbOrderRefundRequest) SetOrderRefundDto(v domain.AlibabaAlscUnionKbOrderRefundOrderRefundDto) *AlibabaAlscUnionKbOrderRefundRequest {
	s.OrderRefundDto = &v
	return s
}

func (req *AlibabaAlscUnionKbOrderRefundRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.OrderRefundDto != nil {
		paramMap["order_refund_dto"] = util.ConvertStruct(*req.OrderRefundDto)
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbOrderRefundRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
