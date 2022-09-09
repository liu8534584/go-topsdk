package ability374

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability374/request"
    "github.com/liu8534584/topsdk/ability374/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability374 struct {
    Client *topsdk.TopClient
}

func NewAbility374(client *topsdk.TopClient) *Ability374{
    return &Ability374{client}
}

/*
    淘宝客-推广者-官方活动转链
*/
func (ability *Ability374) TaobaoTbkActivityInfoGet(req *request.TaobaoTbkActivityInfoGetRequest) (*response.TaobaoTbkActivityInfoGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability374 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.activity.info.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkActivityInfoGetResponse{}
    if(err != nil){
        log.Println("taobaoTbkActivityInfoGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
