package request


type TaobaoTbkScVegasSendStatusRequest struct {
    /*
        会员运营id     */
    SpecialId  *string `json:"special_id,omitempty" required:"false" `
    /*
        渠道管理id     */
    RelationId  *string `json:"relation_id,omitempty" required:"false" `
    /*
        加密后的值(ALIPAY_ID除外)，需用MD5加密，32位小写     */
    DeviceValue  *string `json:"device_value,omitempty" required:"false" `
    /*
        入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE 5. ALIPAY_ID     */
    DeviceType  *string `json:"device_type,omitempty" required:"false" `
}

func (s *TaobaoTbkScVegasSendStatusRequest) SetSpecialId(v string) *TaobaoTbkScVegasSendStatusRequest {
    s.SpecialId = &v
    return s
}
func (s *TaobaoTbkScVegasSendStatusRequest) SetRelationId(v string) *TaobaoTbkScVegasSendStatusRequest {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkScVegasSendStatusRequest) SetDeviceValue(v string) *TaobaoTbkScVegasSendStatusRequest {
    s.DeviceValue = &v
    return s
}
func (s *TaobaoTbkScVegasSendStatusRequest) SetDeviceType(v string) *TaobaoTbkScVegasSendStatusRequest {
    s.DeviceType = &v
    return s
}

func (req *TaobaoTbkScVegasSendStatusRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SpecialId != nil) {
        paramMap["special_id"] = *req.SpecialId
    }
    if(req.RelationId != nil) {
        paramMap["relation_id"] = *req.RelationId
    }
    if(req.DeviceValue != nil) {
        paramMap["device_value"] = *req.DeviceValue
    }
    if(req.DeviceType != nil) {
        paramMap["device_type"] = *req.DeviceType
    }
    return paramMap
}

func (req *TaobaoTbkScVegasSendStatusRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}