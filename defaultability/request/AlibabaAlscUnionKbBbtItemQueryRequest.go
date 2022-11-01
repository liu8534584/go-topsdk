package request

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionKbBbtItemQueryRequest struct {
	/*
	   爆爆团商品查询rquest     */
	QueryRequest *domain.AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionKbBbtItemQueryRequest) SetQueryRequest(v domain.AlibabaAlscUnionKbBbtItemQueryBbtItemQueryRequest) *AlibabaAlscUnionKbBbtItemQueryRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionKbBbtItemQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbBbtItemQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
