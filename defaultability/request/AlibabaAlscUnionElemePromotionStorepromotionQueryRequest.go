package request

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionElemePromotionStorepromotionQueryRequest struct {
	/*
	   查询rquest     */
	QueryRequest *domain.AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest `json:"query_request" required:"true" `
}

func (s *AlibabaAlscUnionElemePromotionStorepromotionQueryRequest) SetQueryRequest(v domain.AlibabaAlscUnionElemePromotionStorepromotionQueryPromotionQueryRequest) *AlibabaAlscUnionElemePromotionStorepromotionQueryRequest {
	s.QueryRequest = &v
	return s
}

func (req *AlibabaAlscUnionElemePromotionStorepromotionQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryRequest != nil {
		paramMap["query_request"] = util.ConvertStruct(*req.QueryRequest)
	}
	return paramMap
}

func (req *AlibabaAlscUnionElemePromotionStorepromotionQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
