package response

import (
    "github.com/liu8534584/topsdk/ability2860/domain"
)

type TaobaoTbkScRidReportGetResponse struct {

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
    Result  domain.TaobaoTbkScRidReportGetResult `json:"result,omitempty" `
}
