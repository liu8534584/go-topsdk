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
   本地生活爆爆团选品筛选集合
*/
func (ability *Defaultability) AlibabaAlscUnionKbBbtItemPromotionFilterList(req *request.AlibabaAlscUnionKbBbtItemPromotionFilterListRequest) (*response.AlibabaAlscUnionKbBbtItemPromotionFilterListResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.bbt.item.promotion.filter.list", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbBbtItemPromotionFilterListResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbBbtItemPromotionFilterList error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地联盟爆爆团商品详情
*/
func (ability *Defaultability) AlibabaAlscUnionKbBbtItemDetailGet(req *request.AlibabaAlscUnionKbBbtItemDetailGetRequest) (*response.AlibabaAlscUnionKbBbtItemDetailGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.bbt.item.detail.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbBbtItemDetailGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbBbtItemDetailGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地联盟爆爆团商品门店关系
*/
func (ability *Defaultability) AlibabaAlscUnionKbBbtItemStoreRelationQuery(req *request.AlibabaAlscUnionKbBbtItemStoreRelationQueryRequest) (*response.AlibabaAlscUnionKbBbtItemStoreRelationQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.bbt.item.store.relation.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbBbtItemStoreRelationQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbBbtItemStoreRelationQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地联盟爆爆团门店详情
*/
func (ability *Defaultability) AlibabaAlscUnionKbBbtItemStoreDetailGet(req *request.AlibabaAlscUnionKbBbtItemStoreDetailGetRequest) (*response.AlibabaAlscUnionKbBbtItemStoreDetailGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.bbt.item.store.detail.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbBbtItemStoreDetailGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbBbtItemStoreDetailGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地生活媒体平台口碑选品筛选项集合
*/
func (ability *Defaultability) AlibabaAlscUnionKbItemPromotionFilterList(req *request.AlibabaAlscUnionKbItemPromotionFilterListRequest) (*response.AlibabaAlscUnionKbItemPromotionFilterListResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.item.promotion.filter.list", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbItemPromotionFilterListResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbItemPromotionFilterList error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地生活媒体推广位查询
*/
func (ability *Defaultability) AlibabaAlscUnionMediaZoneGet(req *request.AlibabaAlscUnionMediaZoneGetRequest) (*response.AlibabaAlscUnionMediaZoneGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.media.zone.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionMediaZoneGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionMediaZoneGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地生活媒体平台口碑选品
*/
func (ability *Defaultability) AlibabaAlscUnionKbItemPromotion(req *request.AlibabaAlscUnionKbItemPromotionRequest) (*response.AlibabaAlscUnionKbItemPromotionResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.item.promotion", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbItemPromotionResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbItemPromotion error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   加密方法
*/
func (ability *Defaultability) AlibabaAlscUnionKbCommonEncrypt(req *request.AlibabaAlscUnionKbCommonEncryptRequest) (*response.AlibabaAlscUnionKbCommonEncryptResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.common.encrypt", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbCommonEncryptResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbCommonEncrypt error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地生活媒体推广位创建
*/
func (ability *Defaultability) AlibabaAlscUnionMediaZoneAdd(req *request.AlibabaAlscUnionMediaZoneAddRequest) (*response.AlibabaAlscUnionMediaZoneAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.media.zone.add", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionMediaZoneAddResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionMediaZoneAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地联盟口碑商品列表
*/
func (ability *Defaultability) AlibabaAlscUnionKbItemQuery(req *request.AlibabaAlscUnionKbItemQueryRequest) (*response.AlibabaAlscUnionKbItemQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.item.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbItemQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbItemQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   openapi预下单订单支付
*/
func (ability *Defaultability) AlibabaAlscUnionKbOrderPay(req *request.AlibabaAlscUnionKbOrderPayRequest) (*response.AlibabaAlscUnionKbOrderPayResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.order.pay", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbOrderPayResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbOrderPay error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地生活媒体推广口碑CPA用户反作弊订单数据报表
*/
func (ability *Defaultability) AlibabaAlscUnionKbcpaPunishOrderGet(req *request.AlibabaAlscUnionKbcpaPunishOrderGetRequest) (*response.AlibabaAlscUnionKbcpaPunishOrderGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kbcpa.punish.order.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbcpaPunishOrderGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbcpaPunishOrderGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   获取用户已开通消息
*/
func (ability *Defaultability) TaobaoTmcUserGet(req *request.TaobaoTmcUserGetRequest) (*response.TaobaoTmcUserGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.user.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcUserGetResponse{}
	if err != nil {
		log.Println("taobaoTmcUserGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地联盟饿了么单店推广店铺列表
*/
func (ability *Defaultability) AlibabaAlscUnionElemePromotionStorepromotionQuery(req *request.AlibabaAlscUnionElemePromotionStorepromotionQueryRequest) (*response.AlibabaAlscUnionElemePromotionStorepromotionQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.promotion.storepromotion.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemePromotionStorepromotionQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemePromotionStorepromotionQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地联盟饿了么单店推广单店铺查询
*/
func (ability *Defaultability) AlibabaAlscUnionElemePromotionStorepromotionGet(req *request.AlibabaAlscUnionElemePromotionStorepromotionGetRequest) (*response.AlibabaAlscUnionElemePromotionStorepromotionGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.promotion.storepromotion.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemePromotionStorepromotionGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemePromotionStorepromotionGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地联盟饿了么推广官方活动查询
*/
func (ability *Defaultability) AlibabaAlscUnionElemePromotionOfficialactivityGet(req *request.AlibabaAlscUnionElemePromotionOfficialactivityGetRequest) (*response.AlibabaAlscUnionElemePromotionOfficialactivityGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.eleme.promotion.officialactivity.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionElemePromotionOfficialactivityGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionElemePromotionOfficialactivityGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   发布单条消息
*/
func (ability *Defaultability) TaobaoTmcMessageProduce(req *request.TaobaoTmcMessageProduceRequest) (*response.TaobaoTmcMessageProduceResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.message.produce", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcMessageProduceResponse{}
	if err != nil {
		log.Println("taobaoTmcMessageProduce error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地生活媒体推广口碑CPA用户维权订单数据报表
*/
func (ability *Defaultability) AlibabaAlscUnionKbcpaRefundOrderGet(req *request.AlibabaAlscUnionKbcpaRefundOrderGetRequest) (*response.AlibabaAlscUnionKbcpaRefundOrderGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kbcpa.refund.order.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbcpaRefundOrderGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbcpaRefundOrderGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地联盟口碑商户列表
*/
func (ability *Defaultability) AlibabaAlscUnionKbStoreQuery(req *request.AlibabaAlscUnionKbStoreQueryRequest) (*response.AlibabaAlscUnionKbStoreQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.store.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbStoreQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbStoreQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地生活媒体创建商品推广链接
*/
func (ability *Defaultability) AlibabaAlscUnionKbItemPromotionShareCreate(req *request.AlibabaAlscUnionKbItemPromotionShareCreateRequest) (*response.AlibabaAlscUnionKbItemPromotionShareCreateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.item.promotion.share.create", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbItemPromotionShareCreateResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbItemPromotionShareCreate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地联盟口碑门店商品列表
*/
func (ability *Defaultability) AlibabaAlscUnionKbStoreItemQuery(req *request.AlibabaAlscUnionKbStoreItemQueryRequest) (*response.AlibabaAlscUnionKbStoreItemQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.store.item.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbStoreItemQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbStoreItemQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   取消用户的消息服务
*/
func (ability *Defaultability) TaobaoTmcUserCancel(req *request.TaobaoTmcUserCancelRequest) (*response.TaobaoTmcUserCancelResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.user.cancel", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcUserCancelResponse{}
	if err != nil {
		log.Println("taobaoTmcUserCancel error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地联盟爆爆团商品列表
*/
func (ability *Defaultability) AlibabaAlscUnionKbBbtItemQuery(req *request.AlibabaAlscUnionKbBbtItemQueryRequest) (*response.AlibabaAlscUnionKbBbtItemQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.bbt.item.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbBbtItemQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbBbtItemQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   为已授权的用户开通消息服务
*/
func (ability *Defaultability) TaobaoTmcUserPermit(req *request.TaobaoTmcUserPermitRequest, session string) (*response.TaobaoTmcUserPermitResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tmc.user.permit", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTmcUserPermitResponse{}
	if err != nil {
		log.Println("taobaoTmcUserPermit error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   openapi订单创建
*/
func (ability *Defaultability) AlibabaAlscUnionKbOrderCreate(req *request.AlibabaAlscUnionKbOrderCreateRequest) (*response.AlibabaAlscUnionKbOrderCreateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.order.create", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbOrderCreateResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbOrderCreate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   openapi订单查询
*/
func (ability *Defaultability) AlibabaAlscUnionKbOrderQuery(req *request.AlibabaAlscUnionKbOrderQueryRequest) (*response.AlibabaAlscUnionKbOrderQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.order.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbOrderQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbOrderQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   openapi订单售中退
*/
func (ability *Defaultability) AlibabaAlscUnionKbOrderRefund(req *request.AlibabaAlscUnionKbOrderRefundRequest) (*response.AlibabaAlscUnionKbOrderRefundResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.order.refund", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbOrderRefundResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbOrderRefund error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地生活媒体推广订单明细查询
*/
func (ability *Defaultability) AlibabaAlscUnionKbcpaOrderDetailsGet(req *request.AlibabaAlscUnionKbcpaOrderDetailsGetRequest) (*response.AlibabaAlscUnionKbcpaOrderDetailsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kbcpa.order.details.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbcpaOrderDetailsGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbcpaOrderDetailsGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地联盟口碑商品详情
*/
func (ability *Defaultability) AlibabaAlscUnionKbItemDetailGet(req *request.AlibabaAlscUnionKbItemDetailGetRequest) (*response.AlibabaAlscUnionKbItemDetailGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.item.detail.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbItemDetailGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbItemDetailGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地联盟口碑商品门店关系
*/
func (ability *Defaultability) AlibabaAlscUnionKbItemStoreRelationQuery(req *request.AlibabaAlscUnionKbItemStoreRelationQueryRequest) (*response.AlibabaAlscUnionKbItemStoreRelationQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.item.store.relation.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbItemStoreRelationQueryResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbItemStoreRelationQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地生活媒体推广订单明细报表查询
*/
func (ability *Defaultability) AlibabaAlscUnionKbcpxPositiveOrderGet(req *request.AlibabaAlscUnionKbcpxPositiveOrderGetRequest) (*response.AlibabaAlscUnionKbcpxPositiveOrderGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kbcpx.positive.order.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbcpxPositiveOrderGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbcpxPositiveOrderGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地联盟口碑门店详情
*/
func (ability *Defaultability) AlibabaAlscUnionKbItemStoreDetailGet(req *request.AlibabaAlscUnionKbItemStoreDetailGetRequest) (*response.AlibabaAlscUnionKbItemStoreDetailGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kb.item.store.detail.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbItemStoreDetailGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbItemStoreDetailGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地生活媒体推广用户维权订单数据报表
*/
func (ability *Defaultability) AlibabaAlscUnionKbcpxRefundOrderGet(req *request.AlibabaAlscUnionKbcpxRefundOrderGetRequest) (*response.AlibabaAlscUnionKbcpxRefundOrderGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kbcpx.refund.order.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbcpxRefundOrderGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbcpxRefundOrderGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   本地生活媒体推广用户反作弊订单数据报表
*/
func (ability *Defaultability) AlibabaAlscUnionKbcpxPunishOrderGet(req *request.AlibabaAlscUnionKbcpxPunishOrderGetRequest) (*response.AlibabaAlscUnionKbcpxPunishOrderGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("alibaba.alsc.union.kbcpx.punish.order.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.AlibabaAlscUnionKbcpxPunishOrderGetResponse{}
	if err != nil {
		log.Println("alibabaAlscUnionKbcpxPunishOrderGet error", err)
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

/*
   获取前台展示的店铺类目
*/
func (ability *Defaultability) TaobaoShopcatsListGet(req *request.TaobaoShopcatsListGetRequest) (*response.TaobaoShopcatsListGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.shopcats.list.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoShopcatsListGetResponse{}
	if err != nil {
		log.Println("taobaoShopcatsListGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   获取前台展示的店铺内卖家自定义商品类目
*/
func (ability *Defaultability) TaobaoSellercatsListGet(req *request.TaobaoSellercatsListGetRequest, session string) (*response.TaobaoSellercatsListGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.sellercats.list.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSellercatsListGetResponse{}
	if err != nil {
		log.Println("taobaoSellercatsListGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
