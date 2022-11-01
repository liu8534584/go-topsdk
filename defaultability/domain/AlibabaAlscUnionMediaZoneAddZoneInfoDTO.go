package domain

type AlibabaAlscUnionMediaZoneAddZoneInfoDTO struct {
	/*
	   推广位名称     */
	Name *string `json:"name,omitempty" `

	/*
	   推广位pid     */
	Pid *string `json:"pid,omitempty" `
}

func (s *AlibabaAlscUnionMediaZoneAddZoneInfoDTO) SetName(v string) *AlibabaAlscUnionMediaZoneAddZoneInfoDTO {
	s.Name = &v
	return s
}
func (s *AlibabaAlscUnionMediaZoneAddZoneInfoDTO) SetPid(v string) *AlibabaAlscUnionMediaZoneAddZoneInfoDTO {
	s.Pid = &v
	return s
}
