package ability406

import (
	"errors"
	"github.com/liu8534584/topsdk"
	"github.com/liu8534584/topsdk/ability406/request"
	"github.com/liu8534584/topsdk/ability406/response"
	"github.com/liu8534584/topsdk/util"
	"log"
)

type Ability406 struct {
	Client *topsdk.TopClient
}

func NewAbility406(client *topsdk.TopClient) *Ability406 {
	return &Ability406{client}
}

/*
   淘宝客-服务商-店铺链接转换
*/
func (ability *Ability406) TaobaoTbkScShopConvert(req *request.TaobaoTbkScShopConvertRequest, session string) (*response.TaobaoTbkScShopConvertResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability406 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.shop.convert", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScShopConvertResponse{}
	if err != nil {
		log.Println("taobaoTbkScShopConvert error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
