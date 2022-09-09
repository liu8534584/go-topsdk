package response

import (
    "github.com/liu8534584/topsdk/ability396/domain"
)

type AlibabaEletkActivitylinkWechatGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        异步获取历史数据接口返回结果
    */
    Result  domain.AlibabaEletkActivitylinkWechatGetResultDto `json:"result,omitempty" `
}
