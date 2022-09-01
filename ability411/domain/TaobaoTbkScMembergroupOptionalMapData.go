package domain

import (
	"topsdk/util"
)

type TaobaoTbkScMembergroupOptionalMapData struct {
	/*
	   组ID     */
	MemberGroupId *int64 `json:"member_group_id,omitempty" `

	/*
	   组内的成员ID     */
	TbNumIds *string `json:"tb_num_ids,omitempty" `

	/*
	   创建时间     */
	CreateTime *util.LocalTime `json:"create_time,omitempty" `

	/*
	   更新时间     */
	UpdateTime *util.LocalTime `json:"update_time,omitempty" `
}

func (s *TaobaoTbkScMembergroupOptionalMapData) SetMemberGroupId(v int64) *TaobaoTbkScMembergroupOptionalMapData {
	s.MemberGroupId = &v
	return s
}
func (s *TaobaoTbkScMembergroupOptionalMapData) SetTbNumIds(v string) *TaobaoTbkScMembergroupOptionalMapData {
	s.TbNumIds = &v
	return s
}
func (s *TaobaoTbkScMembergroupOptionalMapData) SetCreateTime(v util.LocalTime) *TaobaoTbkScMembergroupOptionalMapData {
	s.CreateTime = &v
	return s
}
func (s *TaobaoTbkScMembergroupOptionalMapData) SetUpdateTime(v util.LocalTime) *TaobaoTbkScMembergroupOptionalMapData {
	s.UpdateTime = &v
	return s
}
