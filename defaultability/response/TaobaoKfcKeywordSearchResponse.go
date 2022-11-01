package response

type TaobaoKfcKeywordSearchResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   KFC 关键词过滤匹配结果
	*/
	KfcSearchResult string `json:"kfc_search_result,omitempty" `
}
