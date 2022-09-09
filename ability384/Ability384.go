package ability384

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability384/request"
    "github.com/liu8534584/topsdk/ability384/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability384 struct {
    Client *topsdk.TopClient
}

func NewAbility384(client *topsdk.TopClient) *Ability384{
    return &Ability384{client}
}

/*
    淘宝客-推广者-渠道管理rid推广效果
*/
func (ability *Ability384) TaobaoTbkDgRidReportGet(req *request.TaobaoTbkDgRidReportGetRequest) (*response.TaobaoTbkDgRidReportGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability384 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.dg.rid.report.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkDgRidReportGetResponse{}
    if(err != nil){
        log.Println("taobaoTbkDgRidReportGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
