package ability2817

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability2817/request"
    "github.com/liu8534584/topsdk/ability2817/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability2817 struct {
    Client *topsdk.TopClient
}

func NewAbility2817(client *topsdk.TopClient) *Ability2817{
    return &Ability2817{client}
}

/*
    淘宝客-推广者-渠道管理组团明细数据
*/
func (ability *Ability2817) TaobaoTbkDgCpaReportSubDetail(req *request.TaobaoTbkDgCpaReportSubDetailRequest) (*response.TaobaoTbkDgCpaReportSubDetailResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability2817 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.dg.cpa.report.sub.detail",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkDgCpaReportSubDetailResponse{}
    if(err != nil){
        log.Println("taobaoTbkDgCpaReportSubDetail error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    淘宝客-推广者-渠道管理组团汇总数据
*/
func (ability *Ability2817) TaobaoTbkDgCpaReportDetail(req *request.TaobaoTbkDgCpaReportDetailRequest) (*response.TaobaoTbkDgCpaReportDetailResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability2817 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.dg.cpa.report.detail",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkDgCpaReportDetailResponse{}
    if(err != nil){
        log.Println("taobaoTbkDgCpaReportDetail error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
