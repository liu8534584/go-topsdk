package request

import (
        "github.com/liu8534584/topsdk/util"
    )

type TaobaoShopSellerGetRequest struct {
    /*
        需返回的字段列表。可选值：Shop 结构中的所有字段；多个字段之间用逗号(,)分隔     */
    Fields  *[]string `json:"fields" required:"true" `
}

func (s *TaobaoShopSellerGetRequest) SetFields(v []string) *TaobaoShopSellerGetRequest {
    s.Fields = &v
    return s
}

func (req *TaobaoShopSellerGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    return paramMap
}

func (req *TaobaoShopSellerGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}