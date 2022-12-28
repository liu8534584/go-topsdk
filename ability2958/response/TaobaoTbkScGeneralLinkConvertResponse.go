package response

import (
	"github.com/liu8534584/topsdk/ability2958/domain"
)

type TaobaoTbkScGeneralLinkConvertResponse struct {

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
	Result domain.TaobaoTbkScGeneralLinkConvertResponse `json:"data,omitempty" `
}
