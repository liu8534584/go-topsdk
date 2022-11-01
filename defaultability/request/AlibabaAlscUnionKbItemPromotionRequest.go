package request

type AlibabaAlscUnionKbItemPromotionRequest struct {
	/*
	   页码，默认第一页，取值范围1~50     */
	PageNumber *int64 `json:"page_number" required:"true" `
	/*
	   排序类型 normal-默认排序 reservePrice-折后价从高到低  commission-佣金从高到低 totalSales-月销量从高到低     */
	SortType *string `json:"sort_type,omitempty" required:"false" `
	/*
	   每页返回数据大小，默认20，最大返回20     */
	PageSize *int64 `json:"page_size" required:"true" `
	/*
	   推广参数     */
	Pid *string `json:"pid" required:"true" `
	/*
	   用来分页，翻页时将上一次结果的sessionId带下来     */
	SessionId *string `json:"session_id,omitempty" required:"false" `
	/*
	   推广物料结算模型 1-cpa 2-cps，3spu     */
	SettleType *int64 `json:"settle_type" required:"true" `
	/*
	   类目筛选，多个类目逗号分隔     */
	FilterCategoryIds *string `json:"filter_category_ids,omitempty" required:"false" `
	/*
	   城市id(国标)筛选，多个城市逗号分隔     */
	FilterCityIds *string `json:"filter_city_ids,omitempty" required:"false" `
	/*
	   关键词搜索，多个词逗号分割     */
	SearchKeyword *string `json:"search_keyword,omitempty" required:"false" `
	/*
	   指定itemId查询推广信息，多个逗号分割     */
	HitItemIds *string `json:"hit_item_ids,omitempty" required:"false" `
	/*
	   第三方会员id扩展     */
	Sid *string `json:"sid,omitempty" required:"false" `
	/*
	   商品可售卖的端类型。1支付宝端商品，2微信端商品，3全部     */
	ItemType *int64 `json:"item_type,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionKbItemPromotionRequest) SetPageNumber(v int64) *AlibabaAlscUnionKbItemPromotionRequest {
	s.PageNumber = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionRequest) SetSortType(v string) *AlibabaAlscUnionKbItemPromotionRequest {
	s.SortType = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionRequest) SetPageSize(v int64) *AlibabaAlscUnionKbItemPromotionRequest {
	s.PageSize = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionRequest) SetPid(v string) *AlibabaAlscUnionKbItemPromotionRequest {
	s.Pid = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionRequest) SetSessionId(v string) *AlibabaAlscUnionKbItemPromotionRequest {
	s.SessionId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionRequest) SetSettleType(v int64) *AlibabaAlscUnionKbItemPromotionRequest {
	s.SettleType = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionRequest) SetFilterCategoryIds(v string) *AlibabaAlscUnionKbItemPromotionRequest {
	s.FilterCategoryIds = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionRequest) SetFilterCityIds(v string) *AlibabaAlscUnionKbItemPromotionRequest {
	s.FilterCityIds = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionRequest) SetSearchKeyword(v string) *AlibabaAlscUnionKbItemPromotionRequest {
	s.SearchKeyword = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionRequest) SetHitItemIds(v string) *AlibabaAlscUnionKbItemPromotionRequest {
	s.HitItemIds = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionRequest) SetSid(v string) *AlibabaAlscUnionKbItemPromotionRequest {
	s.Sid = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionRequest) SetItemType(v int64) *AlibabaAlscUnionKbItemPromotionRequest {
	s.ItemType = &v
	return s
}

func (req *AlibabaAlscUnionKbItemPromotionRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.PageNumber != nil {
		paramMap["page_number"] = *req.PageNumber
	}
	if req.SortType != nil {
		paramMap["sort_type"] = *req.SortType
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.Pid != nil {
		paramMap["pid"] = *req.Pid
	}
	if req.SessionId != nil {
		paramMap["session_id"] = *req.SessionId
	}
	if req.SettleType != nil {
		paramMap["settle_type"] = *req.SettleType
	}
	if req.FilterCategoryIds != nil {
		paramMap["filter_category_ids"] = *req.FilterCategoryIds
	}
	if req.FilterCityIds != nil {
		paramMap["filter_city_ids"] = *req.FilterCityIds
	}
	if req.SearchKeyword != nil {
		paramMap["search_keyword"] = *req.SearchKeyword
	}
	if req.HitItemIds != nil {
		paramMap["hit_item_ids"] = *req.HitItemIds
	}
	if req.Sid != nil {
		paramMap["sid"] = *req.Sid
	}
	if req.ItemType != nil {
		paramMap["item_type"] = *req.ItemType
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbItemPromotionRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
