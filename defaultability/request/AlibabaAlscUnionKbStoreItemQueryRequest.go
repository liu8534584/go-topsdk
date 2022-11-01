package request

type AlibabaAlscUnionKbStoreItemQueryRequest struct {
	/*
	   门店ID     */
	StoreId *string `json:"store_id" required:"true" `
	/*
	   场景类型（"kb_natural";）     */
	BizType *string `json:"biz_type" required:"true" `
	/*
	   推广位     */
	Pid *string `json:"pid" required:"true" `
	/*
	   sid     */
	Sid *string `json:"sid,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionKbStoreItemQueryRequest) SetStoreId(v string) *AlibabaAlscUnionKbStoreItemQueryRequest {
	s.StoreId = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreItemQueryRequest) SetBizType(v string) *AlibabaAlscUnionKbStoreItemQueryRequest {
	s.BizType = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreItemQueryRequest) SetPid(v string) *AlibabaAlscUnionKbStoreItemQueryRequest {
	s.Pid = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreItemQueryRequest) SetSid(v string) *AlibabaAlscUnionKbStoreItemQueryRequest {
	s.Sid = &v
	return s
}

func (req *AlibabaAlscUnionKbStoreItemQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.StoreId != nil {
		paramMap["store_id"] = *req.StoreId
	}
	if req.BizType != nil {
		paramMap["biz_type"] = *req.BizType
	}
	if req.Pid != nil {
		paramMap["pid"] = *req.Pid
	}
	if req.Sid != nil {
		paramMap["sid"] = *req.Sid
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbStoreItemQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
