package request


type TaobaoTbkDgPailitaoWidgetConvertRequest struct {
    /*
        mm_xxx_xxx_xxx的第三位     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        渠道id     */
    RelationId  *int64 `json:"relation_id,omitempty" required:"false" `
    /*
        0：代表转官方插件  1：代表转deeplink链接 defalutValue��0    */
    Type  *int64 `json:"type,omitempty" required:"false" `
}

func (s *TaobaoTbkDgPailitaoWidgetConvertRequest) SetAdzoneId(v int64) *TaobaoTbkDgPailitaoWidgetConvertRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkDgPailitaoWidgetConvertRequest) SetRelationId(v int64) *TaobaoTbkDgPailitaoWidgetConvertRequest {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkDgPailitaoWidgetConvertRequest) SetType(v int64) *TaobaoTbkDgPailitaoWidgetConvertRequest {
    s.Type = &v
    return s
}

func (req *TaobaoTbkDgPailitaoWidgetConvertRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.RelationId != nil) {
        paramMap["relation_id"] = *req.RelationId
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    return paramMap
}

func (req *TaobaoTbkDgPailitaoWidgetConvertRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}