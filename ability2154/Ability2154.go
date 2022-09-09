package ability2154

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability2154/request"
    "github.com/liu8534584/topsdk/ability2154/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability2154 struct {
    Client *topsdk.TopClient
}

func NewAbility2154(client *topsdk.TopClient) *Ability2154{
    return &Ability2154{client}
}

/*
    淘宝客-推广者-淘口令解析&转链
*/
func (ability *Ability2154) TaobaoTbkTpwdConvert(req *request.TaobaoTbkTpwdConvertRequest) (*response.TaobaoTbkTpwdConvertResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability2154 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.tpwd.convert",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkTpwdConvertResponse{}
    if(err != nil){
        log.Println("taobaoTbkTpwdConvert error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    淘宝客-推广者-淘口令解析&&转链（临时接口）
*/
func (ability *Ability2154) TaobaoTbkTpwdTemporaryConvert(req *request.TaobaoTbkTpwdTemporaryConvertRequest) (*response.TaobaoTbkTpwdTemporaryConvertResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability2154 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.tpwd.temporary.convert",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkTpwdTemporaryConvertResponse{}
    if(err != nil){
        log.Println("taobaoTbkTpwdTemporaryConvert error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
