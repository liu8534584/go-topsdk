package request

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionKbOrderPayRequest struct {
	/*
	   订单支付对象     */
	OrderPayDto *domain.AlibabaAlscUnionKbOrderPayOrderPayDto `json:"order_pay_dto,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionKbOrderPayRequest) SetOrderPayDto(v domain.AlibabaAlscUnionKbOrderPayOrderPayDto) *AlibabaAlscUnionKbOrderPayRequest {
	s.OrderPayDto = &v
	return s
}

func (req *AlibabaAlscUnionKbOrderPayRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.OrderPayDto != nil {
		paramMap["order_pay_dto"] = util.ConvertStruct(*req.OrderPayDto)
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbOrderPayRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
