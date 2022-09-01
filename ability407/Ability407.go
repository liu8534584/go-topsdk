package ability407

import (
	"errors"
	"log"
	"topsdk"
	"topsdk/ability407/request"
	"topsdk/ability407/response"
	"topsdk/util"
)

type Ability407 struct {
	Client *topsdk.TopClient
}

func NewAbility407(client *topsdk.TopClient) *Ability407 {
	return &Ability407{client}
}

/*
   淘宝客-服务商-淘口令解析&&转链（临时接口）
*/
func (ability *Ability407) TaobaoTbkScTpwdTemporaryConvert(req *request.TaobaoTbkScTpwdTemporaryConvertRequest, session string) (*response.TaobaoTbkScTpwdTemporaryConvertResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability407 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.tpwd.temporary.convert", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScTpwdTemporaryConvertResponse{}
	if err != nil {
		log.Println("taobaoTbkScTpwdTemporaryConvert error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-服务商-淘口令解析&转链
*/
func (ability *Ability407) TaobaoTbkScTpwdConvert(req *request.TaobaoTbkScTpwdConvertRequest, session string) (*response.TaobaoTbkScTpwdConvertResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability407 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.tpwd.convert", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScTpwdConvertResponse{}
	if err != nil {
		log.Println("taobaoTbkScTpwdConvert error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
