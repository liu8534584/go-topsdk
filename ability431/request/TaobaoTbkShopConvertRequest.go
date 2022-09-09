package request


type TaobaoTbkShopConvertRequest struct {
    /*
        广告位ID，区分效果位置     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        需返回的字段列表     */
    Fields  *string `json:"fields" required:"true" `
    /*
        链接形式：1：PC，2：无线，默认：１     */
    Platform  *int64 `json:"platform,omitempty" required:"false" `
    /*
        自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道     */
    Unid  *string `json:"unid,omitempty" required:"false" `
    /*
        卖家ID串，用','分割，从taobao.tbk.shop.get接口获取user_id字段     */
    UserIds  *string `json:"user_ids" required:"true" `
}

func (s *TaobaoTbkShopConvertRequest) SetAdzoneId(v int64) *TaobaoTbkShopConvertRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkShopConvertRequest) SetFields(v string) *TaobaoTbkShopConvertRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoTbkShopConvertRequest) SetPlatform(v int64) *TaobaoTbkShopConvertRequest {
    s.Platform = &v
    return s
}
func (s *TaobaoTbkShopConvertRequest) SetUnid(v string) *TaobaoTbkShopConvertRequest {
    s.Unid = &v
    return s
}
func (s *TaobaoTbkShopConvertRequest) SetUserIds(v string) *TaobaoTbkShopConvertRequest {
    s.UserIds = &v
    return s
}

func (req *TaobaoTbkShopConvertRequest) ToMap() map[string]interface{} {
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
    if(req.Unid != nil) {
        paramMap["unid"] = *req.Unid
    }
    if(req.UserIds != nil) {
        paramMap["user_ids"] = *req.UserIds
    }
    return paramMap
}

func (req *TaobaoTbkShopConvertRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}