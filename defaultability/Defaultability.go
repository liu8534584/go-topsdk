package defaultability

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/defaultability/request"
    "github.com/liu8534584/topsdk/defaultability/response"
    "github.com/liu8534584/topsdk/util"
)

type Defaultability struct {
    Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability{
    return &Defaultability{client}
}

/*
    淘宝客-推广者-商品id升级（临时接口）
*/
func (ability *Defaultability) TaobaoTbkItemidTemporaryTransform(req *request.TaobaoTbkItemidTemporaryTransformRequest) (*response.TaobaoTbkItemidTemporaryTransformResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.itemid.temporary.transform",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkItemidTemporaryTransformResponse{}
    if(err != nil){
        log.Println("taobaoTbkItemidTemporaryTransform error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    淘宝客-推广者-商品id转化（二方）（专有）
*/
func (ability *Defaultability) TaobaoTbkItemidPrivateTransform(req *request.TaobaoTbkItemidPrivateTransformRequest) (*response.TaobaoTbkItemidPrivateTransformResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.itemid.private.transform",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkItemidPrivateTransformResponse{}
    if(err != nil){
        log.Println("taobaoTbkItemidPrivateTransform error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取前台展示的店铺类目
*/
func (ability *Defaultability) TaobaoShopcatsListGet(req *request.TaobaoShopcatsListGetRequest) (*response.TaobaoShopcatsListGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.shopcats.list.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoShopcatsListGetResponse{}
    if(err != nil){
        log.Println("taobaoShopcatsListGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取前台展示的店铺内卖家自定义商品类目
*/
func (ability *Defaultability) TaobaoSellercatsListGet(req *request.TaobaoSellercatsListGetRequest,session string) (*response.TaobaoSellercatsListGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercats.list.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercatsListGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercatsListGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
