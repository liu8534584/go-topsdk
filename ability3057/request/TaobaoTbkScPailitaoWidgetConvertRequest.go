package request


type TaobaoTbkScPailitaoWidgetConvertRequest struct {
    /*
        mm_xxx_xxx_xxx的第三位     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        渠道id     */
    RelationId  *int64 `json:"relation_id,omitempty" required:"false" `
    /*
        mm_xxx_xxx_xxx的第二位     */
    SiteId  *int64 `json:"site_id" required:"true" `
    /*
        0：代表转官方插件  1：代表转deeplink链接 defalutValue��0    */
    Type  *int32 `json:"type,omitempty" required:"false" `
}

func (s *TaobaoTbkScPailitaoWidgetConvertRequest) SetAdzoneId(v int64) *TaobaoTbkScPailitaoWidgetConvertRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkScPailitaoWidgetConvertRequest) SetRelationId(v int64) *TaobaoTbkScPailitaoWidgetConvertRequest {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkScPailitaoWidgetConvertRequest) SetSiteId(v int64) *TaobaoTbkScPailitaoWidgetConvertRequest {
    s.SiteId = &v
    return s
}
func (s *TaobaoTbkScPailitaoWidgetConvertRequest) SetType(v int32) *TaobaoTbkScPailitaoWidgetConvertRequest {
    s.Type = &v
    return s
}

func (req *TaobaoTbkScPailitaoWidgetConvertRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.RelationId != nil) {
        paramMap["relation_id"] = *req.RelationId
    }
    if(req.SiteId != nil) {
        paramMap["site_id"] = *req.SiteId
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    return paramMap
}

func (req *TaobaoTbkScPailitaoWidgetConvertRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}