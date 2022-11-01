package request

type AlibabaAlscUnionKbItemQueryRequest struct {
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
	   推广位     */
	Pid *string `json:"pid" required:"true" `
	/*
	   淘宝二级类目（逗号分隔）     */
	TbCategory2Ids *string `json:"tb_category_2_ids,omitempty" required:"false" `
	/*
	   淘宝三级类目（逗号分隔）     */
	TbCategory3Ids *string `json:"tb_category_3_ids,omitempty" required:"false" `
	/*
	   城市ID     */
	CityId *string `json:"city_id,omitempty" required:"false" `
	/*
	   经度（经纬度、范围配合使用）     */
	Longitude *string `json:"longitude,omitempty" required:"false" `
	/*
	   纬度（经纬度、范围配合使用）     */
	Latitude *string `json:"latitude,omitempty" required:"false" `
	/*
	   范围（单位：米，经纬度、范围配合使用）     */
	Range *int64 `json:"range,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionKbItemQueryRequest) SetPageNumber(v int64) *AlibabaAlscUnionKbItemQueryRequest {
	s.PageNumber = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryRequest) SetPageSize(v int64) *AlibabaAlscUnionKbItemQueryRequest {
	s.PageSize = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryRequest) SetSessionId(v string) *AlibabaAlscUnionKbItemQueryRequest {
	s.SessionId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryRequest) SetBizType(v string) *AlibabaAlscUnionKbItemQueryRequest {
	s.BizType = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryRequest) SetSortType(v string) *AlibabaAlscUnionKbItemQueryRequest {
	s.SortType = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryRequest) SetPid(v string) *AlibabaAlscUnionKbItemQueryRequest {
	s.Pid = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryRequest) SetTbCategory2Ids(v string) *AlibabaAlscUnionKbItemQueryRequest {
	s.TbCategory2Ids = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryRequest) SetTbCategory3Ids(v string) *AlibabaAlscUnionKbItemQueryRequest {
	s.TbCategory3Ids = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryRequest) SetCityId(v string) *AlibabaAlscUnionKbItemQueryRequest {
	s.CityId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryRequest) SetLongitude(v string) *AlibabaAlscUnionKbItemQueryRequest {
	s.Longitude = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryRequest) SetLatitude(v string) *AlibabaAlscUnionKbItemQueryRequest {
	s.Latitude = &v
	return s
}
func (s *AlibabaAlscUnionKbItemQueryRequest) SetRange(v int64) *AlibabaAlscUnionKbItemQueryRequest {
	s.Range = &v
	return s
}

func (req *AlibabaAlscUnionKbItemQueryRequest) ToMap() map[string]interface{} {
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
	if req.Pid != nil {
		paramMap["pid"] = *req.Pid
	}
	if req.TbCategory2Ids != nil {
		paramMap["tb_category_2_ids"] = *req.TbCategory2Ids
	}
	if req.TbCategory3Ids != nil {
		paramMap["tb_category_3_ids"] = *req.TbCategory3Ids
	}
	if req.CityId != nil {
		paramMap["city_id"] = *req.CityId
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
	return paramMap
}

func (req *AlibabaAlscUnionKbItemQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
