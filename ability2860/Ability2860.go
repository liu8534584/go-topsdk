package ability2860

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability2860/request"
    "github.com/liu8534584/topsdk/ability2860/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability2860 struct {
    Client *topsdk.TopClient
}

func NewAbility2860(client *topsdk.TopClient) *Ability2860{
    return &Ability2860{client}
}

/*
    淘宝客-服务商-渠道管理rid推广效果
*/
func (ability *Ability2860) TaobaoTbkScRidReportGet(req *request.TaobaoTbkScRidReportGetRequest,session string) (*response.TaobaoTbkScRidReportGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability2860 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.tbk.sc.rid.report.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTbkScRidReportGetResponse{}
    if(err != nil){
        log.Println("taobaoTbkScRidReportGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
