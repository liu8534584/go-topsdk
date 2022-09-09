package request

import (
        "github.com/liu8534584/topsdk/util"
    )

type TaobaoSellercatsListGetRequest struct {
    /*
        fields参数     */
    Fields  *[]string `json:"fields,omitempty" required:"false" `
}

func (s *TaobaoSellercatsListGetRequest) SetFields(v []string) *TaobaoSellercatsListGetRequest {
    s.Fields = &v
    return s
}

func (req *TaobaoSellercatsListGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    return paramMap
}

func (req *TaobaoSellercatsListGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}