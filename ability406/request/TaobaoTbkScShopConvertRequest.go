package request


type TaobaoTbkScShopConvertRequest struct {
    /*
        mm_xxx_xxx_xxx的第3段数字     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        需返回的字段列表     */
    Fields  *string `json:"fields" required:"true" `
    /*
        链接形式-1：PC，2：无线。默认为１     */
    Platform  *int64 `json:"platform,omitempty" required:"false" `
    /*
        mm_xxx_xxx_xxx的第2段数字     */
    SiteId  *int64 `json:"site_id" required:"true" `
    /*
        (该字段不开放)自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道     */
    Unid  *string `json:"unid,omitempty" required:"false" `
    /*
        卖家id串，用,分割，从taobao.tbk.shop.get接口获取user_id字段     */
    UserIds  *string `json:"user_ids" required:"true" `
}

func (s *TaobaoTbkScShopConvertRequest) SetAdzoneId(v int64) *TaobaoTbkScShopConvertRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkScShopConvertRequest) SetFields(v string) *TaobaoTbkScShopConvertRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoTbkScShopConvertRequest) SetPlatform(v int64) *TaobaoTbkScShopConvertRequest {
    s.Platform = &v
    return s
}
func (s *TaobaoTbkScShopConvertRequest) SetSiteId(v int64) *TaobaoTbkScShopConvertRequest {
    s.SiteId = &v
    return s
}
func (s *TaobaoTbkScShopConvertRequest) SetUnid(v string) *TaobaoTbkScShopConvertRequest {
    s.Unid = &v
    return s
}
func (s *TaobaoTbkScShopConvertRequest) SetUserIds(v string) *TaobaoTbkScShopConvertRequest {
    s.UserIds = &v
    return s
}

func (req *TaobaoTbkScShopConvertRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.Fields != nil) {
        paramMap["fields"] = *req.Fields
    }
    if(req.Platform != nil) {
        paramMap["platform"] = *req.Platform
    }
    if(req.SiteId != nil) {
        paramMap["site_id"] = *req.SiteId
    }
    if(req.Unid != nil) {
        paramMap["unid"] = *req.Unid
    }
    if(req.UserIds != nil) {
        paramMap["user_ids"] = *req.UserIds
    }
    return paramMap
}

func (req *TaobaoTbkScShopConvertRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}