package ability1556

import (
	"errors"
	"log"
	"topsdk"
	"topsdk/ability1556/request"
	"topsdk/ability1556/response"
	"topsdk/util"
)

type Ability1556 struct {
	Client *topsdk.TopClient
}

func NewAbility1556(client *topsdk.TopClient) *Ability1556 {
	return &Ability1556{client}
}

/*
   淘宝客-服务商-物料搜索（临时接口）
*/
func (ability *Ability1556) TaobaoTbkScMaterialTemporaryOptional(req *request.TaobaoTbkScMaterialTemporaryOptionalRequest, session string) (*response.TaobaoTbkScMaterialTemporaryOptionalResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability1556 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.material.temporary.optional", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScMaterialTemporaryOptionalResponse{}
	if err != nil {
		log.Println("taobaoTbkScMaterialTemporaryOptional error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-服务商-物料搜索
*/
func (ability *Ability1556) TaobaoTbkScMaterialOptional(req *request.TaobaoTbkScMaterialOptionalRequest, session string) (*response.TaobaoTbkScMaterialOptionalResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability1556 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.material.optional", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScMaterialOptionalResponse{}
	if err != nil {
		log.Println("taobaoTbkScMaterialOptional error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
