package ability379

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability379/request"
    "github.com/liu8534584/topsdk/ability379/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability379 struct {
    Client *topsdk.TopClient
}

func NewAbility379(client *topsdk.TopClient) *Ability379{
    return &Ability379{client}
}

/*
    淘宝客-推广者-商业化图片生成
*/
func (ability *Ability379) TaobaoTbkDgPicEmbed(req *request.TaobaoTbkDgPicEmbedRequest) (*response.TaobaoTbkDgPicEmbedResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability379 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.dg.pic.embed",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkDgPicEmbedResponse{}
    if(err != nil){
        log.Println("taobaoTbkDgPicEmbed error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
