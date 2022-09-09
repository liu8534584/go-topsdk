package request


type AlibabaEletkActivitylinkWechatGetRequest struct {
    /*
        淘宝联盟小程序地址里面的scene参数     */
    Url  *string `json:"url" required:"true" `
    /*
        生成微信Url Scheme链接固定传3     */
    CodeType  *int64 `json:"code_type" required:"true" `
}

func (s *AlibabaEletkActivitylinkWechatGetRequest) SetUrl(v string) *AlibabaEletkActivitylinkWechatGetRequest {
    s.Url = &v
    return s
}
func (s *AlibabaEletkActivitylinkWechatGetRequest) SetCodeType(v int64) *AlibabaEletkActivitylinkWechatGetRequest {
    s.CodeType = &v
    return s
}

func (req *AlibabaEletkActivitylinkWechatGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Url != nil) {
        paramMap["url"] = *req.Url
    }
    if(req.CodeType != nil) {
        paramMap["code_type"] = *req.CodeType
    }
    return paramMap
}

func (req *AlibabaEletkActivitylinkWechatGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}