package request

import (
        "github.com/liu8534584/topsdk/util"
    )

type TaobaoTbkDgPicEmbedRequest struct {
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
        返现比例，0-10000，最多支持6位小数     */
    UserRate  *string `json:"user_rate,omitempty" required:"false" `
}

func (s *TaobaoTbkDgPicEmbedRequest) SetAdzoneId(v int64) *TaobaoTbkDgPicEmbedRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkDgPicEmbedRequest) SetItemId(v int64) *TaobaoTbkDgPicEmbedRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkDgPicEmbedRequest) SetPicStream(v []byte) *TaobaoTbkDgPicEmbedRequest {
    s.PicStream = &v
    return s
}
func (s *TaobaoTbkDgPicEmbedRequest) SetRelationId(v string) *TaobaoTbkDgPicEmbedRequest {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkDgPicEmbedRequest) SetUserRate(v string) *TaobaoTbkDgPicEmbedRequest {
    s.UserRate = &v
    return s
}

func (req *TaobaoTbkDgPicEmbedRequest) ToMap() map[string]interface{} {
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
    if(req.UserRate != nil) {
        paramMap["user_rate"] = *req.UserRate
    }
    return paramMap
}

func (req *TaobaoTbkDgPicEmbedRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.PicStream != nil {
        fileMap["pic_stream"] = *req.PicStream
    }
    return fileMap
}