package request


type TaobaoTbkItemConvertRequest struct {
    /*
        广告位ID，区分效果位置     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        需返回的字段列表     */
    Fields  *string `json:"fields" required:"true" `
    /*
        商品ID串，用','分割，从taobao.tbk.item.get接口获取num_iid字段，最大40个     */
    NumIids  *string `json:"num_iids" required:"true" `
    /*
        链接形式：1：PC，2：无线，默认：１ defalutValue��1    */
    Platform  *int32 `json:"platform,omitempty" required:"false" `
    /*
        自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道     */
    Unid  *string `json:"unid,omitempty" required:"false" `
    /*
        1表示商品转通用计划链接，其他值或不传表示转营销计划链接     */
    Dx  *string `json:"dx,omitempty" required:"false" `
}

func (s *TaobaoTbkItemConvertRequest) SetAdzoneId(v int64) *TaobaoTbkItemConvertRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkItemConvertRequest) SetFields(v string) *TaobaoTbkItemConvertRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoTbkItemConvertRequest) SetNumIids(v string) *TaobaoTbkItemConvertRequest {
    s.NumIids = &v
    return s
}
func (s *TaobaoTbkItemConvertRequest) SetPlatform(v int32) *TaobaoTbkItemConvertRequest {
    s.Platform = &v
    return s
}
func (s *TaobaoTbkItemConvertRequest) SetUnid(v string) *TaobaoTbkItemConvertRequest {
    s.Unid = &v
    return s
}
func (s *TaobaoTbkItemConvertRequest) SetDx(v string) *TaobaoTbkItemConvertRequest {
    s.Dx = &v
    return s
}

func (req *TaobaoTbkItemConvertRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.Fields != nil) {
        paramMap["fields"] = *req.Fields
    }
    if(req.NumIids != nil) {
        paramMap["num_iids"] = *req.NumIids
    }
    if(req.Platform != nil) {
        paramMap["platform"] = *req.Platform
    }
    if(req.Unid != nil) {
        paramMap["unid"] = *req.Unid
    }
    if(req.Dx != nil) {
        paramMap["dx"] = *req.Dx
    }
    return paramMap
}

func (req *TaobaoTbkItemConvertRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}