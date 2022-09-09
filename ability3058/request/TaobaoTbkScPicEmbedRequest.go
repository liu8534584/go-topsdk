package request

import (
        "github.com/liu8534584/topsdk/util"
    )

type TaobaoTbkScPicEmbedRequest struct {
    /*
        mm_xxx_xxx_xxx的第三位     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        需打码的图片     */
    PicStream  *[]byte `json:"pic_stream,omitempty" required:"false" `
    /*
        渠道管理ID     */
    RelationId  *string `json:"relation_id,omitempty" required:"false" `
    /*
        mm_xxx_xxx_xxx的第二位     */
    SiteId  *int64 `json:"site_id" required:"true" `
}

func (s *TaobaoTbkScPicEmbedRequest) SetAdzoneId(v int64) *TaobaoTbkScPicEmbedRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkScPicEmbedRequest) SetItemId(v int64) *TaobaoTbkScPicEmbedRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkScPicEmbedRequest) SetPicStream(v []byte) *TaobaoTbkScPicEmbedRequest {
    s.PicStream = &v
    return s
}
func (s *TaobaoTbkScPicEmbedRequest) SetRelationId(v string) *TaobaoTbkScPicEmbedRequest {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkScPicEmbedRequest) SetSiteId(v int64) *TaobaoTbkScPicEmbedRequest {
    s.SiteId = &v
    return s
}

func (req *TaobaoTbkScPicEmbedRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.RelationId != nil) {
        paramMap["relation_id"] = *req.RelationId
    }
    if(req.SiteId != nil) {
        paramMap["site_id"] = *req.SiteId
    }
    return paramMap
}

func (req *TaobaoTbkScPicEmbedRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.PicStream != nil {
        fileMap["pic_stream"] = *req.PicStream
    }
    return fileMap
}