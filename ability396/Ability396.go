package ability396

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability396/request"
    "github.com/liu8534584/topsdk/ability396/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability396 struct {
    Client *topsdk.TopClient
}

func NewAbility396(client *topsdk.TopClient) *Ability396{
    return &Ability396{client}
}

/*
    本地生活饿了么淘客报表对外披露
*/
func (ability *Ability396) TaobaoCpsreportAdGet(req *request.TaobaoCpsreportAdGetRequest) (*response.TaobaoCpsreportAdGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability396 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.cpsreport.ad.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoCpsreportAdGetResponse{}
    if(err != nil){
        log.Println("taobaoCpsreportAdGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    饿了么淘客结算数据明细接口
*/
func (ability *Ability396) AlibabaEletkSettlesTagGet(req *request.AlibabaEletkSettlesTagGetRequest) (*response.AlibabaEletkSettlesTagGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability396 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.eletk.settles.tag.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaEletkSettlesTagGetResponse{}
    if(err != nil){
        log.Println("alibabaEletkSettlesTagGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    饿了么淘客结算汇总数据披露
*/
func (ability *Ability396) AlibabaEletkSettlesSumGet(req *request.AlibabaEletkSettlesSumGetRequest) (*response.AlibabaEletkSettlesSumGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability396 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.eletk.settles.sum.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaEletkSettlesSumGetResponse{}
    if(err != nil){
        log.Println("alibabaEletkSettlesSumGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    微信Url Scheme链接获取
*/
func (ability *Ability396) AlibabaEletkActivitylinkWechatGet(req *request.AlibabaEletkActivitylinkWechatGetRequest) (*response.AlibabaEletkActivitylinkWechatGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability396 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.eletk.activitylink.wechat.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaEletkActivitylinkWechatGetResponse{}
    if(err != nil){
        log.Println("alibabaEletkActivitylinkWechatGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
