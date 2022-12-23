package ability2958

import (
	"errors"
	"github.com/liu8534584/topsdk"
	"github.com/liu8534584/topsdk/ability2958/request"
	"github.com/liu8534584/topsdk/ability2958/response"
	"github.com/liu8534584/topsdk/util"
	"log"
)

type Ability2958 struct {
	Client *topsdk.TopClient
}

func NewAbility2958(client *topsdk.TopClient) *Ability2958 {
	return &Ability2958{client}
}

// TaobaoTbkScGeneralLinkConvert 淘宝客-服务商-万能转链
func (ability *Ability2958) TaobaoTbkScGeneralLinkConvert(req *request.TaobaoTbkScGeneralLinkConvertRequest, session string) (*response.TaobaoTbkScGeneralLinkConvertResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2860 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.general.link.convert", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScGeneralLinkConvertResponse{}
	if err != nil {
		log.Println("taobaoTbkScRidReportGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
