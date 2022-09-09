package request


type TaobaoTbkItemidPrivateTransformRequest struct {
    /*
        商品id列表，使用英文逗号拼接     */
    ItemIds  *string `json:"item_ids" required:"true" `
}

func (s *TaobaoTbkItemidPrivateTransformRequest) SetItemIds(v string) *TaobaoTbkItemidPrivateTransformRequest {
    s.ItemIds = &v
    return s
}

func (req *TaobaoTbkItemidPrivateTransformRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemIds != nil) {
        paramMap["item_ids"] = *req.ItemIds
    }
    return paramMap
}

func (req *TaobaoTbkItemidPrivateTransformRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}