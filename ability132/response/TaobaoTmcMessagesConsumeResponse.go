package response

import (
	"github.com/liu8534584/topsdk/ability132/domain"
)

type TaobaoTmcMessagesConsumeResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   消息列表
	*/
	Messages []domain.TaobaoTmcMessagesConsumeTmcMessage `json:"messages,omitempty" `
}
