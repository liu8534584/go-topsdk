package domain

type AlibabaAlscUnionKbItemStoreRelationQueryKbItemShopRelationDto struct {
	/*
	   门店ID（city_id不为空则返回当前城市门店，否则返回全部门店）     */
	StoreId *string `json:"store_id,omitempty" `
}

func (s *AlibabaAlscUnionKbItemStoreRelationQueryKbItemShopRelationDto) SetStoreId(v string) *AlibabaAlscUnionKbItemStoreRelationQueryKbItemShopRelationDto {
	s.StoreId = &v
	return s
}
