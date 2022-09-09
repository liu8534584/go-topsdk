package response

import (
    "github.com/liu8534584/topsdk/ability2817/domain"
)

type TaobaoTbkDgCpaReportSubDetailResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        结果根
    */
    Result  domain.TaobaoTbkDgCpaReportSubDetailRpcResult `json:"result,omitempty" `
}
