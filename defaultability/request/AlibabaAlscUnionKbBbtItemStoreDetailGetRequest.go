package request

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionKbBbtItemStoreDetailGetRequest struct {
	/*
	   门店详情查询rquest     */
	QueryRequest *domain.AlibabaAlscUnionKbBbtItemStoreDetailGetBbtItemShopDetailRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetRequest) SetQueryRequest(v domain.AlibabaAlscUnionKbBbtItemStoreDetailGetBbtItemShopDetailRequest) *AlibabaAlscUnionKbBbtItemStoreDetailGetRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionKbBbtItemStoreDetailGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbBbtItemStoreDetailGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
