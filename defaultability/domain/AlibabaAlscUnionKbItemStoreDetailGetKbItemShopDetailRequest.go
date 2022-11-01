package domain

type AlibabaAlscUnionKbItemStoreDetailGetKbItemShopDetailRequest struct {
	/*
	   门店ID     */
	StoreId *string `json:"store_id,omitempty" `
}

func (s *AlibabaAlscUnionKbItemStoreDetailGetKbItemShopDetailRequest) SetStoreId(v string) *AlibabaAlscUnionKbItemStoreDetailGetKbItemShopDetailRequest {
	s.StoreId = &v
	return s
}
