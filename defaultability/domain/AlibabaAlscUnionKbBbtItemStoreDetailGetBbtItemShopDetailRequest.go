package domain

type AlibabaAlscUnionKbBbtItemStoreDetailGetBbtItemShopDetailRequest struct {
	/*
	   门店ID     */
	StoreId *string `json:"store_id,omitempty" `
}

func (s *AlibabaAlscUnionKbBbtItemStoreDetailGetBbtItemShopDetailRequest) SetStoreId(v string) *AlibabaAlscUnionKbBbtItemStoreDetailGetBbtItemShopDetailRequest {
	s.StoreId = &v
	return s
}
