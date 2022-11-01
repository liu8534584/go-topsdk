package request

type AlibabaAlscUnionMediaZoneAddRequest struct {
	/*
	   推广位名称     */
	ZoneName *string `json:"zone_name" required:"true" `
	/*
	   媒体id，工具商渠道必填     */
	MediaId *string `json:"media_id,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionMediaZoneAddRequest) SetZoneName(v string) *AlibabaAlscUnionMediaZoneAddRequest {
	s.ZoneName = &v
	return s
}
func (s *AlibabaAlscUnionMediaZoneAddRequest) SetMediaId(v string) *AlibabaAlscUnionMediaZoneAddRequest {
	s.MediaId = &v
	return s
}

func (req *AlibabaAlscUnionMediaZoneAddRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.ZoneName != nil {
		paramMap["zone_name"] = *req.ZoneName
	}
	if req.MediaId != nil {
		paramMap["media_id"] = *req.MediaId
	}
	return paramMap
}

func (req *AlibabaAlscUnionMediaZoneAddRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
