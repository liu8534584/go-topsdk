package request

type AlibabaAlscUnionKbItemPromotionFilterListRequest struct {
	/*
	   获取筛选项集合的类型     */
	FilterType *string `json:"filter_type" required:"true" `
	/*
	   1-cpa,2-cps.默认不填为cpa     */
	BizUnit *int64 `json:"biz_unit,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionKbItemPromotionFilterListRequest) SetFilterType(v string) *AlibabaAlscUnionKbItemPromotionFilterListRequest {
	s.FilterType = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionFilterListRequest) SetBizUnit(v int64) *AlibabaAlscUnionKbItemPromotionFilterListRequest {
	s.BizUnit = &v
	return s
}

func (req *AlibabaAlscUnionKbItemPromotionFilterListRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.FilterType != nil {
		paramMap["filter_type"] = *req.FilterType
	}
	if req.BizUnit != nil {
		paramMap["biz_unit"] = *req.BizUnit
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbItemPromotionFilterListRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
