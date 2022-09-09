package response

import (
    "github.com/liu8534584/topsdk/ability384/domain"
)

type TaobaoTbkDgRidReportGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回结果封装对象
    */
    Result  domain.TaobaoTbkDgRidReportGetResult `json:"result,omitempty" `
}
