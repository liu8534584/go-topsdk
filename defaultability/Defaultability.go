package defaultability

import (
	"errors"
	"github.com/liu8534584/topsdk"
	"github.com/liu8534584/topsdk/defaultability/request"
	"github.com/liu8534584/topsdk/defaultability/response"
	"github.com/liu8534584/topsdk/util"
	"log"
)

type Defaultability struct {
	Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability {
	return &Defaultability{client}
}

/*
   关键词过滤匹配
*/
func (ability *Defaultability) TaobaoKfcKeywordSearch(req *request.TaobaoKfcKeywordSearchRequest, session string) (*response.TaobaoKfcKeywordSearchResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.kfc.keyword.search", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoKfcKeywordSearchResponse{}
	if err != nil {
		log.Println("taobaoKfcKeywordSearch error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-推广者-商品id升级（临时接口）
*/
func (ability *Defaultability) TaobaoTbkItemidTemporaryTransform(req *request.TaobaoTbkItemidTemporaryTransformRequest) (*response.TaobaoTbkItemidTemporaryTransformResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.itemid.temporary.transform", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkItemidTemporaryTransformResponse{}
	if err != nil {
		log.Println("taobaoTbkItemidTemporaryTransform error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-推广者-商品id转化（二方）（专有）
*/
func (ability *Defaultability) TaobaoTbkItemidPrivateTransform(req *request.TaobaoTbkItemidPrivateTransformRequest) (*response.TaobaoTbkItemidPrivateTransformResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.itemid.private.transform", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkItemidPrivateTransformResponse{}
	if err != nil {
		log.Println("taobaoTbkItemidPrivateTransform error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
