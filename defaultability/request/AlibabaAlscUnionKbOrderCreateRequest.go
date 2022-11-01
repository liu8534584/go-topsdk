package request

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionKbOrderCreateRequest struct {
	/*
	   订单对象     */
	OrderDto *domain.AlibabaAlscUnionKbOrderCreateOrderDto `json:"order_dto" required:"true" `
}

func (s *AlibabaAlscUnionKbOrderCreateRequest) SetOrderDto(v domain.AlibabaAlscUnionKbOrderCreateOrderDto) *AlibabaAlscUnionKbOrderCreateRequest {
	s.OrderDto = &v
	return s
}

func (req *AlibabaAlscUnionKbOrderCreateRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.OrderDto != nil {
		paramMap["order_dto"] = util.ConvertStruct(*req.OrderDto)
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbOrderCreateRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
