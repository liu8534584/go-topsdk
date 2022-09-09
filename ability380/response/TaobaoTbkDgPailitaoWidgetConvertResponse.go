package response

import (
)

type TaobaoTbkDgPailitaoWidgetConvertResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        根据入参信息，转出开启官方拍立淘插件URL或deeplink链接
    */
    Url  string `json:"url,omitempty" `
}
