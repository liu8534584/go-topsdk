package request

type AlibabaAlscUnionKbItemPromotionShareCreateRequest struct {
	/*
	   推广位pid     */
	Pid *string `json:"pid" required:"true" `
	/*
	   商品ID，默认CPA的品，如果推广其他业务单元的品，请填写对应的biz_unit     */
	ItemId *string `json:"item_id" required:"true" `
	/*
	   业务单元，1-CPA，2-CPS，3-SPU。默认1-CPA defalutValue��1    */
	BizUnit *int64 `json:"biz_unit,omitempty" required:"false" `
	/*
	   废弃 defalutValue��false    */
	IncludeMiniQrCode *bool `json:"include_mini_qr_code,omitempty" required:"false" `
	/*
	   废弃 defalutValue��false    */
	IncludeMiniQrCodeHyaline *bool `json:"include_mini_qr_code_hyaline,omitempty" required:"false" `
	/*
	   废弃 defalutValue��true    */
	IncludeImgUrl *bool `json:"include_img_url,omitempty" required:"false" `
	/*
	   第三方会员id扩展     */
	Sid *string `json:"sid,omitempty" required:"false" `
	/*
	   是否合成微信推广图 defalutValue��true    */
	IncludeWxImgUrl *bool `json:"include_wx_img_url,omitempty" required:"false" `
	/*
	   是否合成支付宝推广图 defalutValue��true    */
	IncludeAlipayImgUrl *bool `json:"include_alipay_img_url,omitempty" required:"false" `
	/*
	   是否返回吱口令 defalutValue��true    */
	IncludeAlipayWathword *bool `json:"include_alipay_wathword,omitempty" required:"false" `
}

func (s *AlibabaAlscUnionKbItemPromotionShareCreateRequest) SetPid(v string) *AlibabaAlscUnionKbItemPromotionShareCreateRequest {
	s.Pid = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateRequest) SetItemId(v string) *AlibabaAlscUnionKbItemPromotionShareCreateRequest {
	s.ItemId = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateRequest) SetBizUnit(v int64) *AlibabaAlscUnionKbItemPromotionShareCreateRequest {
	s.BizUnit = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateRequest) SetIncludeMiniQrCode(v bool) *AlibabaAlscUnionKbItemPromotionShareCreateRequest {
	s.IncludeMiniQrCode = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateRequest) SetIncludeMiniQrCodeHyaline(v bool) *AlibabaAlscUnionKbItemPromotionShareCreateRequest {
	s.IncludeMiniQrCodeHyaline = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateRequest) SetIncludeImgUrl(v bool) *AlibabaAlscUnionKbItemPromotionShareCreateRequest {
	s.IncludeImgUrl = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateRequest) SetSid(v string) *AlibabaAlscUnionKbItemPromotionShareCreateRequest {
	s.Sid = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateRequest) SetIncludeWxImgUrl(v bool) *AlibabaAlscUnionKbItemPromotionShareCreateRequest {
	s.IncludeWxImgUrl = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateRequest) SetIncludeAlipayImgUrl(v bool) *AlibabaAlscUnionKbItemPromotionShareCreateRequest {
	s.IncludeAlipayImgUrl = &v
	return s
}
func (s *AlibabaAlscUnionKbItemPromotionShareCreateRequest) SetIncludeAlipayWathword(v bool) *AlibabaAlscUnionKbItemPromotionShareCreateRequest {
	s.IncludeAlipayWathword = &v
	return s
}

func (req *AlibabaAlscUnionKbItemPromotionShareCreateRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Pid != nil {
		paramMap["pid"] = *req.Pid
	}
	if req.ItemId != nil {
		paramMap["item_id"] = *req.ItemId
	}
	if req.BizUnit != nil {
		paramMap["biz_unit"] = *req.BizUnit
	}
	if req.IncludeMiniQrCode != nil {
		paramMap["include_mini_qr_code"] = *req.IncludeMiniQrCode
	}
	if req.IncludeMiniQrCodeHyaline != nil {
		paramMap["include_mini_qr_code_hyaline"] = *req.IncludeMiniQrCodeHyaline
	}
	if req.IncludeImgUrl != nil {
		paramMap["include_img_url"] = *req.IncludeImgUrl
	}
	if req.Sid != nil {
		paramMap["sid"] = *req.Sid
	}
	if req.IncludeWxImgUrl != nil {
		paramMap["include_wx_img_url"] = *req.IncludeWxImgUrl
	}
	if req.IncludeAlipayImgUrl != nil {
		paramMap["include_alipay_img_url"] = *req.IncludeAlipayImgUrl
	}
	if req.IncludeAlipayWathword != nil {
		paramMap["include_alipay_wathword"] = *req.IncludeAlipayWathword
	}
	return paramMap
}

func (req *AlibabaAlscUnionKbItemPromotionShareCreateRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
