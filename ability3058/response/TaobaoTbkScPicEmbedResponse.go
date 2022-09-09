package response

import (
)

type TaobaoTbkScPicEmbedResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        打码图下载地址，15分钟内有效
    */
    EmbedPicUrl  string `json:"embed_pic_url,omitempty" `
}
