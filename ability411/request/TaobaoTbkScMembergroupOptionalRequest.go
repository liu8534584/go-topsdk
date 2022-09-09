package request


type TaobaoTbkScMembergroupOptionalRequest struct {
    /*
        member组id     */
    MemberGroupId  *int64 `json:"member_group_id,omitempty" required:"false" `
    /*
        淘宝数字id     */
    TbNumIds  *string `json:"tb_num_ids,omitempty" required:"false" `
}

func (s *TaobaoTbkScMembergroupOptionalRequest) SetMemberGroupId(v int64) *TaobaoTbkScMembergroupOptionalRequest {
    s.MemberGroupId = &v
    return s
}
func (s *TaobaoTbkScMembergroupOptionalRequest) SetTbNumIds(v string) *TaobaoTbkScMembergroupOptionalRequest {
    s.TbNumIds = &v
    return s
}

func (req *TaobaoTbkScMembergroupOptionalRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.MemberGroupId != nil) {
        paramMap["member_group_id"] = *req.MemberGroupId
    }
    if(req.TbNumIds != nil) {
        paramMap["tb_num_ids"] = *req.TbNumIds
    }
    return paramMap
}

func (req *TaobaoTbkScMembergroupOptionalRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}