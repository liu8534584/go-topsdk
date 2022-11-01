package request

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionKbItemStoreRelationQueryRequest struct {
	/*
	   商品门店关系查询rquest     */
	QueryRequest *domain.AlibabaAlscUnionKbItemStoreRelationQueryKbItemShopRelationRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionKbItemStoreRelationQueryRequest) SetQueryRequest(v domain.AlibabaAlscUnionKbItemStoreRelationQueryKbItemShopRelationRequest) *AlibabaAlscUnionKbItemStoreRelationQueryRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionKbItemStoreRelationQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbItemStoreRelationQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
