package request

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionKbItemStoreDetailGetRequest struct {
	/*
	   门店详情查询rquest     */
	QueryRequest *domain.AlibabaAlscUnionKbItemStoreDetailGetKbItemShopDetailRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionKbItemStoreDetailGetRequest) SetQueryRequest(v domain.AlibabaAlscUnionKbItemStoreDetailGetKbItemShopDetailRequest) *AlibabaAlscUnionKbItemStoreDetailGetRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionKbItemStoreDetailGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbItemStoreDetailGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
