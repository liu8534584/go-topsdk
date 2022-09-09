package response

import (
    "github.com/liu8534584/topsdk/ability2817/domain"
)

type TaobaoTbkDgCpaReportDetailResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        结果数据
    */
    Result  domain.TaobaoTbkDgCpaReportDetailRpcResult `json:"result,omitempty" `
}
