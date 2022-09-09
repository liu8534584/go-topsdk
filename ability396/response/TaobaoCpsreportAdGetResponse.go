package response

import (
    "github.com/liu8534584/topsdk/ability396/domain"
)

type TaobaoCpsreportAdGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        接口返回model
    */
    Results  domain.TaobaoCpsreportAdGetPageResultDTO `json:"results,omitempty" `
}
