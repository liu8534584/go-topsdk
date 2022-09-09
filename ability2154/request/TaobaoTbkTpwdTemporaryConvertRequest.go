package request


type TaobaoTbkTpwdTemporaryConvertRequest struct {
    /*
        需要解析的淘口令     */
    PasswordContent  *string `json:"password_content" required:"true" `
    /*
        广告位ID     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        1表示商品转通用计划链接，其他值或不传表示优先转营销计划链接     */
    Dx  *string `json:"dx,omitempty" required:"false" `
    /*
        会员人群ID，转链后会自动带上，可用于统计人群推广效果     */
    UcrowdId  *int64 `json:"ucrowd_id,omitempty" required:"false" `
}

func (s *TaobaoTbkTpwdTemporaryConvertRequest) SetPasswordContent(v string) *TaobaoTbkTpwdTemporaryConvertRequest {
    s.PasswordContent = &v
    return s
}
func (s *TaobaoTbkTpwdTemporaryConvertRequest) SetAdzoneId(v int64) *TaobaoTbkTpwdTemporaryConvertRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkTpwdTemporaryConvertRequest) SetDx(v string) *TaobaoTbkTpwdTemporaryConvertRequest {
    s.Dx = &v
    return s
}
func (s *TaobaoTbkTpwdTemporaryConvertRequest) SetUcrowdId(v int64) *TaobaoTbkTpwdTemporaryConvertRequest {
    s.UcrowdId = &v
    return s
}

func (req *TaobaoTbkTpwdTemporaryConvertRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PasswordContent != nil) {
        paramMap["password_content"] = *req.PasswordContent
    }
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.Dx != nil) {
        paramMap["dx"] = *req.Dx
    }
    if(req.UcrowdId != nil) {
        paramMap["ucrowd_id"] = *req.UcrowdId
    }
    return paramMap
}

func (req *TaobaoTbkTpwdTemporaryConvertRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}