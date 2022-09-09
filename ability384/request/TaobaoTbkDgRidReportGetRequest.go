package request

import (
        "github.com/liu8534584/topsdk/ability384/domain"
        "github.com/liu8534584/topsdk/util"
    )

type TaobaoTbkDgRidReportGetRequest struct {
    /*
        入参     */
    SearchOption  *domain.TaobaoTbkDgRidReportGetRidReportOption `json:"search_option,omitempty" required:"false" `
}

func (s *TaobaoTbkDgRidReportGetRequest) SetSearchOption(v domain.TaobaoTbkDgRidReportGetRidReportOption) *TaobaoTbkDgRidReportGetRequest {
    s.SearchOption = &v
    return s
}

func (req *TaobaoTbkDgRidReportGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SearchOption != nil) {
        paramMap["search_option"] = util.ConvertStruct(*req.SearchOption)
    }
    return paramMap
}

func (req *TaobaoTbkDgRidReportGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}