package request


type TaobaoTbkScLiveMaterialGetRequest struct {
    /*
        智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID     */
    DeviceType  *string `json:"device_type,omitempty" required:"false" `
    /*
        智能匹配-设备号加密类型：MD5，类型为OAID时不传     */
    DeviceEncrypt  *string `json:"device_encrypt,omitempty" required:"false" `
    /*
        智能匹配-设备号加密后的值（MD5加密需32位小写），类型为OAID时传原始OAID值     */
    DeviceValue  *string `json:"device_value,omitempty" required:"false" `
    /*
        第几页，从0开始 defalutValue��0    */
    PageNum  *int64 `json:"page_num,omitempty" required:"false" `
    /*
        页大小，默认1，最大不超过20 defalutValue��1    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        渠道关系ID     */
    RelationId  *int64 `json:"relation_id,omitempty" required:"false" `
    /*
        mm_xxx_xxx_xxx的第2段数字     */
    SiteId  *int64 `json:"site_id" required:"true" `
    /*
        mm_xxx_xxx_xxx的第3段数字     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
}

func (s *TaobaoTbkScLiveMaterialGetRequest) SetDeviceType(v string) *TaobaoTbkScLiveMaterialGetRequest {
    s.DeviceType = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetRequest) SetDeviceEncrypt(v string) *TaobaoTbkScLiveMaterialGetRequest {
    s.DeviceEncrypt = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetRequest) SetDeviceValue(v string) *TaobaoTbkScLiveMaterialGetRequest {
    s.DeviceValue = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetRequest) SetPageNum(v int64) *TaobaoTbkScLiveMaterialGetRequest {
    s.PageNum = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetRequest) SetPageSize(v int64) *TaobaoTbkScLiveMaterialGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetRequest) SetRelationId(v int64) *TaobaoTbkScLiveMaterialGetRequest {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetRequest) SetSiteId(v int64) *TaobaoTbkScLiveMaterialGetRequest {
    s.SiteId = &v
    return s
}
func (s *TaobaoTbkScLiveMaterialGetRequest) SetAdzoneId(v int64) *TaobaoTbkScLiveMaterialGetRequest {
    s.AdzoneId = &v
    return s
}

func (req *TaobaoTbkScLiveMaterialGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DeviceType != nil) {
        paramMap["device_type"] = *req.DeviceType
    }
    if(req.DeviceEncrypt != nil) {
        paramMap["device_encrypt"] = *req.DeviceEncrypt
    }
    if(req.DeviceValue != nil) {
        paramMap["device_value"] = *req.DeviceValue
    }
    if(req.PageNum != nil) {
        paramMap["page_num"] = *req.PageNum
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.RelationId != nil) {
        paramMap["relation_id"] = *req.RelationId
    }
    if(req.SiteId != nil) {
        paramMap["site_id"] = *req.SiteId
    }
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    return paramMap
}

func (req *TaobaoTbkScLiveMaterialGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}