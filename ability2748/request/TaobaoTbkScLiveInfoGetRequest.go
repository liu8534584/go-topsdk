package request


type TaobaoTbkScLiveInfoGetRequest struct {
    /*
        以逗号分隔的直播间ID列表，不可超过10个。注意：直播间ID不得与主播昵称（nicknames）同时使用，否则返回参数非法。     */
    LiveIds  *string `json:"live_ids,omitempty" required:"false" `
    /*
        mm_xxx_xxx_xxx的第2段数字     */
    SiteId  *int64 `json:"site_id" required:"true" `
    /*
        mm_xxx_xxx_xxx的第3段数字     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        渠道关系ID     */
    RelationId  *int64 `json:"relation_id,omitempty" required:"false" `
    /*
        以逗号分隔的主播昵称列表，不可超过5个。注意：主播昵称不得与直播间id（live_ids）同时使用，否则返回参数非法。     */
    Nicknames  *string `json:"nicknames,omitempty" required:"false" `
    /*
        指定查询直播间类型-1为自播直播间，此时不得填写主播昵称和直播间ID；2为转播直播间，此时上述两个条件必须至少一个非空。注意必须与入参中的过滤条件对齐。默认为1（自播直播间） defalutValue��1    */
    LiveQueryType  *int64 `json:"live_query_type,omitempty" required:"false" `
    /*
        页面元素数量，从0开始，默认20，不得超过50 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        页码，从0开始 defalutValue��0    */
    PageNum  *int64 `json:"page_num,omitempty" required:"false" `
    /*
        按照直播间状态查询，0为预告，1为开播中，2为已结束，多个状态用英文逗号分隔，默认返回预告及开播中的直播间     */
    RoomStatus  *string `json:"room_status,omitempty" required:"false" `
}

func (s *TaobaoTbkScLiveInfoGetRequest) SetLiveIds(v string) *TaobaoTbkScLiveInfoGetRequest {
    s.LiveIds = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetRequest) SetSiteId(v int64) *TaobaoTbkScLiveInfoGetRequest {
    s.SiteId = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetRequest) SetAdzoneId(v int64) *TaobaoTbkScLiveInfoGetRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetRequest) SetRelationId(v int64) *TaobaoTbkScLiveInfoGetRequest {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetRequest) SetNicknames(v string) *TaobaoTbkScLiveInfoGetRequest {
    s.Nicknames = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetRequest) SetLiveQueryType(v int64) *TaobaoTbkScLiveInfoGetRequest {
    s.LiveQueryType = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetRequest) SetPageSize(v int64) *TaobaoTbkScLiveInfoGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetRequest) SetPageNum(v int64) *TaobaoTbkScLiveInfoGetRequest {
    s.PageNum = &v
    return s
}
func (s *TaobaoTbkScLiveInfoGetRequest) SetRoomStatus(v string) *TaobaoTbkScLiveInfoGetRequest {
    s.RoomStatus = &v
    return s
}

func (req *TaobaoTbkScLiveInfoGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.LiveIds != nil) {
        paramMap["live_ids"] = *req.LiveIds
    }
    if(req.SiteId != nil) {
        paramMap["site_id"] = *req.SiteId
    }
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.RelationId != nil) {
        paramMap["relation_id"] = *req.RelationId
    }
    if(req.Nicknames != nil) {
        paramMap["nicknames"] = *req.Nicknames
    }
    if(req.LiveQueryType != nil) {
        paramMap["live_query_type"] = *req.LiveQueryType
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNum != nil) {
        paramMap["page_num"] = *req.PageNum
    }
    if(req.RoomStatus != nil) {
        paramMap["room_status"] = *req.RoomStatus
    }
    return paramMap
}

func (req *TaobaoTbkScLiveInfoGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}