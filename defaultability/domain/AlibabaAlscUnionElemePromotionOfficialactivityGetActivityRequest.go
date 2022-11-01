package domain

type AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest struct {
	/*
	   渠道PID     */
	Pid *string `json:"pid,omitempty" `

	/*
	   活动ID     */
	ActivityId *string `json:"activity_id,omitempty" `

	/*
	   三方会员id     */
	Sid *string `json:"sid,omitempty" `

	/*
	   是否返回微信推广图片     */
	IncludeWxImg *bool `json:"include_wx_img,omitempty" `
}

func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest) SetPid(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest {
	s.Pid = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest) SetActivityId(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest {
	s.ActivityId = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest) SetSid(v string) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest {
	s.Sid = &v
	return s
}
func (s *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest) SetIncludeWxImg(v bool) *AlibabaAlscUnionElemePromotionOfficialactivityGetActivityRequest {
	s.IncludeWxImg = &v
	return s
}
