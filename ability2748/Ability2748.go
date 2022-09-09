package ability2748

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability2748/request"
    "github.com/liu8534584/topsdk/ability2748/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability2748 struct {
    Client *topsdk.TopClient
}

func NewAbility2748(client *topsdk.TopClient) *Ability2748{
    return &Ability2748{client}
}

/*
    淘宝客-服务商-直播物料精选
*/
func (ability *Ability2748) TaobaoTbkScLiveMaterialGet(req *request.TaobaoTbkScLiveMaterialGetRequest,session string) (*response.TaobaoTbkScLiveMaterialGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability2748 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.tbk.sc.live.material.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTbkScLiveMaterialGetResponse{}
    if(err != nil){
        log.Println("taobaoTbkScLiveMaterialGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    淘宝客-服务商-直播间详情查询
*/
func (ability *Ability2748) TaobaoTbkScLiveInfoGet(req *request.TaobaoTbkScLiveInfoGetRequest,session string) (*response.TaobaoTbkScLiveInfoGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability2748 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.tbk.sc.live.info.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTbkScLiveInfoGetResponse{}
    if(err != nil){
        log.Println("taobaoTbkScLiveInfoGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
