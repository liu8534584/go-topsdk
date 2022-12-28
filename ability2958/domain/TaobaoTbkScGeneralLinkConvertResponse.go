package domain

type TaobaoTbkScGeneralLinkConvertResponse struct {
	MaterialUrlList struct {
		MaterialUrlList *[]struct {
			Msg              *string `json:"msg"`
			Code             *int    `json:"code"`
			PromotionInfoDto *struct {
				CommissionRate *string `json:"commission_rate"`
			} `json:"promotion_info_dto"`
			CouponInfoDto *struct {
				CouponEndTime     *string `json:"coupon_end_time"`
				ActivityId        *string `json:"activity_id"`
				CouponRemainCount *int    `json:"coupon_remain_count"`
				CouponAmount      *string `json:"coupon_amount"`
				CouponStartTime   *string `json:"coupon_start_time"`
				CouponDesc        *string `json:"coupon_desc"`
			} `json:"coupon_info_dto"`
			LinkInfoDto *struct {
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
			} `json:"link_info_dto"`
		} `json:"material_url_list"`
	} `json:"material_url_list"`
	ShopUrlList *struct {
		ShopUrlList *[]struct {
			Msg         *string `json:"msg"`
			Code        *int    `json:"code"`
			LinkInfoDto *struct {
				MaterialType *int    `json:"material_type"`
				SellerId     *string `json:"seller_id"`
				CpsLongUrl   *string `json:"cps_long_url"`
				CpsShortTpwd *string `json:"cps_short_tpwd"`
				CpsShortUrl  *string `json:"cps_short_url"`
				CpsFullTpwd  *string `json:"cps_full_tpwd"`
			} `json:"link_info_dto"`
		} `json:"shop_url_list"`
	} `json:"shop_url_list"`
	EventUrlList *struct {
		EventUrlList *[]struct {
			Msg         *string `json:"msg"`
			Code        *int    `json:"code"`
			LinkInfoDto *struct {
				MaterialType *int    `json:"material_type"`
				PageId       *string `json:"page_id"`
				CpsLongUrl   *string `json:"cps_long_url"`
				CpsShortTpwd *string `json:"cps_short_tpwd"`
				CpsShortUrl  *string `json:"cps_short_url"`
				CpsFullTpwd  *string `json:"cps_full_tpwd"`
			} `json:"link_info_dto"`
		} `json:"event_url_list"`
	} `json:"event_url_list"`
	ItemUrlList *struct {
		ItemUrlList *[]struct {
			Msg              *string `json:"msg"`
			Code             *int    `json:"code"`
			PromotionInfoDto *struct {
				CommissionRate *string `json:"commission_rate"`
			} `json:"promotion_info_dto"`
			CouponInfoDto *struct {
				CouponEndTime     *string `json:"coupon_end_time"`
				ActivityId        *string `json:"activity_id"`
				CouponRemainCount *int    `json:"coupon_remain_count"`
				CouponAmount      *string `json:"coupon_amount"`
				CouponStartTime   *string `json:"coupon_start_time"`
				CouponDesc        *string `json:"coupon_desc"`
			} `json:"coupon_info_dto"`
			LinkInfoDto *struct {
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
			} `json:"link_info_dto"`
		} `json:"item_url_list"`
	} `json:"item_url_list"`
}
