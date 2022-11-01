package request

type AlibabaAlscUnionKbStoreQueryRequest struct {
	/*
	   页码（默认1） defalutValue��1    */
	PageNumber *int64 `json:"page_number" required:"true" `
	/*
	   每页数目（默认10） defalutValue��10    */
	PageSize *int64 `json:"page_size" required:"true" `
	/*
	   会话ID（分页场景首次请求结果返回，后续请求必须携带，服务根据session_id相同请求次数自动翻页返回）     */
	SessionId *string `json:"session_id,omitempty" required:"false" `
	/*
	   场景类型（"kb_natural";）     */
	BizType *string `json:"biz_type" required:"true" `
	/*
	   排序类型，默认normal（"normal"-门店创建时间倒序;"distance_asc"-距离最近）     */
	SortType *string `json:"sort_type,omitempty" required:"false" `
	/*
	   城市ID     */
	CityId *string `json:"city_id,omitempty" required:"false" `
	/*
	   口碑二级类目（逗号分隔）     */
	KbCategory2Ids *string `json:"kb_category_2_ids,omitempty" required:"false" `
	/*
	   口碑三级类目（逗号分隔）     */
	KbCategory3Ids *string `json:"kb_category_3_ids,omitempty" required:"false" `
	/*
	   经度（经纬度、范围配合使用）     */
	Longitude *string `json:"longitude,omitempty" required:"false" `
	/*
	   纬度（经纬度、范围配合使用）     */
	Latitude *string `json:"latitude,omitempty" required:"false" `
	/*
	   范围（单位：米，经纬度、范围配合使用）     */
	Range *int64 `json:"range,omitempty" required:"false" `
	/*
	   口碑一级类目（逗号分隔）     */
	KbCategory1Ids *string `json:"kb_category_1_ids,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionKbStoreQueryRequest) SetPageNumber(v int64) *AlibabaAlscUnionKbStoreQueryRequest {
	s.PageNumber = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryRequest) SetPageSize(v int64) *AlibabaAlscUnionKbStoreQueryRequest {
	s.PageSize = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryRequest) SetSessionId(v string) *AlibabaAlscUnionKbStoreQueryRequest {
	s.SessionId = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryRequest) SetBizType(v string) *AlibabaAlscUnionKbStoreQueryRequest {
	s.BizType = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryRequest) SetSortType(v string) *AlibabaAlscUnionKbStoreQueryRequest {
	s.SortType = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryRequest) SetCityId(v string) *AlibabaAlscUnionKbStoreQueryRequest {
	s.CityId = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryRequest) SetKbCategory2Ids(v string) *AlibabaAlscUnionKbStoreQueryRequest {
	s.KbCategory2Ids = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryRequest) SetKbCategory3Ids(v string) *AlibabaAlscUnionKbStoreQueryRequest {
	s.KbCategory3Ids = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryRequest) SetLongitude(v string) *AlibabaAlscUnionKbStoreQueryRequest {
	s.Longitude = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryRequest) SetLatitude(v string) *AlibabaAlscUnionKbStoreQueryRequest {
	s.Latitude = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryRequest) SetRange(v int64) *AlibabaAlscUnionKbStoreQueryRequest {
	s.Range = &v
	return s
}
func (s *AlibabaAlscUnionKbStoreQueryRequest) SetKbCategory1Ids(v string) *AlibabaAlscUnionKbStoreQueryRequest {
	s.KbCategory1Ids = &v
	return s
}

func (req *AlibabaAlscUnionKbStoreQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.PageNumber != nil {
		paramMap["page_number"] = *req.PageNumber
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.SessionId != nil {
		paramMap["session_id"] = *req.SessionId
	}
	if req.BizType != nil {
		paramMap["biz_type"] = *req.BizType
	}
	if req.SortType != nil {
		paramMap["sort_type"] = *req.SortType
	}
	if req.CityId != nil {
		paramMap["city_id"] = *req.CityId
	}
	if req.KbCategory2Ids != nil {
		paramMap["kb_category_2_ids"] = *req.KbCategory2Ids
	}
	if req.KbCategory3Ids != nil {
		paramMap["kb_category_3_ids"] = *req.KbCategory3Ids
	}
	if req.Longitude != nil {
		paramMap["longitude"] = *req.Longitude
	}
	if req.Latitude != nil {
		paramMap["latitude"] = *req.Latitude
	}
	if req.Range != nil {
		paramMap["range"] = *req.Range
	}
	if req.KbCategory1Ids != nil {
		paramMap["kb_category_1_ids"] = *req.KbCategory1Ids
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbStoreQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
