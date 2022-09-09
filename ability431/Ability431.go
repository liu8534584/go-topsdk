package ability431

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability431/request"
    "github.com/liu8534584/topsdk/ability431/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability431 struct {
    Client *topsdk.TopClient
}

func NewAbility431(client *topsdk.TopClient) *Ability431{
    return &Ability431{client}
}

/*
    淘宝客-推广者-商品链接转换
*/
func (ability *Ability431) TaobaoTbkItemConvert(req *request.TaobaoTbkItemConvertRequest) (*response.TaobaoTbkItemConvertResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability431 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.item.convert",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkItemConvertResponse{}
    if(err != nil){
        log.Println("taobaoTbkItemConvert error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    淘宝客-推广者-店铺链接转换
*/
func (ability *Ability431) TaobaoTbkShopConvert(req *request.TaobaoTbkShopConvertRequest) (*response.TaobaoTbkShopConvertResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability431 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.shop.convert",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkShopConvertResponse{}
    if(err != nil){
        log.Println("taobaoTbkShopConvert error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
