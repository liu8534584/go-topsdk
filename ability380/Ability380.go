package ability380

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability380/request"
    "github.com/liu8534584/topsdk/ability380/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability380 struct {
    Client *topsdk.TopClient
}

func NewAbility380(client *topsdk.TopClient) *Ability380{
    return &Ability380{client}
}

/*
    淘宝客-推广者-拍立淘插件转链
*/
func (ability *Ability380) TaobaoTbkDgPailitaoWidgetConvert(req *request.TaobaoTbkDgPailitaoWidgetConvertRequest) (*response.TaobaoTbkDgPailitaoWidgetConvertResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability380 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.dg.pailitao.widget.convert",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkDgPailitaoWidgetConvertResponse{}
    if(err != nil){
        log.Println("taobaoTbkDgPailitaoWidgetConvert error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
