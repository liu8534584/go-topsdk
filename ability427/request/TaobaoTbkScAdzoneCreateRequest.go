package request


type TaobaoTbkScAdzoneCreateRequest struct {
    /*
        网站ID     */
    SiteId  *int64 `json:"site_id" required:"true" `
    /*
        广告位名称，最大长度64字符     */
    AdzoneName  *string `json:"adzone_name" required:"true" `
}

func (s *TaobaoTbkScAdzoneCreateRequest) SetSiteId(v int64) *TaobaoTbkScAdzoneCreateRequest {
    s.SiteId = &v
    return s
}
func (s *TaobaoTbkScAdzoneCreateRequest) SetAdzoneName(v string) *TaobaoTbkScAdzoneCreateRequest {
    s.AdzoneName = &v
    return s
}

func (req *TaobaoTbkScAdzoneCreateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SiteId != nil) {
        paramMap["site_id"] = *req.SiteId
    }
    if(req.AdzoneName != nil) {
        paramMap["adzone_name"] = *req.AdzoneName
    }
    return paramMap
}

func (req *TaobaoTbkScAdzoneCreateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}