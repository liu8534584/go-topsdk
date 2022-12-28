package domain

type CouponInfoDto struct {
	CouponEndTime     *string `json:"coupon_end_time"`
	ActivityId        *string `json:"activity_id"`
	CouponRemainCount *int    `json:"coupon_remain_count"`
	CouponAmount      *string `json:"coupon_amount"`
	CouponStartTime   *string `json:"coupon_start_time"`
	CouponDesc        *string `json:"coupon_desc"`
}

type LinkInfoDto struct {
	CouponLongUrl   *string `json:"coupon_long_url"`
	MaterialType    *int    `json:"material_type"`
	TpwdOriginUrl   *string `json:"tpwd_origin_url"`
	MaterialId      *string `json:"material_id"`
	CpsLongUrl      *string `json:"cps_long_url"`
	CpsShortTpwd    *string `json:"cps_short_tpwd"`
	TkBizType       *int    `json:"tk_biz_type"`
	CouponShortTpwd *string `json:"coupon_short_tpwd"`
	CpsShortUrl     *string `json:"cps_short_url"`
	CouponShortUrl  *string `json:"coupon_short_url"`
	CouponFullTpwd  *string `json:"coupon_full_tpwd"`
	CpsFullTpwd     *string `json:"cps_full_tpwd"`
}

type ItemLinkInfoDto struct {
	CouponLongUrl   *string `json:"coupon_long_url"`
	MaterialType    *int    `json:"material_type"`
	ItemId          *string `json:"item_id"`
	CpsLongUrl      *string `json:"cps_long_url"`
	CpsShortTpwd    *string `json:"cps_short_tpwd"`
	CouponShortTpwd *string `json:"coupon_short_tpwd"`
	CpsShortUrl     *string `json:"cps_short_url"`
	CouponShortUrl  *string `json:"coupon_short_url"`
	CouponFullTpwd  *string `json:"coupon_full_tpwd"`
	CpsFullTpwd     *string `json:"cps_full_tpwd"`
}

type PromotionInfoDto struct {
	CommissionRate *string `json:"commission_rate"`
}

type MaterialUrl struct {
	CouponInfoDto    *CouponInfoDto    `json:"coupon_info_dto"`
	LinkInfoDto      *LinkInfoDto      `json:"link_info_dto"`
	PromotionInfoDto *PromotionInfoDto `json:"promotion_info_dto"`
}

type ShopLinkInfoDto struct {
	MaterialType *int    `json:"material_type"`
	SellerId     *string `json:"seller_id"`
	CpsLongUrl   *string `json:"cps_long_url"`
	CpsShortTpwd *string `json:"cps_short_tpwd"`
	CpsShortUrl  *string `json:"cps_short_url"`
	CpsFullTpwd  *string `json:"cps_full_tpwd"`
}

type ShopUrlList struct {
	LinkInfoDto *ShopLinkInfoDto `json:"link_info_dto"`
}

type EventLinkInfoDto struct {
	MaterialType *int    `json:"material_type"`
	PageId       *string `json:"page_id"`
	CpsLongUrl   *string `json:"cps_long_url"`
	CpsShortTpwd *string `json:"cps_short_tpwd"`
	CpsShortUrl  *string `json:"cps_short_url"`
	CpsFullTpwd  *string `json:"cps_full_tpwd"`
}

type EventUrlList struct {
	LinkInfoDto *EventLinkInfoDto `json:"link_info_dto"`
}

type ItemUrlList struct {
	CouponInfoDto    *CouponInfoDto    `json:"coupon_info_dto"`
	LinkInfoDto      *ItemLinkInfoDto  `json:"link_info_dto"`
	PromotionInfoDto *PromotionInfoDto `json:"promotion_info_dto"`
}

type TaobaoTbkScGeneralLinkConvertResponse struct {
	MaterialUrlList *[]MaterialUrl  `json:"material_url_list"`
	ShopUrlList     *[]ShopUrlList  `json:"shop_url_list"`
	EventUrlList    *[]EventUrlList `json:"event_url_list"`
	ItemUrlList     *[]ItemUrlList  `json:"item_url_list"`
}
