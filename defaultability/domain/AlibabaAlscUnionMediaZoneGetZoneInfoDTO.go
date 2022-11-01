package domain

type AlibabaAlscUnionMediaZoneGetZoneInfoDTO struct {
	/*
	   资源位名称     */
	Name *string `json:"name,omitempty" `

	/*
	   资源位pid     */
	Pid *string `json:"pid,omitempty" `
}

func (s *AlibabaAlscUnionMediaZoneGetZoneInfoDTO) SetName(v string) *AlibabaAlscUnionMediaZoneGetZoneInfoDTO {
	s.Name = &v
	return s
}
func (s *AlibabaAlscUnionMediaZoneGetZoneInfoDTO) SetPid(v string) *AlibabaAlscUnionMediaZoneGetZoneInfoDTO {
	s.Pid = &v
	return s
}
