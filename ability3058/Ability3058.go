package ability3058

import (
    "log"
    "errors"
    "github.com/liu8534584/topsdk"
    "github.com/liu8534584/topsdk/ability3058/request"
    "github.com/liu8534584/topsdk/ability3058/response"
    "github.com/liu8534584/topsdk/util"
)

type Ability3058 struct {
    Client *topsdk.TopClient
}

func NewAbility3058(client *topsdk.TopClient) *Ability3058{
    return &Ability3058{client}
}

/*
    淘宝客-服务商-商业化图片生成
*/
func (ability *Ability3058) TaobaoTbkScPicEmbed(req *request.TaobaoTbkScPicEmbedRequest,session string) (*response.TaobaoTbkScPicEmbedResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability3058 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.tbk.sc.pic.embed",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTbkScPicEmbedResponse{}
    if(err != nil){
        log.Println("taobaoTbkScPicEmbed error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
