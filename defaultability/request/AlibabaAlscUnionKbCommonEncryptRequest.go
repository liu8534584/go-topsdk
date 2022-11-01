package request

import (
	"github.com/liu8534584/topsdk/defaultability/domain"
	"github.com/liu8534584/topsdk/util"
)

type AlibabaAlscUnionKbCommonEncryptRequest struct {
	/*
	   待加密对象     */
	EncryptModel *domain.AlibabaAlscUnionKbCommonEncryptBlowfishModel `json:"encrypt_model,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionKbCommonEncryptRequest) SetEncryptModel(v domain.AlibabaAlscUnionKbCommonEncryptBlowfishModel) *AlibabaAlscUnionKbCommonEncryptRequest {
	s.EncryptModel = &v
	return s
}

func (req *AlibabaAlscUnionKbCommonEncryptRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.EncryptModel != nil {
		paramMap["encrypt_model"] = util.ConvertStruct(*req.EncryptModel)
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbCommonEncryptRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
