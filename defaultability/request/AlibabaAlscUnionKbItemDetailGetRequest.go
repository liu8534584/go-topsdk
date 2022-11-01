package request

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionKbItemDetailGetRequest struct {
	/*
	   商品详情rquest     */
	QueryRequest *domain.AlibabaAlscUnionKbItemDetailGetKbItemDetailRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionKbItemDetailGetRequest) SetQueryRequest(v domain.AlibabaAlscUnionKbItemDetailGetKbItemDetailRequest) *AlibabaAlscUnionKbItemDetailGetRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionKbItemDetailGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbItemDetailGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
