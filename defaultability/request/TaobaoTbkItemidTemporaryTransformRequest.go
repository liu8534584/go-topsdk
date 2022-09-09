package request


type TaobaoTbkItemidTemporaryTransformRequest struct {
    /*
        商品id列表，使用英文逗号拼接     */
    ItemIds  *string `json:"item_ids" required:"true" `
}

func (s *TaobaoTbkItemidTemporaryTransformRequest) SetItemIds(v string) *TaobaoTbkItemidTemporaryTransformRequest {
    s.ItemIds = &v
    return s
}

func (req *TaobaoTbkItemidTemporaryTransformRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemIds != nil) {
        paramMap["item_ids"] = *req.ItemIds
    }
    return paramMap
}

func (req *TaobaoTbkItemidTemporaryTransformRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}