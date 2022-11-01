package request

type AlibabaAlscUnionKbBbtItemPromotionFilterListRequest struct {
	/*
	   获取筛选项集合的类型。category类目列表，city城市列表     */
	FilterType *string `json:"filter_type" required:"true" `
	/*
	   产品线，固定bbt defalutValue��bbt    */
	BizType *string `json:"biz_type,omitempty" required:"false" `
	/*
	   固定2cps defalutValue��2    */
	BizUnit *int64 `json:"biz_unit,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionKbBbtItemPromotionFilterListRequest) SetFilterType(v string) *AlibabaAlscUnionKbBbtItemPromotionFilterListRequest {
	s.FilterType = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemPromotionFilterListRequest) SetBizType(v string) *AlibabaAlscUnionKbBbtItemPromotionFilterListRequest {
	s.BizType = &v
	return s
}
func (s *AlibabaAlscUnionKbBbtItemPromotionFilterListRequest) SetBizUnit(v int64) *AlibabaAlscUnionKbBbtItemPromotionFilterListRequest {
	s.BizUnit = &v
	return s
}

func (req *AlibabaAlscUnionKbBbtItemPromotionFilterListRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.FilterType != nil {
		paramMap["filter_type"] = *req.FilterType
	}
	if req.BizType != nil {
		paramMap["biz_type"] = *req.BizType
	}
	if req.BizUnit != nil {
		paramMap["biz_unit"] = *req.BizUnit
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbBbtItemPromotionFilterListRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
