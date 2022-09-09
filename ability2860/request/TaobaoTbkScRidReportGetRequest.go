package request

import (
        "github.com/liu8534584/topsdk/ability2860/domain"
        "github.com/liu8534584/topsdk/util"
    )

type TaobaoTbkScRidReportGetRequest struct {
    /*
        入参     */
    SearchOption  *domain.TaobaoTbkScRidReportGetRidReportOption `json:"search_option,omitempty" required:"false" `
}

func (s *TaobaoTbkScRidReportGetRequest) SetSearchOption(v domain.TaobaoTbkScRidReportGetRidReportOption) *TaobaoTbkScRidReportGetRequest {
    s.SearchOption = &v
    return s
}

func (req *TaobaoTbkScRidReportGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SearchOption != nil) {
        paramMap["search_option"] = util.ConvertStruct(*req.SearchOption)
    }
    return paramMap
}

func (req *TaobaoTbkScRidReportGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}