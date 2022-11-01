package request

import (
	"github.com/liu8534584/topsdk/util"
)

type TaobaoShopcatsListGetRequest struct {
	/*
	   需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent defalutValue��cid,parent_cid,name,is_parent    */
	Fields *[]string `json:"fields,omitempty" required:"false" `
}

func (s *TaobaoShopcatsListGetRequest) SetFields(v []string) *TaobaoShopcatsListGetRequest {
	s.Fields = &v
	return s
}

func (req *TaobaoShopcatsListGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Fields != nil {
		paramMap["fields"] = util.ConvertBasicList(*req.Fields)
	}
	return paramMap
}

func (req *TaobaoShopcatsListGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
