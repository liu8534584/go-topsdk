package ability370

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability370/request"
    "github.com/liu8534584/topsdk/ability370/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability370 struct {
    Client *topsdk.TopClient
}

func NewAbility370(client *topsdk.TopClient) *Ability370{
    return &Ability370{client}
}

/*
    淘宝客-推广者-物料搜索
*/
func (ability *Ability370) TaobaoTbkDgMaterialOptional(req *request.TaobaoTbkDgMaterialOptionalRequest) (*response.TaobaoTbkDgMaterialOptionalResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability370 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.dg.material.optional",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkDgMaterialOptionalResponse{}
    if(err != nil){
        log.Println("taobaoTbkDgMaterialOptional error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    淘宝客-推广者-店铺搜索
*/
func (ability *Ability370) TaobaoTbkShopGet(req *request.TaobaoTbkShopGetRequest) (*response.TaobaoTbkShopGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability370 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.shop.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkShopGetResponse{}
    if(err != nil){
        log.Println("taobaoTbkShopGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    淘宝客-推广者-物料搜索（临时接口）
*/
func (ability *Ability370) TaobaoTbkDgMaterialTemporaryOptional(req *request.TaobaoTbkDgMaterialTemporaryOptionalRequest) (*response.TaobaoTbkDgMaterialTemporaryOptionalResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability370 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.dg.material.temporary.optional",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkDgMaterialTemporaryOptionalResponse{}
    if(err != nil){
        log.Println("taobaoTbkDgMaterialTemporaryOptional error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
