package request

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionKbBbtItemStoreRelationQueryRequest struct {
	/*
	   商品门店关系查询rquest     */
	QueryRequest *domain.AlibabaAlscUnionKbBbtItemStoreRelationQueryBbtItemShopRelationRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionKbBbtItemStoreRelationQueryRequest) SetQueryRequest(v domain.AlibabaAlscUnionKbBbtItemStoreRelationQueryBbtItemShopRelationRequest) *AlibabaAlscUnionKbBbtItemStoreRelationQueryRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionKbBbtItemStoreRelationQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbBbtItemStoreRelationQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
