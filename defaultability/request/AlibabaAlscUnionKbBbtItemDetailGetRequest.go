package request

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionKbBbtItemDetailGetRequest struct {
	/*
	   爆爆团商品详情rquest     */
	QueryRequest *domain.AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionKbBbtItemDetailGetRequest) SetQueryRequest(v domain.AlibabaAlscUnionKbBbtItemDetailGetBbtItemDetailRequest) *AlibabaAlscUnionKbBbtItemDetailGetRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionKbBbtItemDetailGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbBbtItemDetailGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
