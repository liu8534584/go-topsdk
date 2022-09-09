package ability3057

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability3057/request"
    "github.com/liu8534584/topsdk/ability3057/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability3057 struct {
    Client *topsdk.TopClient
}

func NewAbility3057(client *topsdk.TopClient) *Ability3057{
    return &Ability3057{client}
}

/*
    淘宝客-服务商-拍立淘插件转链
*/
func (ability *Ability3057) TaobaoTbkScPailitaoWidgetConvert(req *request.TaobaoTbkScPailitaoWidgetConvertRequest,session string) (*response.TaobaoTbkScPailitaoWidgetConvertResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability3057 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.tbk.sc.pailitao.widget.convert",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTbkScPailitaoWidgetConvertResponse{}
    if(err != nil){
        log.Println("taobaoTbkScPailitaoWidgetConvert error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
