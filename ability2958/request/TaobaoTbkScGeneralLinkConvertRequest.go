package request

type TaobaoTbkScGeneralLinkConvertRequest struct {
	/*
	   1-动态ID转链场景，2-消费者比价场景，4-     */
	BizSceneId *string `json:"biz_scene_id,omitempty" required:"false" `
	/*
	   1-自购省，2-推广赚（代理模式专属ID，代理模式必填，其它模式不用填写本字段）     */
	PromotionType *string `json:"promotion_type,omitempty" required:"false" `
	/*
	   物料列表，可以为url或淘口令     */
	MaterialList *string `json:"material_list,omitempty" required:"false" `
	/*
	   渠道管理ID（如是主站选品推广场景，必须入参该字段，且bizSceneId字段需入参2-消费者比价场景，否则二次转链失败）     */
	RelationId *string `json:"relation_id,omitempty" required:"false" `
	/*
	   推广位id，mm_xx_xx_xx pid三段式中的第三段     */
	AdzoneId *int64 `json:"adzone_id,omitempty" required:"true" `

	/* 卖家ID列表 */
	SellerIdList *string `json:"seller_id_list,omitempty" required:"false" `

	/* 推广位id，mm_xx_xx_xx pid三段式中的第二段 */
	SiteId *int64 `json:"site_id" required:"true" `

	/* 商品ID列表 */
	ItemIdList *string `json:"item_id_list,omitempty" required:"false" `

	/* 会场ID列表 */
	PageIdList *string `json:"page_id_list,omitempty" required:"false" `
}

func (req *TaobaoTbkScGeneralLinkConvertRequest) SetBizSceneId(bizSceneId string) *TaobaoTbkScGeneralLinkConvertRequest {
	req.BizSceneId = &bizSceneId
	return req
}

func (req *TaobaoTbkScGeneralLinkConvertRequest) SetPromotionType(promotionType string) *TaobaoTbkScGeneralLinkConvertRequest {
	req.PromotionType = &promotionType
	return req
}

func (req *TaobaoTbkScGeneralLinkConvertRequest) SetMaterialList(materialList string) *TaobaoTbkScGeneralLinkConvertRequest {
	req.MaterialList = &materialList
	return req
}

func (req *TaobaoTbkScGeneralLinkConvertRequest) SetRelationId(relationId string) *TaobaoTbkScGeneralLinkConvertRequest {
	req.RelationId = &relationId
	return req
}

func (req *TaobaoTbkScGeneralLinkConvertRequest) SetAdzoneId(adzoneId int64) *TaobaoTbkScGeneralLinkConvertRequest {
	req.AdzoneId = &adzoneId
	return req
}

func (req *TaobaoTbkScGeneralLinkConvertRequest) SetSellerIdList(sellerIdList string) *TaobaoTbkScGeneralLinkConvertRequest {
	req.SellerIdList = &sellerIdList
	return req
}

func (req *TaobaoTbkScGeneralLinkConvertRequest) SetSiteId(siteId int64) *TaobaoTbkScGeneralLinkConvertRequest {
	req.SiteId = &siteId
	return req
}

func (req *TaobaoTbkScGeneralLinkConvertRequest) SetItemIdList(itemIdList string) *TaobaoTbkScGeneralLinkConvertRequest {
	req.ItemIdList = &itemIdList
	return req
}

func (req *TaobaoTbkScGeneralLinkConvertRequest) SetPageIdList(pageIdList string) *TaobaoTbkScGeneralLinkConvertRequest {
	req.PageIdList = &pageIdList
	return req
}

func (req *TaobaoTbkScGeneralLinkConvertRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.BizSceneId != nil {
		paramMap["biz_scene_id"] = *req.BizSceneId
	}
	if req.PromotionType != nil {
		paramMap["promotion_type"] = *req.PromotionType
	}
	if req.MaterialList != nil {
		paramMap["material_list"] = *req.MaterialList
	}
	if req.RelationId != nil {
		paramMap["relation_id"] = *req.RelationId
	}
	if req.AdzoneId != nil {
		paramMap["adzone_id"] = *req.AdzoneId
	}
	if req.SellerIdList != nil {
		paramMap["seller_id_list"] = *req.SellerIdList
	}
	if req.SiteId != nil {
		paramMap["site_id"] = *req.SiteId
	}
	if req.ItemIdList != nil {
		paramMap["item_id_list"] = *req.ItemIdList
	}
	if req.PageIdList != nil {
		paramMap["page_id_list"] = *req.PageIdList
	}
	return paramMap
}

func (req *TaobaoTbkScGeneralLinkConvertRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
