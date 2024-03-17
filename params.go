package leetcode_api

type GetQuestionsListParams struct {
	Offset       int
	Limit        int
	Filters      Filters
	CategorySlug CategorySlug
}
