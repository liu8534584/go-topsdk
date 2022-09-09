package ability2973

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability2973/request"
    "github.com/liu8534584/topsdk/ability2973/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability2973 struct {
    Client *topsdk.TopClient
}

func NewAbility2973(client *topsdk.TopClient) *Ability2973{
    return &Ability2973{client}
}

/*
    淘宝客-服务商-红包领取状态查询
*/
func (ability *Ability2973) TaobaoTbkScVegasSendStatus(req *request.TaobaoTbkScVegasSendStatusRequest,session string) (*response.TaobaoTbkScVegasSendStatusResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability2973 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.tbk.sc.vegas.send.status",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTbkScVegasSendStatusResponse{}
    if(err != nil){
        log.Println("taobaoTbkScVegasSendStatus error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
