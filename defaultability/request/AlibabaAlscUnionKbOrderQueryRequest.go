package request

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionKbOrderQueryRequest struct {
	/*
	   查询对象     */
	OrderQueryDto *domain.AlibabaAlscUnionKbOrderQueryOrderQueryDto `json:"order_query_dto" required:"true" `
}

func (s *AlibabaAlscUnionKbOrderQueryRequest) SetOrderQueryDto(v domain.AlibabaAlscUnionKbOrderQueryOrderQueryDto) *AlibabaAlscUnionKbOrderQueryRequest {
	s.OrderQueryDto = &v
	return s
}

func (req *AlibabaAlscUnionKbOrderQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.OrderQueryDto != nil {
		paramMap["order_query_dto"] = util.ConvertStruct(*req.OrderQueryDto)
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbOrderQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
