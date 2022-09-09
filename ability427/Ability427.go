package ability427

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability427/request"
    "github.com/liu8534584/topsdk/ability427/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability427 struct {
    Client *topsdk.TopClient
}

func NewAbility427(client *topsdk.TopClient) *Ability427{
    return &Ability427{client}
}

/*
    淘宝客-服务商-创建推广者位
*/
func (ability *Ability427) TaobaoTbkScAdzoneCreate(req *request.TaobaoTbkScAdzoneCreateRequest,session string) (*response.TaobaoTbkScAdzoneCreateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability427 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.tbk.sc.adzone.create",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTbkScAdzoneCreateResponse{}
    if(err != nil){
        log.Println("taobaoTbkScAdzoneCreate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
